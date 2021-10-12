package domain_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"video-encoder/domain"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.Video{
		ID:         "abc",
		ResourceID: "fake",
		FilePath:   "fake",
		CreatedAt:  time.Now(),
	}

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIsValid(t *testing.T) {
	video := domain.Video{
		ID:         uuid.NewV4().String(),
		ResourceID: "fake",
		FilePath:   "fake",
		CreatedAt:  time.Now(),
	}

	err := video.Validate()
	require.Nil(t, err)
}
