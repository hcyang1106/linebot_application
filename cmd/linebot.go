package cmd

import (
	"github.com/hcyang1106/awesomeProject/linebot"
	"github.com/spf13/cobra"
)

var (
	linebotCMD = &cobra.Command{
		Use:           "linebot",
		Short:         "A linebot server",
		Run: func(_ *cobra.Command, _ []string) {
			bot := linebot.NewLineBot()
			bot.Start()
		},
	}
)

