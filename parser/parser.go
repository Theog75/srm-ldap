package parser

import (
	"fmt"
	"os/exec"
	"regexp"
	"srm-ldap/config"
	"strconv"
	"time"
)

func LdapLoop() {
	for {
		delaystr, _ := strconv.Atoi(config.Delay)
		delay := time.Duration(delaystr)
		ldapReader()
		time.Sleep(delay * time.Second)
	}
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

func ldapReader() {
	cmd := "/ldap-query.sh"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Println("Could not read ldap content ... retrying...")
		return
	}
	fmt.Println(string(out))
	// kafkaproducer.ProducerHandler(out)
}
