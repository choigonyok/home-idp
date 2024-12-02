package kube

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (c *GatewayKubeClient) GetSecrets(namespace string) *[]corev1.Secret {
	secrets, err := c.Client.GetSecrets(namespace)
	if err != nil {
		fmt.Println("TEST GET SECRETS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	return secrets
}
