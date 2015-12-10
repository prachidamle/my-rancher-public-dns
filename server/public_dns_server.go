package server

import (
	"github.com/rancher/rancher-public-dns/model"
	"github.com/rancher/rancher-public-dns/util"
	"strings"
	b64 "encoding/base64"
	log "github.com/Sirupsen/logrus"
	//"fmt"
)

var (
	route53Connector     *Route53Connector
)

const baseDomain string = "rancher.io"

func init() {
	route53Connector = &Route53Connector{}
}

func Authenticate(authHeader string) (string, bool) {
	log.Debug("Called Authenticate")
	if (authHeader == "") {
			return "", false
	}
	// header value format will be "Basic encodedstring" for Basic
	// authentication. Example "Basic YWRtaW46YWRtaW4="
	encodedToken := strings.Replace(authHeader, "Basic", "", 1)
	uDec, _ := b64.URLEncoding.DecodeString(encodedToken)
	token := string(uDec)
	
	
	uuid := getUUID(token)
	log.Debug("uuid: %v", uuid)
	if uuid != "" {
		return uuid, true
	} else {
		return "", false
	}
	
}

func GenerateNewToken() string{
	return util.GenerateUUID()
}

func RegisterNewClient() model.RootDomainInfo {
	domainInfo := model.RootDomainInfo{}
	
	uuid := util.GenerateUUID()
	domainInfo.RootDomain = GetRootDomain(uuid)
	
	token := util.GenerateNewToken()
	domainInfo.Token = token
	
	insertUUIDTokenRecord(token, uuid)
	
	return domainInfo
}

func GetRootDomain(uuid string) string {
	return uuid + "." + baseDomain
}

func ListDNSRecords() ([]model.DnsRecord, error){
	//call connector to listRecords
	log.Debug("Called ListDNSRecords")
	return route53Connector.GetRecords()
}

func AddDNSRecord(record model.DnsRecord) error{
	//call connector to listRecords
	log.Debug("Called AddDNSRecord")
	return route53Connector.AddRecord(record)
}

func UpdateDNSRecord(record model.DnsRecord) error{
	//call connector to listRecords
	log.Debug("Called UpdateDNSRecord")
	return route53Connector.UpdateRecord(record)
}

func RemoveDNSRecord(record model.DnsRecord) error{
	//call connector to listRecords
	log.Debug("Called RemoveDNSRecord")
	return route53Connector.RemoveRecord(record)
}