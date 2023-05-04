package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/didip/tollbooth/v7"
)

func IpLook(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	log.Printf(" INFO - from IP: %s", ip)
	response := map[string]string{"ip": ip}
	jsonResp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	http.Handle("/ip", tollbooth.LimitFuncHandler(lmt, IpLook))
	http.ListenAndServe(":8080", nil)

	// log.Fatal(http.ListenAndServe(":8080", throttledHandler))
}
