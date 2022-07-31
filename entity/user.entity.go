package entity

// User struct is represents a users table in database
type User struct {
	Id       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;no null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
