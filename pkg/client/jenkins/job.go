package client

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

type BuildTrigger struct {
	ChildProjects string `xml:"childProjects"`
}

type Publishers struct {
	BuildTrigger []BuildTrigger `xml:"hudson.tasks.BuildTrigger"`
}

type GitUserRemoteConfig struct {
	URL string `xml:"url"`
}

type UserRemoteConfigs struct {
	UserRemoteConfig []GitUserRemoteConfig `xml:"hudson.plugins.git.UserRemoteConfig"`
}

type BranchSpec struct {
	Name string `xml:"name"`
}

type Branches struct {
	BranchSpec []BranchSpec `xml:"hudson.plugins.git.BranchSpec"`
}

type StringParameterDefinition struct {
	Name         string `xml:"name"`
	DefaultValue string `xml:"defaultValue"`
}

type A struct {
	AString []string `xml:"string"`
}

type Choices struct {
	A []A `xml:"a"`
}

type BooleanParameterDefinition struct {
	Name         string `xml:"name"`
	DefaultValue string `xml:"defaultValue"`
}

type ChoiceParameterDefinition struct {
	Name    string    `xml:"name"`
	Choices []Choices `xml:"choices"`
}

type ParameterDefinitions struct {
	StringParameterDefinition  []StringParameterDefinition  `xml:"hudson.model.StringParameterDefinition"`
	BooleanParameterDefinition []BooleanParameterDefinition `xml:"hudson.model.BooleanParameterDefinition"`
	ChoiceParameterDefinition  []ChoiceParameterDefinition  `xml:"hudson.model.ChoiceParameterDefinition"`
}

type ParametersDefinitionProperty struct {
	ParameterDefinitions []ParameterDefinitions `xml:"parameterDefinitions"`
}

type SCM struct {
	BuildTrigger []UserRemoteConfigs `xml:"userRemoteConfigs"`
	Branches     []Branches          `xml:"branches"`
}

type Properties struct {
	ParametersDefinitionProperty []ParametersDefinitionProperty `xml:"hudson.model.ParametersDefinitionProperty"`
}

type Project struct {
	Publishers []Publishers `xml:"publishers"`
	SCM        []SCM        `xml:"scm"`
	Properties []Properties `xml:"properties"`
}

type Client interface {
	List(string)
}

type Pipeline struct {
	Key      string                   `json:"key"`
	Label    string                   `json:"label"`
	Path     string                   `json:"path"`
	Type     string                   `json:"type"`
	Children *map[string]PipelineView `json:"children"`
}

type PipelineView struct {
	Key      string                  `json:"key"`
	Label    string                  `json:"label"`
	Path     string                  `json:"path"`
	Type     string                  `json:"type"`
	Children *map[string]PipelineJob `json:"children"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type PipelineJob struct {
	Key         string      `json:"key"`
	Label       string      `json:"label"`
	Path        string      `json:"path"`
	Type        string      `json:"type"`
	Git         string      `json:"git"`
	Branch      string      `json:"branch"`
	Timestamp   int         `json:"timestamp"`
	DisplayName string      `json:"displayName"`
	Status      string      `json:"status"`
	Env         string      `json:"env"`
	Trigger     string      `json:"trigger"`
	Params      []Parameter `json:"params"`
}

type JenkinsResp struct {
	Jobs []struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Color       string `json:"color"`
		Timestamp   int    `json:"timestamp"`
		DisplayName string `json:"displayName"`
	} `json:"jobs"`
	Views []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"views"`
}

func runJobs(j *Jenkins, jobName string, query map[string][]string) (*http.Response, error) {
	req := &http.Request{}
	if len(query) > 0 {
		req, _ = http.NewRequest(http.MethodPost, j.Host+"/job/"+jobName+"/buildWithParameters", nil)
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, strings.Join(v, ","))
		}
		req.URL.RawQuery = q.Encode()
	} else {
		req, _ = http.NewRequest(http.MethodPost, j.Host+"/job/"+jobName+"/build", nil)
	}

	if j.Token != "" {
		req.Header.Add("Authorization", "Basic "+j.Token)
	}
	return j.Client.Do(req)
}

