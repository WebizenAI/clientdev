# coredns-tailscale-Wails- Webizen-DNS tailscale-client

THIS IS A WORK IN PROGRESS!!! 


Tailscale client using [Wails](https://wails.io).
=================

An Experimental, untested Frankenstein that's likely not working.

The app will likely be updated to add a database. 



# App Manager UI Features

* File upload and reception.
* Account switching (ugly but works).
* Host listing (target is for headscale based solution, so I am open to criticism).
* Select an exit node.
* Host IP listing.
* Some tag description (like OS, or whether it's online or how long since it was last seen).


it integrates the CoreDNS plugin implementation for Tailscale networks that allows for CNAME and TXT records to be defined via Tailscale node tags.

Rationale
---------

The obective is to be able to provide support for Webizen-DNS which requires DNSSEC, DANE, TXT and CNAME records to be defined in DNS. This plugin is an attempt to provide a way to do that in a way that'll work with Tailscale. 

Features
--------
The original plugin for CoreDNS allows the following:

1. Automatically serving an (arbitrary) DNS zone with each Tailscale server in your Tailnet added with A and AAAA records.
2. Allowing CNAME records to be defined via Tailscale node tags that link logical names to Tailscale machines.

This plugin adds the following:
3. Allowing TXT records to be defined via Tailscale node tags that link logical names to Tailscale machines.
4. Allowing DNSSEC and DANE records to be defined via Tailscale node tags that link logical names to Tailscale machines.

Configuration
-------------

```
example.org:53 {
  tailscale example.org
  log 
  errors
}
```
The above configuration will serve the connected Tailnet on the `example.com`. So, for a Tailnet with a machine named `test-machine`, A and AAAA records for `test-machine.example.com` will resolve.

CNAME records via Labels
------------------------

A CNAME record can be added to point to a machine by simply creating a Tailscale machine tag prefixed by `cname-`. Any text in the tag after that prefix will be used to generate the resulting CNAME entry, so for example, the tag `cname-friendly-name` on the above `test-machine` will result in the following DNS records:

```
friendly-name IN CNAME test-machine.example.com.
test-machine  IN A <Tailscale IPv4 Address>
test-machine  IN AAAA <Tailscale IPv6 Address>
```

Tailscale
---------
Note that currently this plugin uses the local machine Tailscale socket to access Tailnet information. As a result, only machines reachable from the hosting Tailscale machine will be configured in DNS. Those machines are the ones output in `tailscale status` output (and the machine itself). This was implemented to avoid the need for managing expiring Tailscale API tokens.


TODO
----
   * lots.
   * currently only defined to work with linux, if indeed it builds.


# Dev Tests - Installation

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
git clone https://github.com/dgrr/tailscale-client
cd tailscale-client
make install
```
