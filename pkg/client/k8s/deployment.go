package client

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8s) CreateDeployment(namespace, name, image string, replicas int32) error {
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": name}},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": name}},
				Spec:       corev1.PodSpec{Containers: []corev1.Container{{Name: name, Image: image}}},
			},
		},
	}
	_, err := k.clientset.AppsV1().Deployments(namespace).Create(context.Background(), dep, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
