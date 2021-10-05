package seeds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

func (s Seed) UserSeed() error {
	s.dbConn.Exec("DELETE FROM users")

	for i := 0; i < 10; i++ {
		fakeUser := &dto.User{Name: faker.Name(), Email: faker.Email()}
		if err := s.dbConn.Create(database.ConvertToUser(fakeUser)).Error; err != nil {
			return err
		}
	}
	return nil
}
