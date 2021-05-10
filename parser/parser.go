package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"srm-dhcp/config"
	"srm-dhcp/kafkaproducer"
	"strconv"
	"strings"
	"time"
)

func FileLoop() {
	for {
		delaystr, _ := strconv.Atoi(config.Delay)
		delay := time.Duration(delaystr)
		FileReader()
		time.Sleep(delay * time.Second)
	}
}
func FileReader() {
	file, err := os.Open("/tmp/" + config.File)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Could file DHCP log file.... retrying...")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		DhcpLine := &config.Dhcp{
			Id:                    s[0],
			Date:                  s[1],
			Time:                  s[2],
			Description:           s[3],
			Ipaddress:             s[4],
			HostName:              s[5],
			MacAddress:            s[6],
			UserName:              s[7],
			TransactionId:         s[8],
			Qresualt:              s[9],
			ProbationTime:         s[10],
			CorrelationId:         s[11],
			DhcId:                 s[12],
			VendorClassHex:        s[13],
			VendorClassAscii:      s[14],
			UserClassHex:          s[15],
			UserClassAscii:        s[16],
			RelayAgentInformation: s[17],
			DnsRegError:           s[18],
		}
		fmt.Println(DhcpLine)
		kafkaproducer.ProducerHandler(DhcpLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
