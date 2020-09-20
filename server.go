package main

/*
Solve problems :
	-support connect
	-resolve respons
	-decode json
	-encode json
	-respons for client
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type User struct {
	Username string
	Password string
	Do       string
}

type ServerConf struct {
	host string
	port int
}

func (conf *ServerConf) init_server() error {
	var err error
	fmt.Println("Server Start", fmt.Sprint(conf.host, ":", conf.port))
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", fmt.Sprint(conf.host, ":", conf.port))
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go stream(pc, addr, buf[:n])
	}
	return err
}

func stream(pc net.PacketConn, addr net.Addr, buf []byte) {
	var m User
	var date_json []byte
	json.Unmarshal(buf, &m)
	fmt.Println(m)
	if err := ldap_conf.test_autch(m.Username, m.Password); err == 0 {
		date_json = ldap_conf.get_value(m.Username, true)
	} else {
		date_json = []byte("fail")
	}

	pc.WriteTo(date_json, addr)
}

func (ServerConf) stop_server() error {
	var err error
	return err
}
