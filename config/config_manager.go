package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type ConfigManager interface {
	ReadEnvInt(key string) int
	ReadEnvFloat64(key string) float64
	ReadEnvString(key string) string
	ReadEnvBoolWithDefault(key string, defaultVal bool) bool
	MustGetStringSlice(key string) []string
	CheckIfSet(key string)
	ReadEnvStringWithDefault(key string, defaultVal string) string
	ReadEnvIntWithDefault(key string, defaultVal int) int
}

type viperConfigManager struct {
	viper *viper.Viper
}

func newViperConfigManager(filePath string) ConfigManager {
	v := viper.New()
	v.SetConfigFile(filePath)
	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Cannot read env | Error: %v", err))
	}

	return &viperConfigManager{
		viper: v,
	}
}

// ReadEnvInt is function to get environment variable with type integer
func (cm *viperConfigManager) ReadEnvInt(key string) int {
	cm.CheckIfSet(key)
	v, err := strconv.Atoi(cm.viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer", key))
	}

	return v
}

// ReadEnvIntWithDefault is function to get environment variable with type int and default value
func (cm *viperConfigManager) ReadEnvIntWithDefault(key string, defaultVal int) int {
	value := cm.viper.GetString(key)
	if value == "" {
		return defaultVal
	}

	v, err := strconv.Atoi(value)

	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid int", key))
	}

	return v
}

// ReadEnvFloat64 is function to get environment variable with type float64
func (cm *viperConfigManager) ReadEnvFloat64(key string) float64 {
	cm.CheckIfSet(key)
	v, err := strconv.ParseFloat(cm.viper.GetString(key), 64)
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid float", key))
	}

	return v
}

// ReadEnvString is function to get environment variable with type string
func (cm *viperConfigManager) ReadEnvString(key string) string {
	cm.CheckIfSet(key)
	return cm.viper.GetString(key)
}

// ReadEnvStringWithDefault is function to get environment variable with type string and default value
func (cm *viperConfigManager) ReadEnvStringWithDefault(key string, defaultVal string) string {
	value := cm.viper.GetString(key)
	if value == "" {
		return defaultVal
	}
	return value
}

// ReadEnvBoolWithDefault is function to get environment variable with type boolean and default value
func (cm *viperConfigManager) ReadEnvBoolWithDefault(key string, defaultVal bool) bool {
	value := cm.viper.GetString(key)
	if value == "" {
		return defaultVal
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal
	}

	return boolVal
}

func (cm *viperConfigManager) MustGetStringSlice(key string) []string {
	cm.CheckIfSet(key)
	s := cm.viper.GetStringSlice(key)

	if len(s) == 0 {
		panic(fmt.Sprintf("key %s have zero length expected a string slice", key))
	}

	return s
}

func (cm *viperConfigManager) CheckIfSet(key string) {
	if !cm.viper.IsSet(key) {
		err := fmt.Errorf("key %s is not set", key)
		panic(err)
	}
}
