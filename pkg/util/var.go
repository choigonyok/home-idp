package util

type Components int

const (
	SecretManager Components = iota
	DeployManager
	RbacManager
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
)
