package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type UserVote int

const (
	ActionUpVote   = UserVote(1)
	ActionDownVote = UserVote(-1)
)

type Vote struct {
	Id   int       `orm:"pk;auto"`
	User *User     `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date time.Time `orm:"auto_now_add"`

	// Action performed by User (down- or upvote)
	Action UserVote

	// Item that is voted on
	Item string

	// Type of the item which is voted on (post,comment etc.)
	Type string
}

func (v *Vote) Insert() error {
	if _, err := orm.NewOrm().Insert(v); err != nil {
		return err
	}
	return nil
}

func (v *Vote) Valid(va *validation.Validation) {
	va.Required(v.User, "user is required")
	if v.Action != ActionUpVote && v.Action != ActionDownVote {
		va.AddError("Action", "action musst be up or downvote")
	}
	va.Required(v.Item, "no item provided")
	va.Required(v.Type, "no type provided")
}
