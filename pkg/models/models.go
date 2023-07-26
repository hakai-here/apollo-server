package models

import "github.com/google/uuid"

type Users struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"-" form:"-"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email" form:"email" `
	Password  string    `json:"password" gorm:"not null" validate:"required,min=6,max=15" form:"password"`
	FirstName string    `json:"firstname" gorm:"not null" validate:"required" form:"firstname"`
	LastName  string    `json:"lastname" gorm:"not null" validate:"required" form:"lastname" `
	CreatedAt int       `json:"created_at" gorm:"not null"`
	UpdatedAt int       `json:"updated_at" gorm:"not null"`
}
