package rvg

import (
	"fmt"
	"io/ioutil"
)

func WriteData(fileName string, array []float64) {

	// f, err := os.Create(fileName)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer f.Close()

	buff := []byte("")

	for _, X := range array {
		buff = append(buff, fmt.Sprintf("%.3f,\n", X)...)
	}

	ioutil.WriteFile(fileName, buff, 0644)

}
