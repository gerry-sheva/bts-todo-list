package checklist

import "github.com/gerry-sheva/bts-todo-list/pkg/util"

type CreateChecklistInput struct {
	Title string `json:"title"`
}

func (i *CreateChecklistInput) validate(v *util.Validator) map[string]string {
	v.Check(i.Title != "", "Title", "Must not be empty")

	return v.Errors
}
