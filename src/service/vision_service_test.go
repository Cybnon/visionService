package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"visionServiceGo/src/model"
	mock "visionServiceGo/src/orm/mock"
)

func initTest(t *testing.T) (Service, *mock.MockVisionORMModel) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockVisionORMModel(ctrl)
	service := Service{orm: m, router: gin.Default()}
	service.SetRoutes()
	return service, m
}

func Test_GetActiveVisionShouldReturnActiveVision(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().GetActive().Times(1).Return(model.Vision{ID: 1, Title: "foo", Body: "bar", Active: true}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	service.GetActive(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func Test_GetAllShouldReturnAllVision(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().GetAll().Times(1).Return([]model.Vision{{ID: 1, Title: "foo", Body: "bar", Active: true}}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	service.GetAll(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func Test_GetByIdShouldReturnVision(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().GetById(uint64(2)).Times(1).Return(model.Vision{ID: 2, Title: "foo", Body: "bar", Active: true}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "2"}}

	service.GetById(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func Test_GetByIdShouldReturnAnErrorIfNoIdIsProvided(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().GetById(2).Times(1).Return(model.Vision{ID: 2, Title: "foo", Body: "bar", Active: true}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	service.GetById(c)
	assert.Equal(t, 500, c.Writer.Status())
}

func Test_CreateShouldReturnAnErrorIfNoVisionIsProvided(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().Create(2).Times(1).Return(model.Vision{ID: 2, Title: "foo", Body: "bar", Active: true}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	service.Create(c)
	assert.Equal(t, 500, c.Writer.Status())
}

func Test_DeleteShouldDeleteVision(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().Delete(uint64(2)).Times(1).Return(nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "2"}}

	service.Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func Test_UpdateShouldReturnAnErrorIfNoVisionIsProvided(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().Update(2).Times(1).Return(model.Vision{ID: 2, Title: "foo", Body: "bar", Active: true}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	service.Update(c)
	assert.Equal(t, 500, c.Writer.Status())
}

func Test_ActivateShouldActivateVision(t *testing.T) {
	service, m := initTest(t)

	m.EXPECT().GetActive().Times(1).Return(model.Vision{ID: 1, Title: "foo", Body: "bar", Active: true}, nil)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	service.GetActive(c)
	assert.Equal(t, 200, c.Writer.Status())
}
