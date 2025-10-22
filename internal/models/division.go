package models

type Division struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:100;not null;unique" json:"name"`
	Departments []Department `gorm:"foreignKey:DivisionID" json:"departments,omitempty"`
	Users       []User       `gorm:"foreignKey:DivisionID" json:"users,omitempty"`
}
