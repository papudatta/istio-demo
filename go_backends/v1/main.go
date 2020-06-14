package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/ammario/ipisp"
)

func main() {
	http.HandleFunc("/", reqhandler)
	// "healthz" handler for kubernetes liveliness and readiness probes
	http.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
}

// technical debt: error handling
// "reqhandler" uses whois service to fetch the details
// in one of the struct types.
// We parse the details and only return asn and prefix
func reqhandler(w http.ResponseWriter, r *http.Request) {
	ipaddress := r.URL.Query().Get("ipaddress")
	fmt.Println(ipaddress)
	client, _ := ipisp.NewDNSClient()
	defer client.Close()
	res, _ := client.LookupIP(net.ParseIP(ipaddress))
	asn := res.ASN.String()
	prefix := res.Range.String()
	resp, _ := json.Marshal(map[string]string{
		"asn": asn, "prefix": prefix})
	w.Write(resp)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "i am cool")
}
