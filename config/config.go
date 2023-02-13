package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ChannelSecret      string
	ChannelAccessToken string
	DbName             string
	CollectionName     string
	MongoDBAddress     string
	Address            string
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		ChannelSecret:      viper.Get("channel_secret").(string),
		ChannelAccessToken: viper.Get("channel_access_token").(string),
		DbName:             viper.Get("db_name").(string),
		CollectionName:     viper.Get("collection_name").(string),
		MongoDBAddress:     viper.Get("mongodb_address").(string),
		Address:            viper.Get("address").(string),
	}
}
