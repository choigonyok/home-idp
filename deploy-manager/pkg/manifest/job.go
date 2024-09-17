package manifest

import (
	"github.com/choigonyok/home-idp/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetKanikoJobManifest(img *docker.Image, repo string) *batchv1.Job {
	return &batchv1.Job{
		TypeMeta: v1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "kaniko-" + img.Name + "-" + img.Version,
			Namespace: env.Get("HOME_IDP_NAMESPACE"),
		},
		Spec: batchv1.JobSpec{
			BackoffLimit:            util.PtrInt32(3),
			TTLSecondsAfterFinished: util.PtrInt32(100),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:      "kaniko-" + img.Name + "-" + img.Version,
					Namespace: env.Get("HOME_IDP_NAMESPACE"),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "kaniko",
							Image: "gcr.io/kaniko-project/executor:v1.23.2",
							Args: []string{
								"--insecure",
								"--skip-tls-verify",
								"--dockerfile=/docker/" + img.Pusher + "/Dockerfile." + img.Name + ":" + img.Version,
								"--context=git://github.com/" + env.Get("HOME_IDP_GIT_USERNAME") + "/" + repo + "#main",
								"--destination=" + env.Get("HOME_IDP_HARBOR_HOST") + ":8080/library/" + img.Name + ":" + img.Version,
								"--cache=true",
							},
							// "--destination=harbor." + env.Get("HOME_IDP_NAMESPACE") + ".svc.cluster.local:80/library/" + img.Name + ":" + img.Version,
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "kaniko-secret",
									MountPath: "/kaniko/.docker",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []corev1.Volume{
						{
							Name: "kaniko-secret",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: "harborcred",
									Items: []corev1.KeyToPath{
										{
											Key:  ".dockerconfigjson",
											Path: "config.json",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
