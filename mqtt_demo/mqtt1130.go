package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func publish(msg string) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetUsername("d")
	opts.SetPassword("123456")
	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("send err")
	}
	if msg == "" {
		msg = "Hello"
	}
	mc.Publish("huati1", 0x00, true, msg)
	time.Sleep(time.Second)
}
func main() {
	publish("hello")
	for {
		var anykey string
		fmt.Scanln(&anykey)
		publish(anykey)
	}
}
