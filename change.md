# t1 操作

1. 创建了 `config` 包，包含 `config.go` 文件，用于处理应用程序配置。

2. 在 `config.go` 中：
   - 定义了 `Config` 结构体，包含应用程序和数据库配置。
   - 创建了 `AppConfig` 全局变量。
   - 实现了 `InitConfig` 函数，用于初始化配置。

3. 在项目根目录创建了 `config/config.yml` 文件，包含应用程序和数据库的配置信息。

4. 更新了 `main.go`：
   - 导入了 `exchangeapp/config` 包。
   - 在 `main` 函数中调用 `config.InitConfig()`。
   - 添加了打印配置信息的代码。

5. 更新了 `go.mod` 和 `go.sum` 文件，添加了 `github.com/spf13/viper` 依赖。

6. 创建了 `笔记.md` 文件，记录了项目使用的工具和技术栈。

这些更改构成了项目的基本结构和配置管理系统。下一步可能需要实现具体的业务逻辑和数据库操作。
