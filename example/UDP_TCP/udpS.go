package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	PORT := ":" + "30000"
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 4096)
	rand.Seed(time.Now().Unix())
	x := 0
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		x += n
		fmt.Println(x, n, addr, err)
		//fmt.Printf("[%v___", g.SliceStr{string(buffer  )})
		//fmt.Printf("[%v___", g.SliceStr{string(buffer[0 : n-2])})

		//if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
		//	fmt.Println("Exiting UDP server!")
		//	return
		//}

		//data := []byte(strconv.Itoa(random(1, 1001)))
		//fmt.Printf("data: %s\n", string(data))
		//_, err = connection.WriteToUDP(data, addr)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
	}
}
