package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
)

type error struct {
	Message string
}

//new mac : 00:1A:C2:7B:00:47

func main() {
	args := getArgs()
	nameInteface, err := checkInterface(args)
	if err.Message!=""{
		log.Fatal(err.Message)
	}

	changeMac(nameInteface, args["-m"])
}


func getArgs() map[string]string{
	myMap := map[string]string{"-i":"", "-m":""}
	args := os.Args
	err := error{}
	if  err.Message = "I need more arguments !!!!" ; len(args) < 5 || (len(args) % 2 == 0) {
		log.Fatal(err.Message)
	}
	for i:=1 ; i < len(args) ; i++{
		if _ , ok := myMap[args[i]] ; ok{
			myMap[args[i]] = args[i+1]
		}
	}
	ok, _ := regexp.MatchString("^([0-9A-F]{2}[:-]){5}([0-9A-F]{2})$", myMap["-m"])
	if !ok{
		log.Fatal("The Mac address is not true !!!! ")
	}
	return myMap
}


func checkInterface(args map[string]string ) (string, error){
	interfaces, err := net.Interfaces()
	if err != nil{
		log.Fatal(err.Error())
	}
	var index int
	for i:=0 ; i < len(interfaces) ; i++{
		if interfaces[i].Name == args["-i"]{
			index = i
			break
		}else if len(interfaces) == i+1{
			return "", error{Message: "Wrong input network interface !!!!"}
		}
	}
	return interfaces[index].Name , error{Message: ""}
}


func changeMac(iFace, mac string){
	err := exec.Command("sudo" , "ifconfig", iFace, "down").Run()
	if err!=nil{
		log.Fatal(err.Error())
	}
	err = exec.Command("sudo" , "ifconfig", iFace, "hw", "ether", mac).Run()
	if err!=nil{
		log.Fatal(err.Error())
	}
	err = exec.Command("sudo" , "ifconfig", iFace, "up").Run()
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Println("The Mac address changed to ", mac)
}
