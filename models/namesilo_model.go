package models

type NamesiloRecordModel struct {
	Namesilo Namesilo `json:"namesilo"`
}
type Request struct {
	Operation string `json:"operation"`
	IP        string `json:"ip"`
}
type ResourceRecord struct {
	RecordID string `json:"record_id"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Value    string `json:"value"`
	TTL      int    `json:"ttl"`
	Distance int    `json:"distance"`
}
type Reply struct {
	Code           int              `json:"code"`
	Detail         string           `json:"detail"`
	ResourceRecord []ResourceRecord `json:"resource_record"`
}
type Namesilo struct {
	Request Request `json:"request"`
	Reply   Reply   `json:"reply"`
}
