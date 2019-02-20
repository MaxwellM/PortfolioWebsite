package visitorCounter

import "time"

type knownIP struct {
	IP        string    `json:"ip"`
	TimeStamp time.Time `json:"timeStamp"`
}

var knownIPs []knownIP

var IpStruct []string

func AddIPToStruct(ip string)(*[]knownIP, error){

	IpStruct = append(IpStruct, ip)

	knownIPs = append(knownIPs, knownIP{IP: ip, TimeStamp: time.Now()})

	return &knownIPs, nil
}
