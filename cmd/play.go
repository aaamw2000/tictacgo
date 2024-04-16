/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/aaamw2000/xogame"
	"github.com/spf13/cobra"
)

var congrat bool

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:     "play",
	Aliases: []string{"p", "ply", "pl"},
	Short:   "inits a new tictactoe game for you to play",
	Long:    `Main command. Let's you play a tictactoe game. Can take player names as input, or you can leave it empty and it will ask for them before starting the game. It can also take players' prefered character as input, but you can also leave those out and it will read them from teh default config file`,
	Example: "tictacgo play <playerX name> <plyaerO name> [--charx <char> --charo <char>]",

	Run: func(cmd *cobra.Command, args []string) {
		game, err := xogame.NewGame()
		if err != nil {
			log.Fatalf("Fatal: %q", err)
		}
		status, err := game.Play()
		if err != nil {
			log.Fatalf("Fatal: %q", err)
		}
		if congrat {
			xogame.Congrat(game, status)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	rootCmd.PersistentFlags().BoolVarP(&congrat, "congrat", "c", true, "Congratualte winner at the end of each game.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
