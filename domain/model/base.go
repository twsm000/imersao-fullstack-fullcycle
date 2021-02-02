package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
)

func init() {
	validator.SetFieldsRequiredByDefault(true)
}

// Base ...
type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdateAt  time.Time `json:"updated_at" valid:"-"`
}
