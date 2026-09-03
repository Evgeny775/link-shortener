package link

import (
	"errors"
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

type LinkService struct {
	Repository *LinkRepository
}

func NewLinkService(repository *LinkRepository) *LinkService {
	return &LinkService{
		Repository: repository,
	}
}

func (s *LinkService) AddLink(url string) (*Link, error) {

	var link *Link
	var hash string
	var err error

	for i := 0; i < 10; i++ {
		hash = getRandString(6)
		link = NewLink(url, hash)

		err = s.Repository.Create(link)

		if err == nil {
			return link, nil
		}

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			continue
		}

		return nil, fmt.Errorf("failed to save link: %w", err)
	}

	return nil, fmt.Errorf("failed to generate unique hash after 10 attempts, last error: %w", err)

}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getRandString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}

func (s *LinkService) UpdateLink(body *LinkUpdateRequest, id int) (*Link, error) {

	hashExist, err := s.Repository.HashExistExeptID(body.Hash, id)
	if err != nil {
		return nil, err
	}

	if hashExist {
		return nil, hashAlreadyExist
	}

	newLink := &Link{
		Model: gorm.Model{ID: uint(id)},
		Url:   body.URL,
		Hash:  body.Hash,
	}

	err = s.Repository.Update(newLink)

	return newLink, err
}

func (s *LinkService) DeleteLink(id int) error {
	idExist, err := s.Repository.IDExist(id)
	if err != nil {
		return err
	}

	if idExist {
		err = s.Repository.Delete(id)
		return err
	}

	return noSuchId
}

func (s *LinkService) GetByHash(hash string) (*Link, error) {
	return s.Repository.GetByHash(hash)

}
