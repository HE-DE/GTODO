package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	fmt.Println("Config Inited!")
}

// 初始化数据库
func InitDB() {
	DBUrl := GetDBUrl()
	fmt.Println("DBUrl: ", DBUrl)
	var err error
	DB, err = gorm.Open(mysql.Open(DBUrl))
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Inited!")
}

// 获取文件配置并生成MySQL路径
func GetDBUrl() string {
	username := viper.GetString("mysql.username")
	pwd := viper.GetString("mysql.pwd")
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	charset := viper.GetString("mysql.charset")
	parseTime := viper.GetString("mysql.parseTime")
	loc := viper.GetString("mysql.loc")
	//生成对应的MySQL路径
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", username, pwd, ip, port, database, charset, parseTime, loc)
}
