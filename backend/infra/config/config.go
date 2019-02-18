package config

import (
	"os"
)

type HttpServer struct {
	Hostname	string		`json:"hostname"`
	Port		string		`json:"port"`
}

type Mqtt struct {
	Host		string		`json:"host"`
	Port		string		`json:"port"`
	User		string		`json:"user"`
	Pass		string		`json:"pass"`
	ClientName	string		`json:"client_name"`
}

func CreateConfig() (*HttpServer, *Mqtt) {
	httpConfig   := &HttpServer{}
	mqttConf := &Mqtt{}

	//Http Config
	httpConfig.Hostname			=	os.Getenv("HTTP_HOST")
	httpConfig.Port 			=	os.Getenv("HTTP_PORT")

	//MQTT Server Config
	mqttConf.Host 				=	os.Getenv("MQTT_HOST")
	mqttConf.Port 				=	os.Getenv("MQTT_PORT")
	mqttConf.User 				=	os.Getenv("MQTT_USER")
	mqttConf.Pass 				=	os.Getenv("MQTT_PASS")
	mqttConf.ClientName			=	os.Getenv("MQTT_CLIENT")

	return httpConfig, mqttConf
}
