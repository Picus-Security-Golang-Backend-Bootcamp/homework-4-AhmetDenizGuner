package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

//Connect create database connection with given parameters
func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=library port=5432" //sslmode=disable TimeZone=Asia/Shanghai
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "table_",
			SingularTable: true,
		},
		/*NowFunc: func() time.Time {
			return time.Now().UTC()
		},*/
	})

	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %s", err.Error()))
	}

	DB = connection

}
