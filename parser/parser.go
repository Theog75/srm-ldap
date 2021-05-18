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
	delaystr, _ := strconv.Atoi(config.Delay)
	delay := time.Duration(delaystr)
	for {
		ldapReader()
		time.Sleep(delay * time.Second)
	}
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\s*\n`).
		Split(strNormalized, -1)

}

func reduceSliceSize(slcs []string) []string {
	res := []string{}
	combinedSlices := ""

	for index, slc := range slcs {
		if combinedSlices != "" {
			combinedSlices += "\r\n"
		}
		combinedSlices += slc
		fmt.Printf("%d\n", config.NUMOFLINESINMSG)

		if (index+1)%config.NUMOFLINESINMSG == 0 {
			res = append(res, combinedSlices)
			combinedSlices = ""
		}
	}

	if combinedSlices != "" {
		res = append(res, combinedSlices)
	}

	return res
}

func ldapReader() {
	cmd := "/ldap-query.sh"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Could not read ldap content ... retrying...")
		return
	}
	for _, msg := range reduceSliceSize(splitByEmptyNewline(string(out))) {
		kafkaproducer.ProducerHandler(msg)

	}
}
