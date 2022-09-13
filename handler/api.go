package handler

import (
	"dealljobs/domain/user"
	"dealljobs/dto/request"
	"dealljobs/dto/response"
	"dealljobs/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ApiHandler struct {
	userSvc user.Service
}

func NewApiHandler(userSvc user.Service) *ApiHandler {
	return &ApiHandler{
		userSvc: userSvc,
	}
}

func (h *ApiHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	err := c.ShouldBindJSON(&req)

	u, signedToken, err := h.userSvc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not registered"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   signedToken,
		"message": "success",
	})
}

func (h *ApiHandler) CheckAuth(c *gin.Context) {
	userInfo, ok := c.Get("userInfo")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userinfo not filled"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    userInfo,
		"message": "success",
	})
}

func (h *ApiHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, err := h.userSvc.CreateUserIfNotAny(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err})
		return
	}

	resp := response.FromUserToResponse(u)
	c.JSON(http.StatusOK, gin.H{
		"data":    resp,
		"message": "success",
	})
}

func (h *ApiHandler) UpdateUser(c *gin.Context) {
	var req request.UpdateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.userSvc.UpdateUser(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *ApiHandler) DeleteUser(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.userSvc.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *ApiHandler) GetUser(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	currentUserInfo, err := auth.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	currentId := int(currentUserInfo["ID"].(float64))
	if currentUserInfo["Role"] != string(user.RoleAdmin) && currentId != id {
		c.JSON(http.StatusForbidden, gin.H{"message": "you are forbidden to access this user data"})
		return
	}

	u := h.userSvc.GetUser(id)
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "this data does not exist"})
		return
	}

	resp := response.FromUserToResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data":    resp,
		"message": "success",
	})
}

func (h *ApiHandler) GetAllUsers(c *gin.Context) {

	users := h.userSvc.GetAllUsers()

	responses := make([]*response.UserResponse, 0)
	for _, u := range users {
		responses = append(responses, response.FromUserToResponse(u))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    responses,
		"message": "success",
	})
}
