package repository

import (
	"errors"
	"soal2/murid/api/entity"

	"gorm.io/gorm"
)

type MuridRepository interface {
	GetAll() ([]entity.MuridEntity)
	GetById(id int) (entity.MuridEntity, error)
	// Create(murid entity.MuridEntity)
	// Update(murid entity.MuridEntity, id int)
	// Delete(id int)
}

type DB struct {
	DB *gorm.DB
}

func (r *DB) GetAll() []entity.MuridEntity{
	var murid []entity.MuridEntity
	res:= r.DB.Find(&murid);
	if res.Error != nil {
		return nil;
	}
	return murid;
}

func (r *DB) GetById(id int) (entity.MuridEntity, error){
	var murid entity.MuridEntity;
	res := r.DB.Where("id == ? ", id).First(&murid);
	if res.Error != nil {
		return entity.MuridEntity{}, errors.New("error query")
	}
	return murid, nil;
}