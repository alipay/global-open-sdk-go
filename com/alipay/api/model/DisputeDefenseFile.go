package model

type DisputeDefenseFile struct {
	// Name of the generated defense material.
	FileName string `json:"fileName,omitempty"`
	// Material format: PDF or WORD.
	FileFormat string `json:"fileFormat,omitempty"`
	// MD5 of the material body as 32 lowercase hexadecimal characters.
	FileMd5 string `json:"fileMd5,omitempty"`
}
