# Changelog

All notable changes to this project will be documented in this file. See https://keepachangelog.com/en/1.1.0/.

## [UNRELEASED] 
### Added 

### Fixed 

### Changed


## [1.0.4] 2025-4-18 
### Added 
- Support for .describeignore and .gitignore as comma-separated ignore sources.
     
### Fixed 
- Ignored directories (.git/, dist/) now fully excluded from structure and file list.
- Prevented inclusion of .describeignore when it's in the scanned directory.
- Fixed issue where structure section still listed ignored folders.

### Changed 
- Output file (codebase.md) is now explicitly deleted before writing.
- Ignore file paths are resolved relative to the inputDir.

## [1.0.3] 2025-2-24
### Changed
- Document binary installation in README
- Make goreleaser create homebrew tap

## [1.0.2] - 2025-2-24
### Added
- Add goreleaser to workflow

## [1.0.1] - 2025-2-24
### Added
- Introduced a `-debug` flag to print debug statements for troubleshooting.
- Automatically creates `.describeignore` if it does not exist, initializing it with `.git/` as a default entry.
- Debug mode now prints which files are ignored.
- CONTRIBUTING.md added.
- MIT License added.

### Fixed
- Resolved issue where `.DS_Store` was being included despite being ignored.
- Improved handling of missing `.describeignore` by ensuring it is created before processing.

### Changed
- Enhanced error messages for better debugging.
- Improved directory traversal efficiency in `getFilesAndStructure` function.

## [1.0.0] - 2025-02-15
### Added
- Initial release of `describe`.
- Supports scanning directories and generating structured Markdown documentation.
- Implements `.describeignore` for excluding files similar to `.gitignore`.
- Allows custom output file names via `-output` flag.
- Provides `-ignore` flag to specify an alternative ignore file.

