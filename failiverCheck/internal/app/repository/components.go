package repository

import (
	"failiverCheck/internal/app/ds"
	"fmt"
)

func (r *Repository) GetComponents() ([]ds.Component, error) {
	var components []ds.Component

	err := r.db.Find(&components).Error
	if err != nil {
		return nil, err
	}
	if len(components) == 0 {
		return nil, fmt.Errorf("massive is empty")
	}
	return components, nil
}

func (r *Repository) GetComponentById(id int) (ds.Component, error) {
	var component ds.Component

	err := r.db.Where("id = ?", id).First(&component).Error
	if err != nil {
		return ds.Component{}, err
	}

	return component, fmt.Errorf("заказ не найден")
}

func (r *Repository) GetComponentsByTitle(title string) ([]ds.Component, error) {
	var components []ds.Component

	err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&components).Error

	if err != nil {
		return nil, err
	}

	return components, nil
}
