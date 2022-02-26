package main

import "mallapi/initialize"

// @title 商城系统后台接口文档
// @version 1.0
// @description gin restful
// @termsOfService http://swagger.io/terms/

// @contact.name Jimmy
// @contact.url http://www.swagger.io/support
// @contact.email tangzhiming90@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 8000
// @BasePath web/api/v1

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Router()
}
