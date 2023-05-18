package requests

type Header struct {
	ProductionVersion       int     `json:"ProductionVersion"`
	Product                 string  `json:"Product"`
	OwnerBusinessPartner    int     `json:"OwnerBusinessPartner"`
	OwnerPlant              string  `json:"OwnerPlant"`
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
