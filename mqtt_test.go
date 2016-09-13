package mqtt_helper_test

import (
	"fmt"
	"github.com/ndphu/espresso.appconfig"
	mqtt_helper "github.com/ndphu/espresso.helper.mqtt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TEST_HOST     = "somewhere.far.far.far.away"
	TEST_PORT     = 1990
	TEST_USERNAME = "tester"
	TEST_PASSWORD = "secret"
	TEST_PROTOCOL = "tcp"
)

func TestCreateClientOpts(t *testing.T) {
	conf := appconfig.New()
	conf.Server.MQTT.Host = TEST_HOST
	conf.Server.MQTT.Port = TEST_PORT
	conf.Server.MQTT.User = TEST_USERNAME
	conf.Server.MQTT.Password = TEST_PASSWORD
	conf.Server.MQTT.Protocol = TEST_PROTOCOL
	opts := mqtt_helper.CreateClientOpts(conf)
	expectedBrokerUrl := fmt.Sprintf("%s://%s:%d", TEST_PROTOCOL, TEST_HOST, TEST_PORT)
	assert.Equal(t, expectedBrokerUrl, opts.Servers[0].String(), "Broker URL mismatch")
	assert.Equal(t, TEST_USERNAME, opts.Username, "Username mismatch")
	assert.Equal(t, TEST_PASSWORD, opts.Password, "Password mismatch")
}
