# Describe

**Describe** is a command-line tool that scans a directory, excludes files based on `.describeignore`, and generates a structured Markdown file (`codebase.md`). It helps document projects efficiently by creating a tree structure and embedding file contents.

---

## **Installation Guide (for Non-Go Users)**

### **1. Install Go (If Not Already Installed)**
If you don’t have Go installed, download and install it from:  
➡️ [Go Official Website](https://go.dev/dl/)

After installing, verify your Go installation by running:
```sh
go version
```
You should see something like:
```sh
go version go1.23.0 darwin/amd64
```

---

### **2. Install Describe Using `go install`**
Run the following command to install `describe`:
```sh
go install github.com/yourusername/describe@latest
```
> Replace `yourusername` with the actual GitHub username or repository location.

This will install `describe` in Go’s **bin directory** (e.g., `$HOME/go/bin/describe`).

---

### **3. Add Go Bin Directory to PATH (If Needed)**  
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
Add `%USERPROFILE%\goin` to your **System Environment Variables**:
1. Open **Control Panel** → **System** → **Advanced system settings**.
2. Click **Environment Variables**.
3. Under **System variables**, find **Path**, then click **Edit**.
4. Add `C:\Users\YourUsername\go\bin` (Replace `YourUsername`).
5. Click **OK**, restart your terminal, and run:
   ```sh
   describe -help
   ```

Now `describe` should work globally! 🎉

---

## **Usage**
After installing, use `describe` like this:

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
describe -help
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

If `.describeignore` is **missing**, all files will be included.
