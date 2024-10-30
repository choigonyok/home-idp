package util

type Components int

const (
	SecretManager Components = iota
	DeployManager
	RbacManager
	TraceManager
	Gateway
	InstallManager
	WebhookManager
)

type Clients int

const (
	GrpcClient Clients = iota
	HelmClient
	StoragePostgresClient
	MailClient
	DockerClient
	KubeClient
	HttpClient
	GitClient
	TraceClient
	GrpcInstallManager
	GrpcDeployManager
	GrpcRbacManager
	Nothing
)

func GetGrpcClient(host string) Clients {
	switch host {
	case "home-idp-rbac-manager":
		return GrpcRbacManager
	case "home-idp-deploy-manager":
		return GrpcDeployManager
	case "home-idp-install-manager":
		return GrpcInstallManager
	default:
		return Nothing
	}
}
