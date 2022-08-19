package gd

import(
	"os"
	"path/filepath"
		
	"github.com/spf13/cobra"
)


type ExportsCmd struct{

}

type exportsExportCmd struct{

}

func (c *ExportsCmd) BuildCommand() *cobra.Command{
	exports := cobra.Command{
		Use:"exports",
		Short:"exports for godot",
		RunE:func(cmd *cobra.Command,args []string)error{
			return c.Run(cmd,args)
		},
	}
	exportCmd := &exportsExportCmd{}
	exports.AddCommand(exportCmd.BuildCommand())

	return &exports
}

func (c *exportsExportCmd) BuildCommand() *cobra.Command{
	return &cobra.Command{
		Use:"export-debug",
		Long:"gd exports export <target> <output>",
		Args:cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command,args []string)error{
			return c.Run(cmd,args)
		},
	}
}

func (c *ExportsCmd) Run(cmd *cobra.Command,args []string)error{
	godot := newGodotCmd()
	return godot.Run(os.Stdout,os.Stderr,"--help")
}

func (c *exportsExportCmd) Run(cmd *cobra.Command,args []string)error{
	godot := newGodotCmd()
	outputPath := args[1]
	err := os.MkdirAll(filepath.Dir(outputPath),0755)
	if err != nil{
		return err
	}
	return godot.Run(os.Stdout,os.Stderr,"--export-debug",args[0],outputPath)
}
