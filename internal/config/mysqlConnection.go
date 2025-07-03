package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "root:20kl45t2@tcp(mysql.mysql.svc.cluster.local:3306)/usuarios"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro na conexão com o Banco de dados.")
	}

}
