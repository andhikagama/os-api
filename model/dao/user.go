package dao

type (
	User struct {
		ID        uint64 `gorm:"primaryKey"`
		Email     string `gorm:"not null"`
		Password  string `gorm:"not null"`
		FirstName string `gorm:"not null"`
		LastName  string `gorm:"not null"`

		Token *string `gorm:"null"`

		Phone *string `gorm:"null"`

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
