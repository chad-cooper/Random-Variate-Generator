package rvg

import (
	"fmt"
	"io/ioutil"
)

func WriteData(fileName string, array []float64) {

	buff := []byte("")

	for _, X := range array {
		buff = append(buff, fmt.Sprintf("%.16f\n", X)...)
	}

	ioutil.WriteFile(fileName, buff, 0644)

}
