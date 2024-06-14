package repository

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
)

type EachBlob struct {
	Shorten string `json:"shorten"`
	Long    string `json:"long"`
}
type BlobStore struct {
	Blob []EachBlob `json:"blob"`
}

type Repository struct {
	Blob BlobStore
}

func (repo *Repository) Init() error {
	file, err := os.ReadFile("blob.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("Log: no blob.json found in current directory")
			return nil
		}
		log.Printf("Error: reading blob.json: %v\n", err)
		return err
	}

	var blob BlobStore
	err = json.Unmarshal(file, &blob)
	if err != nil {
		log.Printf("Error: unmarshaling blob: %v\n", err)
		return err
	}

	repo.Blob = blob
	log.Printf("Blob store initialized with %d elements\n", len(blob.Blob))

	return nil
}

func (repo *Repository) Save(url string) (string, error) {
	short := rand.IntN(89999) + 10000
	shortString := strconv.Itoa(short)
	repo.Blob.Blob = append(repo.Blob.Blob, EachBlob{
		shortString, url},
	)

	file, err := json.MarshalIndent(repo.Blob, "", "    ")
	if err != nil {
		return "", err
	}

	err = os.WriteFile("blob.json", file, 0644)
	if err != nil {
		return "", err
	}

	return shortString, nil
}
