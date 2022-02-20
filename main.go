package main

import "mallapi/initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Router()
}
