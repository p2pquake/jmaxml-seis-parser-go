package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "jmaxml-seis-parser-go",
	Short: "気象庁防災情報 XML 地震火山情報の一部 (VXSE51, VXSE52, VXSE53, VTSE41) のパーサ",
}

func Execute() error {
	return rootCmd.Execute()
}
