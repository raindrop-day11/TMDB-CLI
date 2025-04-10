# TMDB CLI API
在这个项目中，你将构建一个简单的命令行界面CLI应用程序，用于从电影数据库（TMDB）获取数据并且在终端显示。该项目将帮助你练习编程技能，包括使用API，处理JSON数据以及构建简单的CLI应用程序

## Requirements
该应用程序应该通过命令行界面运行，并能够从TMDB API获取并显示流行电影，高分电影，即将上映电影和当前热映电影。用户应能通过向CLI工具传递命令行参数来指定想要查看的电影类型。

这个命令行工具应该可以这样使用：
```bash
tmdb-app --type "playing"
tmdb-app --type "popular"
tmdb-app --type "top"
tmdb-app --type "upcoming"
```

## 注意以下事项
1. 优雅的处理各类错误
2. ReadMe文件中包含：应用程序运行方法

----------------------------------------------------------

### Usage
```bash
./tmdb-app.exe --type "type"
```

```bash
./tmdb-app.exe

# 只运行这个，表示运行API
# http://localhost:port/TMDB/:movie?language=l&page=1
```