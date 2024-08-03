package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/registry"
	"helm.sh/helm/v3/pkg/repo"
)

type HelmClient struct {
	RegistryClient *registry.Client
}

func New() error {
	set := &cli.EnvSettings{}
	providers := getter.All(set)

	// helmRepoFilePath := filepath.Join(os.Getenv("HOME"), "Library", "Preferences", "helm", "repositories.yaml")
	// 저장소 파일 로드
	file := repo.NewFile()

	newEntry := &repo.Entry{
		Name: "tesatest",
		URL:  "https://charts.bitnami.com/bitnami",
	}
	file.Update(newEntry)

	e := file.Get("tesatest")
	fmt.Println("e.String():", e.String())

	// 차트 저장소 설정
	chartRepoURL := "https://charts.bitnami.com/bitnami"
	chartRepo, err := repo.NewChartRepository(&repo.Entry{
		URL:  chartRepoURL,
		Name: "tesatest",
	}, providers)
	if err != nil {
		fmt.Println("Failed to create chart repository:", err)
		return err
	}

	path, err := chartRepo.DownloadIndexFile()
	fmt.Println(path)
	fmt.Println(err)

	// repo.LoadFile()
	indexFile, err := repo.LoadIndexFile(path)
	if err != nil {
		fmt.Println("ERR:", err)
	}

	chartVersion := "9.1.4"
	chartName := "mysql"
	chartInfo, _ := indexFile.Get(chartName, chartVersion)
	fmt.Println(chartInfo.Name)
	fmt.Println(chartInfo.Version)
	fmt.Println("TEST1")
	l := downloader.ChartDownloader{}
	fmt.Println("TEST2")
	fmt.Println(chartRepo.Config.Name + "/" + chartInfo.Name)
	str, _, err := l.DownloadTo(chartRepo.Config.Name+"/"+chartInfo.Name, chartVersion, "Tester")
	// chartInfo.Name+"-"+chartInfo.Version+".tgz"
	fmt.Println("TEST3")
	fmt.Println(str)
	fmt.Println("TEST4")
	fmt.Println("TEST5")
	fmt.Println(err)
	fmt.Println("TEST6")

	// // 차트 로드
	chart, err := loader.Load("Tester")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(chart.Values)
	// // 차트 설치
	// install := action.NewInstall(actionConfig)
	// install.ReleaseName = chartName
	// install.Namespace = settings.Namespace()
	// release, err := install.Run(chart, nil) // 두 번째 인자는 values 파일로, 필요에 따라 설정
	// if err != nil {
	// 	log.Fatalf("Failed to install chart: %s", err)
	// }

	// fmt.Printf("Chart %s has been installed to %s\n", release.Name, release.Namespace)

	return nil
}
