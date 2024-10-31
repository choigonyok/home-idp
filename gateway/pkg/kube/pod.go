package kube

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (c *GatewayKubeClient) GetPods(namespace string) *[]corev1.Pod {
	pods, err := c.Client.GetPods(namespace)
	if err != nil {
		fmt.Println("TEST GET PODS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	// pods := []model.Pod{}

	// for _, p := range *ps {
	// 	age := ""
	// 	t := time.Since(p.GetCreationTimestamp().Time)
	// 	if t.Hours() > 24 {
	// 		age = strconv.Itoa(int(t.Hours()/24)) + "d"
	// 	} else if int(60*t.Hours())+int(t.Minutes())%60 > 120 {
	// 		age = strconv.Itoa(int(t.Hours())) + "h"
	// 	} else if int(60*t.Hours())+int(t.Minutes())%60 < 10 {
	// 		age = strconv.Itoa(int(t.Minutes())) + "m" + strconv.Itoa(int(t.Seconds())%60) + "s"
	// 	} else {
	// 		age = strconv.Itoa(int(60*t.Hours()+t.Minutes())) + "m"
	// 	}

	// 	pods = append(pods, model.Pod{
	// 		Name:   p.Name,
	// 		IP:     p.Status.PodIP,
	// 		Age:    age,
	// 		Status: string(p.Status.Phase),
	// 	})
	// }

	return pods
}
