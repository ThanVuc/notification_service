package settings

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This file create a configuration struct for the application.
	@Note: The configuration struct is used to store the configuration of the application.
	All the configuration is loaded from a YAML file using Viper.
	And it is stored in the global variable `Config` in the `global` package.
*/

type Config struct {
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
}

type Log struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	FileLogPath string `mapstructure:"file_log_path" json:"file_log_path" yaml:"file_log_path"`
	MaxSize     int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxAge      int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress    bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
