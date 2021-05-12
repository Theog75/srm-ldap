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
	cmd := "/ldap-query.sh"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Println("Could not read ldap content ... retrying...")
		return
	}
	fmt.Printf("%s",out)
	// kafkaproducer.ProducerHandler(out)
}
