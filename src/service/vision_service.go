package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
	"visionServiceGo/src/model"
	"visionServiceGo/src/orm"
)

type Service struct {
	orm    orm.VisionORMModel
	router *gin.Engine
}

func NewService(orm orm.VisionORMModel) *Service {
	router := gin.Default()
	service := &Service{orm: orm, router: router}
	service.CORSMiddleware()
	service.SetRoutes()
	return service
}

func (s *Service) Run(addr ...string) error {
	return s.router.Run(addr...)
}

func (s *Service) SetRoutes() {
	s.router.GET("/vision", s.GetAll)
	s.router.GET("/vision/:id", s.GetById)
	s.router.GET("/vision/active", s.GetActive)
	s.router.POST("/vision", s.Create)
	s.router.PUT("/vision/:id/activate", s.Activate)
	s.router.PUT("/vision/:id", s.Update)
	s.router.DELETE("/vision/:id", s.Delete)
}

func (s *Service) CORSMiddleware() {
	s.router.Use(cors.Default())
}

func (s *Service) GetAll(c *gin.Context) {
	visions, err := s.orm.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, visions)
}

func (s *Service) GetById(c *gin.Context) {
	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	vision, err := s.orm.GetById(uId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vision)
}

func (s *Service) GetActive(c *gin.Context) {
	vision, err := s.orm.GetActive()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vision)
}

func (s *Service) Create(c *gin.Context) {
	var vision model.Vision
	err := c.BindJSON(&vision)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	createdVision, err := s.orm.Create(vision)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, createdVision)
}

func (s *Service) Update(c *gin.Context) {
	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var vision model.Vision
	err = c.BindJSON(&vision)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	vision.ID = uId
	updatedVision, err := s.orm.Update(vision)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedVision)
}

func (s *Service) Delete(c *gin.Context) {
	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = s.orm.Delete(uId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Vision deleted successfully"})
}

func (s *Service) Activate(c *gin.Context) {
	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = s.orm.Activate(uId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Vision activated successfully"})
}

func (s *Service) IsActivated(c *gin.Context) {
	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	isActivated, err := s.orm.IsActivated(uId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"isActivated": isActivated})
}
