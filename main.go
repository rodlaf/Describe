package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

var (
	outputFile string
	ignoreFile string
	inputDir   string
)

func init() {
	flag.StringVar(&outputFile, "output", "codebase.md", "Output markdown file path")
	flag.StringVar(&ignoreFile, "ignore", ".describeignore", "Ignore file path (in .gitignore format)")
	flag.Usage = func() {
		fmt.Println("Usage: describe <input-directory> -output <file> -ignore <ignore-file>")
		fmt.Println("\nGenerates a Markdown file documenting the directory structure and file contents.")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
	}
}

func main() {
	// Parse command-line arguments
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Error: Missing input directory.")
		flag.Usage()
		os.Exit(1)
	}
	inputDir = flag.Arg(0)

	// Load ignore rules
	ignoreMatcher := loadIgnoreFile(ignoreFile)

	// Get list of files and directory structure
	files, tree := getFilesAndStructure(inputDir, ignoreMatcher)

	// Generate Markdown
	markdown := generateMarkdown(files, tree)

	// Write to output file
	err := os.WriteFile(outputFile, []byte(markdown), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Markdown file generated:", outputFile)
}

// Reads the ignore file and returns a matcher
func loadIgnoreFile(ignorePath string) *ignore.GitIgnore {
	data, err := os.ReadFile(ignorePath)
	if err != nil {
		fmt.Println("Ignore file not found, proceeding without exclusions.")
		return ignore.CompileIgnoreLines()
	}
	lines := strings.Split(string(data), "\n")
	return ignore.CompileIgnoreLines(lines...)
}

// Returns a list of non-ignored files and a tree structure
func getFilesAndStructure(root string, matcher *ignore.GitIgnore) ([]string, string) {
	var files []string
	treeBuffer := &bytes.Buffer{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(root, path)
		if relPath == "." || matcher.MatchesPath(relPath) {
			return nil
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		indent := strings.Repeat("  ", strings.Count(relPath, string(os.PathSeparator)))
		fmt.Fprintf(treeBuffer, "%s- %s\n", indent, d.Name())

		return nil
	})

	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	return files, treeBuffer.String()
}

// Generates markdown output
func generateMarkdown(files []string, tree string) string {
	var markdown bytes.Buffer

	// Structure
	markdown.WriteString("# Codebase\n\n")
	markdown.WriteString("## Structure\n\n")
	markdown.WriteString("```\n" + tree + "```\n\n")

	// Contents
	markdown.WriteString("## Files\n\n")
	for _, file := range files {
		relPath, _ := filepath.Rel(inputDir, file)
		markdown.WriteString(fmt.Sprintf("### %s\n\n", relPath))
		content, err := os.ReadFile(file)
		if err != nil {
			markdown.WriteString("_(error reading file)_\n\n")
			continue
		}
		markdown.WriteString("```\n" + string(content) + "\n```\n\n")
	}

	return markdown.String()
}
