package util

type Components int

const (
	SecretManager Components = iota
	DeployManager
	RbacManager
	Gateway
)
