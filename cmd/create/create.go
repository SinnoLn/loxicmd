/*
Copyright © 2022 baekgyun jung <backguyn@netlox.io>

*/
package create

import (
	"fmt"

	"loxicmd/pkg/api"

	"github.com/spf13/cobra"
)

func CreateCmd(restOptions *api.RESTOptions) *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called")
		},
	}

	createCmd.AddCommand(NewCreateLoadBalancerCmd(restOptions))
	return createCmd
}
