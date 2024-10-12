package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MandavkarPranjal/gitstat/internal/analysis"

	"github.com/spf13/cobra"
)

var (
	repoPath   string
	outputPath string
	feature    string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "gitstat",
	Short: "GitStat provides insights into your Git repository.",
	Run: func(cmd *cobra.Command, args []string) {
		result := analysis.Analyze(repoPath, feature)
		if err := exportData(result, outputPath); err != nil {
			fmt.Printf("Error exporting data: %v\n", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&repoPath, "repo", "", "Path to the Git repository")
	rootCmd.Flags().StringVar(&outputPath, "output", "output.json", "Path to save JSON output")
	rootCmd.Flags().StringVar(&feature, "feature", "all", "Feature to analyze: changes, commits, all")

	rootCmd.MarkFlagRequired("repo")
}

func exportData(result interface{}, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
