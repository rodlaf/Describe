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
	ignoreList string
	inputDir   string
	debug      bool
)

func init() {
	flag.StringVar(&outputFile, "output", "codebase.md", "Output markdown file path")
	flag.StringVar(&ignoreList, "ignore", ".describeignore,.gitignore", "Comma-separated list of ignore file paths")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")

	flag.Usage = func() {
		fmt.Println("Usage: describe <input-directory> -output <file> -ignore <ignore-files> -debug")
		fmt.Println("\nGenerates a Markdown file documenting the directory structure and file contents.")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Error: Missing input directory.")
		flag.Usage()
		os.Exit(1)
	}
	inputDir = flag.Arg(0)
	outputPath, _ := filepath.Abs(outputFile)

	// Ensure .describeignore exists inside input directory
	defaultIgnorePath := filepath.Join(inputDir, ".describeignore")
	ensureIgnoreFile(defaultIgnorePath)

	// Resolve ignore files relative to inputDir
	var resolvedIgnorePaths []string
	for _, name := range strings.Split(ignoreList, ",") {
		resolved := filepath.Join(inputDir, strings.TrimSpace(name))
		resolvedIgnorePaths = append(resolvedIgnorePaths, resolved)
	}

	// Load ignore rules
	matcher := loadIgnoreFiles(strings.Join(resolvedIgnorePaths, ","))

	// Generate file list and structure
	files, tree := getFilesAndStructure(inputDir, matcher)
	markdown := generateMarkdown(files, tree)

	// Delete old output
	if _, err := os.Stat(outputPath); err == nil {
		if debug {
			fmt.Println("Debug: Removing previous", outputPath)
		}
		if err := os.Remove(outputPath); err != nil {
			fmt.Println("Error removing existing output file:", err)
			os.Exit(1)
		}
	}

	// Write new output
	err := os.WriteFile(outputPath, []byte(markdown), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Markdown file generated:", outputPath)
}

func ensureIgnoreFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if debug {
			fmt.Println("Debug: Creating .describeignore at", path)
		}
		content := ".git/\ndist/\ncodebase.md\n"
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error creating .describeignore:", err)
			os.Exit(1)
		}
	}
}

func loadIgnoreFiles(paths string) *ignore.GitIgnore {
	var patterns []string

	for _, raw := range strings.Split(paths, ",") {
		path := strings.TrimSpace(raw)
		data, err := os.ReadFile(path)
		if err != nil {
			if debug {
				fmt.Println("Debug: Ignore file not found:", path)
			}
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				patterns = append(patterns, line)
			}
		}
	}
	if debug {
		fmt.Println("Debug: Compiled ignore patterns:", patterns)
	}
	return ignore.CompileIgnoreLines(patterns...)
}

func getFilesAndStructure(root string, matcher *ignore.GitIgnore) ([]string, string) {
	var files []string
	treeBuffer := &bytes.Buffer{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(root, path)

		// 🛠 FIX: match both relPath and relPath + "/" to ignore dirs like .git properly
		matched := matcher.MatchesPath(relPath) || matcher.MatchesPath(relPath+"/")
		if relPath != "." && matched {
			if debug {
				fmt.Println("Debug: Ignoring", relPath)
			}
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if relPath != "." {
			indent := strings.Repeat("  ", strings.Count(relPath, string(os.PathSeparator)))
			if d.IsDir() {
				fmt.Fprintf(treeBuffer, "%s- %s/\n", indent, d.Name())
			} else {
				fmt.Fprintf(treeBuffer, "%s- %s\n", indent, d.Name())
				files = append(files, path)
			}
		} else if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	return files, treeBuffer.String()
}

func generateMarkdown(files []string, tree string) string {
	var markdown bytes.Buffer

	markdown.WriteString("# Codebase\n\n")
	markdown.WriteString("## Structure\n\n")
	markdown.WriteString("```\n" + tree + "```\n\n")

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
