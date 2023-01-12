package controllers

import (
	"fmt"
	"net/http"
	"order-service/internal/api/dto"
	"order-service/internal/api/problems"
	models "order-service/internal/pkg/models/users"
	"order-service/internal/pkg/service"
	"strconv"

	"github.com/antonioalfa22/go-rest-template/pkg/crypto"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log" // zerolog for configs
)

type UsersControllerDelegate struct {
	//https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6
	userService service.UserService
}

func NewUsersControllerDelegate(userService service.UserService) *UsersControllerDelegate {
	return &UsersControllerDelegate{userService: userService}
}

func (s *UsersControllerDelegate) GetApiUsersId(c *gin.Context, id int) {
	userId := strconv.Itoa(id)
	if user, err := s.userService.Get(userId); err != nil {
		problem := problems.UserProblem{
			DefaultProblem: *problems.NewProblem(http.StatusNotFound, fmt.Sprintf("/users/%d", id), "User Not Found", fmt.Sprintf("L'utente con ID:%d non esiste", id)),
			Id:             id,
		}
		c.JSON(http.StatusInternalServerError, problem)
		return
	} else {
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("Errore conversione '/api/users/%d' response Dto", id))
			problem := *problems.NewProblem(http.StatusInternalServerError, fmt.Sprintf("/users/%d", id), "Response processing error", "Errore conversione payload risposta")
			c.JSON(http.StatusInternalServerError, problem)
			return
		}
		c.JSON(http.StatusOK, toUserDto(*user))
	}
}

func (s *UsersControllerDelegate) GetApiUsers(c *gin.Context, params dto.GetApiUsersParams) {
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.userService.Query(&q); err != nil {
		log.Error().Err(err).Msg("Errore recupero lista utenti /api/users")
		problem := *problems.NewProblem(http.StatusInternalServerError, fmt.Sprintf("/api/users"), "Response processing error", "Errore recupero lista utenti")
		c.JSON(http.StatusInternalServerError, problem)
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (s *UsersControllerDelegate) CreateUser(c *gin.Context) {
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		Role:      models.UserRole{RoleName: userInput.Role},
	}
	if err := s.userService.Add(&user); err != nil {
		log.Error().Err(err).Msg("Errore creazione utente '/api/users")
		problem := *problems.NewProblem(http.StatusInternalServerError, "/users/%d", "Response error", "Errore creazione utente")
		c.JSON(http.StatusInternalServerError, problem)
		return
	} else {
		if err != nil {
			// error
		}
		c.JSON(http.StatusCreated, user)
	}
}

func (s *UsersControllerDelegate) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.userService.Get(id); err != nil {
		//Gestione Errore
	} else {
		user.Username = userInput.Username
		user.Lastname = userInput.Lastname
		user.Firstname = userInput.Firstname
		user.Role = models.UserRole{RoleName: userInput.Role}
		if err := s.userService.Update(user); err != nil {
			//Gestione Errore
		} else {
			if err != nil {
				// error
			}
			c.JSON(http.StatusOK, user)
		}
	}
}

func (s *UsersControllerDelegate) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.userService.Get(id); err != nil {
		//Gestione errore
	} else {
		if err := s.userService.Delete(user); err != nil {
			//Gestione errore
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}
