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

func TestJobRespositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	assert.NoError(t, err)
	
	repoJob := repositories.JobRepositoryDb{Db:db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	assert.NoError(t, err)
	assert.Equal(t, j.ID, job.ID)
	assert.Equal(t, j.VideoID, video.ID)
}

func TestJobRespositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	assert.NoError(t, err)
	
	repoJob := repositories.JobRepositoryDb{Db:db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	assert.NotEmpty(t, j.ID)
	assert.NoError(t, err)
	assert.Equal(t, j.Status, job.Status)
}
