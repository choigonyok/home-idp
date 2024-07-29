package project

type ProjectOption struct {
	passwordReset bool
	verifyEmail   bool
}

func getProjectOption() *ProjectOption {
	return &ProjectOption{
		passwordReset: true,
		verifyEmail:   true,
	}
}
