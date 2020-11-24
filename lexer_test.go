package golas

import (
	"fmt"
	"os"
	"testing"

	"github.com/oze4/golas/pkg/file"
)

func TestNewLexer(t *testing.T) {
	r, e := os.Open("unwrapped.las")
	if e != nil {
		panic("Unable to open file")
	}

	las := Parse(r)

	printData := func(data []file.Line) {
		for _, line := range data {
			fmt.Println("Mnemonic\t==", line.Mnem, "\nData\t\t==", line.Data, "\nUnits\t\t==", line.Units, "\nDesc\t\t==", line.Description)
		}
	}

	fmt.Println("\nVersion Info")
	printData(las.VersionInformation.Data)
	fmt.Println("\nWell Info")
	printData(las.WellInformation.Data)
	fmt.Println("\nCurve Info")
	printData(las.CurveInformation.Data)
	fmt.Println("\n Param Info")
	printData(las.ParameterInformation.Data)
	fmt.Println("\nOther Info")
	printData(las.Other.Data)
}
