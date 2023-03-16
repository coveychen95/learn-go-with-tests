# 通过 Go 语言学习测试驱动开发
# 来源
GitHub：https://github.com/quii/learn-go-with-tests
Gitbook：https://quii.gitbook.io/learn-go-with-tests


在初始化一个 Go 项目之前，请确保已经在您的系统上安装了 Go。您可以在终端中运行 go version 命令来检查 Go 是否已经安装并且已经配置好 PATH 环境变量。

要初始化一个 Go 项目，请按照以下步骤操作：

1. 创建项目目录。在您选择的位置上创建一个新目录来存放项目文件。例如，您可以在终端中运行以下命令：


```shell
mkdir my-project
```
2. 进入项目目录。在终端中运行以下命令：


```shell
cd my-project
```
3. 初始化项目。在终端中运行以下命令：


```shell
go mod init github.com/your-username/my-project
```

注意将 github.com/your-username/my-project 替换为您的实际项目名称和路径。这个命令将创建一个新的 go.mod 文件，用于管理您的项目依赖。

4. 创建您的 Go 代码文件。在您的项目目录中创建一个新的 Go 源代码文件。例如，您可以在终端中运行以下命令：

```shell
touch main.go
```

这个命令将在您的项目目录中创建一个名为 main.go 的空文件。

5. 在 main.go 文件中添加代码。现在，您可以使用您最喜欢的文本编辑器或 IDE 来打开 main.go 文件，并添加您的 Go 代码。

6. 运行您的代码。在终端中运行以下命令：

```shell
go run main.go
```

这个命令将编译并运行您的 Go 代码文件。

以上步骤将帮助您初始化一个新的 Go 项目，并且在其中添加和运行您的第一个 Go 代码文件。