package models

type User struct {
	ID        uint   `gorm:"column:id" json:"id"`
	Username  string `gorm:"cloumn:user_name" json:"user_name"`
	Name      string `gorm:"column:name" json:"name"`
	Email     string `gorm:"column:email" json:"email"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	Password  string `gorm:"column:password" json:"password"`
}
