package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/p2pquake/jmaxml-vxse-parser-go/converter"
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

		fmt.Println(file.Name())

		v := &vxse.Report{}
		err = xml.Unmarshal(data, &v)
		if err != nil {
			fmt.Printf("%#v\n", err)
			continue
		} else {
			fmt.Printf("%#v\n", v)
		}

		e, err := converter.Vxse2Epsp(*v)
		if err != nil {
			fmt.Printf("%#v\n", err)
			continue
		} else {
			fmt.Printf("%#v\n", e)
		}

		j, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			fmt.Printf("%#v\n", err)
			continue
		} else {
			fmt.Printf("%s\n", j)
		}

		fmt.Println()
	}

}
