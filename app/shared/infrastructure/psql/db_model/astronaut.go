package db_model

import (
	"github.com/google/uuid"
)

type Astronaut struct {
	ID      uuid.UUID `gorm:"column:id;primary_key:unique;not null;type:uuid;default:uuid_generate_v4()"`
	Name    string    `gorm:"column:name;type:varchar(100)"`
	IsPilot bool      `gorm:"column:is_pilot;type:boolean;not null;default:false"`
}
