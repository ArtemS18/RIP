package ds

type Component struct {
	ID          uint32 `gorm:"autoIncrement; primaryKey"`
	Title       string `gorm:"size:256"`
	Type        string `gorm:"size:256"`
	MTBF        uint32 `gorm:"type:integer"`
	MTTR        uint32 `gorm:"type:integer"`
	Available   float32
	Img         string `gorm:"size:512; default: null"`
	Description string
	IsDeleted   bool `gorm:"default:false"`
}
