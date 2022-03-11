package config

// Config 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql Mysql `mapstructure:"mysql"`
	Jwt Jwt `mapstructure:"jwt"`
	Code2Session Code2Session `mapstructure:"code2Session"`
	QiniuStorage QiniuStorage `mapstructure:"qiniuStorage"`
}

// Server 服务启动端口号配置
type Server struct {
	Post string `mapstructure:"post"`
}

// Mysql 数据源配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url string `mapstructure:"url"`
}

// Jwt 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}

// Code2Session 微信小程序相关配置
type Code2Session struct {
	Code string `mapstructure:"code"`
	AppId string `mapstructure:"appId"`
	AppSecret string `mapstructure:"appSecret"`
}

// QiniuStorage 七牛文件存储配置
type QiniuStorage struct {
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	Bucket string `mapstructure:"bucket"`
	QiniuServer string `mapstructure:"qiniuserver"`
}