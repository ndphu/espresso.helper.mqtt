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
	TEST_HOST              = "19november.ddns.net"
	TEST_PORT              = 5384
	TEST_USERNAME          = "someone"
	TEST_PASSWORD          = "secret"
	TEST_PROTOCOL          = "tcp"
	TEST_FIREBASE_APP_NAME = "test-7a4ff"
	TEST_DEVICE_ID         = "0"
	TEST_KEY_FILE          = "key.json"
)

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

func TestCreateClientOptsFromV1(t *testing.T) {
	conf := appconfig.New()
	conf.ParseConfigFile("test_v1config.json")
	opts := mqtt_helper.NewClientOpts(conf)
	expectedBrokerUrl := fmt.Sprintf("%s://%s:%d", TEST_PROTOCOL, TEST_HOST, TEST_PORT)
	assert.Equal(t, expectedBrokerUrl, opts.Servers[0].String(), "Broker URL mismatch")
	assert.Equal(t, TEST_USERNAME, opts.Username, "Username mismatch")
	assert.Equal(t, TEST_PASSWORD, opts.Password, "Password mismatch")
}

func TestCreateClientOptsFromV2(t *testing.T) {
	conf := appconfig.New()
	conf.ParseConfigFile("test_v2config.json")
	opts := mqtt_helper.NewClientOpts(conf)
	expectedBrokerUrl := "wss://iot.eclipse.org/ws"

	assert.Equal(t, expectedBrokerUrl, opts.Servers[0].String(), "Broker URL mismatch")
	assert.Equal(t, TEST_USERNAME, opts.Username, "Username mismatch")
	assert.Equal(t, TEST_PASSWORD, opts.Password, "Password mismatch")
}

func TestCreateClientOptsFromFirebase(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	conf := appconfig.New()
	conf.GetConfigFromFirebase(GetTestFirebaseAppName(), GetTestDeviceId(), GetTestKeyFile())

	client := mqtt_helper.NewClient(conf)
	defer client.Disconnect(500)
	assert.True(t, client.IsConnected(), "Client is not connected")
}

func TestCreateClientFromV1Config(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conf := appconfig.New()
	conf.ParseConfigFile("test_v1config.json")

	client := mqtt_helper.NewClient(conf)
	defer client.Disconnect(500)
	assert.True(t, client.IsConnected(), "Client is not connected")
}

func TestCreateClientFromV2Config(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conf := appconfig.New()
	conf.ParseConfigFile("test_v2config.json")

	client := mqtt_helper.NewClient(conf)
	defer client.Disconnect(500)

	assert.True(t, client.IsConnected(), "Client is not connected")
}
