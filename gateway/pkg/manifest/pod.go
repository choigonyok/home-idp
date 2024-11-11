package manifest

import (
	"encoding/json"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/model"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func GetPodManifest(name, namespace, image string, port int, e []*model.EnvVar, f []*model.File) string {
	harborHost := env.Get("HOME_IDP_HARBOR_HOST") + ":" + env.Get("HOME_IDP_HARBOR_PORT")
	volumes := []corev1.Volume{}
	mnts := []corev1.VolumeMount{}

	m := make(map[string][]*model.File)
	for _, item := range f {
		m[(*item).ConfigMap] = append(m[(*item).ConfigMap], item)
	}

	for cm, files := range m {
		v := corev1.Volume{
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: cm,
					},
					Items: []corev1.KeyToPath{},
				},
			},
			Name: cm,
		}

		kp := []corev1.KeyToPath{}
		for _, file := range files {
			kp = append(kp, corev1.KeyToPath{
				Key:  (*file).Name,
				Path: (*file).MountPath,
			})
		}

		v.VolumeSource.ConfigMap.Items = kp

		volumes = append(volumes, v)

		mnt := corev1.VolumeMount{
			Name:      cm,
			MountPath: "/" + cm,
		}

		mnts = append(mnts, mnt)

	}

	envVars := []corev1.EnvVar{}

	for _, envvar := range e {
		tmp := corev1.EnvVar{
			Name: envvar.Key,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: envvar.Key,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: envvar.Secret,
					},
				},
			},
		}
		envVars = append(envVars, tmp)
	}

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  name,
					Image: harborHost + "/" + namespace + "/" + image,
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: int32(port),
						},
					},
					Env:          envVars,
					VolumeMounts: mnts,
				},
			},
			ImagePullSecrets: []corev1.LocalObjectReference{
				{
					Name: "harborcred",
				},
			},
			Volumes: volumes,
		},
	}

	jsonBytes, _ := json.Marshal(pod)
	yamlBytes, _ := yaml.JSONToYAML(jsonBytes)
	fmt.Println(string(yamlBytes))

	return string(yamlBytes)
}
