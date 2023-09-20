package handler

import (
	"file/models"
	"file/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStaff godoc
// @Router       /staff [POST]
// @Summary      CREATES STAFF
// @Description  CREATES STAFF BASED ON GIVEN DATA
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateStaff  true  "staff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaff(c *gin.Context) {
	var staff models.CreateStaff
	err := c.ShouldBind(&staff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Staff().CreateStaff(&staff)
	if err != nil {
		fmt.Println("error Staff Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "staff successfully created", "id": resp})
}

// GetStaff godoc
// @Router       /staff/{id} [GET]
// @Summary      GET BY ID
// @Description  get staff by ID
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Staff ID" format(uuid)
// @Success      200  {object}  models.Staff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaff(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Staff().GetStaff(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error staff Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})
}

// ListStaffes godoc
// @Router       /staff [GET]
// @Summary      GET  ALL STAFFS
// @Description  gets all staffs based on limit, page and search by name
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllStaff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaff(c *gin.Context) {
	h.log.Info("request GetALLstaff")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.storage.Staff().GetAllStaff(&models.GetAllStaffRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error  Getallstaff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to getAllStaff")
	c.JSON(http.StatusOK, resp)
}

// UpdateStaffs godoc
// @Router       /staff/{id} [PUT]
// @Summary      UPDATE STAFF BY ID
// @Description  UPDATES STAFF BASED ON GIVEN DATA AND ID
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff" format(uuid)
// @Param        data  body      models.CreateStaff  true  "staff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaff(c *gin.Context) {
	var staff models.Staff
	err := c.ShouldBind(&staff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	staff.ID = c.Param("id")
	resp, err := h.storage.Staff().UpdateStaff(&staff)
	if err != nil {
		h.log.Error("error Staff Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff successfully updated", "id": resp})
}

// DeleteStaff godoc
// @Router       /staff/{id} [DELETE]
// @Summary      DELETE STAFF BY ID
// @Description  DELETES STAFF BASED ON ID
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaff(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Staff().DeleteStaff(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting staff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff successfully deleted", "id": resp})

}

// func (h *Handler) ChangeBalance(c *gin.Context) {

// }
