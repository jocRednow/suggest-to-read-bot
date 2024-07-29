package files

import "github.com/jocRednow/suggest-to-read-bot/storage"

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() {}()
	return err
}
