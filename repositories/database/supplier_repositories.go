package database

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
)

type SuppRepose struct {
	storage *Storage
}

func (r *SuppRepose) GetSuppliersByCategory(category string) ([]*models.Suppliers, error) {
	suppModel := &models.Suppliers{}

	rows, err := r.storage.DB.Query("SELECT suppliers.id, supp_name, type_name, image, open_time, close_time\n"+
		"FROM suppliers\n"+
		"INNER JOIN suppliers_types st on suppliers.type_id = st.id\n"+
		"WHERE type_name=?", category)
	if err != nil {
		return nil, err
	}

	rowsArr := make([]*models.Suppliers, 0)

	for rows.Next() {

		err = rows.Scan(
			&suppModel.Id,
			&suppModel.Name,
			&suppModel.Type,
			&suppModel.Image,
			&suppModel.Opening,
			&suppModel.Closing,
		)
		if err != nil {
			return nil, err
		}

		rowsArr = append(rowsArr, &models.Suppliers{
			Id:    suppModel.Id,
			Name:  suppModel.Name,
			Type:  suppModel.Type,
			Image: suppModel.Image,
			WorkingHours: models.WorkingHours{
				Opening: suppModel.Opening,
				Closing: suppModel.Closing,
			},
		})

	}

	logger.Info("The set of suppliers has just been sent to the customer")
	return rowsArr, nil
}

func (r *SuppRepose) AddSupplier(sup *models.Suppliers) (*models.Suppliers, error) {
	_, err := r.storage.DB.Exec("INSERT INTO suppliers (title, type_id, open_time, close_time)\n"+
		"VALUES (?, ?, ?)",
		sup.Name,
		sup.Type,
		sup.WorkingHours.Opening,
		sup.WorkingHours.Closing)
	if err != nil {
		return nil, err
	}
	logger.Info("Supplier with title='" + sup.Name + "' was successfully created!")
	return sup, nil
}

func (r *SuppRepose) DeleteSupplier(suppId string) error {
	_, err := r.storage.DB.Exec("DELETE FROM suppliers WHERE id=?",
		suppId)
	if err != nil {
		return err
	}
	logger.Info("Supplier with id='" + suppId + "' was successfully deleted!")
	return nil
}

func (r *SuppRepose) UpdateSupplier(sup *models.Suppliers) error {
	_, err := r.storage.DB.Exec("UPDATE suppliers SET title=?, type_id=?, opening_time=?, closing_time=?\n"+
		"WHERE id=?",
		sup.Name,
		sup.Type,
		sup.Opening,
		sup.Closing,
		sup.Id)
	if err != nil {
		return err
	}
	logger.Info("Supplier with id='" + sup.Name + "' was successfully updated!")
	return nil
}
