package model

type Attachment struct {
	AttachmentType string `json:"attachmentType,omitempty"`
	File           string `json:"file,omitempty"`
	AttachmentName string `json:"attachmentName,omitempty"`
	FileKey        string `json:"fileKey,omitempty"`
}
