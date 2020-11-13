package main

import (
	//"crypto/md5"
	"encoding/base64"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func client(username string, serverIP string, clientIP string, serverPort int, clientUDPPort int) {
	fmt.Println("Iniciando cliente")
	serverURL := serverIP + ":" + strconv.Itoa(serverPort)
	connTCP, err := net.Dial("tcp", serverURL)
	buffer := make([]byte, 1024)

	if err != nil {
		fmt.Println("Error: No se pudo establecer la conexión con el servidor")
	} else {
		fmt.Println("Autenticando al usuario")
		connTCP.Write([]byte("helloiam " + username))
		_, err := connTCP.Read(buffer)

		if err != nil {
			fmt.Println("Error inesperado en la conexión")
			return
		}

		if strings.Contains(string(buffer), "ok") {
			connTCP.Write([]byte("msglen"))
			_, err := connTCP.Read(buffer)

			if err != nil {
				fmt.Println("Error inesperado en la conexión")
				return
			}

			if strings.Contains(string(buffer), "ok") {
				fmt.Println("Obteniendo mensaje")
				connUDP, er := net.ListenPacket("udp", clientIP+":"+strconv.Itoa(clientUDPPort))

				if er != nil {
					fmt.Println("Error inesperado en la conexión")
				}

				connTCP.Write([]byte("givememsg " + strconv.Itoa(clientUDPPort)))
				connTCP.Read(buffer)

				if strings.Contains(string(buffer), "ok") {
					connUDP.ReadFrom(buffer)
					message, _ := base64.StdEncoding.DecodeString(string(buffer))
					fmt.Println("Mensaje recibido: " + string(message))
				}
			}
		} else {
			fmt.Println("Error: No se pudo autenticar al usuario")
			connTCP.Close()
			return
		}

		connTCP.Write([]byte("bye"))
		connTCP.Close()
	}
}

func main() {
	var username string
	var serverIP string
	var clientIP string
	var serverPort int
	var clientUDPPort int

	fmt.Println("Indique su nombre de usuario:")
	fmt.Scanf("%s", &username)
	fmt.Println("Indique la dirección ip del servidor: (Ej: 192.168.0.1)")
	fmt.Scanf("%s", &serverIP)
	fmt.Println("Indique su dirección ip: (Ej: 192.168.0.1, si es la misma que la del servidor puede escribir 127.0.0.1)")
	fmt.Scanf("%s", &clientIP)
	fmt.Println("Indique el puerto del servidor:")
	fmt.Scanf("%d", &serverPort)
	fmt.Println("Indique puerto por donde quiere recibir el mensaje:")
	fmt.Scanf("%d", &clientUDPPort)

	client(username, serverIP, clientIP, serverPort, clientUDPPort)
}
