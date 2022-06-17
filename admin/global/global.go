package global

import (
	"gin-web-demo/Test/gin-demo/admin/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	ET_DB     *gorm.DB
	ET_CONFIG config.Server
	ET_VP     *viper.Viper
)
