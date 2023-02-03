package db_model

type Mission struct {
	ID          int         `gorm:"column:id;primaryKey:unique;not null;"`
	Title       string      `gorm:"column:title;type:varchar(100);"`
	Description string      `gorm:"column:description;type:varchar(100);"`
	Crew        []Astronaut `gorm:"many2many:astronauts_missions;"`
}
