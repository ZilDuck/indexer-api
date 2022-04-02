package entity

type Metadata struct {
	Uri        string                 `json:"uri"`
	Properties map[string]interface{} `json:"properties"`
	IsIpfs     bool                   `json:"ipfs"`

	Status MetadataStatus `json:"status"`
	Error  string         `json:"error"`
}

type MetadataStatus string
var (
	MetadataPending MetadataStatus = "pending"
	MetadataSuccess MetadataStatus = "success"
	MetadataFailure MetadataStatus = "failure"
)
