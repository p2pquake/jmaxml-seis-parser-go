package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/p2pquake/jmaxml-seis-parser-go/converter"
	"github.com/p2pquake/jmaxml-seis-parser-go/jmaseis"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [FILE]",
	Short: "XML から EPSP JSON 形式への変換",
	Long:  "気象庁防災情報 XML から EPSP JSON (JMAQuake, JMATsunami, JMAEEW) 形式への変換を行います。",
	Args:  cobra.MaximumNArgs(1),
	Run:   convert,
}

var pretty bool
var force bool
var ignoreWarning bool
var tsunami bool
var eew bool

func init() {
	convertCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print")
	convertCmd.Flags().BoolVarP(&force, "force", "f", false, "Ignore validation error")
	convertCmd.Flags().BoolVarP(&ignoreWarning, "ignore-warning", "w", false, "Ignore validation warning")
	convertCmd.Flags().BoolVarP(&tsunami, "tsunami", "t", false, "Parse tsunami forecasts")
	convertCmd.Flags().BoolVarP(&eew, "eew", "e", false, "Parse EEW (Earthquake Early Warning)")

	rootCmd.AddCommand(convertCmd)
}

func convert(cmd *cobra.Command, args []string) {
	// 入力
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

	if tsunami {
		// 変換
		jmaTsunami, err := converter.Vtse2Epsp(*report)
		if err != nil {
			log.Fatalf("%s convert error: %#v", filename, err)
		}

		// 検証
		errors := converter.ValidateJMATsunami(filename, report, jmaTsunami)
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
			data, err = json.MarshalIndent(jmaTsunami, "", "  ")
		} else {
			data, err = json.Marshal(jmaTsunami)
		}

		if err != nil {
			log.Fatalf("%s JSON conversion error: %#v", filename, err)
		}

		fmt.Println(string(data))
	} else if eew {
		// 変換
		jmaEEW, err := converter.Vxse2EpspEEW(*report)
		if err != nil {
			log.Fatalf("%s convert error: %#v", filename, err)
		}

		// 検証
		// FIXME: 未実装

		// 出力
		if pretty {
			data, err = json.MarshalIndent(jmaEEW, "", "  ")
		} else {
			data, err = json.Marshal(jmaEEW)
		}

		if err != nil {
			log.Fatalf("%s JSON conversion error: %#v", filename, err)
		}

		fmt.Println(string(data))
	} else {
		// 変換
		jmaQuake, err := converter.Vxse2EpspQuake(*report)
		if err != nil {
			log.Fatalf("%s convert error: %#v", filename, err)
		}

		// 検証
		errors := converter.ValidateJMAQuake(filename, report, jmaQuake)
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
}
