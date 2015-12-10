package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rancher/external-dns/dns"
	"github.com/rancher/external-dns/providers"
	"github.com/rancher/rancher-public-dns/model"
	//"fmt"
)

const (
	name = "Route53"
)

var (
	provider     *providers.Route53Handler
)

func init() {
	log.Debug("Initializing Rancher route53")
	provider = &providers.Route53Handler{}
}

func GetName() string {
	return name
}

type Route53Connector struct {
}

func (r *Route53Connector) AddRecord(record model.DnsRecord) error {
	awsRecord := dns.DnsRecord{}
	awsRecord.Fqdn = record.Fqdn
	awsRecord.Records = record.Records
	awsRecord.TTL = record.TTL
	awsRecord.Type = record.Type
	
	return provider.AddRecord(awsRecord)
}

func (r *Route53Connector) UpdateRecord(record model.DnsRecord) error {
	awsRecord := dns.DnsRecord{}
	awsRecord.Fqdn = record.Fqdn
	awsRecord.Records = record.Records
	awsRecord.TTL = record.TTL
	awsRecord.Type = record.Type
	
	return provider.UpdateRecord(awsRecord)
}

func (r *Route53Connector) RemoveRecord(record model.DnsRecord) error {
	awsRecord := dns.DnsRecord{}
	awsRecord.Fqdn = record.Fqdn
	awsRecord.Records = record.Records
	awsRecord.TTL = record.TTL
	awsRecord.Type = record.Type
	
	return provider.RemoveRecord(awsRecord)
}

func (r *Route53Connector) GetRecords() ([]model.DnsRecord, error) {
	var records []model.DnsRecord
	awsRecords, err := provider.GetRecords()
	
	log.Debugf("awsRecords %v", awsRecords)
	
	if err != nil {
		return records, err
	}
	
	for _, rec := range awsRecords {
		log.Debugf("rec %v", rec)
		record := model.DnsRecord{Fqdn: rec.Fqdn, Records: rec.Records, Type: rec.Type, TTL: rec.TTL}
		records = append(records, record)
	}
	log.Debugf("records %v", records)
	return records, nil
}