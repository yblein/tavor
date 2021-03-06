// +build example-main

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f := os.Getenv("TAVOR_DD_FILE")

	if f == "" {
		panic("No TAVOR_DD_FILE defined")
	}

	v, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	s := string(v)

	for _, c := range s {
		fmt.Printf("Got %c\n", c)
	}

	os.Exit(0)
}
