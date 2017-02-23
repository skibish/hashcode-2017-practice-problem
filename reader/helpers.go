package reader

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func stringSliceToInt(data []string) (newData []int, err error) {
	for _, v := range data {
		// convert
		var vInt int
		vInt, err = strconv.Atoi(v)
		if err != nil {
			return
		}

		// assign
		newData = append(newData, vInt)
	}

	return
}

func printMarshaled(title string, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("===>> %s: %s\n", title, string(b))
}
