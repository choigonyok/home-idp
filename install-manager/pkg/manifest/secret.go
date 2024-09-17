package manifest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetHarborCredManifest(pw string) *corev1.Secret {
	auth := base64.StdEncoding.EncodeToString([]byte("admin:" + pw))

	m := map[string]interface{}{
		"auths": map[string]interface{}{
			env.Get("HOME_IDP_HARBOR_HOST") + ":8080": map[string]interface{}{
				"username": "admin",
				"password": pw,
				"auth":     auth,
			},
		},
	}

	b, err := json.Marshal(m)
	fmt.Println("JSON ERR:", err)

	return &corev1.Secret{
		TypeMeta: v1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "harborcred",
			Namespace: env.Get("HOME_IDP_NAMESPACE"),
		},
		Type: corev1.SecretTypeDockerConfigJson,
		StringData: map[string]string{
			corev1.DockerConfigJsonKey: string(b),
		},
	}
}
