package logging

import (
	"fmt"
	"property/framework/pkg/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.FileConfigSetting.App.RuntimeRootPath, setting.FileConfigSetting.App.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format(setting.FileConfigSetting.App.TimeFormat),
		setting.FileConfigSetting.App.LogFileExt,
	)
}
