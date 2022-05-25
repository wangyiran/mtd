package conf

import (
	"fmt"
	"myTodoList/model"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	configPath    string = "./conf/config.ini"
	Host          string
	Port          string
	Db            string
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDBName   string
)

func LoadConfig() {
	file, err := ini.Load(configPath)
	if err != nil {
		panic(err)
	}
	//使用file里的参数初始化服务器以及mysql的参数
	LoadServer(file)
	LoadMysql(file)
	connPath := strings.Join([]string{MysqlUser, ":", MysqlPassword, "@tcp(", MysqlHost, MysqlPort, ")/", MysqlDBName, "?charset=utf8mb4&parseTime=True"}, "")
	model.DataBase(connPath)
	//数据库迁移
	if err := model.Migrator(); err != nil {
		panic("数据迁移失败")
	} else {
		fmt.Println("数据迁移成功！")
	}

}
func LoadServer(file *ini.File) {
	Host = file.Section("service").Key("Host").String()
	Port = file.Section("service").Key("Port").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	MysqlHost = file.Section("mysql").Key("MysqlHost").String()
	MysqlPort = file.Section("mysql").Key("MysqlPort").String()
	MysqlUser = file.Section("mysql").Key("MysqlUser").String()
	MysqlPassword = file.Section("mysql").Key("MysqlPassword").String()
	MysqlDBName = file.Section("mysql").Key("MysqlDBName").String()
}
