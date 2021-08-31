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
	"fmt"
	"os"
	"strconv"

	"github.com/k1LoW/oshka/executer"
	"github.com/k1LoW/oshka/runner"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	commands []string
	depth    int
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "execute commands on filesystem/resouces extracted supply chains",
	Long:  `execute commands on filesystem/resouces extracted supply chains.`,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func showResult(cmd *cobra.Command, r *runner.Runner, e *executer.Executer) {
	cmd.Println("")
	cmd.Println("Run results")
	cmd.Println("===========")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Command", "Exit Code", "Hash"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, r := range e.Results() {
		table.Append([]string{r.Target.Name(), r.Target.Type(), r.Command, strconv.Itoa(r.ExitCode), fmt.Sprintf("%s (%s)", r.Target.Hash(), r.Target.HashType())})
	}
	table.Render()
}
