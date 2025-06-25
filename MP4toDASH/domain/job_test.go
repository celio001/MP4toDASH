package domain_test

import (
	"testing"
	"time"

	"github.com/celio001/MP4toDASH/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Convert", video)
	assert.NoError(t, err)
	assert.NotNil(t, job)
}
