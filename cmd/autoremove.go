package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2022
	Description: Apx is a wrapper around multiple package managers to install packages and run commands inside a managed container.
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

func autoRemoveUsage(*cobra.Command) error {
	fmt.Print(`Description: 
	Remove all unused packages automatically.
Usage:
  apx autoremove [options]

Options:
  -h, --help            Show this help message and exit

Usage:
  apx autoremove
`)
	return nil
}

func NewAutoRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: "apx autoremove",
		Use:     "autoremove",
		Short:   "Remove all unused packages automatically",
		RunE:    autoRemove,
	}
	//cmd.SetUsageFunc(autoRemoveUsage)
	return cmd
}

func autoRemove(cmd *cobra.Command, args []string) error {

	command := append([]string{}, container.GetPkgCommand("autoremove")...)
	command = append(command, args...)

	container.Run(command...)

	return nil
}
