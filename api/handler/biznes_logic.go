package handler

import (
	"file/models"
	"file/pkg/logger"
	"fmt"
	"net/http"

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
func (h *Handler) TopStaff(c *gin.Context) {
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
