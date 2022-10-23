/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/kmtym1998/swbotctl/switchbot"
	"github.com/kmtym1998/swbotctl/switchbot/enum"
	"github.com/spf13/cobra"
)

func NewTurnOffCmd(ec *cfg.ExecutionContext, gc *cfg.GlobalCfg) *cobra.Command {
	return &cobra.Command{
		Use:   "turn-off",
		Short: "turn the selected device off",
		Long:  "turn the selected device off",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := switchbot.NewClient(gc.Token, gc.Secret)

			log.Println("レアコイルを消します")
			return c.SendDeviceControlCommands(
				"02-202210162051-61289937",
				switchbot.SendDeviceControlCommandsRequest{
					Command: enum.TurnOff.String(),
				},
			)
		},
	}
}
