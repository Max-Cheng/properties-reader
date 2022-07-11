package propertiesReader

import (
	"fmt"
	"os"
	"testing"
)

func TestReadProptice(t *testing.T) {
	file, err := os.Open("./test.prop")
	if err != nil {
		t.FailNow()
	}
	value, err := ReadProptice(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}
