package parser

import (
	"fmt"
	"os/exec"
	"srm-dhcp/config"
	"strconv"
	"time"
)

func FileLoop() {
	for {
		delaystr, _ := strconv.Atoi(config.Delay)
		delay := time.Duration(delaystr)
		LDAPReader()
		time.Sleep(delay * time.Second)
	}
}
func LDAPReader() {
	fmt.Println("Running command /usr/bin/ldapsearch  -h " + config.LDAPHOST + " -b " + config.LDAPBASEDN + "  -D '" + config.LDAPBINDDN + "' -w " + config.LDAPBINDPASSWORD + " s sub '(objectClass=computer)'")
	out, err := exec.Command("/usr/bin/ldapsearch  -h " + config.LDAPHOST + " -b " + config.LDAPBASEDN + "  -D '" + config.LDAPBINDDN + "' -w " + config.LDAPBINDPASSWORD + " s sub '(objectClass=computer)'").Output()

	if err != nil {
		// log.Fatal(err)
		fmt.Println("Could not read ldap content ... retrying...")
		return
	}
	fmt.Println(out)
	// kafkaproducer.ProducerHandler(out)
}
