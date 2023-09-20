package handler

import (
	"file/models"
	"file/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @Router       /branch [POST]
// @Summary      CREATES BRANCH
// @Description  CREATES BRANCH BASED ON GIVEN DATA
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateBranch  true  "branch data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateBranch(c *gin.Context) {
	var branch models.CreateBranch
	err := c.ShouldBind(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Branch().CreateBranch(&branch)
	if err != nil {
		fmt.Println("error Branch Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Branch successfully created", "id": resp})
}

// GetBranch godoc
// @Router       /branch/{id} [GET]
// @Summary      GET BY ID
// @Description  get branch by ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Branch ID" format(uuid)
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Branch().GetBranch(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Branch Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListBranchs godoc
// @Router       /branch [GET]
// @Summary      GET  ALL BRANCHS
// @Description  gets all branch based on limit, page and search by name
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllBranch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllBranch(c *gin.Context) {
	h.log.Info("request GetAllBranch")
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

	resp, err := h.storage.Branch().GetAllBranch(&models.GetAllBranchRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Branch GetAllBranch:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllBranch")
	c.JSON(http.StatusOK, resp)
}

// UpdateBranch godoc
// @Router       /branch/{id} [PUT]
// @Summary      UPDATE BRANCH BY ID
// @Description  UPDATES BRANCH BASED ON GIVEN DATA AND ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch" format(uuid)
// @Param        data  body      models.CreateBranch  true  "branch data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateBranch(c *gin.Context) {
	var branch models.Branch
	err := c.ShouldBind(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	branch.ID = c.Param("id")
	resp, err := h.storage.Branch().UpdateBranch(&branch)
	if err != nil {
		h.log.Error("error Branch Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update branch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch successfully updated", "id": resp})
}

// DeleteBranch godoc
// @Router       /branch/{id} [DELETE]
// @Summary      DELETE BRANCH BY ID
// @Description  DELETES BRANCH BASED ON ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Branch().DeleteBranch(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting branch:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete branch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch successfully deleted", "id": resp})
}
