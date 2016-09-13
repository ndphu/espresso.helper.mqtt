package mqtt_helper

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	appconfig "github.com/ndphu/espresso.appconfig"
	"strconv"
	"time"
)

func NewClientOpts(appConfig *appconfig.AppConfig) *mqtt.ClientOptions {
	brokerUrl := fmt.Sprintf(appConfig.Server.MQTT.Protocol + "://" + appConfig.Server.MQTT.Host + ":" + strconv.Itoa(appConfig.Server.MQTT.Port))
	opts := mqtt.NewClientOptions().AddBroker(brokerUrl)
	opts.SetUsername(appConfig.Server.MQTT.User)
	opts.SetPassword(appConfig.Server.MQTT.Password)
	opts.SetClientID(appConfig.Device.Id)
	return opts
}

func NewClient(conf *appconfig.AppConfig) mqtt.Client {
	opts := NewClientOpts(conf)
	opts.SetClientID(conf.Device.Id + fmt.Sprintf("%d", time.Now().Nanosecond()))
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
