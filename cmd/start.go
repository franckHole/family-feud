/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/franciscolkdo/family-feud/internal/table"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the game!",
	Long:  `Start the family-feud game.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		_, err := tea.NewProgram(table.New(table.Config{
			Boxes: []table.BoxConfig{{Points: 35, Answer: "Your mom"}, {Points: 30, Answer: "Your dad"}, {Points: 20, Answer: "Your sister"}, {Points: 20, Answer: "Your bro"}, {Points: 5, Answer: "Your step bro"}},
		}), tea.WithMouseCellMotion()).Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
