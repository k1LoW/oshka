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
	"fmt"
	"os"
	"path/filepath"

	"github.com/k1LoW/oshka/analyzer"
	"github.com/k1LoW/oshka/executer"
	tdir "github.com/k1LoW/oshka/target/dir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// filesystemCmd represents the filesystem command
var filesystemCmd = &cobra.Command{
	Use:     "filesystem [DIR]",
	Aliases: []string{"fs"},
	Short:   "(alias: fs) filesystem",
	Long:    `(alias: fs) filesystem.`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dir := args[0]
		e, err := executer.New(commands)
		if err != nil {
			return err
		}
		t, err := tdir.New(dir)
		if err != nil {
			return err
		}
		if err := e.Execute(ctx, t, dir); err != nil {
			return err
		}

		dirs := []string{dir}
		extractRoot := os.TempDir()
		defer func() {
			_ = os.RemoveAll(extractRoot)
		}()

		for i := 0; i < depth; i++ {
			targets, err := analyzer.Analyze(ctx, dirs)
			if err != nil {
				return err
			}
			dirs = []string{}
			for _, t := range targets {
				dest := filepath.Join(extractRoot, t.Dir())
				if _, err := os.Stat(dest); err == nil {
					// already extracted
					continue
				}
				log.Info().Msg(fmt.Sprintf("Extract %s %s to %s", t.Type(), t.Name(), dest))
				if err := t.Extract(ctx, dest); err != nil {
					return err
				}
				if err := e.Execute(ctx, t, dest); err != nil {
					return err
				}

				dirs = append(dirs, dest)
			}
		}

		cmd.Println("")
		cmd.Println("Run results")
		cmd.Println("===========")
		for _, r := range e.Results() {
			cmd.Printf("%s %s %d\n", r.Target.Type(), r.Target.Name(), r.ExitCode)
		}
		return nil
	},
}

func init() {
	runCmd.AddCommand(filesystemCmd)
	filesystemCmd.Flags().StringSliceVarP(&commands, "command", "c", []string{"trivy fs --exit-code 1 ."}, "Command to execute")
	filesystemCmd.Flags().IntVarP(&depth, "depth", "", 2, "Depth of extracting the supply chain")
}
