package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type User struct {
		user   string
		age    int
		active bool
	}

	input := []User{{user: "barney", age: 36, active: true}, {user: "fred", age: 40, active: false}}
	actual := Filter(input, func(item User, index int, slice []User) bool {
		return !item.active
	})
	expected := []User{{user: "fred", age: 40, active: false}}

	assert.Equal(t, expected, actual)
}
