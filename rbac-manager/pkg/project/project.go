package project

import (
	"fmt"

	"github.com/choigonyok/home-idp/rbac-manager/pkg/policy"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/role"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/user"
)

type Project struct {
	Name     string
	Users    []*user.User
	Roles    []*role.Role
	Policies []*policy.Policy
	Options  *ProjectOption
}

func New(name string) *Project {
	fmt.Println("TEST1:", user.GetDefaultUser())
	fmt.Println("TEST2:", user.GetDefaultUser())

	u := user.GetDefaultUser()
	opt := getProjectOption()

	p := &Project{
		Name:     name,
		Users:    []*user.User{u},
		Roles:    u.Roles,
		Policies: u.Roles[0].Policies,
		Options:  opt,
	}
	Store()
	return p
}

func Store() {
	// TODO: STORE TO STORAGE
}

// - policy json 파싱
// - 어드민 policy

// YAML:
// - 어드민
// ㄴ 유저네임
// ㄴ 이메일
// ㄴ 패스워드
// - 패스워드 초기화 기능

// 유저:
// - 외래키 -> 프로젝트

// 프로젝트:
// - 깃허브사용여부 bool
// - 패스워드초기화 bool

// 롤:
// - 외래키 -> 프로젝트
// - policy json

// 깃허브:
// - 외래키 -> 프로젝트

// 유저-롤 매핑:
// - 기본키 -> 유저/롤
// - 외래키 -> 유저, 롤

// policy-롤 매핑:
// - 기본키 -> policy/롤
// - 외래키 -> policy, 롤
