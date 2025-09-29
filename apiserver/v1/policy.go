package v1

import (
	"github.com/jinzhu/gorm"
	"github.com/neee333ko/component-base/pkg/json"
	metav1 "github.com/neee333ko/component-base/pkg/meta/v1"
	"github.com/neee333ko/component-base/pkg/util/idutil"
	"github.com/ory/ladon"
)

type AuthzPolicy struct {
	ladon.DefaultPolicy
}

type Policy struct {
	metav1.ObjectMeta `json:",inline"`
	Username          string      `json:"username" gorm:"username" validate:"omitempty"`
	Policy            AuthzPolicy `json:"policy" gorm:"-" validate:"omitempty"`
	PolicyShadow      string      `json:"policyShadow" gorm:"column:policy_shadow" validate:"omitempty"`
}

type PolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Policy `json:"items"`
}

func (p *Policy) TableName() string {
	return "policy"
}

func (p *Policy) String() string {
	shadow, _ := json.Marshal(p.Policy)
	return string(shadow)
}

func (p *Policy) AfterCreate(tx *gorm.DB) (err error) {
	p.InstanceID = idutil.GetInstanceID(int64(p.ID), "policy-")
	p.Policy.ID = p.InstanceID
	tx.Model(p).Update("instance_id", p.InstanceID)

	return nil
}

func (p *Policy) BeforeCreate(tx *gorm.DB) (err error) {
	if err = p.ObjectMeta.BeforeCreate(tx); err != nil {
		return err
	}

	p.PolicyShadow = p.String()

	return nil
}

func (p *Policy) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = p.ObjectMeta.BeforeUpdate(tx); err != nil {
		return err
	}

	p.PolicyShadow = p.String()

	return nil
}

func (p *Policy) AfterFind(tx *gorm.DB) (err error) {
	err = p.ObjectMeta.AfterFind(tx)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(p.PolicyShadow), &p.Policy)

	return err
}
