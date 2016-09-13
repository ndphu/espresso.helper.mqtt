package mqtt_helper

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	appconfig "github.com/ndphu/espresso.appconfig"
	"strconv"
)

type mqttApplication struct {
	AppConfig *appconfig.AppConfig
	Opts      *mqtt.ClientOptions
}

func (mqttApp *mqttApplication) Init(appConfig *appconfig.AppConfig) {
	mqttApp.Opts = CreateClientOpts(appConfig)
}

func CreateClientOpts(appConfig *appconfig.AppConfig) *mqtt.ClientOptions {
	brokerUrl := fmt.Sprintf(appConfig.Server.MQTT.Protocol + "://" + appConfig.Server.MQTT.Host + ":" + strconv.Itoa(appConfig.Server.MQTT.Port))
	opts := mqtt.NewClientOptions().AddBroker(brokerUrl)
	opts.SetUsername(appConfig.Server.MQTT.User)
	opts.SetPassword(appConfig.Server.MQTT.Password)
	opts.SetClientID(appConfig.Device.Id)
	return opts
}
