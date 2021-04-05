package auth

import (
	"github.com/KenLu0929/flowControlTask/config/auth"
	"github.com/KenLu0929/flowControlTask/config/db"
	"github.com/KenLu0929/flowControlTask/pkg/models"
	"github.com/KenLu0929/flowControlTask/pkg/models/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var database = db.DB

// Login godoc
// @Summary login
// @Description login
// @Tags Auth
// @Version 1.0
// @ID login
// @Accept json
// @Produce json
// @Param input body object{account=string,password=string} true "account and password"
// @Success 200 {object} models.JsonOKResult
// @Failure 400 {object} models.JsonFailedResult
// @Router /login/ [post]
func Login(c *gin.Context) {
	var input struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	var userdata user.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	//check account and password
	check := database.Select("id", "username", "password").Where("account = ?", input.Account).Find(&userdata)

	if check.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   "account doesn't exist",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userdata.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   "wrong password",
		})
		return
	}

	token, err := auth.SignToken(input.Account, userdata.ID, userdata.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, models.Result{
		Success: true,
		Message: "login and sign token",
	})

}
