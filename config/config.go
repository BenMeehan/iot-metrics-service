package config

import (
	"github.com/spf13/viper"
)

type Config struct {
    InfluxDB struct {
        URL    string
        Token  string
        Org    string
        Bucket string
    }
    NATS struct {
        URL string
    }
}

func LoadConfig() Config {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    
    var cfg Config
    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

    err := viper.Unmarshal(&cfg)
    if err != nil {
        panic(err)
    }

    return cfg
}
