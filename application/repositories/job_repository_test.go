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

func TestNewJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource-id"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db: db}
	videoRepo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDb{Db: db}
	jobRepo.Insert(job)

	jobExists, err := jobRepo.Find(job.ID)

	require.Nil(t, err)
	require.NotEmpty(t, jobExists.ID)
	require.Equal(t, jobExists.ID, job.ID)
	require.Equal(t, jobExists.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource-id"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db: db}
	videoRepo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDb{Db: db}
	jobRepo.Insert(job)

	job.Status = "Completed"

	jobRepo.Update(job)

	jobExists, err := jobRepo.Find(job.ID)

	require.Nil(t, err)
	require.NotEmpty(t, jobExists.ID)
	require.Equal(t, jobExists.ID, job.ID)
	require.Equal(t, jobExists.Status, job.Status)
}
