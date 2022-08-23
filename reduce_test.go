package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	type User struct {
		name string
		age  int
	}

	input := []User{{name: "Petr", age: 4}, {name: "Igor", age: 19}, {name: "Vovan", age: 4}, {name: "Matvey", age: 16}}
	actual := Reduce(input, func(accumulator User, value User, index int, slice []User) User {
		if value.age > accumulator.age {
			return value
		}

		return accumulator
	}, input[0])
	expected := User{name: "Igor", age: 19}

	assert.Equal(t, expected, actual)

	assert.Equal(t, 3, Reduce([]int{1, 2}, func(accumulator int, value int, index int, slice []int) int {
		return accumulator + value
	}, 0))
}
