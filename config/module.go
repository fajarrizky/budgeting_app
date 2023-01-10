package config

type ConfigModule struct {
	configService ConfigService
}

func NewConfigModule() *ConfigModule {

	configService := newConfigService("env/app.env")

	return &ConfigModule{
		configService: configService,
	}

}

func (m *ConfigModule) GetConfigService() ConfigService {
	return m.configService
}
