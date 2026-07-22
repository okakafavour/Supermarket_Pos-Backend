package customer

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create Customer
func (r *Repository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

// Update Customer
func (r *Repository) Update(customer *Customer) error {
	return r.db.Save(customer).Error
}

// Soft Delete Customer
func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Customer{}, "id = ?", id).Error
}

// Restore Customer
func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Customer{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

// Permanent Delete Customer
func (r *Repository) PermanentDelete(id string) error {
	return r.db.
		Unscoped().
		Delete(&Customer{}, "id = ?", id).Error
}

// Get Customer By ID
func (r *Repository) GetByID(id string) (*Customer, error) {

	var customer Customer

	err := r.db.
		First(&customer, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// Get All Customers
func (r *Repository) GetAll() ([]Customer, error) {

	var customers []Customer

	err := r.db.
		Order("created_at DESC").
		Find(&customers).Error

	return customers, err
}

// Get Deleted Customers
func (r *Repository) GetDeleted() ([]Customer, error) {

	var customers []Customer

	err := r.db.
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("deleted_at DESC").
		Find(&customers).Error

	return customers, err
}

// Search Customers
func (r *Repository) Search(query string) ([]Customer, error) {

	var customers []Customer

	err := r.db.
		Where(`
			first_name ILIKE ?
			OR last_name ILIKE ?
			OR CONCAT(first_name, ' ', last_name) ILIKE ?
			OR email ILIKE ?
			OR phone ILIKE ?
		`,
			"%"+query+"%",
			"%"+query+"%",
			"%"+query+"%",
			"%"+query+"%",
			"%"+query+"%",
		).
		Order("created_at DESC").
		Find(&customers).Error

	return customers, err
}
