package api

import (
	"auth/db/models"
	"auth/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	req := createUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// TODO: fetch id from payload req
	user, err := server.store.CreateUser(ctx, models.CreateUserParams{
		ID:        uuid.Generate().String(),
		LoginID:   req.LoginID,
		Password:  req.Password,
		Status:    "active", // TODO:
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type loginUserRequest struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	req := loginUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.LoginID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	// TODO: compare hashes
	if user.Password != req.Password {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("invalid password")))
	}

	// TODO: return login token
	ctx.JSON(http.StatusOK, user)
}

type verifyUserRequest struct {
	LoginID string `json:"login_id" binding:"required"`
}

func (server *Server) verifyUser(ctx *gin.Context) {
	req := verifyUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.LoginID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	err = server.store.UpdateUserStatus(ctx, models.UpdateUserStatusParams{
		ID:        user.ID,
		Status:    "verified",
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	req := getUserRequest{}
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	Offset uint16
	Count  uint16
}

func (server *Server) listUsers(ctx *gin.Context) {
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	req := listUsersRequest{}
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	users, err := server.store.ListUsers(ctx, models.ListUsersParams{
		Limit:  int32(req.Count),
		Offset: int32(req.Offset),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type updateUserPasswordRequest struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) updateUserPassword(ctx *gin.Context) {
	req := updateUserPasswordRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// get id from the uri
	updateUserPasswordReq := models.UpdateUserPasswordParams{
		ID:        req.LoginID,
		Password:  req.Password,
		UpdatedAt: time.Now().UTC(),
	}
	err := server.store.UpdateUserPassword(ctx, updateUserPasswordReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type updateUserStatusRequest struct {
	LoginID string `json:"login_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

func (server *Server) updateUserStatus(ctx *gin.Context) {
	req := updateUserStatusRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// get id from the uri
	updateUserStatusReq := models.UpdateUserStatusParams{
		ID:        req.LoginID,
		Status:    req.Status,
		UpdatedAt: time.Now().UTC(),
	}
	err := server.store.UpdateUserStatus(ctx, updateUserStatusReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// func (server *Server) deleteuser(ctx *gin.Context) {
// 	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// }
