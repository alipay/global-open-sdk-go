package model

type EvidenceFile struct {
	FileName    string `json:"fileName,omitempty"`
	FileFormat  string `json:"fileFormat,omitempty"`
	FileContent string `json:"fileContent,omitempty"`
}
