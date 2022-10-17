package proxmoxve

import (
	resty "gopkg.in/resty.v1"
)

// ProxmoxVE open api connection representation
type ProxmoxVE struct {
	// connection parameters
	Username string // root
	password string // must be given
	Realm    string // pam
	Host     string
	Port     int // default 8006

	// not so imported internal stuff
	Node                string // if not present, use first node present
	Prefix              string // if PVE is proxied, this is the added prefix
	CSRFPreventionToken string // filled by the framework
	Ticket              string // filled by the framework

	Version string // ProxmoxVE version of the connected host

	client *resty.Client // resty client
}