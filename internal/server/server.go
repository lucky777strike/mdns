package server

import (
	"log"
	"mdns/internal/resolver"

	"github.com/miekg/dns"
)

type server struct {
	addr     string
	resolver *resolver.Resolver
}

func (s *server) HandleFunc(w dns.ResponseWriter, r *dns.Msg) {
	resp, _, err := s.resolver.Exchange(r)
	if err != nil {
		dns.HandleFailed(w, r)
		return
	}
	w.WriteMsg(resp)
}

func (s *server) Run() {
	udpServer := &dns.Server{Addr: s.addr, Net: "udp"}
	tcpServer := &dns.Server{Addr: s.addr, Net: "tcp"}
	dns.HandleFunc(".", s.HandleFunc)
	go func() {
		if err := udpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		s.resolver.StartServer()
	}()
}

func New(addr string, resolver *resolver.Resolver) *server {
	return &server{addr: addr, resolver: resolver}
}
