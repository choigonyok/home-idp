package kube

import (
	"fmt"

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

func (c *GatewayKubeClient) GetConfigmapMountedService(configmap, namespace string) []string {
	pods, err := c.Client.GetPodsWithConfigmap(configmap, namespace)
	if err != nil {
		fmt.Println("TEST GET CONFIGMAP1: ", err)
		return nil
	}

	service := []string{}
	for i, pod := range pods {
		fmt.Println("-------TEST", i, "-------")
		l := pod.GetLabels()
		fmt.Println("LABELS:", l)
		svc, err := c.Client.GetServicesWithLabels(l, namespace)
		if err != nil {
			fmt.Println("TEST GET CONFIGMAP2: ", err)
			return nil
		}
		for _, s := range *svc {
			service = append(service, s.Name)
			fmt.Println("SERVICE NAME:", s.Name)
		}
	}

	return service
}
