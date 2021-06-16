package main

import ("fmt"
MQTT "github.com/eclipse/paho.mqtt.golang")

type MQTTBlaster struct{
host string
clientID string
client MQTT.Client
}


func (mqtt_struct MQTTBlaster) connect(){

	var onconnect MQTT.OnConnectHandler =func(client MQTT.Client){
		fmt.Println("Connected to broker")
	}

	opt:=MQTT.NewClientOptions()
	opt.AddBroker(mqtt_struct.host)
	opt.SetClientID(mqtt_struct.clientID)
	opt.OnConnect=onconnect

	client := MQTT.NewClient(opt)
	fmt.Println("Connecting to ", mqtt_struct.host, "with clientID ", mqtt_struct.clientID)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (mqtt_struct MQTTBlaster) publish(){
	message := "Test Message"
	token := mqtt_struct.client.Publish("topic/test",0, false, message)
	token.Wait()
}