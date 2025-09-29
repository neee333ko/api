package v1

import (
	"github.com/jinzhu/gorm"
	metav1 "github.com/neee333ko/component-base/pkg/meta/v1"
	"github.com/neee333ko/component-base/pkg/util/idutil"
)

type Secret struct {
	metav1.ObjectMeta `json:",inline"`
	Username          string `json:"username" gorm:"column:username" validate:"omitempty"`
	SecretID          string `json:"secretID" gorm:"column:secret-id" validate:"omitempty"`
	SecretKey         string `json:"secretKey" gorm:"column:secret-key" validate:"omitempty"`
	Expires           int64  `json:"expires" gorm:"column:expires" validate:"required"`
	Description       string `json:"description" gorm:"column:description" validate:"description"`
}

type SecretList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Secret `json:"items"`
}

func (s *Secret) TableName() string {
	return "secret"
}

func (s *Secret) AfterCreate(tx *gorm.DB) (err error) {
	s.InstanceID = idutil.GetInstanceID(int64(s.ID), "secret-")
	tx.Model(s).Update("instance_id", s.InstanceID)

	return nil
}
