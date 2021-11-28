package seeds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

func (s Seed) PostSeed() error {
	s.dbConn.Exec("DELETE FROM posts")

	for i := 0; i < 10; i++ {
		fakerPost := &dto.Post{Title: faker.Name(), Body: faker.Name(), Published: &[]bool{true}[0]}
		if err := s.dbConn.Create(fakerPost).Error; err != nil {
			return err
		}
	}
	return nil
}
