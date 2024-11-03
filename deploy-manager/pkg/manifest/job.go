package manifest

import (
	"fmt"
	"strings"

	"github.com/choigonyok/home-idp/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetKanikoJobManifest(img *docker.Image, repo string) *batchv1.Job {
	i := strings.LastIndex(repo, "/")
	repoName := repo[i+1:]
	repoNameWithoutDotGit, _, _ := strings.Cut(repoName, ".")
	fmt.Println("REPO:", repo)
	fmt.Println("REPONAME:", repoName)
	fmt.Println("REPONAME:", repoNameWithoutDotGit)

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
					InitContainers: []corev1.Container{
						{
							Name:  "git-sync-dockerfile-repo",
							Image: "k8s.gcr.io/git-sync/git-sync:v3.2.2",
							Env: []corev1.EnvVar{
								{
									Name:  "GIT_SYNC_REPO",
									Value: "https://github.com/" + env.Get("HOME_IDP_GIT_USERNAME") + "/" + env.Get("HOME_IDP_GIT_REPO") + ".git",
								},
								{
									Name:  "GIT_SYNC_BRANCH",
									Value: "main",
								},
								{
									Name:  "GIT_SYNC_ROOT",
									Value: "/tmp/git/dockerfile",
								},
								{
									Name:  "GIT_SYNC_DEPTH",
									Value: "1",
								},
								{
									Name:  "GIT_SYNC_ONE_TIME",
									Value: "true",
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "share-dockerfile",
									MountPath: "/tmp/git/dockerfile",
								},
							},
						},
						{
							Name:  "git-sync-source-repo",
							Image: "k8s.gcr.io/git-sync/git-sync:v3.2.2",
							Env: []corev1.EnvVar{
								{
									Name:  "GIT_SYNC_REPO",
									Value: repo,
								},
								{
									Name:  "GIT_SYNC_BRANCH",
									Value: "main",
								},
								{
									Name:  "GIT_SYNC_ROOT",
									Value: "/tmp/git/source",
								},
								{
									Name:  "GIT_SYNC_DEPTH",
									Value: "1",
								},
								{
									Name:  "GIT_SYNC_ONE_TIME",
									Value: "true",
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "share-source",
									MountPath: "/tmp/git/source",
								},
							},
						},
						{
							Name:  "git-sync-merge-1",
							Image: "busybox:latest",
							Command: []string{
								"/bin/sh",
								"-c",
								"mkdir -p /workspace/" + repoNameWithoutDotGit + " && cp -rL --remove-destination /tmp/git/source/" + repoName + " /workspace/" + repoNameWithoutDotGit + "/source",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "share-source",
									MountPath: "/tmp/git/source",
								},
								{
									Name:      "share-merge",
									MountPath: "/workspace",
								},
							},
						},
						{
							Name:  "git-sync-merge-2",
							Image: "busybox:latest",
							Command: []string{
								"/bin/sh",
								"-c",
								"mv /tmp/git/dockerfile/" + env.Get("HOME_IDP_GIT_REPO") + ".git/docker/" + img.Pusher + "/Dockerfile." + img.Name + ":" + img.Version + " /workspace/" + repoNameWithoutDotGit + "/source/Dockerfile." + img.Name + ":" + img.Version,
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "share-dockerfile",
									MountPath: "/tmp/git/dockerfile",
								},
								{
									Name:      "share-merge",
									MountPath: "/workspace",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "kaniko",
							Image: "gcr.io/kaniko-project/executor:v1.23.2",
							Args: []string{
								"--insecure",
								"--skip-tls-verify",
								"--context=dir:///workspace/" + repoNameWithoutDotGit + "/source",
								"--dockerfile=Dockerfile." + img.Name + ":" + img.Version,
								"--destination=" + env.Get("HOME_IDP_HARBOR_HOST") + ":" + env.Get("HOME_IDP_HARBOR_PORT") + "/library/" + img.Name + ":" + img.Version,
								"--cache=true",
							},

							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "kaniko-secret",
									MountPath: "/kaniko/.docker",
								},
								{
									Name:      "share-merge",
									MountPath: "/workspace",
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
						{
							Name: "share-source",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "share-dockerfile",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "share-merge",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
					},
				},
			},
		},
	}
}
