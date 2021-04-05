package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8/internal"
	"github.com/stretchr/testify/assert"
)

//acc = "tokentest" pass = "tokentest"
var testToken string = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoidG9rZW50ZXN0IiwiaWQiOjEsInVzZXJuYW1lIjoidG9rZW50ZXN0IiwiYXVkIjoidG9rZW50ZXN0IiwianRpIjoidG9rZW50ZXN0dG9rZW50ZXN0MTYxNzYxNjgwNCIsImlhdCI6MTYxNzYxNjgwNCwibmJmIjoxNjE3NjE2ODA0LCJzdWIiOiJ0b2tlbnRlc3QifQ.kYlq-ohTZn3EeHz9msuG9boxOU4ypymWa3zqe00Jgto"

func TestUserCreated(t *testing.T) {
	reqBody := "{\"account\":\"test\",\"password\":\"test\",\"username\":\"test\"}"

	router := SetRouter()
	r, _ := http.NewRequest("POST", "/api/user/", bytes.NewBuffer([]byte(reqBody)))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestUserCreatedFailAccountExists(t *testing.T) {
	reqBody := "{\"account\":\"test1\",\"password\":\"test1\",\"username\":\"test1\"}"

	router := SetRouter()
	r, _ := http.NewRequest("POST", "/api/user/", bytes.NewBuffer([]byte(reqBody)))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Nil(t, err)
	assert.Equal(t, "Account already exists", response["error"])
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestGetUserOk(t *testing.T) {
	payload := map[string]interface{}{
		"account":  "tokentest",
		"birthday": nil,
		"gender":   "",
		"password": "",
		"username": "tokentest",
	}

	router := SetRouter()
	r, _ := http.NewRequest("GET", "/api/user/", nil)
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	r.Header.Set("Authorization", testToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	token := w.HeaderMap.Get("Authorization")
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.Equal(t, payload, response["payload"])
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestLoginOk(t *testing.T) {
	reqBody := "{\"account\":\"test1\",\"password\":\"test1\"}"

	router := SetRouter()
	r, _ := http.NewRequest("POST", "/api/login/", bytes.NewBuffer([]byte(reqBody)))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	token := w.HeaderMap.Get("Authorization")
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestLoginAccountFail(t *testing.T) {
	reqBody := "{\"account\":\"testerr\",\"password\":\"test\"}"

	router := SetRouter()
	r, _ := http.NewRequest("POST", "/api/login/", bytes.NewBuffer([]byte(reqBody)))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	token := w.HeaderMap.Get("Authorization")
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, token)
	assert.Nil(t, err)
	assert.Equal(t, "account doesn't exist", response["error"])
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestLoginPasswordFail(t *testing.T) {
	reqBody := "{\"account\":\"test1\",\"password\":\"test123\"}"

	router := SetRouter()
	r, _ := http.NewRequest("POST", "/api/login/", bytes.NewBuffer([]byte(reqBody)))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	token := w.HeaderMap.Get("Authorization")
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, token)
	assert.Nil(t, err)
	assert.Equal(t, "wrong password", response["error"])
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestGetPostFailNoExists(t *testing.T) {
	router := SetRouter()
	r, _ := http.NewRequest("GET", "/api/post/", nil)
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	r.Header.Set("Authorization", testToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	token := w.HeaderMap.Get("Authorization")
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestPostCreated(t *testing.T) {
	reqBody := "{\"title\":\"testTitle\",\"content\":\"testContent\"}"

	router := SetRouter()
	r, _ := http.NewRequest("POST", "/api/post/", bytes.NewBuffer([]byte(reqBody)))
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	r.Header.Set("Authorization", testToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	token := w.HeaderMap.Get("Authorization")
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
}

func TestGetPostOk(t *testing.T) {
	payload := []interface{}{map[string]interface{}{
		"content": "testContent",
		"title":   "testTitle",
		"user_id": float64(1),
	}}

	router := SetRouter()
	r, _ := http.NewRequest("GET", "/api/post/", nil)
	r.Header.Set("X-Forwarded-For", "0.0.0.0")
	r.Header.Set("Authorization", testToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	token := w.HeaderMap.Get("Authorization")
	count := w.HeaderMap.Get("X-Ratelimit-Remaining")
	reset := w.HeaderMap.Get("X-Ratelimit-Reset")

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, reset)
	assert.Nil(t, err)
	assert.Equal(t, payload, response["payload"])
}
