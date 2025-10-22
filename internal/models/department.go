package models

type Department struct {
	ID             uint     `gorm:"primaryKey" json:"id"`
	DepartmentName string   `gorm:"size:100;not null" json:"department_name"`
	DivisionID     uint     `gorm:"not null" json:"division_id"`
	Division       Division `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"division"`
	Users          []User   `gorm:"foreignKey:DepartmentID" json:"users,omitempty"`
}
