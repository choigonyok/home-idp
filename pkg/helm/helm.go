package helm

import (
	"fmt"
	"log"
	"os"
	"strings"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

type HelmClient struct {
	Downloader *downloader.ChartDownloader
	Setting    *cli.EnvSettings
	Repository map[string]*repo.ChartRepository
}

// func NewActionSetting

func New() *HelmClient {
	settings := cli.New()

	downloader := &downloader.ChartDownloader{
		Out:              os.Stdout,
		RepositoryConfig: settings.RepositoryConfig,
		RepositoryCache:  settings.RepositoryCache,
		Getters:          getter.All(settings),
	}

	return &HelmClient{
		Downloader: downloader,
		Setting:    settings,
		Repository: make(map[string]*repo.ChartRepository),
	}
}

func (c *HelmClient) AddRepository(repoName, repoUrl string, public bool) error {
	repoEntry := &repo.Entry{
		Name:               repoName,
		URL:                repoUrl,
		PassCredentialsAll: public,
	}

	chartRepo, err := repo.NewChartRepository(repoEntry, getter.All(c.Setting))

	if err != nil {
		log.Fatalf("Failed to create chart repository: %s", err)
		return err
	}

	_, err = chartRepo.DownloadIndexFile()
	if err != nil {
		log.Fatalf("Failed to download index file: %s", err)
		return err
	}

	c.Repository[repoEntry.Name] = chartRepo

	return nil
}

func (c *HelmClient) Find(releaseName, namespace string) *chart.Chart {
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(c.Setting.RESTClientGetter(), c.Setting.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Failed to initialize Helm action configuration: %s", err)
	}

	list := action.NewList(actionConfig)
	releases, _ := list.Run()

	for _, r := range releases {
		if r.Name == releaseName && r.Namespace == namespace {
			return r.Chart
		}
	}

	return nil
}

func (c *HelmClient) Upgrade(repoChartVersion, namespace, releaseName string, values map[string]interface{}) error {
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(c.Setting.RESTClientGetter(), c.Setting.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Failed to initialize Helm action configuration: %s", err)
	}
	upgrade := action.NewUpgrade(actionConfig)
	release, err := upgrade.Run("", c.Find(releaseName, namespace), values)
	if err != nil {
		return err
	}

	fmt.Printf("Chart %s has been upgraded to %s\n", release.Name, release.Namespace)

	return nil
}

// chartPath is like bitnami/mysql:10.2.1
func (c *HelmClient) Install(repoChartVersion, namespace, releaseName string, values map[string]interface{}) error {
	repoName, after, _ := strings.Cut(repoChartVersion, "/")
	chartName, versionName, found := strings.Cut(after, ":")
	if !found {
		versionName = "latest"
	}

	c.Setting.SetNamespace(namespace)

	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(c.Setting.RESTClientGetter(), c.Setting.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Failed to initialize Helm action configuration: %s", err)
	}
	install := action.NewInstall(actionConfig)

	chartURL, err := repo.FindChartInAuthRepoURL(c.Repository[repoName].Config.URL, "", "", chartName, versionName, "", "", "", getter.All(c.Setting))
	if err != nil {
		log.Fatalf("Failed to find chart URL: %s", err)
	}

	chartPath, _, err := c.Downloader.DownloadTo(chartURL, versionName, ".")
	if err != nil {
		log.Fatalf("Failed to download chart: %s", err)
	}

	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Fatalf("Failed to load chart: %s", err)
	}

	release, err := install.Run(chart, values)
	if err != nil {
		log.Fatalf("Failed to install chart: %s", err)
	}

	fmt.Printf("Chart %s has been installed to %s\n", release.Name, release.Namespace)

	return nil
}

func (c *HelmClient) Uninstall(releaseName, namespace string) error {

	actionConfig := new(action.Configuration)
	c.Setting.SetNamespace(namespace)

	if err := actionConfig.Init(c.Setting.RESTClientGetter(), c.Setting.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Failed to initialize Helm action configuration: %s", err)
	}
	uninstall := action.NewUninstall(actionConfig)

	uninstall.Run(releaseName)

	return nil
}
