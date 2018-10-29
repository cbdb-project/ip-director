package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oschwald/geoip2-golang"
)

// convert ip address to country name
func queryIPToCountry(ipAddr string) string {
	//db, err := geoip2.Open("GeoIP2-City.mmdb")
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	// ip := net.ParseIP("45.32.47.180")
	ip := net.ParseIP(ipAddr)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(record)
	//fmt.Printf("%+v\n", record.Country.Names["en"])
	return record.Country.Names["en"]
}

func readConfig(fileName string) []string {
	var output []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func countryToURL(countryName string) string {
	var output string
	directURLs := readConfig("directInfo.txt")
	if countryName == "China" {
		output = directURLs[len(directURLs)-1]
	} else {
		output = directURLs[0]
	}
	// fmt.Println(directURLs)
	return output
}

func ipToCountry(w http.ResponseWriter, r *http.Request) {
	var countryName, clinetIP, output string
	//for local
	//clinetIP, _, _ = net.SplitHostPort(r.RemoteAddr)
	//for remote nginx
	clinetIP = r.Header.Get("X-real-ip")
	//fmt.Println(clinetIP)
	ipAddr := clinetIP
	countryName = queryIPToCountry(ipAddr)
	output = countryToURL(countryName)

	// fmt.Println(ipAddr)
	// fmt.Println(countryName)
	json.NewEncoder(w).Encode(output)
}

func ipToCountrySubmitAddr(w http.ResponseWriter, r *http.Request) {
	var countryName, output string
	vars := mux.Vars(r)
	ipAddr := vars["ipAddr"]
	countryName = queryIPToCountry(ipAddr)
	output = countryToURL(countryName)
	json.NewEncoder(w).Encode(output)
}

//test:
//127.0.0.1:8012/ipToCountry
//127.0.0.1:8012/ipToCountrySubmitAddr/45.32.47.180
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ipToCountry)
	r.HandleFunc("/ipToCountry", ipToCountry)
	r.HandleFunc("/ipToCountrySubmitAddr/{ipAddr}", ipToCountrySubmitAddr)
	log.Fatal(http.ListenAndServe(":8012", r))
	return

}
