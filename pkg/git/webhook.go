package git

import (
	"context"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/google/go-github/v63/github"
)

func CreateWebhook(url string, gc *GitClient) error {
	env.Set("IDP_GIT_WEBHOOK_CONTENT_TYPE", "json") // REMOVE LATER
	h := &github.Hook{
		Active: github.Bool(true),
		Config: &github.HookConfig{
			ContentType: github.String(env.Get("IDP_GIT_WEBHOOK_CONTENT_TYPE")),
			URL:         github.String(url),
		},
	}
	_, resp, err := gc.Client.Repositories.CreateHook(context.TODO(), gc.Owner, gc.Repository.GetName(), h)
	if err != nil {
		fmt.Println("FAIL TO CREATE WEBHOOK: ", resp.StatusCode)
		return err
	}

	fmt.Println("SUCCEED TO CREATE WEBHOOK")
	return nil
}
