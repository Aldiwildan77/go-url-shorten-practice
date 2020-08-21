package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"url-shortener/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var response utils.Responses

type UsersController struct {
	repo UsersRepository
}

func NewUsersController(repo UsersRepository) UsersController {
	return UsersController{repo}
}

func (c *UsersController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			c.Users(w, r)
		} else {
			c.User(w, r)
		}
	case http.MethodPost:
		c.Create(w, r)
	case http.MethodPut:
		c.Update(w, r)
	case http.MethodDelete:
		c.Delete(w, r)
	default:
		response.ResponseWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (c *UsersController) Users(w http.ResponseWriter, r *http.Request) {
	users := c.repo.Users()

	var uu []interface{}

	for _, user := range users {
		uResult := map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}
		uu = append(uu, uResult)
	}

	response.ResponseWithJSON(w, http.StatusOK, uu)
}

func (c *UsersController) User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	u := c.repo.User(id)

	uResult := map[string]interface{}{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
	}

	response.ResponseWithJSON(w, http.StatusOK, uResult)
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var u Users
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := c.repo.Create(u)

	response.ResponseWithJSON(w, http.StatusCreated, user)
}

func (c *UsersController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var u Users
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	c.repo.Update(id, u)
	user := c.repo.User(id)

	response.ResponseWithJSON(w, http.StatusOK, user)
}

func (c *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	ok := c.repo.Delete(id)

	if ok {
		response.ResponseWithJSON(w, http.StatusOK, ok)
	}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (c *UsersController) SignIn(w http.ResponseWriter, r *http.Request) {
	var credential Credentials

	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, "Your request is rejected because the body is wrong")
		return
	}

	user := c.repo.FindUserByEmail(credential.Email)
	if user.ID == 0 {
		response.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized account")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password)); err != nil {
		response.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	expTime := time.Now().Add(5 * time.Minute)

	// Set Payload
	claims := &Claims{
		Email: credential.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_KEY")))
	if err != nil {
		response.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})
}

func (c *UsersController) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
