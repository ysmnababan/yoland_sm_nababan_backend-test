package service

import (
	"soal2/murid/api/entity"
	"soal2/murid/api/repository"
)

type Service struct {
	repository.MuridRepository
}

type MuridService interface {
	GetAll() []entity.MuridEntity
}

func (r *Service) GetAll() []entity.MuridEntity{
	return r.MuridRepository.GetAll();
}

