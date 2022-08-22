package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type User struct {
		user   string
		age    int
		active bool
	}

	input := []User{{user: "barney", age: 36, active: true}, {user: "fred", age: 40, active: false}}
	actual := Map(input, func(item User, index int, slice []User) string {
		return item.user
	})
	expected := []string{"barney", "fred"}

	assert.Equal(t, expected, actual)
}
