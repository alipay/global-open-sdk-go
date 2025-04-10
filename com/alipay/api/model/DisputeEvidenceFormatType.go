package model

type DisputeEvidenceFormatType string

const (
	DisputeEvidenceFormatType_PDF  DisputeEvidenceFormatType = "PDF"
	DisputeEvidenceFormatType_WORD DisputeEvidenceFormatType = "WORD"
	DisputeEvidenceFormatType_ZIP  DisputeEvidenceFormatType = "ZIP"
	DisputeEvidenceFormatType_JPG  DisputeEvidenceFormatType = "JPG"
)

type DisputeEvidenceType string

const (
	DisputeEvidenceType_DISPUTE_EVIDENCE_TEMPLATE DisputeEvidenceType = "DISPUTE_EVIDENCE_TEMPLATE"
	DisputeEvidenceType_DISPUTE_EVIDENCE_FILE     DisputeEvidenceType = "DISPUTE_EVIDENCE_FILE"
)
