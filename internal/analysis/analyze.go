package analysis

import (
	"log"

	"github.com/MandavkarPranjal/gitstat/internal/git"
)

// AnalysisResult holds the results of the analysis
type AnalysisResult struct {
	Changes []git.Change `json:"changes,omitempty"`
	Commits []git.Commit `json:"commits,omitempty"`
}

// Analyze performs analysis based on the specified feature
func Analyze(repoPath, feature string) AnalysisResult {
	var result AnalysisResult

	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	log.Println("Repository opened successfully")

	switch feature {
	case "changes":
		result.Changes = AnalyzeChanges(repo)
	case "commits":
		result.Commits = AnalyzeCommits(repo)
	case "all":
		result.Changes = AnalyzeChanges(repo)
		result.Commits = AnalyzeCommits(repo)
	default:
		log.Fatalf("Unknown feature: %s", feature)
	}

	log.Printf("Analysis result: %+v\n", result)
	return result
}

func AnalyzeChanges(repo *git.Repository) []git.Change {
	changes, err := git.GetFileChanges(repo)
	if err != nil {
		log.Fatalf("Failed to analyze changes: %v", err)
	}
	// log.Printf("File changes: %+v\n", changes)
	log.Printf("Number of changes: %d", len(changes))
	return changes
}

func AnalyzeCommits(repo *git.Repository) []git.Commit {
	commits, err := git.GetCommitStats(repo)
	if err != nil {
		log.Fatalf("Failed to analyze commits: %v", err)
	}
	// log.Printf("Commit stats: %+v\n", commits)
	log.Printf("Number of commits: %d", len(commits))
	return commits
}

// func exportData(result interface{}, outputPath string) error {
// 	file, err := os.Create(outputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
//
// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ")
// 	err = encoder.Encode(result)
// 	if err != nil {
// 		log.Fatalf("Failed to encode JSON: %v", err)
// 		return err
// 	}
//
// 	return nil
// }
