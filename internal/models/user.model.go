package models

type UserModel struct {
	BaseModel
	FirstName *string `gorm:"type:varchar(200)" json:"first_name"`
	LastName  *string `gorm:"type:varchar(200)" json:"last_name"`
	Email     string  `gorm:"type:varchar(100);unique;not null" json:"email"`
	Gender    *string `gorm:"varchar(10)" json:"gender,omitempty"`
	Password  string  `gorm:"varchar(255);not null" json:"-"`
}

func (u UserModel) TableName() string {
	return "users"
}
