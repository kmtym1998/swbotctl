/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/kmtym1998/swbotctl/switchbot"
	"github.com/kmtym1998/swbotctl/switchbot/enum"
	"github.com/spf13/cobra"
)

func NewListCmd(ec *cfg.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:   "リスト",
		Short: "turn the selected device on",
		Long:  "turn the selected device on",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ec.SwitchBotAPIClient.SendDeviceControlCommands(
				"02-202210162051-61289937",
				switchbot.SendDeviceControlCommandsRequest{
					Command: enum.TurnOn.String(),
				},
			)
		},
	}
}
