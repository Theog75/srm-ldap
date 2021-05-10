package config

import (
	"os"
	"strings"
)

const (
	EnvPort       = "PORT"
	EnvFile       = "LOGFILE"
	EnvDelay      = "DELAY"
	EnvkafkaURL   = "KAFKAURL"
	Envkafkatopic = "KAFKATOPIC"

	EnvLDAPBINDDN       = "LDAPBINDDN"
	EnvLDAPBINDPASSWORD = "LDAPBINDPASSWORD"
	EnvLDAPBASEDN       = "LDAPBASEDN"
	EnvLDAPHOST         = "LDAPHOST"
)

var (
	Port       = ""
	File       = ""
	Delay      = ""
	KafkaURL   = ""
	KafkaTopic = ""

	LDAPBINDDN       = ""
	LDAPBINDPASSWORD = ""
	LDAPBASEDN       = ""
	LDAPHOST         = ""
)

func ReadEnv() {
	Port = GetEnv(EnvPort, "8080")
	File = GetEnv(EnvFile, "DhcpSrvLog-Mon.log")
	Delay = GetEnv(EnvDelay, "10")
	KafkaURL = GetEnv(EnvkafkaURL, "kafka-0.kafka")
	KafkaTopic = GetEnv(Envkafkatopic, "collector.dhcp")

	LDAPBINDDN = GetEnv(EnvLDAPBINDDN, "")
	LDAPBINDPASSWORD = GetEnv(EnvLDAPBINDPASSWORD, "")
	LDAPBASEDN = GetEnv(EnvLDAPBASEDN, "")
	LDAPHOST = GetEnv(EnvLDAPHOST, "")

}

func GetEnv(key string, def string) string {

	result := os.Getenv(key)
	if result == "" {
		return def
	}

	return strings.TrimSpace(result)
}
