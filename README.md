# Describe

**Describe** is a command-line tool that scans a directory, excludes files based on `.describeignore`, and generates a structured Markdown file (`codebase.md`). It helps document projects efficiently by creating a tree structure and embedding file contents.

---

## **Installation Guide (for Non-Go Users)**

### **Option 1: Download Executable (Recommended)**
If you don't have Go installed, you can download a precompiled binary:

1. Visit the [Releases](https://github.com/yourusername/describe/releases) page.
2. Download the correct binary for your system:
   - macOS (Apple Silicon): `describe-macos-arm64.tar.gz`
   - macOS (Intel): `describe-macos-amd64.tar.gz`
   - Linux (x86_64): `describe-linux-amd64.tar.gz`
   - Linux (ARM64): `describe-linux-arm64.tar.gz`
3. Extract and move it to `/usr/local/bin`:
   ```sh
   tar -xvzf describe-macos-arm64.tar.gz
   mv describe /usr/local/bin/
   chmod +x /usr/local/bin/describe
   ```
4. Run:
   ```sh
   describe --help
   ```

### **Option 2: Install with Go**
If you already have Go 1.23+ installed, you can install `describe` with:
```sh
go install github.com/yourusername/describe@latest
```

This will install `describe` in Go’s **bin directory** (e.g., `$HOME/go/bin/describe`).

If you get a "command not found" error after running `describe`, add Go’s bin directory to your system’s **PATH**.

#### **For macOS/Linux**
Add this to your `~/.bashrc`, `~/.bash_profile`, or `~/.zshrc`:
```sh
export PATH=$HOME/go/bin:$PATH
```
Then run:
```sh
source ~/.zshrc  # or source ~/.bashrc
```

#### **For Windows**
Add `%USERPROFILE%\go\bin` to your **System Environment Variables**:
1. Open **Control Panel** → **System** → **Advanced system settings**.
2. Click **Environment Variables**.
3. Under **System variables**, find **Path**, then click **Edit**.
4. Add `C:\Users\YourUsername\go\bin` (Replace `YourUsername`).
5. Click **OK**, restart your terminal, and run:
   ```sh
   describe --help
   ```

Now `describe` should work globally! 🎉

---

## **Usage**

### **Basic Usage**
```sh
describe <input-directory>
```
- Scans `<input-directory>`.
- Uses `.describeignore` (if present).
- Outputs `codebase.md`.

### **Custom Output File**
```sh
describe <input-directory> -output docs.md
```
- Saves Markdown to `docs.md`.

### **Custom Ignore File**
```sh
describe <input-directory> -ignore .gitignore
```
- Uses `.gitignore` instead of `.describeignore`.

### **Help**
```sh
describe --help
```

---

## **How `.describeignore` Works**
The `.describeignore` file works **just like `.gitignore`**, specifying files or directories to exclude from the output.

### **Example `.describeignore`**
```
# Ignore all `.log` and `.tmp` files
*.log
*.tmp

# Ignore entire directories
node_modules/
build/
```

- **Wildcards (`*`)** match patterns.
- **Folder names** exclude whole directories.
- **Specific file names** can be ignored.

If `.describeignore` is **missing**, one will be created automatically with `.git/` as a default entry.

---

## **License**
This project is licensed under the [MIT License](LICENSE.md).

