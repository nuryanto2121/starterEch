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
// var ServerSetting = &Server{}

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
// var DatabaseSetting = &Database{}

// App :
type App struct {
	JwtSecret       string   `mapstructure:"jwt_secret"`
	PageSize        int      `mapstructure:"page_size"`
	PrefixURL       string   `mapstructure:"prefix_url"`
	RuntimeRootPath string   `mapstructure:"runtime_root_path"`
	ImageSavePath   string   `mapstructure:"image_save_path"`
	ImageMaxSize    int      `mapstructure:"image_size"`
	ImageAllowExts  []string `mapstructure:"image_allow_ext"`
	ExportSavePath  string   `mapstructure:"export_save_path"`
	QrCodeSavePath  string   `mapstructure:"qr_code"`
	LogSavePath     string   `mapstructure:"log_save_path"`
	LogSaveName     string   `mapstructure:"log_save_name"`
	LogFileExt      string   `mapstructure:"log_file_ext"`
	TimeFormat      string   `mapstructure:"time_format"`
	Issuer          string   `mapstructure:"issuer"`
}

// AppSetting interface pointer
// var AppSetting = &App{}

// FileConfig :
type FileConfig struct {
	Debug    bool      `mapstructure:"debug"`
	Server   *Server   `mapstructure:"server"`
	App      *App      `mapstructure:"app"`
	Database *Database `mapstructure:"database"`
}

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
