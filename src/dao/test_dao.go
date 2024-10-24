package dao

import (
	"ChaoXing_SignIn_GoRebuild/src/model"
	"github.com/jinzhu/gorm"
)

type TestDao interface {
	AddTest(test *model.Test) error
	UpdateTest(test *model.Test) error
	DeleteTest(test *model.Test) error
	ListTestById(id uint) (*model.Test, error)
}

type TestDaoImpl struct {
	DB *gorm.DB
}

func (b *TestDaoImpl) AddTest(Test *model.Test) error {
	if createResult := b.DB.Create(Test); createResult.Error != nil {
		return createResult.Error
	}
	return nil
}

func (b *TestDaoImpl) UpdateTest(Test *model.Test) error {
	if saveResult := b.DB.Save(Test); saveResult.Error != nil {
		return saveResult.Error
	}
	return nil
}

func (b *TestDaoImpl) DeleteTest(Test *model.Test) error {
	if deleteResult := b.DB.Delete(Test); deleteResult.Error != nil {
		return deleteResult.Error
	}
	return nil
}

func (b *TestDaoImpl) ListTestById(id uint) (*model.Test, error) {
	var Test model.Test
	if listResult := b.DB.Where("id = ?", id).First(&Test); listResult.Error != nil {
		return nil, listResult.Error
	}
	return &Test, nil
}
