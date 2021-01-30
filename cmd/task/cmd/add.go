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
	"errors"
	"strings"

	"github.com/movaua/task/pkg/model"
	"github.com/movaua/task/pkg/store"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("provide a task to add")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		task := model.Task{
			Name:    strings.Join(args, " "),
			State:   model.Task_TODO,
			Created: timestamppb.Now(),
		}
		s, err := store.New(dbFile)
		if err != nil {
			return err
		}
		return s.Create(&task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
