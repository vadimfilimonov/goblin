package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	type User struct {
		user   string
		age    int
		active bool
	}

	input := []User{{user: "barney", age: 36, active: true}, {user: "fred", age: 40, active: false}}
	actual, ok := Find(input, func(item User, index int, slice []User) bool {
		return !item.active
	})
	expected := User{user: "fred", age: 40, active: false}

	assert.True(t, ok)
	assert.Equal(t, expected, actual)

	input2 := []int64{1, 2, 3}
	actual2, ok2 := Find(input2, func(item int64, index int, slice []int64) bool {
		return item > 50
	})

	assert.False(t, ok2)
	assert.Equal(t, int64(0), actual2)
}
