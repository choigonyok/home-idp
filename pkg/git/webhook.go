package git

import (
	"context"
	"fmt"

	"github.com/google/go-github/v63/github"
)

func CreateWebhook(url string, gc *GitClient) error {
	h := &github.Hook{
		Active: github.Bool(true),
		Config: &github.HookConfig{
			ContentType: github.String("json"),
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
