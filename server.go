package main

// /*
// Solve problems :
// 	-support connect
// 	-resolve respons
// 	-decode json
// 	-encode json
// 	-respons for client
// */
// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net"
// )

// type User struct {
// 	Username string
// 	Password string
// 	Do       string
// }

// type ServerConf struct {
// 	host  string
// 	port  int
// 	Debug bool
// }

// func (conf *ServerConf) init_server() error {
// 	var err error
// 	if conf.Debug {
// 		fmt.Println("Server Start")
// 	}
// 	// listen to incoming udp packets
// 	pc, err := net.ListenPacket("udp", fmt.Sprint("%s:%v", conf.host, conf.port))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer pc.Close()

// 	for {
// 		buf := make([]byte, 1024)
// 		n, addr, err := pc.ReadFrom(buf)
// 		if err != nil {
// 			continue
// 		}
// 		go stream(pc, addr, buf[:n])
// 	}

// 	return err
// }

// func stream(pc net.PacketConn, addr net.Addr, buf []byte) {
// 	var m User
// 	json.Unmarshal(buf, &m)
// 	fmt.Println(m)
// 	switch t := m.Do; t {
// 	case "0":
// 		err := test_autch(m.Username, m.Password)
// 		fmt.Println(err)
// 	case "1":
// 		_, err := create_user(m.Username, m.Password)
// 		fmt.Println(err)
// 	default:
// 		fmt.Println("Fail Do")
// 	}
// 	pc.WriteTo([]byte("Ok"), addr)
// }

// func (ServerConf) stop_server() error {
// 	var err error
// 	return err
// }
