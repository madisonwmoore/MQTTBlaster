package main

import ("fmt"
 "bufio"
  "os"
"strings"
// "net"
MQTT "github.com/eclipse/paho.mqtt.golang"
"time"
)



var onconnect MQTT.OnConnectHandler =func(client MQTT.Client){
	fmt.Println("Connected to broker")
}


func main() {

	opt:=MQTT.NewClientOptions()
	opt.AddBroker("tcp://test.mosquitto.org:1883")
	opt.SetClientID("MQTTblaster")
	opt.OnConnect=onconnect

	client := MQTT.NewClient(opt)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	fmt.Println("/********/")
	fmt.Println("MQTTBlaster Version 1.0")
	fmt.Println("/********/")
	
	reader := bufio.NewReader(os.Stdin)
	
	command,_:=reader.ReadString('\n')
	command = strings.TrimSpace(command)

	switch command{
	case "run":
		fmt.Println("Nothing to run")
	}
}