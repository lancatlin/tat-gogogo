package controller

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
)

type loginController struct {
	studentID string
	password  string
}

/*
LoginController handle login related task
*/
type LoginController interface {
	Login() (*model.Result, error)
	LoginCurriculum() (bool, error)
}

/*
NewLoginController get a new LoginHandler
*/
func NewLoginController(studentID, password string) LoginController {
	return &loginController{
		studentID: studentID,
		password:  password,
	}
}

/*
Login will login the school system
*/
func (c *loginController) Login() (*model.Result, error) {
	loginResultRepo := repository.NewResultRepository()
	loginResultService := service.NewResultService(loginResultRepo)
	loginResultUsecase := usecase.NewResultUseCase(loginResultRepo, loginResultService)

	result, err := loginResultUsecase.LoginResult(c.studentID, c.password)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return result, nil
}

/*
LoginCurriculum will login school curriculum system
*/
func (c *loginController) LoginCurriculum() (bool, error) {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.LoginCurriculum()
}
