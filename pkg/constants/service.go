package constants

type ServiceDocuments struct {
	Docs []ServiceDocument `json:"docs"`
}

type ServiceDocument struct {
	Id      string `json:"_id"`
	Program string `json:"program"`
	Type    string `json:"type"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Service string `json:"service"`
}

type ServiceUpdateDocuments struct {
	UpdateDocs ServiceUpdateDocument `json:"doc"`
}

type ServiceBulkUpdate struct {
	Docs []ServiceUpdateDocument `json:"docs"`
}

type ServiceUpdateDocument struct {
	Id      string `json:"_id"`
	Rev     string `json:"_rev"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Service string `json:"service"`
	Program string `json:"program"`
	Type    string `json:"type"`
}

type ServiceCurrentDocuments struct {
	Rows []ServiceUpdateDocuments `json:"rows"`
}
