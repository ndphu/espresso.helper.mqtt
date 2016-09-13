package mqtt_helper_test

import (
	"fmt"
	"github.com/ndphu/espresso.appconfig"
	mqtt_helper "github.com/ndphu/espresso.helper.mqtt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	TEST_HOST              = "somewhere.far.far.far.away"
	TEST_PORT              = 1990
	TEST_USERNAME          = "tester"
	TEST_PASSWORD          = "secret"
	TEST_PROTOCOL          = "tcp"
	TEST_FIREBASE_APP_NAME = "test-7a4ff"
	TEST_DEVICE_ID         = "0"
	TEST_KEY_FILE          = "key.json"
)

func TestCreateClientOpts(t *testing.T) {
	conf := appconfig.New()
	conf.Server.MQTT.Host = TEST_HOST
	conf.Server.MQTT.Port = TEST_PORT
	conf.Server.MQTT.User = TEST_USERNAME
	conf.Server.MQTT.Password = TEST_PASSWORD
	conf.Server.MQTT.Protocol = TEST_PROTOCOL
	opts := mqtt_helper.NewClientOpts(conf)
	expectedBrokerUrl := fmt.Sprintf("%s://%s:%d", TEST_PROTOCOL, TEST_HOST, TEST_PORT)
	assert.Equal(t, expectedBrokerUrl, opts.Servers[0].String(), "Broker URL mismatch")
	assert.Equal(t, TEST_USERNAME, opts.Username, "Username mismatch")
	assert.Equal(t, TEST_PASSWORD, opts.Password, "Password mismatch")
}

func GetTestFirebaseAppName() string {
	firebaseAppName, exists := os.LookupEnv("FIREBASE_APP")
	if !exists {
		return TEST_FIREBASE_APP_NAME
	}
	return firebaseAppName
}

func GetTestDeviceId() string {
	deviceId, exists := os.LookupEnv("DEVICE_ID")
	if !exists {
		return TEST_DEVICE_ID
	}
	return deviceId
}
func GetTestKeyFile() string {
	keyFile, exists := os.LookupEnv("KEY_FILE")
	if !exists {
		return TEST_KEY_FILE
	}
	return keyFile
}

func TestCreateClientOptsFromFirebase(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	conf := appconfig.New()
	conf.GetConfigFromFirebase(GetTestFirebaseAppName(), GetTestDeviceId(), GetTestKeyFile())

	client := mqtt_helper.NewClient(conf)

	assert.True(t, client.IsConnected(), "Client is not connected")

	//conf := appconfig.New()
	//conf.GetConfigFromFirebase(deviceId, authToken)

}
