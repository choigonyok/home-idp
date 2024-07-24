package secret

import (
	"context"
	"fmt"
)

func (sc *SecretClient) GetSecretV2(token, rootPath, path, key string) (string, error) {

	secret, err := sc.Client.KVv2(rootPath).Get(context.Background(), path)
	if err != nil {
		return "", fmt.Errorf("unable to read secret: %w", err)
	}

	value, ok := secret.Data[key].(string)
	if !ok {
		return "", fmt.Errorf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	return value, nil
}

func (sc *SecretClient) PutSecretV2(roleID, rootPath, path string, kv map[string]any) error {
	_, err := sc.Client.KVv2(rootPath).Put(context.Background(), path, kv)
	if err != nil {
		return fmt.Errorf("unable to write secret: %w", err)
	}
	return nil
}
