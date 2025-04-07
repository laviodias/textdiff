package main

import (
	"fmt"

	"github.com/laviodias/textdiff/diff"
)

func main() {
	a := "my name is John and I like apples"
	b := "my name is Josh and I love oranges"

	result := diff.CompareTexts(a, b)
	fmt.Println(result)

}
