package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pp "github.com/husamettinarabaci/go-packcal/core/application/presentation/port"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	dto "github.com/husamettinarabaci/go-packcal/pkg/presentation/dto"
)

type RestAPI struct {
	engine         *gin.Engine
	commandHandler pp.CommandPort
	queryHandler   pp.QueryPort
}

func NewRestAPI(qh pp.QueryPort, ch pp.CommandPort) *RestAPI {
	api := &RestAPI{
		commandHandler: ch,
		queryHandler:   qh,
	}
	api.engine = gin.New()
	api.engine.POST("/api/calc", api.Calculate)
	return api
}

func (api *RestAPI) Serve(debug bool, port string) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	api.engine.Run(":" + port)
}

func (api *RestAPI) Calculate(c *gin.Context) {
	var calcRequestDto dto.CalcRequest
	if err := c.ShouldBindJSON(&calcRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	calcRequestDto.FillPackSizes()
	if err := calcRequestDto.IsValid(); err == nil {
		calcResponse, err := api.commandHandler.Calculate(c, calcRequestDto.ToCalcRequestEntity())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.FromResponseObject(calcResponse.Response))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": mo.ErrInvalidInput.Error()})
	}
}
