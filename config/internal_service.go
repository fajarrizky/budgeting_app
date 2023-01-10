package config

import (
)

type InternalServiceConfig struct {
	ServiceConfig ServiceConfig
}

type ServiceConfig struct {
	
}

func getInternalServicesConfig(cm ConfigManager) InternalServiceConfig {
	return InternalServiceConfig{
		ServiceConfig: ServiceConfig{
		},
	}
}
