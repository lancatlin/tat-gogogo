package usecase

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"

	"log"
)

/*
CurriculumUseCase contains the functions for curriculum usecase
*/
type CurriculumUseCase interface {
	GetCurriculums(targetStudentID string) ([]model.Curriculum, error)
	LoginCurriculum() (bool, error)
	IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool
}

type curriculumUseCase struct {
	repo    repository.CurriculumRepository
	service *service.CurriculumService
}

/*
NewCurriculumUseCase init a new curriculum usecase
*/
func NewCurriculumUseCase(repo repository.CurriculumRepository, service *service.CurriculumService) CurriculumUseCase {
	return &curriculumUseCase{repo: repo, service: service}
}

/*
LoginCurriculum login curriculum system
*/
func (c *curriculumUseCase) LoginCurriculum() (bool, error) {
	return c.service.IsLoginCurriculum()
}

/*
IsSameYearAndSemBy judge is same year and semester
*/
func (c *curriculumUseCase) IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool {
	for _, curriculum := range curriculums {
		if curriculum.Year == year && curriculum.Semester == semester {
			return true
		}
	}
	return false
}

/*
GetCurriculums get []model.Curriculum
*/
func (c *curriculumUseCase) GetCurriculums(targetStudentID string) ([]model.Curriculum, error) {
	doc, err := c.service.GetCurriculumDocument(targetStudentID)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return c.repo.ParseCurriculums(doc), nil
}
