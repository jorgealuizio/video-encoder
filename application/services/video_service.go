package services

import (
	"context"
	"video-encoder/domain"
	"video-encoder/application/repositories"
	"cloud.google.com/go/storage"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type VideoService struct {
	Video *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {

	cxt := context.BackGround()
	client, err := storage.NewClient(cxt)
	if err != nil {
		return err
	}

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(v.Video.FilePath)

	reader, err := obj.NewReader(cxt)
	if err != nil {
		return err
	}

	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	file, err := os.Create(os.GetEnv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err := file.Write(body)
	if err != nil {
		return err
	}

	defer file.Close()

	log.Printf("Video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {
	err := os.Mkdir(os.GetEnv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID, os.ModePerm)
	if err != nil {
		return err
	}

	source := os.GetEnv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".mp4"
	target := os.GetEnv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	printOutput(output)

	return nil

}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("========> Output: %s\n", string(out))
	}
}