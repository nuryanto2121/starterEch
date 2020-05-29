package logging

import (
	"fmt"
	"log"
	"property/framework/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // bridge
)

var conn *gorm.DB

type auditLog struct {
	AuditLogID int       `json:"audit_log_id" gorm:"primary_key"`
	Level      string    `json:"level"`
	UUID       string    `json:"uuid"`
	FuncName   string    `json:"func_name"`
	FileName   string    `json:"file_name"`
	Line       int       `json:"line"`
	Time       string    `json:"time"`
	Message    string    `json:"message"`
	CreatedBy  string    `json:"created_by" gorm:"type:varchar(20)"`
	CreatedAt  time.Time `json:"crated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdateBy   string    `json:"updated_by" gorm:"type:varchar(20)"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (a *auditLog) saveAudit() {
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
	conn, err = gorm.Open(setting.FileConfigSetting.Database.Type, connectionstring)

	if err != nil {
		log.Printf("connection.setup err : %v", err)
		panic(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.FileConfigSetting.Database.TablePrefix + defaultTableName
	}
	conn.SingularTable(true)
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(100)

	go autoMigrate()

	err = conn.Create(&a).Error
	if err != nil {
		log.Printf("%s", err)
	}
	// defer conn.Close()
}

// autoMigrate : create or alter table from struct
func autoMigrate() {
	// Add auto migrate bellow this line
	log.Println("STARTING AUTO MIGRATE LOG")
	conn.AutoMigrate(
		auditLog{},
	)

	log.Println("FINISHING AUTO MIGRATE LOG")
}
