package config

import "github.com/spf13/viper"
import "log"

func InitConfigs() {
	viper.SetConfigType("env")  // Устанавливаем тип файла как .env
	viper.SetConfigFile(".env") // Устанавливаем имя файла .env
	viper.AddConfigPath(".")    // Добавляем путь к текущей директории
	viper.AutomaticEnv()        // Включаем автоматическое чтение переменных окружения

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Не удалось прочитать файл конфигурации .env: %v", err)
	}

	//// Чтение файла YAML
	//viper.SetConfigType("yaml")    // Устанавливаем тип файла как YAML
	//viper.SetConfigName("appconf") // Устанавливаем имя файла YAML
	//viper.AddConfigPath(".")       // Добавляем путь к текущей директории
	//
	//err = viper.MergeInConfig() // Слияние значений с предыдущей конфигурацией
	//if err != nil {
	//	log.Fatalf("Не удалось прочитать файл конфигурации YAML: %v", err)
	//}
}

func GetConfig(key string) string {
	return viper.GetString(key)
}

func GetList(key string) []string {
	return viper.GetStringSlice(key)
}
