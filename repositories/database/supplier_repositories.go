package database

import "github.com/Heroin-lab/nixProject/models"

type SuppRepose struct {
	storage *Storage
}

func (r *SuppRepose) GetSuppliersByCategory(category string) ([]*models.SuppliersForSelect, error) {
	suppModel := models.Suppliers{}

	rows, err := r.storage.DB.Query("SELECT title, type_name, working_time\n"+
		"FROM suppliers\n"+
		"INNER JOIN suppliers_type st on suppliers.type_id = st.id\n"+
		"WHERE type_name=?", category)
	if err != nil {
		return nil, err
	}

	rowsArr := make([]*models.SuppliersForSelect, 0)

	for rows.Next() {

		err = rows.Scan(
			&suppModel.Title,
			&suppModel.Type,
			&suppModel.Working_time,
		)
		if err != nil {
			return nil, err
		}

		rowsArr = append(rowsArr, &models.SuppliersForSelect{
			Title:        suppModel.Title,
			Type:         suppModel.Type,
			Working_time: suppModel.Working_time,
		})
	}
	return rowsArr, nil
}

func (r *SuppRepose) InsertSupplier(product *models.Products) *models.Products {
	return nil
}

func (r *SuppRepose) DeleteSupplier(product *models.Products) error {
	return nil
}
