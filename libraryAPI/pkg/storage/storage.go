package storage

import "os"

type Storage struct {
	File *os.File
} 

func NewStorage(pathToFile string) (*Storage, error) {
	file, err := os.OpenFile(pathToFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &Storage{File: file}, nil
}

func (s *Storage) ClearPointer() error {
	_, err := s.File.Seek(0, 0)

	return err
}
