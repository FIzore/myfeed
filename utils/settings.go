package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	RunMode    string
	ServerPort string
	Db         string
	DbAddress  string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func main() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("配置文件加载失败", err)
	}
	loadServer(file)
	loadDatabase(file)
}
func loadServer(file *ini.File) {
	RunMode = file.Section("server").Key("RunMode").MustString("debug")
	ServerPort = file.Section("server").Key("ServerPort").MustString("3000")
}
func loadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbAddress = file.Section("database").Key("DbAddress").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("6676")
	DbUser = file.Section("database").Key("DbUser").MustString("myfeed")
	DbPassword = file.Section("database").Key("DbPassword").MustString("myfeed")
}
