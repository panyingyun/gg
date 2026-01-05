package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 各种语言的.gitignore模板
var templates = map[string]string{
	"go": `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# Build output
dist/
build/
bin/
`,
	"react": `# Dependencies
node_modules/
/.pnp
.pnp.js

# Testing
/coverage

# Production
/build
/dist

# Misc
.DS_Store
.env.local
.env.development.local
.env.test.local
.env.production.local

npm-debug.log*
yarn-debug.log*
yarn-error.log*

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`,
	"c++": `# Compiled Object files
*.slo
*.lo
*.o
*.obj

# Precompiled Headers
*.gch
*.pch

# Compiled Dynamic libraries
*.so
*.dylib
*.dll

# Fortran module files
*.mod
*.smod

# Compiled Static libraries
*.lai
*.la
*.a
*.lib

# Executables
*.exe
*.out
*.app

# Build directories
build/
dist/
cmake-build-*/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`,
	"c": `# Compiled Object files
*.o
*.obj

# Compiled Dynamic libraries
*.so
*.dylib
*.dll

# Compiled Static libraries
*.a
*.lib

# Executables
*.exe
*.out
*.app

# Build directories
build/
dist/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`,
	"matlab": `# MATLAB
*.asv
*.m~
*.autosave

# MATLAB Compiler
*.ctf

# MATLAB figures
*.fig

# MATLAB data files
*.mat

# MATLAB code generation
codegen/
slprj/

# Simulink
*.slxc
*.slx.autosave

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`,
}

// 查找.gitignore文件，从当前目录向上查找
func findGitignoreFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		gitignorePath := filepath.Join(dir, ".gitignore")
		if _, err := os.Stat(gitignorePath); err == nil {
			return gitignorePath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// 已经到达根目录
			break
		}
		dir = parent
	}

	// 如果没找到，返回当前目录的路径
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentDir, ".gitignore"), nil
}

// 读取.gitignore文件内容
func readGitignore(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", nil
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 写入.gitignore文件
func writeGitignore(path, content string) error {
	return os.WriteFile(path, []byte(content), 0o644)
}

// 检查内容是否已存在于.gitignore中
func containsIgnore(content, pattern string) bool {
	lines := strings.Split(content, "\n")
	pattern = strings.TrimSpace(pattern)
	for _, line := range lines {
		if strings.TrimSpace(line) == pattern {
			return true
		}
	}
	return false
}

// 生成模板的.gitignore
func generateTemplate(lang string) error {
	template, exists := templates[strings.ToLower(lang)]
	if !exists {
		return fmt.Errorf("不支持的语言模板: %s\n支持的语言: go, react, c++, c, matlab", lang)
	}

	gitignorePath, err := findGitignoreFile()
	if err != nil {
		return fmt.Errorf("查找.gitignore文件失败: %v", err)
	}

	existingContent, err := readGitignore(gitignorePath)
	if err != nil {
		return fmt.Errorf("读取.gitignore文件失败: %v", err)
	}

	// 如果文件已存在且有内容，检查是否已包含该模板
	if existingContent != "" {
		// 检查是否已经包含这个模板的内容（简单检查）
		if strings.Contains(existingContent, template[:50]) {
			fmt.Printf("警告: .gitignore文件已存在且可能包含%s模板内容\n", lang)
			fmt.Printf("文件位置: %s\n", gitignorePath)
			return nil
		}
		// 追加模板内容
		newContent := existingContent
		if !strings.HasSuffix(newContent, "\n") {
			newContent += "\n"
		}
		newContent += "\n# " + lang + " template\n" + template
		err = writeGitignore(gitignorePath, newContent)
	} else {
		// 创建新文件
		err = writeGitignore(gitignorePath, template)
	}

	if err != nil {
		return fmt.Errorf("写入.gitignore文件失败: %v", err)
	}

	fmt.Printf("成功生成%s模板的.gitignore文件\n", lang)
	fmt.Printf("文件位置: %s\n", gitignorePath)
	return nil
}

// 添加文件或文件夹到.gitignore
func addToGitignore(path string) error {
	// 检查路径是否存在
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("获取绝对路径失败: %v", err)
	}

	// 获取相对于.gitignore文件的路径
	gitignorePath, err := findGitignoreFile()
	if err != nil {
		return fmt.Errorf("查找.gitignore文件失败: %v", err)
	}

	gitignoreDir := filepath.Dir(gitignorePath)
	relPath, err := filepath.Rel(gitignoreDir, absPath)
	if err != nil {
		// 如果无法计算相对路径，使用绝对路径
		relPath = absPath
	}

	// 如果是目录，添加斜杠
	info, err := os.Stat(absPath)
	if err == nil && info.IsDir() {
		if !strings.HasSuffix(relPath, "/") {
			relPath += "/"
		}
	}

	// 读取现有内容
	existingContent, err := readGitignore(gitignorePath)
	if err != nil {
		return fmt.Errorf("读取.gitignore文件失败: %v", err)
	}

	// 检查是否已存在
	if containsIgnore(existingContent, relPath) {
		fmt.Printf("路径已存在于.gitignore中: %s\n", relPath)
		fmt.Printf("文件位置: %s\n", gitignorePath)
		return nil
	}

	// 添加新路径
	newContent := existingContent
	if newContent != "" && !strings.HasSuffix(newContent, "\n") {
		newContent += "\n"
	}
	newContent += relPath + "\n"

	err = writeGitignore(gitignorePath, newContent)
	if err != nil {
		return fmt.Errorf("写入.gitignore文件失败: %v", err)
	}

	fmt.Printf("成功添加路径到.gitignore: %s\n", relPath)
	fmt.Printf("文件位置: %s\n", gitignorePath)
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: %s [选项] [语言/路径]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "选项:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n示例:\n")
		fmt.Fprintf(os.Stderr, "  %s go              # 生成Go模板的.gitignore\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s react           # 生成React模板的.gitignore\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s c++             # 生成C++模板的.gitignore\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s c               # 生成C模板的.gitignore\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s matlab          # 生成MATLAB模板的.gitignore\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -f dir          # 添加忽略文件夹\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -f file         # 添加忽略文件\n", os.Args[0])
	}

	fileFlag := flag.Bool("f", false, "添加文件或文件夹到.gitignore")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	if *fileFlag {
		// 添加文件或文件夹
		if flag.NArg() < 1 {
			fmt.Fprintf(os.Stderr, "错误: -f选项需要指定文件或文件夹路径\n")
			os.Exit(1)
		}
		path := flag.Arg(0)
		err = addToGitignore(path)
	} else {
		// 生成模板
		lang := flag.Arg(0)
		err = generateTemplate(lang)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}
