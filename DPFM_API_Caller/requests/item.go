package requests

type Item struct {
	ProductionVersion       int     `json:"ProductionVersion"`
	ProductionVersionItem   int     `json:"ProductionVersionItem"`
	Product                 string  `json:"Product"`
	BusinessPartner         int     `json:"BusinessPartner"`
	Plant                   string  `json:"Plant"`
	BillOfMaterial          int     `json:"BillOfMaterial"`
	Operations              int     `json:"Operations"`
	ProductionVersionText   *string `json:"ProductionVersionText"`
	ProductionVersionStatus *string `json:"ProductionVersionStatus"`
	ValidityStartDate       *string `json:"ValidityStartDate"`
	ValidityEndDate         *string `json:"ValidityEndDate"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsLocked                *bool   `json:"IsLocked"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}
