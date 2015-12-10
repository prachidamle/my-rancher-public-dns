package model

import "github.com/rancher/go-rancher/client"

//RootDomainInfo structure contains the new version info
type RootDomainInfo struct {
	client.Resource
	RootDomain  string            `json:"rootDomain"`
	Token string 				   `json:"token"`
}
