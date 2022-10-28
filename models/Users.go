package models

type Users struct {
	UserID   int    `gorm:"column:user_id;primaryKey"`
	UserName string `gorm:"user_name"`
	RoleName string
}

func (this *Users) TableName() string {
	return "users"
}
