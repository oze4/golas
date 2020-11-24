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

	printData := func(data []Line) {
		for _, line := range data {
			fmt.Println("Mnemonic\t==", line.Mnem, "\nData\t\t==", line.Data, "\nUnits\t\t==", line.Units, "\nDesc\t\t==", line.Description)
			fmt.Println("-----------------------------------------------------------------------------------")
		}
	}

	fmt.Printf("\n\n===========\nVersion Info\n===========\n\n")
	printData(las.VersionInformation.Data)

	fmt.Printf("\n\n===========\nWell Info\n===========\n\n")
	printData(las.WellInformation.Data)

	fmt.Printf("\n\n===========\nCurve Info\n===========\n\n")
	printData(las.CurveInformation.Data)

	fmt.Printf("\n\n===========\nParam Info\n===========\n\n")
	printData(las.ParameterInformation.Data)

	fmt.Printf("\n\n===========\nOther Info\n===========\n\n")
	printData(las.Other.Data)

	for _, customSectn := range las.CustomSections {
		fmt.Printf("\n\n===========\n%s\n===========\n\n", customSectn.Name)
		printData(customSectn.Data)
	}
}
