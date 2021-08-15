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

type UpdateDocuments struct {
	UpdateDocs UpdateDocument `json:"doc"`
}

type BulkUpdate struct {
	Docs []UpdateDocument `json:"docs"`
}

type UpdateDocument struct {
	Id      string   `json:"_id"`
	Rev     string   `json:"_rev"`
	Program string   `json:"program"`
	Type    string   `json:"type"`
	Ips     []string `json:"ips"`
	Source  string   `json:"source"`
}

type CurrentDocuments struct {
	Rows []UpdateDocuments `json:"rows"`
}

type Scopes struct {
	Rows []ScopeDocs `json:"rows"`
}

type ScopeDocs struct {
	Doc ScopeDoc `json:"doc"`
}

type ScopeDoc struct {
	Inscope  []string `json:"inscope"`
	Outscope []string `json:"outscope"`
}

var conf BBRFConfig
