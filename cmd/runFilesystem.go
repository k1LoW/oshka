/*
Copyright © 2021 Ken'ichiro Oyama <k1lowxb@gmail.com>

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
package cmd

import (
	"context"
	"os"
	"strconv"

	"github.com/k1LoW/oshka/executer"
	"github.com/k1LoW/oshka/runner"
	"github.com/k1LoW/oshka/target"
	"github.com/k1LoW/oshka/target/local"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// filesystemCmd represents the filesystem command
var filesystemCmd = &cobra.Command{
	Use:     "filesystem [DIR]",
	Aliases: []string{"fs"},
	Short:   "(alias: fs) execute commands starting from local filesystem",
	Long:    `(alias: fs) execute commands starting from local filesystem.`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		e, err := executer.New(commands)
		if err != nil {
			return err
		}
		t, err := local.New(args[0])
		if err != nil {
			return err
		}
		targets := []target.Target{t}
		r, err := runner.New(targets, e)
		if err != nil {
			return err
		}
		if err := r.Run(ctx, depth); err != nil {
			return err
		}

		cmd.Println("")
		cmd.Println("Run results")
		cmd.Println("===========")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Type", "Command", "Exit Code"})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for _, r := range e.Results() {
			table.Append([]string{r.Target.Name(), r.Target.Type(), r.Command, strconv.Itoa(r.ExitCode)})
		}
		table.Render()

		os.Exit(e.ExitCode())
		return nil
	},
}

func init() {
	runCmd.AddCommand(filesystemCmd)
	filesystemCmd.Flags().StringSliceVarP(&commands, "command", "c", []string{"trivy fs --exit-code 1 ."}, "Command to execute")
	filesystemCmd.Flags().IntVarP(&depth, "depth", "", 1, "Depth of extracting supply chains")
}
