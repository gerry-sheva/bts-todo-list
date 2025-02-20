package auth

import "github.com/gerry-sheva/bts-todo-list/pkg/util"

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *AuthInput) validate(v *util.Validator) map[string]string {
	v.Check(i.Username != "", "Username", "Must not be empty")
	v.Check(i.Password != "", "Password", "Must not be empty")

	return v.Errors
}
