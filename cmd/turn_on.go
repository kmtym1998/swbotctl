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

func NewTurnOnCmd(ec *cfg.ExecutionContext, gc *cfg.GlobalCfg) *cobra.Command {
	return &cobra.Command{
		Use:   "turn-on",
		Short: "turn the selected device on",
		Long:  "turn the selected device on",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := switchbot.NewClient(gc.Token, gc.Secret)

			log.Println("レアコイルをつけます")
			return c.SendDeviceControlCommands(
				"02-202210162051-61289937",
				switchbot.SendDeviceControlCommandsRequest{
					Command: enum.TurnOn.String(),
				},
			)
		},
	}
}