func listJobs(j *Jenkins) (*Pipeline, error) {
	list := JenkinsResp{}
	if err := listJenkinsViews(j, &list); err != nil {
		return nil, err
	}

	if err := listJenkinsJobs(j, &list); err != nil {
		return nil, err
	}

	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	mutex1 := sync.Mutex{}
	mutex2 := sync.Mutex{}
	mutex3 := sync.Mutex{}
	mutex4 := sync.Mutex{}
	mutex5 := sync.Mutex{}

	jobViewsMap := make(map[string][]string)
	m := make(map[string]PipelineView)
	p := Pipeline{
		Key:      "pipelines",
		Label:    "🔁 파이프라인",
		Path:     "/pipelines",
		Type:     "dir",
		Children: &m,
	}

	for _, v := range list.Views {
		wg1.Add(1)
		go func(jenkins *Jenkins, viewName string, jobViewsMap map[string][]string) error {
			defer wg1.Done()
			mutex1.Lock()
			m := make(map[string]PipelineJob)
			(*p.Children)[v.Name] = PipelineView{
				Type:     "dir",
				Key:      viewName,
				Label:    viewName,
				Path:     "/pipelines/" + viewName,
				Children: &m,
			}
			mutex1.Unlock()
			r, err := listJobsInView(jenkins, viewName)
			if err != nil {
				fmt.Println(err)
				return err
			}

			for _, j := range r.Jobs {
				mutex2.Lock()

				if j.Color != "disabled" {
					(*(*p.Children)[v.Name].Children)[j.Name] = PipelineJob{
						Key:    j.Name,
						Label:  j.Name,
						Type:   "file",
						Path:   "/pipelines/" + viewName + "/" + j.Name,
						Status: j.Color,
						Params: []Parameter{},
					}
				}
				mutex2.Unlock()

				mutex3.Lock()
				if j.Color != "disabled" {
					jobViewsMap[j.Name] = append(jobViewsMap[j.Name], v.Name)
				}
				mutex3.Unlock()
			}

			return nil
		}(j, v.Name, jobViewsMap)
	}

	wg1.Wait()

	for _, job := range list.Jobs {
		wg2.Add(1)
		go func(j *Jenkins, r *JenkinsResp, jobName string) {
			defer wg2.Done()
			displayName, timestamp, err := getJobLastBuild(j, jobName)
			proj := getJobConfigXML(j, jobName)
			if err != nil {
				fmt.Println(err)
			}

			mutex4.Lock()
			for _, view := range jobViewsMap[jobName] {
				tmp := (*(*p.Children)[view].Children)[jobName]
				tmp.Timestamp = timestamp
				tmp.DisplayName = displayName
				if len(proj.SCM) > 0 {
					if len(proj.SCM[0].Branches) > 0 {
						if len(proj.SCM[0].Branches) > 0 {
							if len(proj.SCM[0].Branches[0].BranchSpec) > 0 {
								tmp.Branch = proj.SCM[0].Branches[0].BranchSpec[0].Name
							}
						}
					}
					if len(proj.SCM[0].BuildTrigger) > 0 {
						if len(proj.SCM[0].BuildTrigger[0].UserRemoteConfig) > 0 {
							tmp.Git = proj.SCM[0].BuildTrigger[0].UserRemoteConfig[0].URL
						}
					}
				}
				if len(proj.Publishers) > 0 {
					if len(proj.Publishers[0].BuildTrigger) > 0 {
						tmp.Trigger = proj.Publishers[0].BuildTrigger[0].ChildProjects
					}
				}
				if len(proj.Properties) > 0 {
					if len(proj.Properties[0].ParametersDefinitionProperty) > 0 {
						if len(proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions) > 0 {
							if len(proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].StringParameterDefinition) > 0 {
								for i := range proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].StringParameterDefinition {
									mutex5.Lock()
									tmp.Params = append(tmp.Params, Parameter{
										Key:   proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].StringParameterDefinition[i].Name,
										Value: proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].StringParameterDefinition[i].DefaultValue,
										Type:  "string",
									})
									mutex5.Unlock()
									if proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].StringParameterDefinition[i].Name == "ENVIRONMENT" {
										tmp.Env = proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].StringParameterDefinition[i].DefaultValue
									}
								}
							}
							if len(proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].ChoiceParameterDefinition) > 0 {
								for i := range proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].ChoiceParameterDefinition {
									mutex5.Lock()
									values := []string{}
									if len(proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].ChoiceParameterDefinition[i].Choices[0].A) > 0 {
										values = append(values, proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].ChoiceParameterDefinition[i].Choices[0].A[0].AString...)
									}
									tmp.Params = append(tmp.Params, Parameter{
										Key:   proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].ChoiceParameterDefinition[i].Name,
										Value: strings.Join(values, ","),
										Type:  "choice",
									})
									mutex5.Unlock()
								}
							}
							if len(proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].BooleanParameterDefinition) > 0 {
								for _, b := range proj.Properties[0].ParametersDefinitionProperty[0].ParameterDefinitions[0].BooleanParameterDefinition {
									mutex5.Lock()
									tmp.Params = append(tmp.Params, Parameter{
										Key:   b.Name,
										Value: b.DefaultValue,
										Type:  "boolean",
									})
									mutex5.Unlock()
								}
							}
						}
					}
				}
				(*(*p.Children)[view].Children)[jobName] = tmp
			}
			mutex4.Unlock()
		}(j, &list, job.Name)
	}

	wg2.Wait()
	return &p, nil
}

