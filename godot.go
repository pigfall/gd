package gd

import(
	"io"

	"github.com/pigfall/gosdk/process"
)

type godotCmd struct{}

func newGodotCmd() *godotCmd{
	return &godotCmd{}
}

func (c *godotCmd) Run(out io.Writer,errOut io.Writer,args ...string)error{
	cmd,args := c.buildCommand(args...)
	return process.ExecOutput(out,errOut,cmd,args...)
}

func (c *godotCmd) buildCommand(argsInput ...string)(cmd string,argsOutput []string){
	argsOutput  = append([]string{"--no-window"},argsInput...)
	return "godot",argsOutput
}
