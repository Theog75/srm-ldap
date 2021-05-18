package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EnvPort       = "PORT"
	EnvDelay      = "DELAY"
	EnvkafkaURL   = "KAFKA_HOSTS"
	Envkafkatopic = "KAFKATOPIC"

	EnvNUMOFLINESINMSG  = "NUMOFLINESINMSG"
	EnvLDAPBINDDN       = "LDAPBINDDN"
	EnvLDAPBINDPASSWORD = "LDAPBINDPASSWORD"
	EnvLDAPBASEDN       = "LDAPBASEDN"
	EnvLDAPHOST         = "LDAPHOST"
	EnvLDAPOBJCLASS     = "OBJCLASS"
)

var (
	Port       = ""
	Delay      = ""
	KafkaURL   = ""
	KafkaTopic = ""

	NUMOFLINESINMSG  = 0
	LDAPBINDDN       = ""
	LDAPBINDPASSWORD = ""
	LDAPBASEDN       = ""
	LDAPHOST         = ""
	LDAPOBJCLASS     = ""
)

func ReadEnv() {
	Port = GetEnv(EnvPort, "8080")
	Delay = GetEnv(EnvDelay, "10")
	KafkaURL = GetEnv(EnvkafkaURL, "kafka-0.kafka")
	KafkaTopic = "collector." + GetEnv(Envkafkatopic, "")

	NUMOFLINESINMSG, _ = strconv.Atoi(GetEnv(EnvNUMOFLINESINMSG, "500"))
	LDAPBINDDN = GetEnv(EnvLDAPBINDDN, "")
	LDAPBINDPASSWORD = GetEnv(EnvLDAPBINDPASSWORD, "")
	LDAPBASEDN = GetEnv(EnvLDAPBASEDN, "")
	LDAPHOST = GetEnv(EnvLDAPHOST, "")
	LDAPOBJCLASS = GetEnv(EnvLDAPOBJCLASS, "")

}

func GetEnv(key string, def string) string {

	result := os.Getenv(key)
	fmt.Println(key)
	fmt.Println(result)
	if result == "" {
		return def
	}

	return strings.TrimSpace(result)
}
