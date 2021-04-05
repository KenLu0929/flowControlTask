package post

import (
	"github.com/KenLu0929/flowControlTask/config/db"
	"github.com/KenLu0929/flowControlTask/pkg/models"
	"github.com/KenLu0929/flowControlTask/pkg/models/post"

	"net/http"

	"github.com/gin-gonic/gin"
)

var database = db.DB

// CreatePost godoc
// @Summary create new post
// @Description create new post by token
// @Tags Post
// @Version 1.0
// @ID create-post
// @Accept json
// @Produce json
// @Param Authorization header string true "token (Bearer+' '+token)"
// @Param data body object{title=string,content=string} true "post"
// @Success 201 {object} models.JsonOKResult
// @Failure 400 {object} models.JsonFailedResult
// @Router /post/ [post]
func CreatePost(c *gin.Context) {
	var postdata post.Post
	postdata.UserID = c.MustGet("id").(uint)

	if err := c.ShouldBindJSON(&postdata); err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	//create
	if err := database.Create(&postdata).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Result{
		Success: true,
		Message: "post created",
	})
}

// GetPost godoc
// @Summary get user post
// @Description get  user post by token
// @Tags Post
// @Version 1.0
// @ID get-post
// @Produce json
// @Param Authorization header string true "token (Bearer+' '+token)"
// @Success 200 {object} models.JsonResult{payload=[]post.Post}
// @Failure 400 {object} models.JsonFailedResult
// @Router /post/ [get]
func GetUserPost(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	var postList []post.Post
	var count int64

	find := database.Model(&post.Post{}).Where("user_id = ?", userID)
	find.Count(&count)

	if count == 0 {
		c.JSON(http.StatusBadRequest, models.Result{
			Success: false,
			Error:   "no any post",
		})
		return
	}

	find.Scan(&postList)

	c.JSON(http.StatusOK, models.Result{
		Success: true,
		Payload: postList,
	})
}
