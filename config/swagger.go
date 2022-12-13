package config

import (
	"fmt"

	"github.com/emicklei/go-restful"
	swagger "github.com/emicklei/go-restful-swagger12"
	"github.com/sirupsen/logrus"
)

const (
	apiPath         = "/apidocs"
	swaggerPath     = "/swagger/"
	swaggerFilePath = "config/swagger-ui/dist"
)

func (c *Config) RegisterSwagger(container *restful.Container) {
	logrus.Debugf("registing swagger: http://%s:%s", c.BindAddress, c.Port)
	config := swagger.Config{
		WebServices:    container.RegisteredWebServices(),
		WebServicesUrl: fmt.Sprintf("http://%s:%s", c.BindAddress, c.Port),
		ApiPath:        apiPath,

		// Optionally, specify where the UI is located
		SwaggerPath:     swaggerPath,
		SwaggerFilePath: swaggerFilePath,
	}
	swagger.RegisterSwaggerService(config, container)
}
