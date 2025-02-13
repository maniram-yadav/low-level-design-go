package test

import (
	"fmt"
	"lld/cqrs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCQRS(t *testing.T) {

	eventBus := cqrs.NewEventBus()
	commandRepo := cqrs.NewProductCommandRepository(eventBus)
	queryRepo := cqrs.NewProductQueryRepository(eventBus)

	// t.Logf("%s", fmt.Sprintf("%T", cqrs.ProductCreateEvent{}))

	eventBus.Subscribe(fmt.Sprintf("%T", cqrs.ProductCreateEvent{}), queryRepo.ApplyEvent)
	eventBus.Subscribe(fmt.Sprintf("%T", cqrs.ProductDeleteEvent{}), queryRepo.ApplyEvent)
	eventBus.Subscribe(fmt.Sprintf("%T", cqrs.ProductNameUpdateEvent{}), queryRepo.ApplyEvent)

	id1, err1 := commandRepo.CreateProduct("Smartphone", 1000)
	id2, err2 := commandRepo.CreateProduct("Watch", 100)
	assert.NotNil(t, id1)
	assert.NotNil(t, id2)
	if err1 != nil || err2 != nil {
		t.Error("error in creating product in cqrs api test")
	}
	time.Sleep(2 * time.Second)
	product1, err3 := queryRepo.GetProductById(id1)
	assert.Nil(t, err3)
	assert.Equal(t, product1.Id, id1)

	products := queryRepo.GetAllProducts()
	assert.Equal(t, len(products), 2)
}
