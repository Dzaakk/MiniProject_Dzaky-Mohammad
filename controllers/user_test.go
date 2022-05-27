package controllers

import (
	"MiniProject/models"
	"MiniProject/service"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPersonControllerAdd(t *testing.T) {
	e := echo.New()

	newPersonJson, _ := json.Marshal(map[string]string{
		"Username": "test1",
		"Password": "pass",
		"Name":     "usertest1",
		"Email":    "usertest@gmail.com",
	})
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPersonJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	ps := service.NewMockPersonService()
	pc := NewPersonController(ps)
	pc.Add(c)

	users, err := ps.Get()
	if err != nil {
		t.Error(err)
	}
	if len(users) != 1 {
		t.Errorf("Expecting len(persons) to be 1, get %d", len(users))
	}
	if users[0].Username != "test1" {
		t.Errorf("Expecting users[0].Username to be test1, get %s", users[0].Username)
	}
	if users[0].Password != "pass" {
		t.Errorf("Expecting users[0].Password to be pass, get %s", users[0].Password)
	}
	if users[0].Name != "usertest1" {
		t.Errorf("Expecting users[0].Name to be usertest1, get %s", users[0].Name)
	}
	if users[0].Email != "usertest@gmail.com" {
		t.Errorf("Expecting users[0].Email to be usertest@gmail.com, get %s", users[0].Email)
	}

}

type PersonController struct {
	ps service.PersonService
}

func (pc PersonController) Get(c echo.Context) error {
	persons, err := pc.ps.Get()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get persons"})
	}
	return c.JSON(http.StatusOK, persons)
}

func (pc PersonController) Add(c echo.Context) error {
	var newPerson models.Users
	c.Bind(&newPerson)
	person, err := pc.ps.Add(newPerson)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get persons"})
	}
	return c.JSON(http.StatusOK, person)
}

func NewPersonController(ps service.PersonService) PersonController {
	return PersonController{
		ps: ps,
	}
}
