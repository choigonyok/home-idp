package manifest

import (
	"encoding/json"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func GetPodManifest(name, image string, port int) string {
	harborHost := env.Get("HOME_IDP_HARBOR_HOST") + ":" + env.Get(env.Get("HOME_IDP_HARBOR_PORT"))

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: env.Get("HOME_IDP_NAMESPACE"),
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  name,
					Image: harborHost + "/library/" + image,
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: int32(port),
						},
					},
				},
			},
			ImagePullSecrets: []corev1.LocalObjectReference{
				{
					Name: "harborcred",
				},
			},
		},
	}

	// b, _ := yaml.Marshal(pod)

	jsonBytes, _ := json.Marshal(pod)
	yamlBytes, _ := yaml.JSONToYAML(jsonBytes)
	fmt.Println(string(yamlBytes))

	return string(yamlBytes)
}
