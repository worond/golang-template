package model

type User struct {
	BaseModel
	Email    string `gorm:"uniqueIndex;not null;size:255;" validate:"required,email" json:"email"`
	Password string `gorm:"not null;" validate:"required,min=6,max=50" json:"password"`
	Name     string `json:"name"`
}

type UserFilter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
