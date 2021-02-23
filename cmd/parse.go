package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/p2pquake/jmaxml-seis-parser-go/jmaseis"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse [FILE]",
	Short: "XML のパース",
	Long:  "気象庁防災情報 XML をパースし、 JSON 形式で出力します。",
	Args:  cobra.MaximumNArgs(1),
	Run:   parse,
}

var prettyPrint bool

func init() {
	parseCmd.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "Pretty print")

	rootCmd.AddCommand(parseCmd)
}

func parse(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("STDIN is not supported.")
		cmd.Help()
		os.Exit(1)
	}

	filename := args[0]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%s read error: %#v", filename, err)
	}

	report := &jmaseis.Report{}
	err = xml.Unmarshal(data, &report)
	if err != nil {
		log.Fatalf("%s parse error: %#v", filename, err)
	}

	if pretty {
		data, err = json.MarshalIndent(report, "", "  ")
	} else {
		data, err = json.Marshal(report)
	}

	if err != nil {
		log.Printf("%#v", report)
		log.Fatalf("%s JSON conversion error: %#v", filename, err)
	}

	fmt.Println(string(data))
}
