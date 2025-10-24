package v1

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neee333ko/component-base/pkg/auth"
	metav1 "github.com/neee333ko/component-base/pkg/meta/v1"
	"github.com/neee333ko/component-base/pkg/util/idutil"
)

type User struct {
	metav1.ObjectMeta `json:",inline"`
	Status            int `json:"status" gorm:"column:status" validate:"omitempty"`
	// required
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required"`
	// required
	Password string `json:"password" gorm:"column:password" validate:"required"`
	// required
	Email       string    `json:"email" gorm:"column:email" validate:"required;email;min=1;max=30"`
	Phone       string    `json:"phone,omitempty" gorm:"column:phone" validate:"omitempty"`
	IsAdmin     int       `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`
	TotalPolicy int64     `json:"totalPolicy" gorm:"column:total_policy" validate:"omitempty"`
	LoginedAt   time.Time `json:"loginedAt" gorm:"column:logined_at" validate:"omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idutil.GetInstanceID(int64(u.ID), "user-")
	tx.Model(u).Update("instance_id", u.InstanceID)

	return nil
}

type UserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}
