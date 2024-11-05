package kube

import (
	"fmt"
	"strings"

	"github.com/choigonyok/home-idp/pkg/model"
	corev1 "k8s.io/api/core/v1"
)

func (c *GatewayKubeClient) GetConfigmaps(namespace string) *[]corev1.ConfigMap {
	configmaps, err := c.Client.GetConfigmaps(namespace)
	if err != nil {
		fmt.Println("TEST GET CONFIGMAPS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	// configmaps := []model.Configmap{}

	// for _, c := range *cms {
	// 	age := ""
	// 	t := time.Since(c.GetCreationTimestamp().Time)
	// 	if t.Hours() > 24 {
	// 		age = strconv.Itoa(int(t.Hours()/24)) + "d"
	// 	} else if int(60*t.Hours())+int(t.Minutes())%60 > 120 {
	// 		age = strconv.Itoa(int(t.Hours())) + "h"
	// 	} else if int(60*t.Hours())+int(t.Minutes())%60 < 10 {
	// 		age = strconv.Itoa(int(t.Minutes())) + "m" + strconv.Itoa(int(t.Seconds())%60) + "s"
	// 	} else {
	// 		age = strconv.Itoa(int(60*t.Hours()+t.Minutes())) + "m"
	// 	}

	// 	configmaps = append(configmaps, model.Configmap{
	// 		Name: c.Name,
	// 		Data: len(c.Data),
	// 		Age:  age,
	// 	})
	// }

	return configmaps
}

func (c *GatewayKubeClient) GetConfigmap(name, namespace string) *map[string]string {
	cms, err := c.Client.GetConfigmap(name, namespace)

	if err != nil {
		fmt.Println("TEST GET CONFIGMAP "+name+" FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	return &cms.Data
}

func (c *GatewayKubeClient) GetConfigmapFiles(namespace string) []model.ConfigMap {
	cms, _ := c.Client.GetConfigmaps(namespace)
	datas := []model.ConfigMap{}

	for _, cm := range *cms {
		data := model.ConfigMap{}
		data.Name = cm.Name
		data.Creator, _ = strings.CutPrefix(cm.Name, "configmap-")
		data.Namespace = namespace
		files := []model.File{}
		for fileName, content := range cm.Data {
			f := model.File{}
			f.Name = fileName
			f.Content = content
			fileMountedServices := []string{}
			labels := c.Client.GetConfigMapFileMountedPodLabels(namespace, fileName)
			for _, l := range labels {
				svc, _ := c.Client.GetServicesWithLabels(l, namespace)
				for _, s := range *svc {
					fileMountedServices = append(fileMountedServices, s.Name)
				}
				fmt.Println("[MOUNT SERVICES]: ", fileMountedServices)
			}
			f.MountServices = fileMountedServices
			files = append(files, f)
			fmt.Println("[FILE]: ", f)
		}
		data.Files = files
		datas = append(datas, data)
		fmt.Println("[DATA]: ", data)
	}
	return datas
}
