package domain_test

import (
	"testing"

	"github.com/celio001/MP4toDASH/domain"
	"github.com/go-playground/assert/v2"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/assert"
)

func TestValidadeIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	assert.Error(err)
}
