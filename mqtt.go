package main

import ("fmt"
MQTT "github.com/eclipse/paho.mqtt.golang")



func connect(host string, clientID string){

	var onconnect MQTT.OnConnectHandler =func(client MQTT.Client){
		fmt.Println("Connected to broker")
	}


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

func publish(client MQTT.Client){
	message := "Test Message"
	token := client.Publish("topic/test",0, flase,message)
	token.Wait()
}