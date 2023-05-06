package delivery

import (
	"fmt"
	"profile/internal/service"
	"profile/pkg/tools/tokenManager"

	"github.com/gin-gonic/gin"
	"github.com/dmitryavdonin/gtools/logger"
)

type Delivery struct {
	services *service.Service
	router   *gin.Engine
	logger   logger.Interface
	port     int
	tm       tokenManager.TokenManager

	options Options
}

type Options struct{}

func New(services *service.Service, tm tokenManager.TokenManager, port int, logger logger.Interface, options Options) (*Delivery, error) {

	var d = &Delivery{
		services: services,
		logger:   logger,
		port:     port,
		tm:       tm,
	}

	d.SetOptions(options)

	d.router = d.initRouter()
	return d, nil
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf(":%d", d.port))
}
