package Go_lang_Generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	FirstName T
	LastName  T
}

func (d *Data[_]) SayHallo(name string) string {
	return "Hello " + name
}

func (d *Data[T]) changeFistName(firstName T) T {
	d.FirstName = firstName
	return firstName
}

func TestMethodData(t *testing.T) {
	data := Data[string]{
		FirstName: "Aolia",
		LastName:  "Rachman",
	}

	// testing untuk test changeName
	changeName := data.changeFistName("Aolia")
	assert.Equal(t, "Aolia", changeName)

	// testing untuk sayhalleo
	hello := data.SayHallo("Aolia")
	assert.Equal(t, "Hello Aolia", hello)
}

func TestData(t *testing.T) {
	data := Data[string]{
		FirstName: "Aolia",
		LastName:  "Rachman",
	}

	fmt.Println(data)

	data2 := Data[int]{
		FirstName: 10,
		LastName:  200,
	}
	fmt.Println(data2)
}
