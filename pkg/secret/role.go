package secret

func (sc *SecretClient) ApplyAppRoleWithPolicies(role string, policies []string) (string, error) {
	data := map[string]interface{}{
		"policies": policies,
	}
	if _, err := sc.Client.Logical().Write("auth/approle/role/"+role, data); err != nil {
		return "", err
	}

	return sc.applySecretIDAgainstRole(role)
}

func (sc *SecretClient) applySecretIDAgainstRole(role string) (string, error) {
	secret, err := sc.Client.Logical().Write("auth/approle/role/"+role+"/secret-id", nil)
	if err != nil {
		return "", err
	}
	return secret.Data["secret_id"].(string), nil
}
