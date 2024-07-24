package secret

import (
	"strings"
)

// 새 경로에 Vault Secret이 생성될 때마다 자동으로 Policy 생성
// Role에 해당 경로에 대한 Read/Write Policy를 관리자가 할당하도록 설정
func (sc *SecretClient) AddPolicy(policyName, path string, capabilities []string) error {
	if err := sc.Client.Sys().PutPolicy(policyName, generatePolicyFormat(path, capabilities)); err != nil {
		return err
	}
	return nil
}

func generatePolicyFormat(path string, capabilities []string) string {
	for i, cap := range capabilities {
		capabilities[i] = `"` + cap + `"`
	}
	caps := strings.Join(capabilities, ", ")

	return `path ` + `"` + path + `"` + ` {
	capabilities = [` + caps + `]
}`
}
