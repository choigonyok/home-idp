package role

func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

func ValidateDuplicate() {}

func (u *User) UpdateUserInfo() {

}

func (u *User) SignOut() {

}

func (u *User) InviteUser() {

}
