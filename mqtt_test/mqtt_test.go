package mqtt_client

import (
	"fmt"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func TestMqttClient(t *testing.T) {
	mqttClient := &MqttClient{
		IP: "ssl://10.122.48.78:1883",
		CA: "myCA\\ca.crt",
		Cert       :"myCA\\client-crt.pem",
		PrivateKey: "myCA\\client-key.pem",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}

	_ = mqttClient.Subscribe("/hdf/test", msgHandle)

	select {}
}

func msgHandle(client mqtt.Client, message mqtt.Message) {
	fmt.Println(message)
	fmt.Println(string(message.Payload()))
}