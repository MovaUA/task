/*
Copyright © 2020 Valeriy Molchanov <valeriy.molchanov.77@gmail.com>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/movaua/task/pkg/model"
	"github.com/movaua/task/pkg/store"
	"github.com/spf13/cobra"
)

var doCmdID uint64

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := store.New(dbFile)
		if err != nil {
			return err
		}
		t, err := s.Read(doCmdID)
		if err != nil {
			return err
		}
		t.State = model.Task_DONE
		return s.Update(t)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	doCmd.Flags().Uint64VarP(&doCmdID, "id", "i", 0, "id of the task to mark as complete")
}
