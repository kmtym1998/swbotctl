/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func NewInitCmd(ec *cfg.ExecutionContext, cfgFilePath string) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "初期化する。設定ファイルを作る。",
		Long:  "初期化する。設定ファイルを作る。",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfgFilePath == "" {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					return err
				}

				cfgDirName := ".swbotctl"
				cfgFileName := "config.json"
				cfgFilePath = filepath.Join(homeDir, cfgDirName, cfgFileName)
			}

			tokenInputPrompt := promptui.Prompt{
				Label:       "トークンを入力してください",
				HideEntered: true,
				Mask:        []rune("*")[0],
			}
			tokenInput, err := tokenInputPrompt.Run()
			if err != nil {
				return err
			}

			fmt.Println("✅ トークンの入力完了")

			clientSecretInputPrompt := promptui.Prompt{
				Label:       "クライアントシークレットを入力してください",
				HideEntered: true,
				Mask:        []rune("*")[0],
			}
			clientSecretInput, err := clientSecretInputPrompt.Run()
			if err != nil {
				return err
			}

			fmt.Println("✅ クライアントシークレットの入力完了")

			cfg := cfg.GlobalCfg{
				Token:          tokenInput,
				Secret:         clientSecretInput,
				DeviceListSize: 10,
			}

			b, err := json.Marshal(cfg)
			if err != nil {
				return err
			}

			if err := os.WriteFile(cfgFilePath, b, os.ModePerm); err != nil {
				return err
			}

			fmt.Printf("✅ %s に設定ファイルが作成されました", cfgFilePath)

			return nil
		},
	}
}
