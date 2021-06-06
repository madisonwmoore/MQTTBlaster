package main

import ("fmt"
 "bufio"
  "os"
"strings")


func main() {
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