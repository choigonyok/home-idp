package manifest

import (
	"sigs.k8s.io/yaml"
)

type ArgoCDManifest struct {
	Name      string
	Namespace string
}

func GetArgoCDManifest(argocd *ArgoCDManifest) []byte {
	m := map[string]interface{}{
		"apiVersion": "argoproj.io/v1alpha1",
		"kind":       "Application",
		"metadata": map[string]interface{}{
			"name":      argocd.Name,
			"namespace": argocd.Namespace,
		},
		"spec": map[string]interface{}{
			"destination": map[string]interface{}{
				"name":      "",
				"namespace": argocd.Namespace,
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
