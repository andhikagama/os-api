package dao

import (
	"github.com/andhikagama/os-api/shared/utils"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint64  `gorm:"primaryKey"`
		Phone     string  `gorm:"not null"`
		Password  string  `gorm:"not null"`
		Email     *string `gorm:"null"`
		FirstName string  `gorm:"not null"`
		LastName  string  `gorm:"not null"`
		Token     *string `gorm:"null"`

		GormModel
		BaseModelSoftDelete
	}

	Users []User
)

// TableName .
func (User) TableName() string {
	return "users"
}

// ModelName .
func (User) ModelName() string {
	return "User"
}

func (usr *User) BeforeCreate(tx *gorm.DB) error {
	md5Password := utils.NewMD5Hash(usr.Password)
	bfPassword, err := utils.NewEncryptedString(md5Password)
	if err != nil {
		return err
	}

	usr.Password = bfPassword

	return nil
}
