package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net/http"
	"os"
	//"net/http"
)

type HealthController struct {
	*echo.Echo
}

func SetupHealthController(e *echo.Echo) {
	controller := &HealthController{e}
	fmt.Println("Setup Health controller...")

	e.GET("/health", controller.checkHealth)
}

func (pc *HealthController) checkHealth(c echo.Context) error {
	services := map[string]string{
		"post-service":    os.Getenv("POST_SERVICE_URL"),
		"comment-service": os.Getenv("COMMENT_SERVICE_URL"),
		"log-service":     os.Getenv("LOG_SERVICE_URL"),
		// Add more services as needed
	}

	healthCheckResult := make(map[string]string)

	for serviceName, serviceURL := range services {
		// Set up a connection to the server.
		conn, err := grpc.Dial(serviceURL, grpc.WithInsecure())
		if err != nil {
			log.Printf("Unable to connect to %s: %v", serviceName, err)
			healthCheckResult[serviceName] = "Unable to connect to service"
			continue
		}
		defer conn.Close()

		// Create a Health client
		healthClient := grpc_health_v1.NewHealthClient(conn)

		// Call Check method to check the health of the service
		response, err := healthClient.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			log.Printf("Health check failed for %s: %v", serviceName, err)
			healthCheckResult[serviceName] = "Health check failed"
			continue
		}

		// Save the health status to the healthCheckResult map
		healthCheckResult[serviceName] = response.GetStatus().String()

		// Print the health status
		//fmt.Printf("Service Health Status of %s: %s\n", serviceName, response.GetStatus())
	}

	return c.JSON(http.StatusOK, healthCheckResult)
}
