package config

import (
	"os"
	"strings"
)

const (
	EnvPort       = "PORT"
	EnvDelay      = "DELAY"
	EnvkafkaURL   = "KAFKA_HOSTS"
	Envkafkatopic = "KAFKATOPIC"

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

	LDAPBINDDN = GetEnv(EnvLDAPBINDDN, "")
	LDAPBINDPASSWORD = GetEnv(EnvLDAPBINDPASSWORD, "")
	LDAPBASEDN = GetEnv(EnvLDAPBASEDN, "")
	LDAPHOST = GetEnv(EnvLDAPHOST, "")
	LDAPOBJCLASS = GetEnv(EnvLDAPOBJCLASS, "")

}

func GetEnv(key string, def string) string {

	result := os.Getenv(key)
	if result == "" {
		return def
	}

	return strings.TrimSpace(result)
}
