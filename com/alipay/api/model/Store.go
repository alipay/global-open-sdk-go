package model

type Store struct {
	ReferenceStoreId string   `json:"referenceStoreId,omitempty"`
	StoreName        string   `json:"storeName,omitempty"`
	StoreMCC         string   `json:"storeMCC,omitempty"`
	StoreDisplayName string   `json:"storeDisplayName,omitempty"`
	StoreTerminalId  string   `json:"storeTerminalId,omitempty"`
	StoreOperatorId  string   `json:"storeOperatorId,omitempty"`
	StoreAddress     *Address `json:"storeAddress,omitempty"`
	StorePhoneNo     string   `json:"storePhoneNo,omitempty"`
}
