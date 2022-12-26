package artist

import "github.com/dezhishen/photo-album-metadata/pkg/model"

func Query(name string, pageNum, pageSize int) ([]model.Artist, int64, error) {
	return nil, 0, nil
}

func Get(id int64) (*model.Artist, error) {
	return nil, nil
}

func List(name string) ([]model.Artist, error) {
	return nil, nil
}
