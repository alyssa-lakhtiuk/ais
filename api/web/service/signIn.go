package service

import (
	"ais/entities"
	"ais/repository"
)

type roleService struct {
	repo repository.RoleRepo
}

func NewRoleService(repo repository.RoleRepo) *roleService {
	return &roleService{repo: repo}
}

func (r *roleService) GetByPhone(phone string) (entities.SignIn, error) {
	return r.repo.GetRoleByPhone(phone)
}

func (r *roleService) GetByIdEmployee(id string) (entities.SignIn, error) {
	return r.repo.GetRoleByIdEmployee(id)
}
