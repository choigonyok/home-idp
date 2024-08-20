package manifest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretSpec struct {
	Name      string
	Namespace string
	Type      string
	Data      *map[string]string
}

func (s *SecretSpec) Get() string {
	keys := []string{}
	values := []string{}

	for k, v := range *s.Data {
		keys = append(keys, k)
		values = append(values, v)
	}

	data := ""
	for i := 0; i < len(*s.Data); i++ {
		data += fmt.Sprintf("\n  %s: %s", keys[i], values[i])
	}

	return fmt.Sprintf("apiVersion: v1\nkind: Secret\nmetadata:\n  name: %s\n  namespace: %s\ntype: %s\ndata: %s", s.Name, s.Namespace, s.Type, data)
}

func GetHarborCredManifest(pw string) *corev1.Secret {
	auth := base64.RawStdEncoding.EncodeToString([]byte("admin:" + pw))

	m := map[string]interface{}{
		"auths": map[string]interface{}{
			"harbor.idp-system.svc.cluster.local:80": map[string]interface{}{
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
			Name:      "harborcred-tt",
			Namespace: env.Get("GLOBAL_NAMESPACE"),
		},
		Type: "kubernetes.io/dockerconfigjson",
		StringData: map[string]string{
			".dockerconfigjson": string(b),
		},
	}
}
