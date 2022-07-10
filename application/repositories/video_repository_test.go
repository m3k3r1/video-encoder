package repositories_test

import (
	"github.com/m3k3r1/video-encoder/application/repositories"
	"github.com/m3k3r1/video-encoder/domain"
	"github.com/m3k3r1/video-encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource-id"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	videoExists, err := repo.Find(video.ID)

	require.Nil(t, err)
	require.NotEmpty(t, videoExists.ID)
	require.Equal(t, videoExists.ID, video.ID)
}
