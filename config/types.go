package config

type Dhcp struct {
	Id                    string `json:id,omitempty`
	Date                  string `json:date,omitempty`
	Time                  string `json:time,omitempty`
	Description           string `json:description,omitempty`
	Ipaddress             string `json:address,omitempty`
	HostName              string `json:hostName,omitempty`
	MacAddress            string `json:macAddress,omitempty`
	UserName              string `json:userName,omitempty`
	TransactionId         string `json:transactionId,omitempty`
	Qresualt              string `json:qresualt,omitempty`
	ProbationTime         string `json:probationTime,omitempty`
	CorrelationId         string `json:correlationId,omitempty`
	DhcId                 string `json:dhcId,omitempty`
	VendorClassHex        string `json:vendorClassHex,omitempty`
	VendorClassAscii      string `json:vendorClassAscii,omitempty`
	UserClassHex          string `json:userClassHex,omitempty`
	UserClassAscii        string `json:userClassAscii,omitempty`
	RelayAgentInformation string `json:relayAgentInformation,omitempty`
	DnsRegError           string `json:dnsRegError,omitempty`
}
