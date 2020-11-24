package golas

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNewLexer(t *testing.T) {
	r, e := os.Open("unwrapped.las")
	if e != nil {
		panic("Unable to open file")
	}

	las := Parse(r)

	v := reflect.ValueOf(las)
	tp := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Println("\n", tp.Field(i).Name)
		for _, i := range las.VersionInformation.Data {
			fmt.Println("MNEM :", i.Mnem, "\nDATA :", i.Data, "\nUNITS :", i.Units, "\nDESC :", i.Description)
		}
	}
}
