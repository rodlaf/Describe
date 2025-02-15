# Changelog

All notable changes to this project will be documented in this file. See https://keepachangelog.com/en/1.1.0/.

## [Unreleased]
### Added
- Add goreleaser to workflow

### Fixed

### Changed

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

