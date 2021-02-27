package departments

import (
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"gorm.io/gorm"
)

type repo struct {
	gormDB *gorm.DB
}

// NewRepo creates a new departments repo.
func NewRepo(gormDB *gorm.DB) Repo {
	return &repo{
		gormDB: gormDB,
	}
}

func (r *repo) Create(newDepartmentData model.NewDepartment) (*model.Department, error) {
	department := Department{
		Name: newDepartmentData.Name,
	}

	result := r.gormDB.Create(
		&department,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return r.GetByID(uint64(department.ID))
}

func (r *repo) GetByID(id uint64) (*model.Department, error) {
	var department Department

	result := r.gormDB.First(&department, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_department := model.Department{
		ID:   int64(id),
		Name: department.Name,
	}

	return &_department, nil
}

func (r *repo) GetAll() ([]*model.Department, error) {
	var departments []*Department

	result := r.gormDB.Find(&departments)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	var _departments []*model.Department
	for _, department := range departments {
		_department := model.Department{
			ID:   int64(department.ID),
			Name: department.Name,
		}
		_departments = append(_departments, &_department)
	}

	return _departments, nil
}

func (r *repo) DeleteByID(id uint64) (*model.Department, error) {
	var department Department

	result := r.gormDB.Delete(&department, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_department := model.Department{
		ID:   int64(id),
		Name: department.Name,
	}

	return &_department, nil
}
