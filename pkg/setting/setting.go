package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Server :
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting :
var ServerSetting = &Server{}

// Database :
type Database struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table_prefix"`
}

// DatabaseSetting :
var DatabaseSetting = &Database{}

// Setup Load config.json and map to struct
func Setup() {
	now := time.Now()
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	err = viper.Unmarshal(DatabaseSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to Unmarshal 'config.json': %v", err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
}
