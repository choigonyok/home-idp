package model

type Role struct {
	ID       int
	Name     string
	Policies []*Policy
}

func GetDefaultRole() *Role {
	r := &Role{
		Name:     "admin",
		Policies: []*Policy{GetDefaultPolicy()},
	}

	Store(r)
	return r
}

func Store(r *Role) error {
	// STORE TO STORAGE
	return nil
}
