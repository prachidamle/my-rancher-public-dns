package service

import (
	"net/http"
	"github.com/rancher/go-rancher/api"
	"github.com/rancher/rancher-public-dns/server"
	"github.com/rancher/rancher-public-dns/model"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"encoding/json"
	"io"
)

func GetRootDomain(w http.ResponseWriter, r *http.Request) {
	//check the auth header. If token is passed, check if it exists in the db. 
	//otherwise generate new token-uuid pair and return.
	log.Debug("Called GetRootDomain")
	apiContext := api.GetApiContext(r)
	authHeader := r.Header.Get("Authorization")
	
	if (authHeader != "") {
		uuid,ok := (server.Authenticate(authHeader)) 
		if ok {
			rootDomain := server.GetRootDomain(uuid)
			apiContext.Write(&rootDomain)	
		} else{
			//failed to authenticate
			log.Debug("Failed to authenticate")
			returnAuthError(w,r)
		}
	} else {
		//generate new token and uuid
		rootDomain := server.RegisterNewClient()
		log.Debug("Called GetRootDomain %v", rootDomain)
		apiContext.Write(&rootDomain)	
	}
}


//GetDNSRecords is a handler for route /dnsrecords and returns the dns records
func GetDNSRecords(w http.ResponseWriter, r *http.Request) {
	apiContext := api.GetApiContext(r)
	dnsRecords, err := server.ListDNSRecords()
	resp := model.DnsRecordCollection{}
	if err == nil {
		for _, value := range dnsRecords {
			fmt.Printf("dnsRecord %v", value)
			resp.Data = append(resp.Data, value)
		}
	}
	fmt.Printf("resp.Data %v", resp.Data)
	/*w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        panic(err)
    }*/
	apiContext.Write(&resp) //does not work
}

//AddDNSRecords is a handler for route POST /dnsrecords and adds the dns record
func AddDNSRecord(w http.ResponseWriter, r *http.Request) {
	//apiContext := api.GetApiContext(r)
	var rec model.DnsRecord
	dec := json.NewDecoder(r.Body)
    for {
        if err := dec.Decode(&rec); err == io.EOF {
            break
        } else if err != nil {
            log.Error(err)
            returnBadRequestError(w,r)
            return
        }
    }
    
    err := server.AddDNSRecord(rec)
    
    if err != nil {
	    returnBadRequestError(w,r)
        return
    }

}

//UpdateDNSRecord is a handler for route PUT /dnsrecords/{recordId} and updates the dns record
func UpdateDNSRecord(w http.ResponseWriter, r *http.Request) {
	//apiContext := api.GetApiContext(r)
	var rec model.DnsRecord
	dec := json.NewDecoder(r.Body)
    for {
        if err := dec.Decode(&rec); err == io.EOF {
            break
        } else if err != nil {
            log.Error(err)
            returnBadRequestError(w,r)
            return
        }
    }
    
    err := server.UpdateDNSRecord(rec)
    
    if err != nil {
	    returnBadRequestError(w,r)
        return
    }

}

//RemoveDNSRecord is a handler for route DELETE /dnsrecords/{recordId} and removes the dns record
func RemoveDNSRecord(w http.ResponseWriter, r *http.Request) {
	//apiContext := api.GetApiContext(r)
	var rec model.DnsRecord
	dec := json.NewDecoder(r.Body)
    for {
        if err := dec.Decode(&rec); err == io.EOF {
            break
        } else if err != nil {
            log.Error(err)
            returnBadRequestError(w,r)
            return
        }
    }
    
    err := server.RemoveDNSRecord(rec)
    
    if err != nil {
	    returnBadRequestError(w,r)
        return
    }
}

func returnBadRequestError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request, Please check the request content"))
}

func returnAuthError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid Token"))
}