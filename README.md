# gg

一个简单易用的 `.gitignore` 文件管理工具，支持快速生成各种语言的 gitignore 模板，以及添加自定义文件/文件夹到忽略列表。

## 功能特性

- 🚀 快速生成多种语言的 `.gitignore` 模板
- 📁 智能查找 `.gitignore` 文件（自动向上查找）
- ➕ 便捷添加文件/文件夹到忽略列表
- 🔍 自动检测重复，避免重复添加
- 📝 支持相对路径和绝对路径

## 支持的语言模板

- `go` - Go 语言
- `react` - React 项目
- `c++` - C++ 项目
- `c` - C 语言
- `matlab` - MATLAB 项目

可以支持更多模板，当前工作太多没时间加：https://github.com/github/gitignore

## 安装

### 方法一：下载最新的release包 （推荐）

- https://github.com/panyingyun/gg/releases

### 方法二：使用 go install（推荐）

```bash
go install github.com/panyingyun/gg@latest
```
安装完成后，确保 `$GOPATH/bin` 或 `$HOME/go/bin` 在你的 `PATH` 环境变量中，然后就可以直接使用 `gg` 命令了。

### 方法三：从源码编译

```bash
git clone https://github.com/panyingyun/gg.git
cd gg
make build
sudo cp gg /usr/local/bin
```

### 从 Release 下载

访问 [Releases](https://github.com/panyingyun/gg/releases) 页面下载对应平台的二进制文件。

## 使用方法

### 生成语言模板

```bash
# 生成 Go 模板
gg go

# 生成 React 模板
gg react

# 生成 C++ 模板
gg c++

# 生成 C 模板
gg c

# 生成 MATLAB 模板
gg matlab
```

### 添加文件/文件夹到 .gitignore

```bash
# 添加文件
gg -f filename.txt

# 添加文件夹
gg -f directory/

# 使用相对路径或绝对路径
gg -f ./path/to/file
gg -f /absolute/path/to/file
```

## 工作原理

1. **查找 .gitignore 文件**：工具会从当前目录开始，向上查找 `.gitignore` 文件
2. **创建或更新**：如果找到 `.gitignore` 文件，则追加内容；如果未找到，则在当前目录创建新文件
3. **智能路径处理**：添加文件/文件夹时，会自动计算相对路径，确保路径正确

## 示例

```bash
# 在 Go 项目中生成 .gitignore
$ gg go
成功生成go模板的.gitignore文件
文件位置: /path/to/project/.gitignore

# 添加构建目录到忽略列表
$ gg -f build/
成功添加路径到.gitignore: build/
文件位置: /path/to/project/.gitignore

# 添加特定文件
$ gg -f config.local.yaml
成功添加路径到.gitignore: config.local.yaml
文件位置: /path/to/project/.gitignore
```

## 许可证

本项目采用 GPLv3 许可证，详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 作者

[panyingyun](https://github.com/panyingyun)

