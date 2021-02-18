package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
)

func main() {
	dir := "./data"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		v := &vxse.Report{}
		xml.Unmarshal(data, &v)

		fmt.Println(file.Name())
		fmt.Printf("%#v\n", v)
	}

}
