package property

import (
	"github.com/devlopersabbir/juan_don82-server/api/property/core"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func Store(property *core.Property) error {
	return database.DB.Create(property).Error
}

func FindAll() ([]core.Property, error) {
	var properties []core.Property
	if err := database.DB.Find(&properties).Error; err != nil {
		return nil, err
	}
	return properties, nil
}

func FindByID(id int) (*core.Property, error) {
	var property core.Property
	if err := database.DB.First(&property, id).Error; err != nil {
		return nil, err
	}
	return &property, nil
}

func Update(property *core.Property) error {
	return database.DB.Save(property).Error
}

func Delete(id int) error {
	return database.DB.Delete(&core.Property{}, id).Error
}
