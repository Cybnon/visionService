package orm

import (
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
	"time"
	"visionServiceGo/src/model"
)

var visionEntry1 = model.Vision{
	Title:     "Test Vision 1",
	Body:      "Test Vision 1 Body",
	Active:    true,
	CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
}

var visionEntry2 = model.Vision{
	Title:     "Test Vision 2",
	Body:      "Test Vision 2 Body",
	Active:    false,
	CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
}

func Test_GetAll(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	_, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	visions, err := visionORM.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(visions))
	require.Equal(t, visionEntry1.Title, visions[0].Title)
	require.Equal(t, visionEntry1.Body, visions[0].Body)
	require.Equal(t, visionEntry1.Active, visions[0].Active)
	require.Equal(t, visionEntry1.CreatedAt, visions[0].CreatedAt)
	require.Equal(t, visionEntry1.UpdatedAt, visions[0].UpdatedAt)
}

func Test_GetById(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	createdVision, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	createdVision2, err := visionORM.Create(visionEntry2)
	require.NoError(t, err)
	visions, err := visionORM.GetById(createdVision.ID)
	require.NoError(t, err)
	require.Equal(t, visionEntry1.Title, visions.Title)
	require.Equal(t, visionEntry1.Body, visions.Body)
	require.Equal(t, visionEntry1.Active, visions.Active)
	require.Equal(t, visionEntry1.CreatedAt, visions.CreatedAt)
	require.Equal(t, visionEntry1.UpdatedAt, visions.UpdatedAt)
	visions, err = visionORM.GetById(createdVision2.ID)
	require.NoError(t, err)
	require.Equal(t, visionEntry2.Title, visions.Title)
	require.Equal(t, visionEntry2.Body, visions.Body)
	require.Equal(t, visionEntry2.Active, visions.Active)
	require.Equal(t, visionEntry2.CreatedAt, visions.CreatedAt)
	require.Equal(t, visionEntry2.UpdatedAt, visions.UpdatedAt)
}

func Test_Create(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	createdVision, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	require.Equal(t, visionEntry1.Title, createdVision.Title)
	require.Equal(t, visionEntry1.Body, createdVision.Body)
	require.Equal(t, visionEntry1.Active, createdVision.Active)
	require.Equal(t, visionEntry1.CreatedAt, createdVision.CreatedAt)
	require.Equal(t, visionEntry1.UpdatedAt, createdVision.UpdatedAt)
}

func Test_Update(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	createdVision, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	createdVision.Title = "Updated Title"
	updatedVision, err := visionORM.Update(createdVision)
	require.NoError(t, err)
	require.Equal(t, createdVision.Title, updatedVision.Title)
	require.Equal(t, createdVision.Body, updatedVision.Body)
	require.Equal(t, createdVision.Active, updatedVision.Active)
	require.Equal(t, createdVision.CreatedAt, updatedVision.CreatedAt)
	require.NotEqual(t, createdVision.UpdatedAt, updatedVision.UpdatedAt)
}

func Test_Delete(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	createdVision, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	err = visionORM.Delete(createdVision.ID)
	require.NoError(t, err)
	_, err = visionORM.GetById(createdVision.ID)
	require.Error(t, err)
}

func Test_deactivateVision(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	createdVision, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	require.True(t, createdVision.Active)
	createdVision2, err := visionORM.Create(visionEntry2)
	require.NoError(t, err)
	require.False(t, createdVision2.Active)
	createdVision2.Active = true
	updatedVision, err := visionORM.Update(createdVision2)
	require.NoError(t, err)
	isActive, err := visionORM.IsActivated(updatedVision.ID)
	require.NoError(t, err)
	require.True(t, isActive)
	isActive, err = visionORM.IsActivated(createdVision.ID)
	require.NoError(t, err)
	require.False(t, isActive)
}

func Test_deactivateVisionIfNewActiveOneIsCreatesd(t *testing.T) {
	tmpDir := t.TempDir()

	t.Cleanup(func() {
		_ = os.RemoveAll(tmpDir)
	})

	db := createTestDb(filepath.Join(tmpDir, "test.db"))
	visionORM := VisionORM{db: db}

	createdVision, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	require.True(t, createdVision.Active)
	createdVision2, err := visionORM.Create(visionEntry1)
	require.NoError(t, err)
	require.True(t, createdVision2.Active)
	isActive, err := visionORM.IsActivated(createdVision2.ID)
	require.NoError(t, err)
	require.True(t, isActive)
	isActive, err = visionORM.IsActivated(createdVision.ID)
	require.NoError(t, err)
	require.False(t, isActive)
}

func createTestDb(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&model.Vision{})
	return db
}
