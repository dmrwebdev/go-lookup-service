package lookup

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetIp(ip string, w http.ResponseWriter) {

	url := fmt.Sprintf("https://get.geojs.io/v1/ip/geo/%s.json", ip)
	fmt.Println(url)
	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		io.WriteString(w, "404 returned from get request- IP most likely invalid or GeoJS is down")
	}

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}