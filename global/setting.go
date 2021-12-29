package global

import (
    "github.com/spark8899/go-blog-service/pkg/logger"
    "github.com/spark8899/go-blog-service/pkg/setting"
)

var (
    ServerSetting   *setting.ServerSettingS
    AppSetting      *setting.AppSettingS
    EmailSetting    *setting.EmailSettingS
    JWTSetting      *setting.JWTSettingS
    DatabaseSetting *setting.DatabaseSettingS
    Logger          *logger.Logger
)
