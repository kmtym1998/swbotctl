/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/kmtym1998/swbotctl/prompter"
	"github.com/kmtym1998/swbotctl/switchbot"
	"github.com/kmtym1998/swbotctl/switchbot/enum"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func NewTurnOnCmd(ec *cfg.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:   "turn-on",
		Short: "turn the selected device on",
		Long:  "turn the selected device on",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := ec.SwitchBotAPIClient.ListDevices()
			if err != nil {
				return err
			}

			turnOnRequestParam := switchbot.SendDeviceControlCommandsRequest{
				Command: enum.TurnOn.String(),
			}

			if len(args) != 0 {
				inputDeviceName := args[0]

				selectedDevice, found := lo.Find(data.Body.DeviceList, func(device switchbot.Device) bool {
					return device.DeviceName == inputDeviceName
				})
				if found {
					return ec.SwitchBotAPIClient.SendDeviceControlCommands(
						selectedDevice.DeviceID,
						turnOnRequestParam,
					)
				}

				selectedInfraredRemote, found := lo.Find(data.Body.InfraredRemoteList, func(device switchbot.InfraredRemote) bool {
					return device.DeviceName == inputDeviceName
				})
				if found {
					return ec.SwitchBotAPIClient.SendDeviceControlCommands(
						selectedInfraredRemote.DeviceID,
						turnOnRequestParam,
					)
				}
			}

			deviceSelectionList := append(
				lo.Map(data.Body.DeviceList, func(device switchbot.Device, _ int) prompter.PromptSelection {
					return prompter.PromptSelection{
						DisplayName: device.DeviceName,
						Value:       device.DeviceID,
					}
				}),
				lo.Map(data.Body.InfraredRemoteList, func(device switchbot.InfraredRemote, _ int) prompter.PromptSelection {
					return prompter.PromptSelection{
						DisplayName: device.DeviceName,
						Value:       device.DeviceID,
					}
				})...,
			)

			deviceSelectionList = append(deviceSelectionList, prompter.PromptSelection{
				DisplayName: "DONE", Value: "DONE",
			})

			doMultiSelect, err := cmd.Flags().GetBool("multi-select")
			if err != nil {
				return err
			}

			var selectedOpts []*prompter.PromptSelection
			for {
				selected, err := prompter.GetInputFromPrompt(
					deviceSelectionList,
					&prompter.PromptSelectionOpts{
						Label: "Select device you want to turn on",
						Size:  ec.Cfg.DeviceListSize,
					},
				)
				if err != nil {
					return err
				}

				if !doMultiSelect || selected.Value == "DONE" {
					break
				}

				selectedOpts = append(selectedOpts, selected)
			}

			for _, selected := range selectedOpts {
				if err := ec.SwitchBotAPIClient.SendDeviceControlCommands(
					selected.Value,
					turnOnRequestParam,
				); err != nil {
					return err
				}
			}

			return nil
		},
	}
}
