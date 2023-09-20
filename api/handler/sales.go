package handler

import (
	"file/models"
	"file/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSale godoc
// @Router       /sale [POST]
// @Summary      CREATES SALE
// @Description  CREATES SALE BASED ON GIVEN DATA
// @Tags         SALE
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateSales  true  "branch data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateSale(c *gin.Context) {
	var sale models.CreateSales
	err := c.ShouldBind(&sale)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Sales().CreateSale(&sale)
	if err != nil {
		fmt.Println("error sale Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "id": resp})
}

// GetSale godoc
// @Router       /sale/{id} [GET]
// @Summary      GET BY ID
// @Description  get sale by ID
// @Tags         SALE
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Sale ID" format(uuid)
// @Success      200  {object}  models.Sales
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetSale(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Sales().GetSale(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error sale get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})
}

// ListSales godoc
// @Router       /sale [GET]
// @Summary      GET  ALL SALES
// @Description  gets all sales based on limit, page and search by name
// @Tags         SALE
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllSalesResponse
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllSale(c *gin.Context) {
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

	resp, err := h.storage.Sales().GetAllSale(&models.GetAllSalesRequest{
		Page:        page,
		Limit:       limit,
		Client_name: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error sale getall:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateSale godoc
// @Router       /sale/{id} [PUT]
// @Summary      UPDATE SALE BY ID
// @Description  UPDATES SALE BASED ON GIVEN DATA AND ID
// @Tags         SALE
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of sale" format(uuid)
// @Param        data  body      models.Sales  true  "sale data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateSale(c *gin.Context) {
	var sale models.Sales
	err := c.ShouldBind(&sale)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sale.Id = c.Param("id")
	resp, err := h.storage.Sales().UpdateSale(&sale)
	if err != nil {
		h.log.Error("error sale update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sale successfully updated", "id": resp})
}

// DeleteSale godoc
// @Router       /sale/{id} [DELETE]
// @Summary      DELETE SALE BY ID
// @Description  DELETES SALE BASED ON ID
// @Tags         SALE
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of sale" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteSale(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Sales().DeleteSale(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting sale:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tariff successfully deleted", "id": resp})
}

// func (h *Handler) GetTopSaleBranch() {
// 	resp, err := h.strg.Sales().GetTopSaleBranch()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	branches, _ := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{})
// 	branchName := make(map[string]string)

// 	for _, v := range branches.Branches {
// 		branchName[v.Id] = v.Name
// 	}
// 	for _, structs := range resp {
// 		fmt.Printf("%s - %s - %f\n", structs.Day, branchName[structs.BranchId], structs.SalesAmount)
// 	}
// }

// func (h *Handler) CancelSale(id string) {
// 	resp, err := h.strg.Sales().CancelSale(models.IdRequest{Id: id})
// 	if err != nil {
// 		fmt.Println("error from CreateSale:", err.Error())
// 		return
// 	}

// 	// shop assistant change balance
// 	sale, err := h.strg.Sales().GetSale(models.IdRequest{Id: id})
// 	if err != nil {
// 		fmt.Println("error while read data", err)
// 		return
// 	}
// 	shopAsistant, err := h.strg.Staff().GetStaff(models.IdRequest{Id: sale.Shop_asissent_id})
// 	if err == nil {
// 		amount := 0.0
// 		tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: shopAsistant.TariffId})
// 		if err != nil {
// 			fmt.Println("error while get staff tarif")
// 			fmt.Println(err)
// 			return
// 		}

// 		if tarif.Type == config.Fixed {
// 			if sale.Payment_Type == 2 {
// 				amount = tarif.AmountForCash
// 			} else {
// 				amount = tarif.AmountForCard
// 			}
// 		} else if tarif.Type == config.Percent {
// 			if sale.Payment_Type == 2 {
// 				amount = sale.Price * tarif.AmountForCash / 100
// 			} else {
// 				amount = sale.Price * tarif.AmountForCard / 100
// 			}
// 		}

// 		// shop assitant uchun transaction
// 		_, err = h.strg.Transaction().CreateTransaction(models.CreateTransaction{
// 			Sale_id:     resp,
// 			Staff_id:    shopAsistant.Id,
// 			Type:        "shop_assistant",
// 			Source_type: "sales",
// 			Amount:      int(sale.Price),
// 			Text:        fmt.Sprintf("transcatiion cancelled, summa: %.2f", sale.Price),
// 		})
// 		if err != nil {
// 			fmt.Println("Error while create transaction")
// 			return
// 		}

// 		_, err = h.strg.Staff().ChangeBalance(models.ChangeBalance{Id: shopAsistant.Id, Balance: -amount})
// 		if err != nil {
// 			fmt.Println("Error while change balance")
// 			return
// 		}
// 	} else {
// 		fmt.Println("shopAssistant not found")
// 	}

// 	cashier, err := h.strg.Staff().GetStaff(models.IdRequest{Id: sale.Cashier_id})
// 	if err == nil {
// 		amount := 0.0
// 		tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: cashier.TariffId})
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		if tarif.Type == config.Fixed {
// 			if sale.Payment_Type == 2 {
// 				amount = tarif.AmountForCash
// 			} else {
// 				amount = tarif.AmountForCard
// 			}
// 		} else if tarif.Type == config.Percent {
// 			if sale.Payment_Type == 2 {
// 				amount = sale.Price * tarif.AmountForCash / 100
// 			} else {
// 				amount = sale.Price * tarif.AmountForCard / 100
// 			}
// 		}

// 		// cashier uchun transaction create qilish
// 		_, err = h.strg.Transaction().CreateTransaction(models.CreateTransaction{
// 			Sale_id:     resp,
// 			Staff_id:    cashier.Id,
// 			Type:        "cashier",
// 			Source_type: "sales",
// 			Amount:      int(sale.Price),
// 			Text:        fmt.Sprintf("transcatiion cancelled, summa: %.2f", sale.Price),
// 		})
// 		if err != nil {
// 			fmt.Println("Error while create transaction")
// 			return
// 		}
// 		// cashier change balance
// 		_, err = h.strg.Staff().ChangeBalance(models.ChangeBalance{Id: cashier.Id, Balance: -amount})
// 		if err != nil {
// 			fmt.Println("Error while change balance")
// 			return
// 		}
// 	} else {
// 		fmt.Println("error while get cashier data")
// 		return
// 	}

// 	fmt.Println(resp)
// }

// func (h *Handler) GetSaleCountBranch() {
// 	resp, err := h.strg.Sales().GetSaleCountBranch()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	branches, _ := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{})
// 	branchName := make(map[string]string)

// 	for _, v := range branches.Branches {
// 		branchName[v.Id] = v.Name
// 	}
// 	var sortedSlice []models.SaleCountSumBranch

// 	for _, structs := range resp {
// 		sortedSlice = append(sortedSlice, structs)
// 	}
// 	sort.Slice(sortedSlice, func(i, j int) bool {
// 		return sortedSlice[i].SalesAmount > sortedSlice[j].SalesAmount
// 	})

// 	for _, v := range sortedSlice {
// 		fmt.Printf("%s - %f - %d\n", branchName[v.BranchId], v.SalesAmount, v.Count)

// 	}

// }
