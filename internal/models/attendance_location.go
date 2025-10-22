package models

type AttendanceLocation struct {
	ID           uint     `gorm:"primaryKey" json:"id"`
	Name         string   `gorm:"type:varchar(100);not null" json:"name"`
	Latitude     float64  `gorm:"not null" json:"latitude"`
	Longitude    float64  `gorm:"not null" json:"longitude"`
	RadiusMeters float64  `gorm:"not null" json:"radius_meters"` 
	IsActive     bool     `gorm:"default:true" json:"is_active"`
}