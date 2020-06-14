package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/ammario/ipisp"
)

func main() {
	http.HandleFunc("/", Reqhandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
}

func Reqhandler(w http.ResponseWriter, r *http.Request) {
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
