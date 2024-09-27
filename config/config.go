// config 包用于处理应用程序的配置
package config

// 导入 viper 包，用于读取和解析配置文件
import "github.com/spf13/viper"

// Config 结构体定义了应用程序的配置结构
type Config struct {
	// App 结构体包含应用程序的基本配置
	App struct {
		Name string // 应用程序名称
		Port int    // 应用程序运行的端口
	}
	// Database 结构体包含数据库的配置信息
	Database struct {
		Host     string // 数据库主机地址
		Port     int    // 数据库端口
		User     string // 数据库用户名
		Password string // 数据库密码
		Name     string // 数据库名称
	}
}

// AppConfig 是一个指向 Config 结构体的指针变量
// 它用于存储应用程序的全局配置信息
// 声明为包级变量，可以在整个包中访问
// 初始值为 nil，需要在程序中进行初始化才能使用
var AppConfig *Config

// InitConfig 函数用于初始化配置
func InitConfig() {
	// 设置配置文件的名称
	viper.SetConfigName("config")
	// 设置配置文件的类型为 YAML
	viper.SetConfigType("yaml")
	// 添加配置文件的搜索路径
	viper.AddConfigPath("./config")
	// 读取配置文件
	err := viper.ReadInConfig()
	// 如果读取配置文件出错，则触发 panic
	if err != nil {
		panic("读取配置文件时出错")
	}

	// 创建一个新的 Config 结构体实例，并将其地址赋值给 AppConfig 变量
	// & 操作符用于获取 Config{} 实例的内存地址
	// 这为后续的 viper.Unmarshal() 函数准备了一个可以填充配置数据的目标结构体
	AppConfig = &Config{}

	// 将配置文件的内容解析到 AppConfig 结构体中
	err = viper.Unmarshal(&AppConfig)
	// 如果解析出错，则触发 panic
	if err != nil {
		panic("解析配置文件时出错")
	}
}
