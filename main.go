package main

import (
	"fmt"
	"io/ioutil"
	"log"

	auth "github.com/korylprince/go-ad-auth"
	gonfig "github.com/tkanos/gonfig"
)

type Configuration struct {
	Server     string
	Port       int
	BaseDN     string
	BindDN     string
	PasswdDN   string
	Debug      bool
	Listenserv string
	Listenport int
}

// set State output log
var Debug = true
var ldap_conf = ldap_suppor{}
var server_conf = ServerConf{}

func set_config(path string, security auth.SecurityType) error {
	configuration := Configuration{}
	err := gonfig.GetConf(path, &configuration)
	Debug = configuration.Debug
	ldap_conf.config = auth.Config{
		Server:   configuration.Server,
		Port:     configuration.Port,
		BaseDN:   configuration.BaseDN,
		Security: security,
	}
	ldap_conf.Adm.BindDN = configuration.BindDN
	ldap_conf.Adm.PasswordDN = configuration.PasswdDN
	server_conf.host = configuration.Listenserv
	server_conf.port = configuration.Listenport
	return err
}
func write_dump(file_name string, d1 []byte) error {

	err := ioutil.WriteFile("./dump/"+file_name, d1, 0644)
	if err != nil {
		log.Panic(err)
	}
	return err
}
func init() {
	set_config("./conf.json", auth.SecurityNone)
	ldap_conf.conn_init()
}

// var attr []string{"sn",}
func main() {
	fmt.Println("Start")
	server_conf.init_server()
}
