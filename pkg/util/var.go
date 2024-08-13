package util

type Components int

const (
	SecretManager Components = iota
	DeployManager
	RbacManager
	Gateway
	InstallManager
)

type Clients int

const (
	GrpcClient Clients = iota
	HelmClient
	StorageClient
	MailClient
	GrpcInstallManagerClient
	GrpcGatewayClient
	GrpcRbacManagerClient
	GrpcDeployManagerClient
	GrpcSecretManagerClient
	DockerClient
	KubeClient
)
