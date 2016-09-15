package mqtt_helper

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	appconfig "github.com/ndphu/espresso.appconfig"
	"strconv"
	"time"
)

func NewClientOpts(conf *appconfig.AppConfig) *mqtt.ClientOptions {
	var brokerUrl string
	switch {
	case conf.Schema == "1.0":
		brokerUrl = fmt.Sprintf(conf.Server.MQTT.Protocol + "://" + conf.Server.MQTT.Host + ":" + strconv.Itoa(conf.Server.MQTT.Port))
	case conf.Schema == "2.0":
		brokerUrl = conf.Server.MQTT.BrokerUrl
	}
	opts := mqtt.NewClientOptions().AddBroker(brokerUrl)
	opts.SetUsername(conf.Server.MQTT.User)
	opts.SetPassword(conf.Server.MQTT.Password)
	opts.SetClientID(conf.Device.Id)
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