func getJobLastBuild(j *Jenkins, jobName string) (displayName string, timestamp int, err error) {
	req, _ := http.NewRequest(http.MethodGet, j.Host+"/job/"+jobName+"/api/json", nil)

	q := req.URL.Query()
	q.Add("tree", "lastBuild[displayName,timestamp]")
	req.URL.RawQuery = q.Encode()

	if j.Token != "" {
		req.Header.Add("Authorization", "Basic "+j.Token)
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		return "", 0, err
	}

	b, _ := io.ReadAll(resp.Body)

	tmp := struct {
		LastBuild struct {
			Timestamp   int    `json:"timestamp"`
			DisplayName string `json:"displayName"`
		} `json:"lastBuild"`
	}{}
	json.Unmarshal(b, &tmp)

	return tmp.LastBuild.DisplayName, tmp.LastBuild.Timestamp, nil
}

func listJenkinsJobs(j *Jenkins, r *JenkinsResp) error {
	req, _ := http.NewRequest(http.MethodGet, j.Host+"/api/json", nil)
	q := req.URL.Query()
	q.Add("tree", "jobs[name,url,color,displayName,timestamp]")
	req.URL.RawQuery = q.Encode()
	if j.Token != "" {
		req.Header.Add("Authorization", "Basic "+j.Token)
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		return err
	}

	b, _ := io.ReadAll(resp.Body)
	json.Unmarshal(b, r)

	return nil
}

func getJobConfigXML(j *Jenkins, jobName string) *Project {
	req, _ := http.NewRequest(http.MethodGet, j.Host+"/job/"+jobName+"/config.xml", nil)

	if j.Token != "" {
		req.Header.Add("Authorization", "Basic "+j.Token)
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		// return "", 0, err
		fmt.Println(err)
	}

	b, _ := io.ReadAll(resp.Body)

	xmlData := strings.Replace(string(b), "version='1.1'", "version='1.0'", 1)
	xmlData = strings.Replace(xmlData, "version=\"1.1\"", "version=\"1.0\"", 1)

	var tmp Project
	err = xml.Unmarshal([]byte(xmlData), &tmp)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &tmp
}
