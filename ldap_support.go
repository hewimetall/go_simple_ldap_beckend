package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	auth "github.com/korylprince/go-ad-auth"
)

/*
	File resolve problems:
		-autch
			-simple
			-test +
		-search +
*/
// Set type conf ldap
var Config auth.Config

type AdminDN struct {
	BindDN     string
	PasswordDN string
}

type ldap_suppor struct {
	config auth.Config
	Adm    AdminDN
	conn   *auth.Conn
	flag_t bool
}

func (ld *ldap_suppor) test_conn() {
	ld.flag_t = true
	for ld.flag_t {
		time.Sleep(time.Minute * 3)
		_, err := ld.conn.Search(fmt.Sprintf("(userPrincipalName=%s)", ld.Adm.BindDN), []string{"sAMAccountName"}, 0)
		if err != nil {
			log.Println("Valid search: Expected err to be nil but got:", err)
			ld.conn_init()
		}
		log.Println("Work")
	}
}

func (ld *ldap_suppor) test_autch(username string, password string) int {
	status, err := auth.Authenticate(&ld.config, username, password)

	if err != nil {
		//handle err
		fmt.Println("handle err")
		return 1
	}

	if !status {
		//handle failed authentication
		fmt.Println("failed authentication")
		return 2
	}
	return 0
}

func (ld *ldap_suppor) search(upn string, short bool) map[string]string {
	m := make(map[string]string)
	atr_full := []string{
		"sAMAccountName", "givenName", "cn", "initials", "displayName", "memberOf", "department", "mail", "telephoneNumber", "description",
	}
	atr_shor := []string{
		"sAMAccountName", "givenName", "cn", "initials", "displayName", "memberOf", "department",
	}
	var atr []string
	if short == true {
		atr = atr_shor
	} else {
		atr = atr_full
	}

	ent, err := ld.conn.Search(fmt.Sprintf("(userPrincipalName=%s)", upn), atr, 0)
	if err != nil {
		log.Panic("Valid search: Expected err to be nil but got:", err)
		ld.conn_init()
		ld.search(upn, short)
		// return err
	}

	for _, entry := range ent {
		for _, attr := range entry.Attributes {
			m[attr.Name] = strings.Join(attr.Values, "")
		}

	}
	return m
}

func (ld *ldap_suppor) conn_init() error {
	// init_conn
	// sheduler_init
	// awaiting requests
	// response to requests
	// awaiting requests
	//simple
	config := &ld.config
	// set Bind setting for simple autch
	upn := ld.Adm.BindDN // userPrincipalName
	password := ld.Adm.PasswordDN

	// Connect to ldap server
	conn, err := config.Connect()
	if err != nil {
		log.Panic("Error connecting to server:", err)
		// return err
	}
	// set glob conn
	ld.conn = conn
	// TODO: kastil
	// defer conn.Conn.Close()
	//  Bind to ldap server for search date
	status, err := conn.Bind(upn, password)
	if err != nil {
		log.Panic("Error binding to server:", err)
		return err
	}
	if !status {
		//handle failed authentication
		log.Panic("failed authentication root", status)
		return err
	}

	ent, err := conn.Search(fmt.Sprintf("(userPrincipalName=%s)", upn), []string{"sAMAccountName"}, 0)
	if err != nil {
		log.Panic("Valid search: Expected err to be nil but got:", err)
		return err
	}
	for _, entry := range ent {
		entry.Print()
	}
	go ld.test_conn()
	return err
}
func (ldap_suppor) convers(m map[string]string) []byte {
	jsonString, err := json.Marshal(m)
	if err != nil {
		log.Panic("Not convers map to json:", err)
	}

	return jsonString
}

func (ldap_suppor) convers_str(m []string) []byte {
	jsonString, err := json.Marshal(m)
	if err != nil {
		log.Panic("Not convers map to json:", err)
	}

	return jsonString
}
func (l *ldap_suppor) get_value(username string, short bool) []byte {
	upn, _ := l.config.UPN(username)
	return l.convers(l.search(upn, short))
}

func (l *ldap_suppor) get_group_dn(username string, short bool) []byte {
	upn, _ := l.config.UPN(username)
	m := make(map[string]string)
	ent, _ := l.conn.Search(fmt.Sprintf("(userPrincipalName=%s)", upn), []string{"memberof"}, 0)

	for _, entry := range ent {
		for _, attrs := range entry.Attributes {
			if attrs.Name == "memberOf" {
				for _, attr := range attrs.Values {
					cn, err := l.conn.GroupDN(string(attr))
					if err == nil {
						cn = string(cn[3:support_util(cn)])
						m[cn] = string(l.get_group_attr(cn))
					}
				}
			}
		}
	}
	return l.convers(m)
}
func support_util(st string) int {
	var K int
	for k, c := range st {
		if c == ',' {
			K = k
			break
		}
	}
	return K
}

func (l *ldap_suppor) get_group_attr(gr_cn string) string {
	ent, _ := l.conn.Search("(&(CN="+gr_cn+")(objectClass=group))", []string{"description"}, 0)
	for _, entry := range ent {
		for _, attr := range entry.Attributes {
			if attr.Name == "description" {
				return attr.Values[0]
			}
		}
	}
	return ""
}

func (l *ldap_suppor) get_group_all(short bool) []byte {
	m := make(map[string]string)
	ent, _ := l.conn.Search("(&(memberof=*)(objectClass=person))", []string{"memberof"}, 0)

	for _, entry := range ent {
		for _, attrs := range entry.Attributes {
			if attrs.Name == "memberOf" {
				for _, attr := range attrs.Values {
					cn, err := l.conn.GroupDN(string(attr))
					if err == nil {
						cn = string(cn[3:support_util(cn)])
						m[cn] = string(l.get_group_attr(cn))
					}
				}
			}
		}
	}
	return l.convers(m)

}

func (l *ldap_suppor) get_group_user_member_all(username string, short bool) []byte {
	ent, _ := l.conn.Search("(&(member=*)(objectClass=group))", []string{"member"}, 0)

	for _, entry := range ent {
		for _, attrs := range entry.Attributes {
			if attrs.Name == "memberOf" {
				for _, attr := range attrs.Values {
					_, cn := l.conn.GroupDN(string(attr))
					log.Println(string(attr), cn)
				}
			}
		}
	}
	return []byte("{test}")
}
