package main

import(
	"net"
	"os"
	"fmt"
) 

var clientsConnected map[int]*net.UDPAddr
var globalNext int

func ErrorHandle(err error){
    if err  != nil {
        fmt.Println("We have a error: " , err)
        fmt.Println("Exitting") 
		os.Exit(0)
    }
}
 
func PrintErrorIfExists(err error){
    if err != nil {
        fmt.Println("We have a error: ",err)
    } 
}

func Reverse(s string) string {    
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }        
    return string(runes)
}

func ConnectClientHandle(cliAddr *net.UDPAddr){
    clientsConnected[globalNext] = cliAddr
    globalNext += 1
}

func main() {

    globalNext = 0
    clientsConnected = make( map[int]*net.UDPAddr)

    fmt.Println("Server Initiated...")

	ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
	ErrorHandle(err)
	fmt.Println("Address solved...")
	
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
	ErrorHandle(err)
	fmt.Println("Server Working...")
    defer ServerConn.Close()
    defer fmt.Println("Exiting")
 
    buf := make([]byte, 1024)
	fmt.Println("Waiting Mensages...")
    for {

        n,addr,err := ServerConn.ReadFromUDP(buf)
        PrintErrorIfExists(err)
        
        ConnectClientHandle(addr)
    
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
        fmt.Println("Sending Response to ", addr)
                
        reverseString := Reverse(string(buf[0:n]))
        messageBack := []byte(reverseString)

        for _, value := range clientsConnected {
            _, err = ServerConn.WriteToUDP(messageBack, value)
            PrintErrorIfExists(err)
        }
        

    }
}

