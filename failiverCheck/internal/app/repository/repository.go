package repository

import (
	"fmt"
	"strings"
)

type Repository struct {
}

type Component struct {
	ID          int
	Title       string
	Type        string
	MTBF        int
	MTTR        int
	Available   float32
	Replication int
	Img         string
	Description string
}

// type AvailabilityCalculation struct {
// 	ID   int
// 	Name string
// }

type CalculationToComponent struct {
	Component        Component
	ReplicationCount int
}

type Calculation struct {
	Name       string
	Components []CalculationToComponent
}

var components = []Component{
	{
		ID:          1,
		Title:       "Одиночный сервер базы данных",
		Type:        "Одиночный сервер",
		MTBF:        8760,
		MTTR:        2,
		Available:   0.999771,
		Img:         "http://localhost:9000/failivercheck/svg/database.svg",
		Description: "Один экземпляр сервера, содержащий базу данных.  Подвержен простоям при отказах и обслуживании.",
	},
	{
		ID:          2,
		Title:       "Балансировщик нагрузки (активный/пассивный)",
		Type:        "Балансировщик",
		MTBF:        43800,
		MTTR:        4,
		Available:   0.999909,
		Img:         "http://localhost:9000/failivercheck/svg/loader.svg",
		Description: "Два балансировщика нагрузки, один активный, другой пассивный. При отказе активного, пассивный автоматически занимает его место.",
	},
	{
		ID:          3,
		Title:       "Балансировщик нагрузки (геораспределенный)",
		Type:        "Балансировщик",
		MTBF:        52560,
		MTTR:        3,
		Available:   0.999943,
		Img:         "http://localhost:9000/failivercheck/svg/loader.svg",
		Description: "Три балансировщика нагрузки, размещенных в разных географических регионах. Распределяют нагрузку и обеспечивают отказоустойчивость даже при выходе из строя целого региона.",
	},
	{
		ID:          4,
		Title:       "Балансировщик нагрузки (VRRP)",
		Type:        "Балансировщик",
		MTBF:        43800,
		MTTR:        4,
		Available:   0.999909,
		Img:         "http://localhost:9000/failivercheck/svg/loader.svg",
		Description: "Два балансировщика нагрузки, использующие VRRP (Virtual Router Redundancy Protocol) для обеспечения отказоустойчивости. Один активен, другой - в режиме ожидания.",
	},
	{
		ID:          7,
		Title:       "Веб-сервер (кластер)",
		Type:        "Сервер",
		MTBF:        8760,
		MTTR:        2,
		Available:   0.999771,
		Img:         "http://localhost:9000/failivercheck/svg/cluster.svg",
		Description: "Несколько веб-серверов, работающих вместе для обработки запросов. Если один сервер выходит из строя, другие продолжают работать.",
	},
	{
		ID:          9,
		Title:       "Кэширующий сервер (распределенный)",
		Type:        "Сервер",
		MTBF:        26280, // Увеличил MTBF
		MTTR:        3,
		Available:   0.999885,
		Img:         "http://localhost:9000/failivercheck/svg/database.svg",
		Description: "Три кэширующих сервера, объединенных в кластер. Данные распределены между серверами, обеспечивая отказоустойчивость и масштабируемость.",
	},
	{
		ID:          10,
		Title:       "Очередь сообщений (кластер)",
		Type:        "Сервис",
		MTBF:        35040,
		MTTR:        4,
		Available:   0.999886,
		Img:         "http://localhost:9000/failivercheck/svg/database.svg",
		Description: "Четыре сервера очереди сообщений, работающих в кластере. Обеспечивают высокую доступность и надежность доставки сообщений.",
	},
}

var AvailabilityCalculation = map[int]Calculation{
	1: {
		Name: "System 1",
		Components: []CalculationToComponent{
			{Component: components[0], ReplicationCount: 1},
			{Component: components[1], ReplicationCount: 2},
			{Component: components[2], ReplicationCount: 4},
		},
	},
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) GetComponentsInApplication(id int) (Calculation, error) {
	return AvailabilityCalculation[id], nil
}

// func (r *Repository) PushUserList(id int, c Component) ([]Component, error) {
// 	if id >= len(UserList) {
// 		r.PostUserList(c)
// 		return UserList[len(UserList)], nil
// 	}
// 	UserList[id] = append(UserList[id], c)
// 	return UserList[id], nil
// }

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
