package sql

import "github.com/FranklinThree/male-vocal-dev/config"

var mysqlDsn string

func mysqlDsnInit() {

	username := config.SqlConfig.GetString("mysql.username")
	password := config.SqlConfig.GetString("mysql.password")
	host := config.SqlConfig.GetString("mysql.host")
	port := config.SqlConfig.GetString("mysql.port")
	dbname := config.SqlConfig.GetString("mysql.Dbname")
	timeout := config.SqlConfig.GetString("mysql.timeout")

	//username:password@tcp(host:port)/Dbname?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s
	mysqlDsn = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf-8&parseTime=true&loc=Local&timeout=" + timeout + "&readTimeout=30s&writeTimeout=60s"

}
