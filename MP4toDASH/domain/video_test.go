package domain_test

import (
	"testing"
	"time"

	"github.com/celio001/MP4toDASH/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidadeIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	assert.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T) {

	video := domain.NewVideo()
	video.ID = "abc"
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	assert.Error(t, err)

}
func TestVideoValidation(t *testing.T) {

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	assert.NoError(t, err)

}
