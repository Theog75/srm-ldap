package parser

import (
	"fmt"
	"os/exec"
	"regexp"
	"srm-ldap/config"
	"srm-ldap/kafkaproducer"
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

func reduceSliceSize(slcs []string) []string {
	res := []string{}
	combinedSlices := ""
	for i, slc := range slcs {
		if combinedSlices == "" {
			combinedSlices += slc
		} else {
			combinedSlices += "\r\n\r\n" + slc
		}

		if i%1000 == 0 {
			res = append(res, combinedSlices)
			combinedSlices = ""
		}
	}

	return res
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
	for _, s := range reduceSliceSize(splitByEmptyNewline(string(out))) {
		kafkaproducer.ProducerHandler(s)

	}
}
