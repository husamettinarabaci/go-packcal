package presentation

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	pp "github.com/husamettinarabaci/go-packcal/core/application/presentation/port"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	dto "github.com/husamettinarabaci/go-packcal/pkg/presentation/dto"
)

type WebServer struct {
	engine         *gin.Engine
	commandHandler pp.CommandPort
	queryHandler   pp.QueryPort
}

func NewWebServer(qh pp.QueryPort, ch pp.CommandPort) *WebServer {
	srv := &WebServer{
		commandHandler: ch,
		queryHandler:   qh,
	}
	srv.engine = gin.New()
	isLocal := false
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			isLocal = true
			srv.engine.LoadHTMLGlob("pkg/presentation/controller/web/views/*")
		}
	}
	if !isLocal {
		srv.engine.LoadHTMLGlob("views/*")
	}
	srv.engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	srv.engine.POST("/calculate", srv.Calculate)
	return srv
}

func (srv *WebServer) Calculate(c *gin.Context) {
	var calcRequestDto dto.CalcRequest
	if err := c.Bind(&calcRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"results": dto.CalcResponse{}})
		return
	}
	calcRequestDto.FillPackSizes()
	if err := calcRequestDto.IsValid(); err == nil {
		calcResponse, err := srv.commandHandler.Calculate(c, calcRequestDto.ToCalcRequestEntity())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"results": dto.CalcResponse{}})
			return
		}
		results := dto.FromResponseObject(calcResponse.Response)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"error":   "",
			"results": results,
		})
	} else {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error":   mo.ErrInvalidInput.Error(),
			"results": dto.CalcResponse{},
		})
	}
}

func (srv *WebServer) Serve(debug bool, port string) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	srv.engine.Run(":" + port)
}
