package manifest

import (
	"sigs.k8s.io/yaml"
)

func GetArgoCDManifest(username, namespace string) []byte {
	m := map[string]interface{}{
		"apiVersion": "argoproj.io/v1alpha1",
		"kind":       "Application",
		"metadata": map[string]interface{}{
			"name":      "app" + username,
			"namespace": namespace,
		},
		"spec": map[string]interface{}{
			"destination": map[string]interface{}{
				"name":      "",
				"namespace": namespace,
				"server":    "https://kubernetes.default.svc",
			},
			"project": "default",
			"sources": []interface{}{
				map[string]interface{}{
					"repoURL":        "",
					"path":           "",
					"targetRevision": "HEAD",
				},
			},
		},
		"syncPolicy": map[string]interface{}{
			"automated": map[string]interface{}{
				"prune":    true,
				"selfHeal": true,
			},
		},
	}

	b, _ := yaml.Marshal(m)
	return b
}
