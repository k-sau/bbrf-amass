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
