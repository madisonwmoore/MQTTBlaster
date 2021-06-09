package main

import ("fmt"
MQTT "github.com/eclipse/paho.mqtt.golang")

var onconnect MQTT.OnConnectHandler =func(client MQTT.Client){
	fmt.Println("Connected to broker")
}

func connect(host string, clientID string){
	opt:=MQTT.NewClientOptions()
	opt.AddBroker(host)
	opt.SetClientID(clientID)
	opt.OnConnect=onconnect

	client := MQTT.NewClient(opt)
	fmt.Println("Connecting to ", host, "with clientID ", clientID)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}