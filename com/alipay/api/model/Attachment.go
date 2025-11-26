package model

type Attachment struct {
	AttachmentType model.AttachmentType `json:"attachmentType,omitempty"`
	File           string               `json:"file,omitempty"`
	AttachmentName string               `json:"attachmentName,omitempty"`
	FileKey        string               `json:"fileKey,omitempty"`
}
