package cmd

import (
	"fmt"
	"strconv"

	"github.com/movaua/task/pkg/store"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "rm removes specifed tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids := make([]uint64, 0, len(args))
		for _, v := range args {
			id, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ids = append(ids, id)
		}

		s, err := store.New(dbFile)
		if err != nil {
			return err
		}

		for _, id := range ids {
			if err := s.Delete(id); err != nil {
				fmt.Printf("could not delete %d: %v", id, err)
				continue
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
