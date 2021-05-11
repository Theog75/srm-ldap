package parser

import (
	"fmt"
	"os/exec"
	"srm-ldap/config"
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
	fmt.Println("Running command /usr/bin/ldapsearch  -h " + config.LDAPHOST + " -b " + config.LDAPBASEDN + "  -D '" + config.LDAPBINDDN + "' -w " + config.LDAPBINDPASSWORD + " s sub '(objectClass=" + config.LDAPOBJCLASS + ")'")
	cmd := "/usr/bin/ldapsearch"
	//args := " -h " + config.LDAPHOST + " -b " + config.LDAPBASEDN + "  -D '" + config.LDAPBINDDN + "' -w " + config.LDAPBINDPASSWORD + " s sub '(objectClass=" + config.LDAPOBJCLASS + ")'"
	args := []string{" -h " + config.LDAPHOST, "-b " + config.LDAPBASEDN , " -D '" + config.LDAPBINDDN + "'", "-w " + config.LDAPBINDPASSWORD, " s sub '(objectClass=" + config.LDAPOBJCLASS + ")'"}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Println("Could not read ldap content ... retrying...")
		return
	}
	fmt.Println(out)
	// kafkaproducer.ProducerHandler(out)
}
