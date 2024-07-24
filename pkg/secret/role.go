package secret

func (sc *SecretClient) createAppRole(role string, policies []string) error {
	data := map[string]interface{}{
		"policies": policies,
	}
	_, err := sc.Client.Logical().Write("auth/approle/role/"+role, data)
	return err
}
