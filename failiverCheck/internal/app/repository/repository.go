package repository

import (
	"fmt"
	"strings"
)

type Repository struct {
}

type Component struct {
	ID    int
	Title string
}

var components = []Component{
	{
		ID:    1,
		Title: "Одиночный сервер базы данных ",
	},
	{
		ID:    8,
		Title: "Одиночный сервер базы данных ",
	},
	{
		ID:    2,
		Title: "Далансировщик нагрузки",
	},
	{
		ID:    3,
		Title: "Балансировщик нагрузки",
	},
	{
		ID:    4,
		Title: "Балансировщик нагрузки",
	},
	{
		ID:    5,
		Title: "Балансировщик нагрузки",
	},
	{
		ID:    6,
		Title: "Балансировщик нагрузки",
	},
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) GetComponents() ([]Component, error) {
	if len(components) == 0 {
		return nil, fmt.Errorf("massiv is empty")
	}
	return components, nil
}

func (r *Repository) GetComponent(id int) (Component, error) {
	orders, err := r.GetComponents()
	if err != nil {
		return Component{}, err
	}

	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}
	return Component{}, fmt.Errorf("заказ не найден")
}

func (r *Repository) GetComponentsByTitle(title string) ([]Component, error) {
	orders, err := r.GetComponents()
	if err != nil {
		return []Component{}, err
	}

	var result []Component
	for _, order := range orders {
		if strings.Contains(strings.ToLower(order.Title), strings.ToLower(title)) {
			result = append(result, order)
		}
	}

	return result, nil
}
