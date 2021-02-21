package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/p2pquake/jmaxml-vxse-parser-go/converter"
	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [FILE]",
	Short: "XML から EPSP JSON 形式への変換",
	Long:  "気象庁防災情報 XML から EPSP JSON (JMAQuake) 形式への変換を行います。\nファイル未指定の場合は標準入力から変換します。",
	Args:  cobra.MaximumNArgs(1),
	Run:   convert,
}

var pretty bool
var force bool
var ignoreWarning bool

func init() {
	convertCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print")
	convertCmd.Flags().BoolVarP(&force, "force", "f", false, "Ignore validation error")
	convertCmd.Flags().BoolVarP(&ignoreWarning, "ignore-warning", "w", false, "Ignore validation warning")

	rootCmd.AddCommand(convertCmd)
}

func convert(cmd *cobra.Command, args []string) {
	// 入力
	if len(args) == 0 {
		log.Fatalln("STDIN is not supported.")
	}

	filename := args[0]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%s read error: %#v", filename, err)
	}

	report := &vxse.Report{}
	err = xml.Unmarshal(data, &report)
	if err != nil {
		log.Fatalf("%s parse error: %#v", filename, err)
	}

	// 変換
	jmaQuake, err := converter.Vxse2Epsp(*report)
	if err != nil {
		log.Fatalf("%s convert error: %#v", filename, err)
	}

	// 検証
	errors := converter.Validate(filename, jmaQuake)
	for _, err := range errors {
		_, ok := err.(converter.ValidationError)
		if ok && !force {
			log.Fatalf("%s has validation errors: %#v", filename, errors)
		}
	}

	for _, err := range errors {
		_, ok := err.(converter.ValidationWarning)
		if ok && !force && !ignoreWarning {
			log.Fatalf("%s has validation warnings: %#v", filename, errors)
		}
	}

	// 出力
	if pretty {
		data, err = json.MarshalIndent(jmaQuake, "", "  ")
	} else {
		data, err = json.Marshal(jmaQuake)
	}

	if err != nil {
		log.Fatalf("%s JSON conversion error: %#v", filename, err)
	}

	fmt.Println(string(data))
}
