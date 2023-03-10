package tailscale

import (
	"context"
	"net"
	"strings"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("tailscale")

const (
	TypeAll = iota
	TypeA
	TypeAAAA
)

// ServeDNS implements the plugin.Handler interface. This method gets called when tailscale is used
// in a Server.

func (t *Tailscale) resolveA(domainName string, msg *dns.Msg) {

	name := strings.Split(domainName, ".")[0]
	entry, ok := t.entries[name]["A"]
	if ok {
		log.Debugf("Found an v4 entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domainName, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP(entry),
		})
	} else {
		// There's no A record, so see if a CNAME exists
		log.Debug("No v4 entry after lookup, so trying CNAME")
		t.resolveCNAME(domainName, msg, TypeA)
	}

}

func (t *Tailscale) resolveAAAA(domainName string, msg *dns.Msg) {

	name := strings.Split(domainName, ".")[0]
	entry, ok := t.entries[name]["AAAA"]
	if ok {
		log.Debugf("Found a v6 entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.AAAA{
			Hdr:  dns.RR_Header{Name: domainName, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 60},
			AAAA: net.ParseIP(entry),
		})
	} else {
		// There's no AAAA record, so see if a CNAME exists
		log.Debug("No v6 entry after lookup, so trying CNAME")
		t.resolveCNAME(domainName, msg, TypeAAAA)
	}

}

func (t *Tailscale) resolveCNAME(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["CNAME"]
	if ok {
		log.Debugf("Found a CNAME entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.CNAME{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("CNAME record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("CNAME record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) resolveTXT(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["TXT"]
	if ok {
		log.Debugf("Found a TXT entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.CNAME{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("TXT record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("TXT record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) resolveRRSIG(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["RRSIG"]
	if ok {
		log.Debugf("Found a RRSIG entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.RRSIG{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeRRSIG, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("RRSIG record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("RRSIG record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) resolveCDNSKEY(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["CDNSKEY"]
	if ok {
		log.Debugf("Found a CDNSKEY entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.CDNSKEY{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeCDNSKEY, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("CDNSKEY record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("CDNSKEY record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

// CDS
func (t *Tailscale) resolveCDS(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["CDS"]
	if ok {
		log.Debugf("Found a CDS entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.CDS{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeCDS, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("CDS record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("CDS record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) resolveDS(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["CDS"]
	if ok {
		log.Debugf("Found a DS entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.CDS{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeDS, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("DS record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("DS record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) resolveDNSKEY(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["DNSKEY"]
	if ok {
		log.Debugf("Found a DNSKEY entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.DNSKEY{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeDNSKEY, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("DNSKEY record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("DNSKEY record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) resolveDNSKEY(domainName string, msg *dns.Msg, lookupType int) {

	name := strings.Split(domainName, ".")[0]
	target, ok := t.entries[name]["TLSA"]
	if ok {
		log.Debugf("Found a TLSA entry after lookup for: %s", name)
		msg.Answer = append(msg.Answer, &dns.TLSA{
			Hdr:    dns.RR_Header{Name: domainName, Rrtype: dns.TypeTLSA, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		})

		// Resolve local zone A or AAAA records if they exist for the referenced target
		if lookupType == TypeAll || lookupType == TypeA {
			log.Debug("TLSA record found, lookup up local recursive A")
			t.resolveA(target, msg)
		}
		if lookupType == TypeAll || lookupType == TypeAAAA {
			log.Debug("TLSA record found, lookup up local recursive AAAA")
			t.resolveAAAA(target, msg)
		}
	}

}

func (t *Tailscale) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debugf("Received request for name: %v", r.Question[0].Name)
	log.Debugf("Tailscale peers list has %d entries", len(t.entries))

	msg := dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	switch r.Question[0].Qtype {

	case dns.TypeA:
		log.Debug("Handling A record lookup")
		t.resolveA(r.Question[0].Name, &msg)

	case dns.TypeAAAA:
		log.Debug("Handling AAAA record lookup")
		t.resolveAAAA(r.Question[0].Name, &msg)

	case dns.TypeCNAME:
		log.Debug("Handling CNAME record lookup")
		t.resolveCNAME(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeTXT:
		log.Debug("Handling TXT record lookup")
		t.resolveTXT(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeRRSIG:
		log.Debug("Handling RRSIG record lookup")
		t.resolveRRSIG(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeCDNSKEY:
		log.Debug("Handling CDNSKEY record lookup")
		t.resolveCDNSKEY(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeCDS:
		log.Debug("Handling CDS record lookup")
		t.resolveCDS(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeDS:
		log.Debug("Handling DS record lookup")
		t.resolveDS(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeDNSKEY:
		log.Debug("Handling DNSKEY record lookup")
		t.resolveDNSKEY(r.Question[0].Name, &msg, TypeAll)

	case dns.TypeTLSA:
		log.Debug("Handling TLSA record lookup")
		t.resolveTLSA(r.Question[0].Name, &msg, TypeAll)
	}

	// If we have a response, write it back to the client.

	if len(msg.Answer) > 0 {
		log.Debugf("Writing response: %+v", msg)
		w.WriteMsg(&msg)
		return dns.RcodeSuccess, nil
	}

	// Export metric with the server label set to the current server handling the request.
	//requestCount.WithLabelValues(metrics.WithServer(ctx)).Inc()

	// Call next plugin (if any).
	return plugin.NextOrFailure(t.Name(), t.Next, ctx, w, r)
}
