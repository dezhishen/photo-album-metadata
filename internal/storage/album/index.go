package storage

import "github.com/dezhishen/photo-album-metadata/pkg/model"

func Query(name string, pageNum, pageSize int) ([]model.Album, int64, error) {
	return nil, 0, nil
}

func Get(id int64) (*model.Album, error) {
	return nil, nil
}

func List(name string) ([]model.Album, error) {
	return nil, nil
}
