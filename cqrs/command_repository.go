package cqrs

import "sync"

type ProductCommandRepository struct {
	eventStore       []interface{}
	mu               sync.Mutex
	productIdCounter int
	EventBus         *EventBus
}

func NewProductCommandRepository(eventBus *EventBus) *ProductCommandRepository {
	return &ProductCommandRepository{EventBus: eventBus, eventStore: make([]interface{}, 0), productIdCounter: 0}
}

func (repo *ProductCommandRepository) CreateProduct(name string, price float64) (int, error) {

	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.productIdCounter++
	productCreateEvent := &ProductCreateEvent{Name: name, Price: price, Id: repo.productIdCounter}
	repo.eventStore = append(repo.eventStore, productCreateEvent)
	repo.EventBus.Publish(*productCreateEvent)
	return productCreateEvent.Id, nil

}

func (repo *ProductCommandRepository) UpdateProductName(id int, name string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	productNameUpdateEvent := &ProductNameUpdateEvent{Name: name, Id: id}
	repo.eventStore = append(repo.eventStore, productNameUpdateEvent)
	repo.EventBus.Publish(*productNameUpdateEvent)
	return nil
}

func (repo *ProductCommandRepository) DeleteProduct(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	productDeleteEvent := &ProductDeleteEvent{Id: id}
	repo.eventStore = append(repo.eventStore, productDeleteEvent)
	repo.EventBus.Publish(*productDeleteEvent)
	return nil
}
