package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Server :
type Server struct {
	RunMode      string        `mapstructure:"run_mode"`
	HTTPPort     int           `mapstructure:"http_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
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

// FileConfig :
type FileConfig struct {
	Debug    bool      `mapstructure:"debug"`
	Server   *Server   `mapstructure:"server"`
	Database *Database `mapstructure:"database"`
}

// TestContextType :
type TestContextType struct {
	Provider string

	ClusterLoader struct {
		Projects []struct {
			Number    int    `mapstructure:"num"`
			BaseName  string `mapstructure:"basename"`
			Tuning    string `mapstructure:"tuning"`
			Templates []struct {
				Number int    `mapstructure:"num"`
				File   string `mapstructure:"file"`
			} `mapstructure:"templates"`
		} `mapstructure:"projects"`
	}
}

// TestContextTypeSetting :
var TestContextTypeSetting = &TestContextType{}

// FileConfigSetting :
var FileConfigSetting = &FileConfig{}

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

	err = viper.Unmarshal(FileConfigSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to Unmarshal 'config.json': %v", err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
}
