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

/*
	DO
		-0 get short atr for user
		-1 get full atr for user
		-2 get cn group for user
		-3 get full group
*/
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

func cont(buf []byte) []byte {
	/*
		DO
			-0 get short atr for user
			-1 get full atr for user
			-2 get cn group for user
			-3 get full group
	*/
	var m User
	var date_json []byte
	err := json.Unmarshal(buf, &m)
	if err != nil {
		log.Println("stream", err)
		date_json = []byte("{}")
		return date_json
	}
	if err := ldap_conf.test_autch(m.Username, m.Password); err != 0 && m.Do != "2" {
		date_json = []byte("{}")
		return date_json
	}
	switch os := m.Do; os {
	case "0":
		date_json = ldap_conf.get_value(m.Username, true)
	case "1":
		date_json = ldap_conf.get_value(m.Username, true)
	case "2":
		date_json = ldap_conf.get_group_dn(m.Username, true)
	case "3":
		date_json = ldap_conf.get_group_all(true)
		write_dump("get_group_all", date_json)
		date_json = []byte("Ok")
	default:
		date_json = []byte("{}")
	}
	return date_json
}

func stream(pc net.PacketConn, addr net.Addr, buf []byte) {
	var date_json []byte
	date_json = cont(buf)
	pc.WriteTo(date_json, addr)
}

func (ServerConf) stop_server() error {
	var err error
	return err
}
