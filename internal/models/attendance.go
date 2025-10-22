package models

import (
	"time"
//	"gorm.io/gorm"
)


type AttendanceType string

const (
	AttendanceWFO AttendanceType = "WFO"
	AttendanceWFH AttendanceType = "WFH"
)

type Attendance struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `json:"user_id"`
	User          User           `gorm:"foreignKey:UserID"`
	Type          AttendanceType `gorm:"type:enum('WFO','WFH');not null" json:"type"`
	CheckInTime   *time.Time     `json:"check_in_time"`
	CheckOutTime  *time.Time     `json:"check_out_time"`
	Latitude      *float64       `json:"latitude,omitempty"`
	Longitude     *float64       `json:"longitude,omitempty"`
	LocationName  *string        `json:"location_name,omitempty"`
	Date          time.Time      `gorm:"type:date;not null" json:"date"`
	
}
