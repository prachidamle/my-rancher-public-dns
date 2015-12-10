package model

import "github.com/rancher/go-rancher/client"


//DnsRecord structure to respond on rest api
type DnsRecord struct {
	client.Resource
	Fqdn    string		`json:"fqdn"`
	Records []string	`json:"records"`
	Type    string		`json:"type"`
	TTL     int			`json:"ttl"`
	uuid	string 		`json:"uuid"`
}

//DnsRecordCollection holds a collection of templates
type DnsRecordCollection struct {
	client.Collection
	Data []DnsRecord `json:"data,omitempty"`
}