package kube

// import (
// 	"context"

// 	"google.golang.org/grpc/credentials"
// 	"istio.io/istio/pkg/version"
// 	v1 "k8s.io/api/core/v1"
// )

// type CLIClient interface {
// 	Client
// 	// Revision of the Istio control plane.
// 	Revision() string

// 	// EnvoyDo makes a http request to the Envoy in the specified pod.
// 	EnvoyDo(ctx context.Context, podName, podNamespace, method, path string) ([]byte, error)

// 	// EnvoyDoWithPort makes a http request to the Envoy in the specified pod and port.
// 	EnvoyDoWithPort(ctx context.Context, podName, podNamespace, method, path string, port int) ([]byte, error)

// 	// AllDiscoveryDo makes a http request to each Istio discovery instance.
// 	AllDiscoveryDo(ctx context.Context, namespace, path string) (map[string][]byte, error)

// 	// GetIstioVersions gets the version for each Istio control plane component.
// 	GetIstioVersions(ctx context.Context, namespace string) (*version.MeshInfo, error)

// 	// PodsForSelector finds pods matching selector.
// 	PodsForSelector(ctx context.Context, namespace string, labelSelectors ...string) (*v1.PodList, error)

// 	// GetIstioPods retrieves the pod objects for Istio deployments
// 	GetIstioPods(ctx context.Context, namespace string, opts metav1.ListOptions) ([]v1.Pod, error)

// 	// GetProxyPods retrieves all the proxy pod objects: sidecar injected pods and gateway pods.
// 	GetProxyPods(ctx context.Context, limit int64, token string) (*v1.PodList, error)

// 	// PodExecCommands takes a list of commands and the pod data to run the commands in the specified pod.
// 	PodExecCommands(podName, podNamespace, container string, commands []string) (stdout string, stderr string, err error)

// 	// PodExec takes a command and the pod data to run the command in the specified pod.
// 	PodExec(podName, podNamespace, container string, command string) (stdout string, stderr string, err error)

// 	// PodLogs retrieves the logs for the given pod.
// 	PodLogs(ctx context.Context, podName string, podNamespace string, container string, previousLog bool) (string, error)

// 	// NewPortForwarder creates a new PortForwarder configured for the given pod. If localPort=0, a port will be
// 	// dynamically selected. If localAddress is empty, "localhost" is used.
// 	NewPortForwarder(podName string, ns string, localAddress string, localPort int, podPort int) (PortForwarder, error)

// 	// ApplyYAMLFiles applies the resources in the given YAML files.
// 	ApplyYAMLFiles(namespace string, yamlFiles ...string) error

// 	// ApplyYAMLContents applies the resources in the given YAML strings.
// 	ApplyYAMLContents(namespace string, yamls ...string) error

// 	// ApplyYAMLFilesDryRun performs a dry run for applying the resource in the given YAML files
// 	ApplyYAMLFilesDryRun(namespace string, yamlFiles ...string) error

// 	// DeleteYAMLFiles deletes the resources in the given YAML files.
// 	DeleteYAMLFiles(namespace string, yamlFiles ...string) error

// 	// DeleteYAMLFilesDryRun performs a dry run for deleting the resources in the given YAML files.
// 	DeleteYAMLFilesDryRun(namespace string, yamlFiles ...string) error

// 	// CreatePerRPCCredentials creates a gRPC bearer token provider that can create (and renew!) Istio tokens
// 	CreatePerRPCCredentials(ctx context.Context, tokenNamespace, tokenServiceAccount string, audiences []string,
// 		expirationSeconds int64) (credentials.PerRPCCredentials, error)

// 	// UtilFactory returns a kubectl factory
// 	UtilFactory() PartialFactory

// 	// InvalidateDiscovery invalidates the discovery client, useful after manually changing CRD's
// 	InvalidateDiscovery()
// }
