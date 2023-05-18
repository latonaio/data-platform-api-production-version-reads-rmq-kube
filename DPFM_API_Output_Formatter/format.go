package dpfm_api_output_formatter

import (
	"data-platform-api-production-version-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.ProductionVersion,
			&pm.Product,
			&pm.OwnerBusinessPartner,
			&pm.OwnerPlant,
			&pm.BillOfMaterial,
			&pm.Operations,
			&pm.ProductionVersionText,
			&pm.ProductionVersionStatus,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsLocked,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			ProductionVersion:       data.ProductionVersion,
			Product:                 data.Product,
			OwnerBusinessPartner:    data.OwnerBusinessPartner,
			OwnerPlant:              data.OwnerPlant,
			BillOfMaterial:          data.BillOfMaterial,
			Operations:              data.Operations,
			ProductionVersionText:   data.ProductionVersionText,
			ProductionVersionStatus: data.ProductionVersionStatus,
			ValidityStartDate:       data.ValidityStartDate,
			ValidityEndDate:         data.ValidityEndDate,
			CreationDate:            data.CreationDate,
			LastChangeDate:          data.LastChangeDate,
			IsLocked:                data.IsLocked,
			IsMarkedForDeletion:     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.BillOfMaterial,
			&pm.Operations,
			&pm.ProductionVersionText,
			&pm.ProductionVersionStatus,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsLocked,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			ProductionVersion:       data.ProductionVersion,
			ProductionVersionItem:   data.ProductionVersionItem,
			Product:                 data.Product,
			BusinessPartner:         data.BusinessPartner,
			Plant:                   data.Plant,
			BillOfMaterial:          data.BillOfMaterial,
			Operations:              data.Operations,
			ProductionVersionText:   data.ProductionVersionText,
			ProductionVersionStatus: data.ProductionVersionStatus,
			ValidityStartDate:       data.ValidityStartDate,
			ValidityEndDate:         data.ValidityEndDate,
			CreationDate:            data.CreationDate,
			LastChangeDate:          data.LastChangeDate,
			IsLocked:                data.IsLocked,
			IsMarkedForDeletion:     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
