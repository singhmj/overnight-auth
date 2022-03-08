package api

import (
	"auth/db/models"
	"auth/utils"
	"net/http"
	"time"

	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
)

type createProfileRequest struct {
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	Dob            time.Time `json:"dob" binding:"required"`
	AddressLine1   string    `json:"address_line1"`
	AddressLine2   string    `json:"address_line2"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	PostalCode     string    `json:"postal_code"`
	PrimaryPhone   string    `json:"primary_phone"`
	SecondaryPhone string    `json:"secondary_phone"`
	PrimaryEmail   string    `json:"primary_email"`
	SecondaryEmail string    `json:"secondary_email"`
}

func (server *Server) createProfile(ctx *gin.Context) {
	req := createProfileRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// TODO: fetch id from payload req
	profile, err := server.store.CreateProfile(ctx, models.CreateProfileParams{
		UserID:         uuid.Generate().String(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Dob:            req.Dob,
		AddressLine1:   req.AddressLine1,
		AddressLine2:   req.AddressLine2,
		City:           req.City,
		State:          req.State,
		Country:        req.Country,
		PostalCode:     req.PostalCode,
		PrimaryPhone:   req.PrimaryPhone,
		SecondaryPhone: req.SecondaryPhone,
		PrimaryEmail:   req.PrimaryEmail,
		SecondaryEmail: req.SecondaryEmail,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

type getProfileRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getProfile(ctx *gin.Context) {
	req := getProfileRequest{}
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	profile, err := server.store.GetProfile(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

type listProfilesRequest struct {
	Offset uint16
	Count  uint16
}

func (server *Server) listProfiles(ctx *gin.Context) {
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	req := listProfilesRequest{}
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	profiles, err := server.store.ListProfiles(ctx, models.ListProfilesParams{
		Limit:  int32(req.Count),
		Offset: int32(req.Offset),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profiles)
}

type updateProfileRequest struct {
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	Dob            time.Time `json:"dob" binding:"required"`
	AddressLine1   string    `json:"address_line1"`
	AddressLine2   string    `json:"address_line2"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	PostalCode     string    `json:"postal_code"`
	PrimaryPhone   string    `json:"primary_phone"`
	SecondaryPhone string    `json:"secondary_phone"`
	PrimaryEmail   string    `json:"primary_email"`
	SecondaryEmail string    `json:"secondary_email"`
}

func (server *Server) updateProfile(ctx *gin.Context) {
	req := updateProfileRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// get id from the uri
	updateProfileReq := models.UpdateProfileParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Dob:            req.Dob,
		AddressLine1:   req.AddressLine1,
		AddressLine2:   req.AddressLine2,
		City:           req.City,
		State:          req.State,
		Country:        req.Country,
		PostalCode:     req.PostalCode,
		PrimaryPhone:   req.PrimaryPhone,
		SecondaryPhone: req.SecondaryPhone,
		PrimaryEmail:   req.PrimaryEmail,
		SecondaryEmail: req.SecondaryEmail,
		UpdatedAt:      time.Now().UTC(),
	}
	err := server.store.UpdateProfile(ctx, updateProfileReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// func (server *Server) deleteProfile(ctx *gin.Context) {
// 	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// }
