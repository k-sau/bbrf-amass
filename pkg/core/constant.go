package core

type Output struct {
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Tag       string    `json:"tag"`
	Sources   []string  `json:"sources"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Ip   string `json:"ip"`
	Cidr string `json:"cidr"`
	Asn  int    `json:"asn"`
	Desc string `json:"desc"`
}

type Documents struct {
	Docs []Document `json:"docs"`
}
type Document struct {
	Id      string   `json:"_id"`
	Program string   `json:"program"`
	Type    string   `json:"type"`
	Ips     []string `json:"ips"`
	Source  string   `json:"source"`
}

type BBRFConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	CouchDB  string `json:"couchdb"`
}

type Keys struct {
	Key []string `json:"keys"`
}

var conf BBRFConfig
