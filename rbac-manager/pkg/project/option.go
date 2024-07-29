package project

type ProjectOption struct {
	passwordReset bool
}

func getProjectOption() *ProjectOption {
	return &ProjectOption{
		passwordReset: true,
	}
}
