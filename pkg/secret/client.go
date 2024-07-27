package secret

import (
	"context"
	"fmt"
	"log"

	"github.com/choigonyok/home-idp/pkg/env"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

type SecretClient struct {
	Client  *vault.Client
	ctx     context.Context
	version string
}

func NewClient() *SecretClient {
	env.Set("IDP_VAULT_ADDRESS", "vault.slexn.com") // REMOVE LATER
	env.Set("IDP_VAULT_SCHEME", "https")            // REMOVE LATER
	env.Set("IDP_VAULT_PORT", "8200")               // REMOVE LATER
	env.Set("IDP_VAULT_VERSION", "v2")              // REMOVE LATER
	env.Set("IDP_VAULT_ROOT_TOKEN", "")             // REMOVE LATER

	vaultAddr := env.Get("IDP_VAULT_SCHEME") + "://" + env.Get("IDP_VAULT_ADDRESS")

	config := vault.DefaultConfig()
	config.Address = vaultAddr

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	return &SecretClient{
		Client:  client,
		ctx:     context.TODO(),
		version: env.Get("IDP_VAULT_VERSION"),
	}
}

func (sc *SecretClient) LogInWithRoot() {
	sc.Client.SetToken(env.Get("IDP_VAULT_ROOT_TOKEN"))
}

func (sc *SecretClient) LogInWithAppRole(roleID, secretID string) error {
	sid := &auth.SecretID{FromString: secretID}
	appRoleAuth, err := auth.NewAppRoleAuth(
		roleID,
		sid,
	)
	if err != nil {
		return fmt.Errorf("unable to initialize AppRole auth method: %w", err)
	}

	authData, err := sc.Client.Auth().Login(context.Background(), appRoleAuth)
	if err != nil {
		return fmt.Errorf("unable to login to AppRole auth method: %w", err)
	}
	if authData == nil {
		return fmt.Errorf("no auth info was returned after login")
	}
	sc.Client.SetToken(authData.Auth.ClientToken)
	return nil
}

func (sc *SecretClient) LogOut() {
	sc.Client.ClearToken()
}
