package boostrap

import (
	"flag"
	"fmt"
	"github.com/go-tempest/tempest/config"
	"github.com/go-tempest/tempest/db/gorm"
	"github.com/go-tempest/tempest/db/sql"
	"github.com/go-tempest/tempest/log"
	"github.com/go-tempest/tempest/register"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type bootstrap struct {
	sync.Once
}

func (b *bootstrap) start() {
	b.Do(func() {
		new(register.Registration).StartIfNecessary()
		// TODO 启动
	})
}

func getConfigPath() *string {
	var configPath = flag.String("c", "", "Specifies the path used to search the configuration file")
	flag.Parse()
	return configPath
}

func parseYaml(v *viper.Viper) {
	if err := v.Unmarshal(&config.TempestConfig); err != nil {
		fmt.Println("初始化配置失败") // TODO 后续替换成通用日志组件
		os.Exit(-1)
	}
}

func init() {
	initViper()
	log.Initialize()
	gorm.Initialize()
	sql.Initialize()
	new(bootstrap).start()
}

func initViper() {
	viper.AutomaticEnv()
	viper.SetConfigType(config.DefaultConfigType)
	viper.SetConfigName(config.DefaultConfigName)
	viper.AddConfigPath(*getConfigPath())

	v := viper.GetViper()
	parseYaml(v)
}
