package portal

import (
	"net/http"

	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/util"
	"github.com/gin-gonic/gin"
)

// Me godoc
// @Summary Retrieve my information
// @Description Retrieve my information
// @id getMe
// @Tags User
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} MeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /me [get]
func (h Handler) Me(c *gin.Context) {
	const spanName = "meHandler"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	rs, err := h.userCtrl.Me(ctx)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.MeResponse{
		Data: view.Me{
			ID:       rs.ID,
			Email:    rs.Email,
			FullName: rs.FullName,
			Avatar:   rs.Avatar,
		},
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @id updateUser
// @Tags User
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param body body UpdateUserRequest true "Update user"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [put]
func (h Handler) UpdateUser(c *gin.Context) {
	const spanName = "updateUserHandler"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.UpdateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	rs, err := h.userCtrl.UpdateUser(
		ctx,
		model.UpdateUserRequest{
			FullName: req.FullName,
			Avatar:   req.Avatar,
		})
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.UserResponse{
		Data: view.User{
			ID:       rs.ID,
			Email:    rs.Email,
			FullName: rs.FullName,
			Avatar:   rs.Avatar,
		},
	})
}

// UpdatePassword godoc
// @Summary Update user's password
// @Description Update user's password
// @id updatePassword
// @Tags User
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param body body UpdatePasswordRequest true "Update user"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/password [put]
func (h Handler) UpdatePassword(c *gin.Context) {
	const spanName = "updatePasswordHandler"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.UpdatePasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	err = h.userCtrl.UpdatePassword(
		ctx,
		model.UpdatePasswordRequest{
			NewPassword: req.NewPassword,
			OldPassword: req.OldPassword,
		})
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.MessageResponse{
		Data: view.Message{
			Message: "success",
		},
	})
}
