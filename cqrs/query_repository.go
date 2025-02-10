package cqrs

import (
	"errors"
	"fmt"
	"sync"
)

type ProductView struct {
	Id    int
	Name  string
	Price float64
}
type ProductQueryRepository struct {
	productStore map[int]*ProductView
	EventBus     *EventBus
	mu           sync.Mutex
}

func NewProductQueryRepository(eventBus *EventBus) *ProductQueryRepository {
	return &ProductQueryRepository{EventBus: eventBus, productStore: make(map[int]*ProductView)}
}

func (repo *ProductQueryRepository) ApplyEvent(event interface{}) {

	repo.mu.Lock()
	defer repo.mu.Unlock()
	fmt.Print("\nInside apply method ")

	switch e := event.(type) {
	case ProductCreateEvent:
		fmt.Print("\nReceived ProductCreateEvent")
		repo.productStore[e.Id] = &ProductView{Id: e.Id, Name: e.Name, Price: e.Price}
		fmt.Print(repo.productStore)
	case ProductNameUpdateEvent:
		fmt.Print("\nReceived ProductNameUpdateEvent")
		if productView, ok := repo.productStore[e.Id]; ok {
			productView.Name = e.Name
		}
		return
	case ProductDeleteEvent:
		fmt.Print("\nReceived ProductDeleteEvent")
		delete(repo.productStore, e.Id)
	}
}

func (repo *ProductQueryRepository) GetProductById(id int) (*ProductView, error) {

	repo.mu.Lock()
	defer repo.mu.Unlock()
	fmt.Printf("Inside query method for id %d", id)
	fmt.Print(repo.productStore)
	if p, ok := repo.productStore[id]; ok {
		return p, nil
	}
	return nil, errors.New("product not found")
}

func (repo *ProductQueryRepository) GetAllProducts() []*ProductView {

	repo.mu.Lock()
	defer repo.mu.Unlock()
	allProducts := make([]*ProductView, 0)
	for _, p := range repo.productStore {
		allProducts = append(allProducts, p)
	}
	return allProducts
}
