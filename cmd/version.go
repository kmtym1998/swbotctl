/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/spf13/cobra"
)

func NewVersionCmd(ec *cfg.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print cli version",
		Long:  "Print cli version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version called")
		},
	}
}
