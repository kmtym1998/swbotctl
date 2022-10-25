/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/spf13/cobra"
)

func NewListCmd(ec *cfg.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list", "リスト", "デバイス一覧"},
		Short:   "登録されているデバイスのリストを表示する",
		Long:    "登録されているデバイスのリストを表示する",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := ec.SwitchBotAPIClient.ListDevices()
			if err != nil {
				return err
			}

			for _, item := range data.Body.DeviceList {
				fmt.Println(item.DeviceName)
			}

			for _, item := range data.Body.InfraredRemoteList {
				fmt.Println(item.DeviceName)
			}

			return nil
		},
	}
}
