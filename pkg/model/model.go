package model

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RoleID     string `json:"role_id"`
	CreateTime string `json:"create_time"`
}

type Project struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}

type Policy struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Json string `json:"json"`
}
