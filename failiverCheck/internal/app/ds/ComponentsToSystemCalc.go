package ds

type ComponentsToSystemCalc struct {
	ComponentID         uint              `gorm:"primaryKey; autoIncrement:false"`
	SystemCalculationID uint              `gorm:"primaryKey; autoIncrement:false"`
	ReplicationCount    uint              `gorm:"default:1"`
	Component           Component         `gorm:"foreignKey:ComponentID"`
	SystemCalculation   SystemCalculation `gorm:"foreignKey:SystemCalculationID"`
}
