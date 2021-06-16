package main

import ("fmt"
"flag"
//  "bufio"
//   "os"
// "strings"
// "net"
// MQTT "github.com/eclipse/paho.mqtt.golang"
// "time"

)






func main() {


	// opt:=MQTT.NewClientOptions()
	// opt.AddBroker("tcp://test.mosquitto.org:1883")
	// opt.SetClientID("MQTTblaster")
	// opt.OnConnect=onconnect

	// client := MQTT.NewClient(opt)

	// if token := client.Connect(); token.Wait() && token.Error() != nil {
	// 	panic(token.Error())
	// }
	fmt.Println("/********/")
	fmt.Println("MQTTBlaster Version 1.0")
	fmt.Println("/********/")

	broker := flag.String("b","","Address of the broker")
	clientID := flag.String("c","MQTTblaster","ClientID to be used with connection")
	flag.Parse()

	if(*broker==""){
		fmt.Println("Enter Broker Address:")
	}
	
	fmt.Println("Broker is ",*broker,*clientID)

	blaster := MQTTBlaster{host:*broker,clientID: *clientID}

	blaster.connect()


	
	// reader := bufio.NewReader(os.Stdin)
	
	// command,_:=reader.ReadString('\n')
	// command = strings.TrimSpace(command)

	// switch command{
	// case "run":
	// 	fmt.Println("Nothing to run")
	// }
}