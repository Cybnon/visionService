package orm

import (
	"gorm.io/gorm"
	"visionServiceGo/src/model"
)

type VisionORM struct {
	db *gorm.DB
}

func NewVisionORM(db *gorm.DB) *VisionORM {
	return &VisionORM{db}
}

func (v *VisionORM) GetAll() ([]model.Vision, error) {
	var visions []model.Vision
	result := v.db.Find(&visions)
	return visions, result.Error
}

func (v *VisionORM) GetById(id uint64) (model.Vision, error) {
	var vision model.Vision
	result := v.db.First(&vision, id)
	return vision, result.Error
}

func (v *VisionORM) Create(vision model.Vision) (model.Vision, error) {
	err := v.deactivateOldVisionIfActive(vision)
	if err != nil {
		return model.Vision{}, err
	}
	result := v.db.Create(&vision)
	return vision, result.Error
}

func (v *VisionORM) Update(vision model.Vision) (model.Vision, error) {
	err := v.deactivateOldVisionIfActive(vision)
	if err != nil {
		return model.Vision{}, err
	}
	result := v.db.Save(&vision)
	return vision, result.Error
}

func (v *VisionORM) deactivateOldVisionIfActive(vision model.Vision) error {
	if vision.Active {
		err := v.DeactivateActiveVision()
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VisionORM) Delete(id uint64) error {
	result := v.db.Delete(&model.Vision{}, id)
	return result.Error
}

func (v *VisionORM) Activate(id uint64) error {
	err := v.DeactivateActiveVision()
	if err != nil {
		return err
	}
	result := v.db.Model(&model.Vision{}).Where("id = ?", id).Update("active", true)
	return result.Error
}

func (v *VisionORM) DeactivateActiveVision() error {
	result := v.db.Model(&model.Vision{}).Where("active = ?", true).Update("active", false)
	return result.Error
}

func (v *VisionORM) IsActivated(id uint64) (bool, error) {
	var vision model.Vision
	result := v.db.First(&vision, id)
	return vision.Active, result.Error
}

func (v *VisionORM) GetActive() (model.Vision, error) {
	var vision model.Vision
	result := v.db.First(&vision, "active = ?", true)
	return vision, result.Error
}
