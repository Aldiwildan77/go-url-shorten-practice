package links

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"url-shortener/users"
	"url-shortener/utils"

	"github.com/gorilla/mux"
)

var response utils.Responses

type LinksController struct {
	linksRepo LinksRepository
	usersRepo users.UsersRepository
}

func NewLinksController(linksRepo LinksRepository, usersRepo users.UsersRepository) LinksController {
	return LinksController{linksRepo, usersRepo}
}

func (c *LinksController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			c.Links(w, r)
		} else {
			c.Link(w, r)
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

func (c *LinksController) Shorten(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	links := c.linksRepo.FindByShortenLink(params["translated_url"])

	fmt.Println(links)

	http.Redirect(w, r, links.OriginalURL, http.StatusPermanentRedirect)
	return
}

func (c *LinksController) Links(w http.ResponseWriter, r *http.Request) {
	links := c.linksRepo.Links()

	var ll []Links

	for _, link := range links {
		ll = append(ll, link)
	}

	response.ResponseWithJSON(w, http.StatusOK, ll)
}

func (c *LinksController) Link(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	l := c.linksRepo.Link(id)

	response.ResponseWithJSON(w, http.StatusOK, l)
}

func (c *LinksController) Create(w http.ResponseWriter, r *http.Request) {
	var l Links
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	u := c.usersRepo.User(int(l.UserID))
	if u.ID == 0 {
		response.ResponseWithError(w, http.StatusNotFound, "Cant find user by id, please register first")
		return
	}

	l.TranslatedURL = c.linksRepo.GenerateShortenLink()

	link := c.linksRepo.Create(l)

	response.ResponseWithJSON(w, http.StatusCreated, link)
}

func (c *LinksController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var l Links
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	c.linksRepo.Update(id, l)
	link := c.linksRepo.Link(id)

	response.ResponseWithJSON(w, http.StatusOK, link)
}

func (c *LinksController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	ok := c.linksRepo.Delete(id)

	if ok {
		response.ResponseWithJSON(w, http.StatusOK, ok)
	}
}
