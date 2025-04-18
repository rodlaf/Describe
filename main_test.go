package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeTempFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}
	return path
}

func TestEnsureIgnoreFile_CreatesFile(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, ".describeignore")

	ensureIgnoreFile(path)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !strings.Contains(string(data), ".git/") {
		t.Errorf("Expected '.git/' in created ignore file; got: %s", string(data))
	}
}

func TestLoadIgnoreFiles_MergesPatterns(t *testing.T) {
	tmp := t.TempDir()
	f1 := writeTempFile(t, tmp, ".describeignore", ".git/\nvendor/")
	f2 := writeTempFile(t, tmp, ".gitignore", "build/\ndist/")

	list := f1 + "," + f2
	m := loadIgnoreFiles(list)

	for path, want := range map[string]bool{
		".git/":     true,
		"vendor/":   true,
		"build/app": true,
		"dist/bin":  true,
		"main.go":   false,
	} {
		got := m.MatchesPath(path)
		if got != want {
			t.Errorf("MatchesPath(%q) = %v, want %v", path, got, want)
		}
	}
}

func TestGetFilesAndStructure_RespectsIgnore(t *testing.T) {
	tmp := t.TempDir()
	config := t.TempDir()

	// Create input files
	_ = writeTempFile(t, tmp, "keep.txt", "hello")
	_ = writeTempFile(t, tmp, "ignore.txt", "secret")

	// Write ignore config outside input dir
	ignorePath := writeTempFile(t, config, ".describeignore", "ignore.txt")

	// Load matcher
	m := loadIgnoreFiles(ignorePath)

	// Run function
	files, tree := getFilesAndStructure(tmp, m)

	// Assert only 'keep.txt' is included
	if len(files) != 1 || !strings.HasSuffix(files[0], "keep.txt") {
		t.Errorf("Expected only keep.txt, got: %v", files)
	}

	if strings.Contains(tree, "ignore.txt") {
		t.Error("Tree should not contain ignored file")
	}
}

func TestGenerateMarkdown_ContainsContent(t *testing.T) {
	tmp := t.TempDir()
	file := writeTempFile(t, tmp, "example.txt", "sample content")
	files := []string{file}
	tree := "- example.txt\n"

	inputDir = tmp // for correct relative paths in generateMarkdown()
	md := generateMarkdown(files, tree)

	if !strings.Contains(md, "sample content") {
		t.Error("Markdown output missing file content")
	}
	if !strings.Contains(md, "example.txt") {
		t.Error("Markdown output missing file name/section")
	}
	if !strings.Contains(md, "## Structure") {
		t.Error("Markdown output missing Structure section")
	}
}

func TestLoadIgnoreFiles_HandlesMissingFilesGracefully(t *testing.T) {
	// Path doesn't exist
	m := loadIgnoreFiles("nofile1,nofile2")

	if m.MatchesPath("anyfile.txt") {
		t.Error("Expected matcher to allow everything when ignore files are missing")
	}
}

func TestDontIgnoreRoot(t *testing.T) {
	tmp := t.TempDir()
	_ = writeTempFile(t, tmp, "README.md", "value")
	_ = writeTempFile(t, tmp, ".describeignore", "README.md")

	m := loadIgnoreFiles(filepath.Join(tmp, ".describeignore"))

	// Mock root
	relPath, _ := filepath.Rel(tmp, tmp)
	if m.MatchesPath(relPath) {
		t.Error("Root (.) should not be ignored")
	}
}
