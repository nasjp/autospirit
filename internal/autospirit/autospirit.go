package autospirit

import (
	"fmt"

	"github.com/NasSilverBullet/autospirit/internal/teamspirit"
	"github.com/spf13/cobra"
)

func Exec() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "autospirit",
		Short: "cli tool to insert timestamps on teamspirit automatically",
		Long:  ``,
	}
	return cmd
}

func Hello() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Going to work",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := teamspirit.NewWebDrriverRepository()
			if err != nil {
				return err
			}
			p, err := r.Start()
			if err != nil {
				return err
			}
			defer func() {
				err = r.Stop()
			}()
			if err := r.Login(p); err != nil {
				return err
			}
			if err := r.Go(p); err != nil {
				return err
			}
			fmt.Println("Hello!")
			return nil
		},
	}
	return cmd
}

func Bye() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bye",
		Short: "Leave the office",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := teamspirit.NewWebDrriverRepository()
			if err != nil {
				return err
			}
			p, err := r.Start()
			if err != nil {
				return err
			}
			defer func() {
				err = r.Stop()
			}()
			if err := r.Login(p); err != nil {
				return err
			}
			if err := r.Out(p); err != nil {
				return err
			}
			fmt.Println("Bye!")
			return nil
		},
	}
	return cmd
}
