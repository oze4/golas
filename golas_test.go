package golas

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestNewLexer(t *testing.T) {
	lasReader, _ := os.Open("samples/unwrapped.las")
	las := Parse(lasReader)
	prettyPrintStructAsJSON(las)
}

func prettyPrintStructAsJSON(v interface{}) {
	if j, e := json.MarshalIndent(v, "", "    "); e != nil {
		fmt.Printf("Error : %s \n", e.Error())
	} else {
		fmt.Printf("%s\n", string(j))
	}
}
