package models

import (
	"time"
//	"gorm.io/gorm"
)


type LeaveStatus string

const (
	LeavePending  LeaveStatus = "PENDING"
	LeaveApproved LeaveStatus = "APPROVED"
	LeaveRejected LeaveStatus = "REJECTED"
)

type Leave struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `json:"user_id"`
	User        User        `gorm:"foreignKey:UserID"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Reason      string      `gorm:"type:text;not null" json:"reason"`
	Status      LeaveStatus `gorm:"type:enum('PENDING','APPROVED','REJECTED');default:'PENDING'" json:"status"`
	ApprovedBy  *uint       `json:"approved_by,omitempty"`
	Approver    *User       `gorm:"foreignKey:ApprovedBy" json:"approver,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
}
