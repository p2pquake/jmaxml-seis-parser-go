package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ビルド時に設定
var (
	Version = "develop"
	Commit  = "unknown"
	Date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:     "jmaxml-seis-parser-go",
	Short:   "気象庁防災情報 XML 地震火山情報の一部 (VXSE43, VXSE51, VXSE52, VXSE53, VTSE41) のパーサ",
	Version: fmt.Sprintf("%s (commit %s, built at %s)", Version, Commit, Date),
}

func Execute() error {
	return rootCmd.Execute()
}
