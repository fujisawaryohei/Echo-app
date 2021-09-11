package seeds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/fujisawaryohei/echo-app/database/dao"
	"github.com/fujisawaryohei/echo-app/web/dto"
)

func (s Seed) UserSeed() error {
	s.dbConn.Exec("DELETE FROM users")

	for i := 0; i < 10; i++ {
		var userDAO dao.User
		fakeUser := &dto.User{Name: faker.Name(), Email: faker.Email()}
		if err := s.dbConn.Create(userDAO.ConvertToDAO(fakeUser)).Error; err != nil {
			return err
		}
	}
	return nil
}
