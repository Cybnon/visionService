package service

import (
	"visionServiceGo/src/model"
)

type Model interface {
	GetAll() ([]model.Vision, error)
	GetById(id uint) (model.Vision, error)
	GetActive() (model.Vision, error)
	Create(vision model.Vision) (model.Vision, error)
	Update(vision model.Vision) (model.Vision, error)
	Delete(id uint) error
	Activate(id uint) error
	IsActivated(id uint) (bool, error)
}
