/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/spf13/cobra"
)

func NewInitCmd(ec *cfg.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "初期化する。設定ファイルを作る。",
		Long:  "初期化する。設定ファイルを作る。",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
