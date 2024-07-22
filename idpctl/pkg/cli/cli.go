package cli

import (
	ptr "github.com/choigonyok/home-idp/pkg/pointer"
	"github.com/spf13/pflag"
	v1 "k8s.io/api/core/v1"
)

type RootFlags struct {
	namespace        *string
	config           *string
	defaultNamespace string
}

// type Context interface {
// 	// CLIClient returns a client for the default revision
// 	CLIClient() (kube.CLIClient, error)
// 	// CLIClientWithRevision returns a client for the given revision
// 	CLIClientWithRevision(rev string) (kube.CLIClient, error)
// 	// InferPodInfoFromTypedResource returns the pod name and namespace for the given typed resource
// 	InferPodInfoFromTypedResource(name, namespace string) (pod string, ns string, err error)
// 	// InferPodsFromTypedResource returns the pod names and namespace for the given typed resource
// 	InferPodsFromTypedResource(name, namespace string) ([]string, string, error)
// 	// Namespace returns the namespace specified by the user
// 	Namespace() string
// 	// IstioNamespace returns the Istio namespace specified by the user
// 	IstioNamespace() string
// 	// NamespaceOrDefault returns the namespace specified by the user, or the default namespace if none was specified
// 	NamespaceOrDefault(namespace string) string
// }

// type instance struct {
// 	clients map[string]kube.CLIClient
// 	RootFlags
// }

func AddRootFlags(flags *pflag.FlagSet) *RootFlags {
	rootFlag := &RootFlags{
		// namespace: ptr.Of[string](""),
		namespace: ptr.Of[string](""),
		config:    ptr.Of[string](""),
	}
	// flags.StringVar(f.configContext, FlagContext, "", "Kubernetes configuration context")
	flags.StringVarP(rootFlag.namespace, "namespace", "n", v1.NamespaceAll, "Kubernetes namespace")
	flags.StringVarP(rootFlag.config, "config", "f", v1.NamespaceAll, "home-idp config file")

	return rootFlag
}

func NewCLIContext(rootFlags *RootFlags) *RootFlags {
	if rootFlags == nil {
		rootFlags = &RootFlags{
			namespace:        ptr.Of[string](""),
			config:           ptr.Of[string](""),
			defaultNamespace: "",
		}
	}
	return rootFlags
}
