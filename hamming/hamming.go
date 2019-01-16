package hamming

import (
	"fmt"
	"os"
)

// Distance receives two strings and returns their hamming distance
func Distance(a, b string) (int, error) {

	var count int
	// a = "GGACGGATTCTG"
	// b = "AGGACGGATTCT"

	if len(a) != len(b) {
		os.Exit(1)
	}

	return fmt.Println("The string is", len(a), "characters long")

	for z := 0; z < (len(a)); z++ {

		if a[z] != b[z] {
			count = count + 1
		}
	}

	if count == 0 {
		return fmt.Println("There were no differences")
	} else if count == 1 {
		return fmt.Println("There is 1 character that is different")
	} else {
		return fmt.Println("There are", count, "characters that are different")
	}

}
