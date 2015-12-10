package service

import (
	"net/http"
	//log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/rancher/go-rancher/api"
	"github.com/rancher/go-rancher/client"
	"github.com/rancher/rancher-public-dns/model"
)


//Route defines the properties of a go mux http route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes array of Route defined
type Routes []Route

//NewRouter creates and configures a mux router
func NewRouter() *mux.Router {
	schemas := &client.Schemas{}

	// ApiVersion
	apiVersion := schemas.AddType("apiVersion", client.Resource{})
	apiVersion.CollectionMethods = []string{}

	// Schema
	schemas.AddType("schema", client.Schema{})
	
	// rootDomainInfo
	rootDomainInfo := schemas.AddType("rootDomainInfo", model.RootDomainInfo{})
	rootDomainInfo.CollectionMethods = []string{}

	// DnsRecord
	dnsRecord := schemas.AddType("dnsRecord", model.DnsRecord{})
	dnsRecord.CollectionMethods = []string{}

	// API framework routes
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Handler(api.VersionsHandler(schemas, "v1-rancher-dns"))
	router.Methods("GET").Path("/v1-rancher-dns/schemas").Handler(api.SchemasHandler(schemas))
	router.Methods("GET").Path("/v1-rancher-dns/schemas/{id}").Handler(api.SchemaHandler(schemas))
	router.Methods("GET").Path("/v1-rancher-dns").Handler(api.VersionHandler(schemas, "v1-rancher-dns"))

	// Application routes
	
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(api.ApiHandler(schemas, route.HandlerFunc))
	}

	return router
}

var routes = Routes{
	Route{
		"GetRootDomain",
		"GET",
		"/v1-rancher-dns/rootDomainInfo",
		GetRootDomain,
	},
	Route{
		"GetDNSRecords",
		"GET",
		"/v1-rancher-dns/dnsRecords",
		GetDNSRecords,
	},
	Route{
		"AddDNSRecord",
		"POST",
		"/v1-rancher-dns/dnsRecords",
		AddDNSRecord,
	},
	Route{
		"UpdateDNSRecord",
		"PUT",
		"/v1-rancher-dns/dnsRecords",
		UpdateDNSRecord,
	},
	Route{
		"RemoveDNSRecord",
		"DELETE",
		"/v1-rancher-dns/dnsRecords",
		RemoveDNSRecord,
	},		
}

