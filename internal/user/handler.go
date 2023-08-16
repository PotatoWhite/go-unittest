package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func NewHandler(db *gorm.DB) Handler {
	userRepo := NewRepository(db)
	userQuery := NewQuery(userRepo)
	userCommand := NewCommand(userRepo)

	return &handler{
		query:   userQuery,
		command: userCommand,
	}
}

type Handler interface {
	RouteGroup(prefix string, r *gin.Engine) *gin.RouterGroup
	GetUserById(c *gin.Context)
	AddUser(c *gin.Context)
}

type handler struct {
	query   Query
	command Command
}

func (h handler) RouteGroup(prefix string, r *gin.Engine) *gin.RouterGroup {
	userRouter := r.Group(prefix)
	userRouter.GET("/:id", h.GetUserById)
	userRouter.POST("/", h.AddUser)

	return userRouter
}

func (h handler) GetUserById(c *gin.Context) {
	_id := c.Param("id")
	if _id == "" {
		handleBadRequest(c, errors.New("id is required"))
		return
	}

	// to uuid
	id := uuid.MustParse(_id)

	user, err := h.query.GetUserById(c, id)
	if err != nil {
		handleInternalError(c, err)
		return
	}

	c.JSON(200, user)
}

func (h handler) AddUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleBadRequest(c, err)
		return
	}

	// if id is not empty, create new uuid
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	if user.CreateAt.IsZero() {
		user.CreateAt = time.Now()
	}

	if err := h.command.AddUser(c, user); err != nil {
		handleInternalError(c, err)
		return
	}

	c.JSON(200, ToUserDto(&user))
}

func ToUserDto(user *User) UserDto {
	return UserDto{
		ID:       user.ID,
		Name:     user.Name,
		CreateAt: user.CreateAt.Format(time.RFC3339),
	}
}
