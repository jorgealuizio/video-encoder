package repositories_test

import (
	"video-encoder/application/repositories"
	"video-encoder/framework/database"
	"video-encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "Path"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db:db}
	videoRepo.Insert(video)

	job, err := domain.NewJob("output_path", "pending", video)
	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDb{Db:db}
	jobRepo.Insert(job)

	j, err := jobRepo.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "Path"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db:db}
	videoRepo.Insert(video)

	job, err := domain.NewJob("output_path", "pending", video)
	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDb{Db:db}
	jobRepo.Insert(job)

	job.Status = "Complete"
	jobRepo.Update(job)

	j, err := jobRepo.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}