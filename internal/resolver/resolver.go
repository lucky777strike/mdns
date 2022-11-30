package resolver

import (
	"fmt"
	"math/rand"
	"mdns/internal/dnslog"
	"time"

	"github.com/miekg/dns"
)

type server struct {
	addr string
	log  *dnslog.Dnslog
}

type Resolver struct {
	servers []server
}

func (res *Resolver) Exchange(m *dns.Msg) (r *dns.Msg, rtt time.Duration, err error) { //TODO status handling
	res.PrintStat()
	c := &dns.Client{Net: "udp"}
	server := res.pickRandomServer()
	resp, rtt, err := c.Exchange(m, server.addr)
	server.log.Allreq.Inc()
	if err != nil {
		server.log.Err.Inc()
		return nil, rtt, err
	}
	fmt.Println(resp.MsgHdr.Rcode)
	return resp, rtt, err
}

func (res *Resolver) pickRandomServer() *server {
	return &res.servers[rand.Intn(len(res.servers))]
}

func (res *Resolver) PrintStat() {
	for _, i := range res.servers {
		fmt.Println(i.addr, i.log)
	}
}

func New(addrs []string) *Resolver {
	rand.Seed(time.Now().Unix())
	servers := make([]server, 0, len(addrs))
	for _, i := range addrs {
		log := new(dnslog.Dnslog)
		servers = append(servers, server{addr: i, log: log})
	}
	return &Resolver{servers: servers}
}

func HelloWorld() {
	fmt.Println("Hello world")
}
