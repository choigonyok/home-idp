package client

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/choigonyok/idp/pkg/model"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func (k *K8s) Version(context string) (string, error) {
	info, err := k.clients[context].Discovery().ServerVersion()
	if err != nil {
		return "", err
	}

	return info.String(), nil
}

func (k *K8s) SetContext(c model.Cluster) error {
	kubeconfigPath := filepath.Join("", "/Users/choigonyok/.kube/config")
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		fmt.Printf("Error loading kubeconfig file: %v\n", err)
		os.Exit(1)
	}

	contextName := c.Context
	clusterName := c.Name
	userName := c.User
	clusterEndpoint := c.APIServerEndpoint
	certificateAuthorityData := c.CertificateAuthorityData
	clientCertificateData := c.ClientCertificateData
	clientKeyData := c.ClientKeyData

	caData := Base64LikeEcho(certificateAuthorityData)
	ccData := Base64LikeEcho(clientCertificateData)
	ckData := Base64LikeEcho(clientKeyData)

	config.Clusters[clusterName] = &api.Cluster{
		Server:                   clusterEndpoint,
		CertificateAuthorityData: []byte(caData),
	}

	// 사용자 정보 추가
	config.AuthInfos[userName] = &api.AuthInfo{
		ClientCertificateData: []byte(ccData),
		ClientKeyData:         []byte(ckData),
	}

	// 컨텍스트 정보 추가
	config.Contexts[contextName] = &api.Context{
		Cluster:  clusterName,
		AuthInfo: userName,
	}

	// 5. 현재 컨텍스트를 새로운 컨텍스트로 설정 (선택 사항)
	config.CurrentContext = contextName

	// 6. 변경된 Config 객체를 kubeconfig 파일에 다시 저장
	err = clientcmd.WriteToFile(*config, kubeconfigPath)
	if err != nil {
		fmt.Printf("Error writing to kubeconfig file: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func Base64LikeEcho(input string) string {
	inputWithNewline := input + "\n"
	return inputWithNewline
}
