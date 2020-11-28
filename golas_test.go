package golas

import (
	"fmt"
	"os"
	"testing"
)

func TestNewLexer(t *testing.T) {
	r, e := os.Open("samples/unwrapped.las")
	if e != nil {
		panic("Unable to open file")
	}

	las := Parse(r)

	if len(las.Sections) > 5 {
		t.Fatalf("expected 5 sections : got %d", len(las.Sections))
	}

	for _, sectn := range las.Sections {
		fmt.Printf("\n\n===========\n%s\n===========\n\n", sectn.Name)
		printData(sectn.Data)
	}
}

func printData(data []Line) {
	for _, line := range data {
		fmt.Println("Mnemonic\t==", line.Mnem, "\nData\t\t==", line.Data, "\nUnits\t\t==", line.Units, "\nDesc\t\t==", line.Description)
		fmt.Println("-----------------------------------------------------------------------------------")
	}
}
