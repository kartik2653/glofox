package class

import (
	"fmt"

	"gorm.io/gorm"
)

type ClassRepository interface {
	Create(class *Class) error
	FindClasses(query map[string]interface{}, limit int, offset int) ([]*Class, int, int, int, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Create(class *Class) error {
	return r.db.Create(class).Error
}

func (r *classRepository) FindClasses(
	query map[string]interface{},
	limit int,
	offset int,
) ([]*Class, int, int, int, error) {

	var classes []*Class
	fmt.Println("Query in repo:", query) // Debugging statement to check the query being passed
	err := r.db.
		Where(query).
		Limit(limit).
		Offset(offset).
		Find(&classes).Error

	if err != nil {
		return nil, 0, 0, 0, err
	}

	var total int64

	if err := r.db.
		Model(&Class{}).
		Where(query).
		Count(&total).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return classes, int(total), limit, offset, nil
}
