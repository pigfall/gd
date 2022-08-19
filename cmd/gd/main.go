package main

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/pigfall/gd"
)

func main() {
	rootCmd := cobra.Command{
		Use:"gd",
	}

	// {
	exports := gd.ExportsCmd{}
	rootCmd.AddCommand(exports.BuildCommand())
	// }
	
	err := rootCmd.Execute()
	if err != nil{
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
}
