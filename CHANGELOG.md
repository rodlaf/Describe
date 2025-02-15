# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]
### Added
- Introduced a `-debug` flag to print debug statements for troubleshooting.
- Automatically creates `.describeignore` if it does not exist, initializing it with `.git/` as a default entry.
- Debug mode now prints which files are ignored.

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

