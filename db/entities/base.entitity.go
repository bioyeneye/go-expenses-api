package entities

import (
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//func (base *Base) BeforeCreate(scope *gorm.Scope) error {
//	base.CreatedAt = time.Now()
//	return nil
//}