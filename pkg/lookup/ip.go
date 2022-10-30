package lookup

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Lookup interface {
	GetIp()
	GetAll()
}

type ipLookup struct {
	writer http.ResponseWriter
	cache *LocalCache
}

// Create a new IP lookup
func NewLookup(w http.ResponseWriter, c *LocalCache) *ipLookup {
	return &ipLookup{
		writer: w,
		cache: c,
	}
}

// Check cache for IP address. If not present request from GeoJS
func (il *ipLookup) GetIp(ip string) {
	cip, err := il.cache.ReadCache(ip)

	if err == errIpNotInCache {
		fmt.Printf("%v, fetching IP now \n", err)
		fetchIpInfo(ip, il)
	} else if err != nil {
		fmt.Printf("%v \n", err)
	} else {
		fmt.Println("IP was cached!")
		json.NewEncoder(il.writer).Encode(&cip)
	}
}

type IpLocation map[string]string

// Request filtered results from cache and return if present
func (il *ipLookup) GetFiltered(loc IpLocation) {
	ips := il.cache.Filter(loc)

	if len(ips) == 0 {
		io.WriteString(il.writer, "Sorry, no results for the specified city and/or country found in the cache!")
		return 
	}

	json.NewEncoder(il.writer).Encode(&ips)
}

// Request all values from cache and return if present
func (il *ipLookup) GetAll() {
	ips := il.cache.ReadAll()

	if len(ips) == 0 {
		io.WriteString(il.writer, "No Ip's have been cached, search for one!")
		return 
	}

	json.NewEncoder(il.writer).Encode(&ips)
}

type IpData struct {
	Ip string `json:"ip"`
	City string `json:"city"`
	Country string `json:"country"`
}

// Fetch IP info from GeoJS
func fetchIpInfo(ip string, il *ipLookup) {
	url := fmt.Sprintf("https://get.geojs.io/v1/ip/geo/%s.json", ip)
	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		io.WriteString(il.writer, "404 returned from GeoJs- IP most likely invalid or service is down")
		return
	}

	if err != nil {
		log.Fatalln(err)
	}

	var parRes IpData
	err = json.NewDecoder(resp.Body).Decode(&parRes)
	
	if err != nil {
		http.Error(il.writer, err.Error(), http.StatusBadRequest)
		return
	}

	il.cache.UpdateCache(parRes)
	json.NewEncoder(il.writer).Encode(&parRes)
}