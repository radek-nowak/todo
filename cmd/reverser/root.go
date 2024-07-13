package reverser 

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "test",
	Args: cobra.ExactArgs(1),
	Short: "short test",
	Long: "Long test",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops %s", err)
		os.Exit(1)
	}

}
