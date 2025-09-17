package ds

import "time"

type SystemCalculation struct {
	ID                   uint       `gorm:"autoIncrement; primaryKey"`
	SystemName           string     `gorm:"size:256"`
	AvailableCalculation float32    `gorm:"defaul:null"`
	UserID               uint       `gorm:"not null"`
	Status               enumStatus `gorm:"type:enum_status; not null"`
	DateCreated          time.Time  `gorm:"autoCreateTime; not null"`
	DateFormed           time.Time  `gorm:"default:null"`
	DateAcceped          time.Time  `gorm:"default:null"`
	ModeratorID          uint

	User      User `gorm:"foreignKey:UserID"`
	Moderator User `gorm:"foreignKey:ModeratorID"`
}
