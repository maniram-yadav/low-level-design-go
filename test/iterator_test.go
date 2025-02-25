package test

import (
	"lld/iterator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	userList := iterator.NewList()
	userList.Add(iterator.User{Name: "user1", Age: 17})
	userList.Add(iterator.User{Name: "user2", Age: 25})
	userList.Add(iterator.User{Name: "user3", Age: 35})
	it := userList.CreateIterator()
	count := 0
	for it.HasNext() {
		user := it.GetNext()
		t.Logf("%s\n", user.Name)
		count++
	}
	assert.Equal(t, count, 3)
}
