package orm

import "visionServiceGo/src/model"

type VisionORMModel interface {
	GetAll() ([]model.Vision, error)
	GetById(id uint64) (model.Vision, error)
	GetActive() (model.Vision, error)
	Create(vision model.Vision) (model.Vision, error)
	Update(vision model.Vision) (model.Vision, error)
	Delete(id uint64) error
	Activate(id uint64) error
	IsActivated(id uint64) (bool, error)
}
