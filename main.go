package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
)

func main() {
	data, err := ioutil.ReadFile("./data/20210218101232_0_VXSE53_270000.xml")
	if err != nil {
		panic(err)
	}

	v := &vxse.Report{}
	xml.Unmarshal(data, &v)
	fmt.Printf("%#v", v)
}
