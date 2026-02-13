package class

import (
	apimodel "glofox/internal/model/api_model"
	"glofox/pkg"
	"time"
)

type ClassService interface {
	CreateClass(*apimodel.Class) error
	ListClasses(query map[string]interface{}, skip int, limit int) ([]*Class, int, int, int, error)
}

type classService struct {
	classRepo ClassRepository
}

func NewClassService(classRepo ClassRepository) ClassService {
	return &classService{classRepo: classRepo}
}

func (s *classService) CreateClass(class *apimodel.Class) error {

	if class.ClassName == "" {
		return pkg.ErrClassNameRequired
	}
	if len(class.ClassName) < 3 || len(class.ClassName) > 100 {
		return pkg.ErrInvalidClassName
	}
	if class.StartDate == "" {
		return pkg.ErrStartDateRequired
	}
	if class.EndDate == "" {
		return pkg.ErrEndDateRequired
	}
	if class.Capacity <= 0 {
		return pkg.ErrInvalidCapacity
	}
	if class.InstructorID == 0 {
		return pkg.ErrInstructorRequired
	}

	layout := "2006-01-02"

	parsedStartDate, err := time.Parse(layout, class.StartDate)
	if err != nil {
		return pkg.ErrInvalidDateFormat
	}

	parsedEndDate, err := time.Parse(layout, class.EndDate)
	if err != nil {
		return pkg.ErrInvalidDateFormat
	}

	internalClass := &Class{
		ClassName:    class.ClassName,
		StartDate:    parsedStartDate,
		EndDate:      parsedEndDate,
		Capacity:     class.Capacity,
		InstructorID: class.InstructorID,
	}

	if err := s.classRepo.Create(internalClass); err != nil {
		return err // propagate repo error
	}

	return nil
}

func (s *classService) ListClasses(query map[string]interface{}, skip int, limit int) ([]*Class, int, int, int, error) {
	classes, total, limit, offset, err := s.classRepo.FindClasses(query, limit, skip)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	return classes, total, limit, offset, nil
}
