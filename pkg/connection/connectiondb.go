package connection

import (
	"fmt"
	"log"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // add database driver bridge
)

// Connections :
// type Connections struct {
// 	db *gorm.DB
// }

// // Conn :
// var Conn = &Connections{}

// Conn connection ke DB
var Conn *gorm.DB

// Setup connection to DB
func Setup() {
	now := time.Now()
	var err error
	fmt.Print(setting.FileConfigSetting.Database)
	connectionstring := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		setting.FileConfigSetting.Database.User,
		setting.FileConfigSetting.Database.Password,
		setting.FileConfigSetting.Database.Name,
		setting.FileConfigSetting.Database.Host,
		setting.FileConfigSetting.Database.Port)
	fmt.Printf("%s", connectionstring)
	Conn, err = gorm.Open(setting.FileConfigSetting.Database.Type, connectionstring)
	if err != nil {
		log.Printf("connection.setup err : %v", err)
		panic(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.FileConfigSetting.Database.TablePrefix + defaultTableName
	}
	Conn.SingularTable(true)
	Conn.DB().SetMaxIdleConns(10)
	Conn.DB().SetMaxOpenConns(100)

	go autoMigrate()

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)

}

// autoMigrate : create or alter table from struct
func autoMigrate() {
	// Add auto migrate bellow this line
	Conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	log.Println("STARTING AUTO MIGRATE ")
	Conn.AutoMigrate(
		sa_models.SaUser{},
		sa_models.SaRole{},
		sa_models.SaClient{},
		sa_models.SaBranch{},
		sa_models.SaCompany{},
		sa_models.SaUserBranch{},
		sa_models.SaUserCompany{},
		sa_models.SaRoleMenu{},
		sa_models.SaFileUpload{},
		sa_models.SaMenu{},
	)

	log.Println("FINISHING AUTO MIGRATE ")
}
