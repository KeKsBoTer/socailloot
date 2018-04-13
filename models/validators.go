package models

import (
	"github.com/astaxie/beego/validation"
)

func (form *SignUpForm) Valid(v *validation.Validation) {
	if form.Password != form.PasswordRe {
		v.AddError("Repassword", "Does not matched password, repassword")
	}
}

func (form *SubmitForm) Valid(v *validation.Validation) {
	if form.Type == PostTypeLink || form.Type == PostTypeText {
		if len(form.Content) < 1 {
			v.AddError("Content", "Content cannot be emtpy")
		}
	} else if form.Type != PostTypeImage {
		v.AddError("Type", "Invalid post type")
	}

}