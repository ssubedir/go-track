package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Services struct {
}

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

var netClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}

func NewServices() *Services {
	return &Services{}
}

func (s *Services) CanadaPost(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["track"]

	now := time.Now()
	fmt.Println("now:", now.Format("2006/01/02"))
	then := now.AddDate(0, -2, 0)
	fmt.Println("then:", then.Format("2006/01/02"))
	url := "https://www.canadapost.ca/trackweb/rs/track/json/package?refNbrs=" + id + "&mailedFromDate=" + then.Format("2006/01/02") + "&mailedToDate=" + now.Format("2006/01/02")
	fmt.Println(url)
	// Setup
	body := strings.NewReader("")
	req, err := http.NewRequest("GET", url, body)

	if err != nil {
		log.Println(" Error creating http request to canada post, make sure request data is valid")
		return
	}

	// Send Request
	response, err := netClient.Do(req)
	if err != nil {

		if strings.Contains(err.Error(), "no such host") {
			log.Println(" - No Such host, Check Connection / DNS")
		} else if strings.Contains(err.Error(), "timeout") {
			log.Println(" - Http timed out")
		} else {
			log.Println(" - Error Sending http request")
		}
		return
	}

	// close body
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func (s *Services) FedEx(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["track"]
	var d = `data={"TrackPackagesRequest":{"trackingInfoList":[{"trackNumberInfo":{"trackingNumber":"##########","trackingQualifier":"","trackingCarrier":""}}]}}&action=trackpackages&locale=en_US&version=1&format=json`
	// Setup
	body := strings.NewReader(strings.Replace(d, "##########", id, -1))

	req, err := http.NewRequest("POST", "https://www.fedex.com/trackingCal/track", body)

	if err != nil {
		log.Println(" Error creating http request to canada post, make sure request data is valid")
		return
	}

	//headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// Send Request
	response, err := netClient.Do(req)
	if err != nil {

		if strings.Contains(err.Error(), "no such host") {
			log.Println(" - No Such host, Check Connection / DNS")
		} else if strings.Contains(err.Error(), "timeout") {
			log.Println(" - Http timed out")
		} else {
			log.Println(" - Error Sending http request")
		}
		return
	}

	// close body
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
