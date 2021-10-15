package constants

type UnResolvedDocuments struct {
	Docs []UnResolvedDocument `json:"docs"`
}

type UnResolvedDocument struct {
	Id      string   `json:"_id"`
	Program string   `json:"program"`
	Type    string   `json:"type"`
	IPs     []string `json:"ips"`
}
