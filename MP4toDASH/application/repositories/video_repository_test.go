package repositories_test

import (
	"testing"
	"time"

	"github.com/celio001/MP4toDASH/application/repositories"
	"github.com/celio001/MP4toDASH/domain"
	"github.com/celio001/MP4toDASH/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestVideoRepositoryDbInsert(t *testing.T){
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	assert.NoError(t, err)
	assert.NotNil(t, v.ID)
	assert.Equal(t, v.ID, video.ID)
}