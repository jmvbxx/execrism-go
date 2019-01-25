package hamming

import "errors"

// Distance receives two strings and returns their hamming distance
func Distance(a, b string) (int, error) {

	var count int

	if len(a) != len(b) {
		return 0, errors.New("not equal length")
	}

	for z := 0; z < (len(a)); z++ {
		if a[z] != b[z] {
			count++
		}
	}

	return count, nil

}
