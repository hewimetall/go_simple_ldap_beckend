package main

import (
	"fmt"

	auth "github.com/korylprince/go-ad-auth"
	gonfig "github.com/tkanos/gonfig"
)

type Configuration struct {
	Server   string
	Port     int
	BaseDN   string
	BindDN   string
	PasswdDN string
	Debug    bool
}

// set State output log
var Debug = true
var ldap_conf = ldap_suppor{}

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
	// config =  = &auth.Config{
	// 	Server:   "192.38.1.200",
	// 	Port:     389,
	// 	BaseDN:   "DC=agp,DC=ru",
	// 	Security: security,
	// }

	return err
}

func init() {
	set_config("./conf.json", auth.SecurityNone)
	ldap_conf.conn_init()
	// ldap_conf.get_value("KodolovaE@agp.ru", true)

}

// var attr []string{"sn",}
func main() {
	fmt.Println("Start")
}
