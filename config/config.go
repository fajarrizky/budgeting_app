package config

import (
	"budgetapp/module/consts"
)

type config struct {
	environment               consts.Environment
	serverPort                string
	serverShutdownGracePeriod int
	dbconfig                  DatabaseConfig
	internalServiceConfig     InternalServiceConfig
}

type ConfigService interface {
	GetEnv() consts.Environment
	GetServerPort() string
	GetServerShutdownGracePeriod() int
	GetDbConfig() DatabaseConfig
	GetInternalServicesConfig() InternalServiceConfig
}

type configService struct {
	config *config
}

func newConfigService(filePath string) ConfigService {
	var c config

	load(filePath, func(cm ConfigManager) {
		c = config{
			environment:               consts.Environment(cm.ReadEnvString("ENVIRONMENT")),
			serverPort:                cm.ReadEnvString("SERVER_PORT"),
			serverShutdownGracePeriod: cm.ReadEnvInt("SERVER_SHUTDOWN_GRACE_PERIOD_IN_SECS"),
			dbconfig:                  newDatabaseConfig(cm),
			internalServiceConfig:     getInternalServicesConfig(cm),
		}
	})

	return &configService{
		config: &c,
	}
}

func load(filePath string, fn func(cm ConfigManager)) {

	v := newViperConfigManager(filePath)

	fn(v)

}

func (c *configService) GetInternalServicesConfig() InternalServiceConfig {
	return c.config.internalServiceConfig
}

func (c *configService) GetEnv() consts.Environment {
	return c.config.environment
}

func (c *configService) GetDbConfig() DatabaseConfig {
	return c.config.dbconfig
}

func (c *configService) GetServerPort() string {
	return c.config.serverPort
}

func (c *configService) GetServerShutdownGracePeriod() int {
	return c.config.serverShutdownGracePeriod
}
