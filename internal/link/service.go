package link

import (
	"math/rand"
	"fmt"
)

type LinkService struct{
	Repository *LinkRepository
}

func NewLinkService (repository *LinkRepository) *LinkService{
	return &LinkService{
		Repository: repository,
	}
}

func (s *LinkService) AddLink(url string) (*Link, error){
	
	var exist = true
	var hash string
	var err error
	
	for exist{
		hash = getRandString(6)
		exist, err = s.Repository.IsHashExist(hash)
		
		if err != nil{
			return nil, fmt.Errorf("failed to check hash existence: %w", err)
		}
	}

	link := NewLink(url, hash)
	err = s.Repository.Create(link)
	
	return link, err

}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getRandString(n int) string {
	b := make([]byte, n)
    
	for i := range b {
        b[i] = alphabet[rand.Intn(len(alphabet))]
    }

	return string(b)
}