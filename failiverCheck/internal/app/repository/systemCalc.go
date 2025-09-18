package repository

import (
	"failiverCheck/internal/app/ds"

	"gorm.io/gorm"
)

func (r *Repository) CreateSystemCalc(userId uint) (ds.SystemCalculation, error) {
	newSystemCalc := ds.SystemCalculation{
		UserID: userId,
	}
	createErr := r.db.Create(&newSystemCalc).Error
	if createErr != nil {
		return ds.SystemCalculation{}, createErr
	}
	return newSystemCalc, nil
}

func (r *Repository) GetSystemCalc(userId uint) (ds.SystemCalculation, error) {
	var exist_calc ds.SystemCalculation
	findErr := r.db.Where("user_id = ? AND status = ?", userId, ds.DRAFT).First(&exist_calc).Error
	if findErr != nil {
		return ds.SystemCalculation{}, findErr
	}
	return exist_calc, nil

}

func (r *Repository) CreateOrGetSystemCalc(userId uint) (ds.SystemCalculation, error) {
	exist_calc, findErr := r.GetSystemCalc(userId)
	if findErr != nil {
		if findErr == gorm.ErrRecordNotFound {
			return r.CreateSystemCalc(userId)
		} else {
			return ds.SystemCalculation{}, findErr
		}
	}

	return exist_calc, nil

}

func (r *Repository) AddComponentInSystemCalc(componentID uint, userId uint) error {
	systemCal, err := r.CreateOrGetSystemCalc(userId)
	if err != nil {
		return err
	}

	var existing ds.ComponentsToSystemCalc
	check := r.db.Where("component_id = ? AND system_calculation_id = ?", componentID, systemCal.ID).First(&existing)
	if check.Error == nil {
		return nil
	}
	if check.Error != nil && check.Error != gorm.ErrRecordNotFound {
		return check.Error
	}

	componentsToSystemCalc := ds.ComponentsToSystemCalc{
		ComponentID:         componentID,
		SystemCalculationID: systemCal.ID,
	}
	err = r.db.Create(&componentsToSystemCalc).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetComponentsInSystemCalc(systemCalcId uint) ([]ds.Component, error) {
	var systemCalc ds.SystemCalculation

	err := r.db.Preload("Components", "is_deleted = ?", false).Where("id = ?", systemCalcId).First(&systemCalc).Error
	if err != nil {
		return nil, err
	}

	return systemCalc.Components, nil
}

func (r *Repository) GetCountComnponents(userId uint) (int, error) {
	var components []ds.Component
	systemCalc, err := r.GetSystemCalc(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	components, err = r.GetComponentsInSystemCalc(systemCalc.ID)
	if err != nil {
		return 0, err
	}
	return len(components), nil

}
