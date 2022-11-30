package resolver

import (
	"fmt"
	"math/rand"
	"mdns/internal/dnslog"
	"net/http"
	"strings"
	"time"

	"github.com/miekg/dns"
)

type server struct {
	addr string
	log  *dnslog.Dnslog
}

type Resolver struct {
	servers []server
	logport string
}

func (res *Resolver) Exchange(m *dns.Msg) (r *dns.Msg, rtt time.Duration, err error) { //TODO status handling
	c := &dns.Client{Net: "udp"}
	server := res.pickRandomServer()
	resp, rtt, err := c.Exchange(m, server.addr)
	server.log.Allreq.Inc()
	if err != nil {
		server.log.Err.Inc()
		return nil, rtt, err
	}
	return resp, rtt, err
}

func (res *Resolver) pickRandomServer() *server {
	return &res.servers[rand.Intn(len(res.servers))]
}

func (res *Resolver) ShowStat() string {
	var sb strings.Builder
	sb.WriteString("Allrequests RcodeSuccess RcodeServerFailure RcodeNameError Err \n")
	for _, i := range res.servers {
		sb.WriteString(i.addr)
		sb.WriteString(" ")
		sb.WriteString(fmt.Sprint(i.log))
		sb.WriteString("\n")
	}
	return sb.String()
}

func New(addrs []string, logport string) *Resolver {
	rand.Seed(time.Now().Unix())
	servers := make([]server, 0, len(addrs))
	for _, i := range addrs {
		log := new(dnslog.Dnslog)
		servers = append(servers, server{addr: i, log: log})
	}
	return &Resolver{servers: servers, logport: logport}
}

func (res *Resolver) StartServer() {
	http.HandleFunc("/", res.LogHandler)
	http.ListenAndServe(res.logport, nil)
}
