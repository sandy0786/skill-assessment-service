package service

import (
	"context"
	"log"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	roleDao "github.com/sandy0786/skill-assessment-service/dao/role"
	roleResponse "github.com/sandy0786/skill-assessment-service/response/role"

	"github.com/jinzhu/copier"
)

type RoleService interface {
	GetAllRoles(context.Context) ([]roleResponse.Role, error)
}

// authentication and authorisation servvice
type roleService struct {
	config configuration.ConfigurationInterface
	dao    roleDao.RoleDAO
}

func NewRoleService(c configuration.ConfigurationInterface, dao roleDao.RoleDAO) *roleService {
	return &roleService{
		config: c,
		dao:    dao,
	}
}

func (r *roleService) GetAllRoles(context.Context) ([]roleResponse.Role, error) {
	log.Println("Inside GetAllRoles")
	var rolesResponse []roleResponse.Role
	roles, err := r.dao.GetAllRoles()
	copier.Copy(&rolesResponse, &roles)
	return rolesResponse, err
}
