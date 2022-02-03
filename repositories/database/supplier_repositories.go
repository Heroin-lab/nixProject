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

	rows, err := r.storage.DB.Query("SELECT suppliers.id, title, type_name, working_time\n"+
		"FROM suppliers\n"+
		"INNER JOIN suppliers_type st on suppliers.type_id = st.id\n"+
		"WHERE type_name=?", category)
	if err != nil {
		return nil, err
	}

	rowsArr := make([]*models.Suppliers, 0)

	for rows.Next() {

		err = rows.Scan(
			&suppModel.Id,
			&suppModel.Title,
			&suppModel.Type,
			&suppModel.Working_time,
		)
		if err != nil {
			return nil, err
		}

		rowsArr = append(rowsArr, &models.Suppliers{
			Id:           suppModel.Id,
			Title:        suppModel.Title,
			Type:         suppModel.Type,
			Working_time: suppModel.Working_time,
		})
	}

	logger.Info("The set of suppliers has just been sent to the customer")
	return rowsArr, nil
}

func (r *SuppRepose) AddSupplier(sup *models.Suppliers) (*models.Suppliers, error) {
	_, err := r.storage.DB.Exec("INSERT INTO suppliers (title, type_id, working_time)\n"+
		"VALUES (?, ?, ?)",
		sup.Title,
		sup.Type,
		sup.Working_time)
	if err != nil {
		return nil, err
	}
	logger.Info("Supplier with title='" + sup.Title + "' was successfully created!")
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
	_, err := r.storage.DB.Exec("UPDATE suppliers SET title=?, type_id=?, working_time=?\n"+
		"WHERE id=?",
		sup.Title,
		sup.Type,
		sup.Working_time,
		sup.Id)
	if err != nil {
		return err
	}
	logger.Info("Supplier with id='" + sup.Id + "' was successfully updated!")
	return nil
}
