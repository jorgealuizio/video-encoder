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

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "Path"
	video.CreatedAt = time.Now()

	repository := repositories.VideoRepositoryDb{Db:db}
	repository.Insert(video)

	v, err := repository.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}