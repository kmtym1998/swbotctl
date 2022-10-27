/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"log"

	"github.com/kmtym1998/swbotctl/cfg"
	"github.com/kmtym1998/swbotctl/cmd"
)

var cfgFilePath string
var version string = "0.0.0"

func main() {
	gc := cfg.NewGlobalConfig()
	if err := gc.Prepare(cfgFilePath); err != nil {
		log.Fatal(err)
	}
	ec := cfg.NewExecutionContext()
	if err := ec.Prepare(version, gc.Token, gc.Secret, &gc); err != nil {
		log.Fatal(err)
	}

	rootCmd := cmd.NewRootCmd()

	turnOnCmd := cmd.NewTurnOnCmd(ec)
	turnOffCmd := cmd.NewTurnOffCmd(ec)
	turnOnCmd.Flags().BoolP("multi-select", "m", true, "デバイスを複数選択する")
	turnOffCmd.Flags().BoolP("multi-select", "m", true, "デバイスを複数選択する")

	rootCmd.AddCommand(
		cmd.NewVersionCmd(ec),
		turnOnCmd,
		turnOffCmd,
		cmd.NewListCmd(ec),
		cmd.NewInitCmd(ec, cfgFilePath),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
