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
	"encoding/json"
	"os"

	"github.com/movaua/task/pkg/store"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := store.New(dbFile)
		if err != nil {
			return err
		}

		tasks, err := s.List()
		if err != nil {
			return err
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		return enc.Encode(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
