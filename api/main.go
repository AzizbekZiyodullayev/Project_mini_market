package api

import (
	_ "file/api/docs"
	"file/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branch", h.GetAllBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/tariff", h.CreateStaffTarif)
	r.GET("/tariff/:id", h.GetStaffTarif)
	r.GET("/tariff", h.GetAllStaffTarif)
	r.PUT("/tariff/:id", h.UpdateStaffTarif)
	r.DELETE("/tariff/:id", h.DeleteStaffTarif)

	r.POST("/staff", h.CreateStaff)
	r.GET("/staff/:id", h.GetStaff)
	r.GET("/staff", h.GetAllStaff)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetSale)
	r.GET("/sale", h.GetAllSale)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)

	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransaction)
	r.GET("/transaction", h.GetAllTransaction)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
