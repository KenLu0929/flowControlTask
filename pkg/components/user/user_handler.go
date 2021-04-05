package user

import (
	"github.com/KenLu0929/flowControlTask/config/db"
	"github.com/KenLu0929/flowControlTask/pkg/models"
	"github.com/KenLu0929/flowControlTask/pkg/models/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var database = db.DB

// CreateUser godoc
// @Summary create new user
// @Description create new user
// @Tags User
// @Version 1.0
// @ID create-user
// @Accept json
// @Produce json
// @Param data body user.User true "user data"
// @Success 201 {object} models.JsonOKResult
// @Failure 400 {object} models.JsonFailedResult
// @Router /user/ [post]
func CreateUser(c *gin.Context) {
	var userdata user.User

	if err := c.ShouldBindJSON(&userdata); err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	//check Acc
	findAccount := database.Where("account = ?", userdata.Account).Find(&userdata)
	if findAccount.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   "Account already exists",
		})
		return
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(userdata.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	userdata.Password = string(hash)

	//create
	if err := database.Create(&userdata).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Result{
		Success: true,
		Message: "Account created",
	})
}

// GetUser godoc
// @Summary get user data
// @Description get a user data by token
// @Tags User
// @Version 1.0
// @ID get-user
// @Produce json
// @Param Authorization header string true "token (Bearer+' '+token)"
// @Success 200 {object} models.JsonResult{payload=user.User}
// @Failure 400 {object} models.JsonFailedResult
// @Router /user/ [get]
func GetUser(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	var userdata user.User

	findID := database.Where("id = ?", userID).Find(&userdata)

	if findID.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   "id doesn't exist",
		})
		return
	}

	c.JSON(http.StatusOK, models.Result{
		Success: true,
		Payload: user.User{
			Account:  userdata.Account,
			Username: userdata.Username,
			Gender:   userdata.Gender,
			Birthday: userdata.Birthday,
		},
	})
}
