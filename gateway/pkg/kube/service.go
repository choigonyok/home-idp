package kube

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (c *GatewayKubeClient) GetServices(namespace string) *[]corev1.Service {
	services, err := c.Client.GetServices(namespace)
	if err != nil {
		fmt.Println("TEST GET SERVICES FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	// services := []model.Service{}

	// for _, svc := range *svcs {

	// 	ports := []string{}
	// 	for _, port := range svc.Spec.Ports {
	// 		ports = append(ports, fmt.Sprintf("%d:%s", port.Port, port.TargetPort.String()))
	// 	}

	// 	age := ""
	// 	t := time.Since(svc.GetCreationTimestamp().Time)
	// 	if t.Hours() > 24 {
	// 		age = strconv.Itoa(int(t.Hours()/24)) + "d"
	// 	} else if int(60*t.Hours())+int(t.Minutes())%60 > 120 {
	// 		age = strconv.Itoa(int(t.Hours())) + "h"
	// 	} else if int(60*t.Hours())+int(t.Minutes())%60 < 10 {
	// 		age = strconv.Itoa(int(t.Minutes())) + "m" + strconv.Itoa(int(t.Seconds())%60) + "s"
	// 	} else {
	// 		age = strconv.Itoa(int(60*t.Hours()+t.Minutes())) + "m"
	// 	}

	// 	services = append(services, model.Service{
	// 		Name:     svc.Name,
	// 		Type:     string(svc.Spec.Type),
	// 		Selector: svc.Spec.Selector,
	// 		Age:      age,
	// 		Port:     ports,
	// 		IP:       svc.Spec.ClusterIP,
	// 	})
	// }

	return services
}
