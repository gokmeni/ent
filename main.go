package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gokmeni/ent/ds"
	"github.com/gokmeni/ent/ent"
	"github.com/gokmeni/ent/ent/customer"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	httpServerPort string = ":6789"
)

const (
	getCustomerByIdRoute          string = "/api/customer/:customerID"
	getCustomerByIdRouteParameter string = "customerID"
	getCustomerByIdNotFoundError  string = "Customer Not Found."
	getCustomerByIdSystemError    string = "System Error."
)

var entClient *ent.Client

func main() {
	entClient = ds.GetConnection()

	defer closeConnection()

	startServer()
}

func closeConnection() {
	err := entClient.Close()

	log.Println("connection pool closing ...")

	if err != nil {
		fmt.Printf("connection close error: %v\n", err)
	}

	log.Println("connection pool closed.")
}

func startServer() {
	router := gin.Default()

	router.GET(getCustomerByIdRoute, getCustomerByID)

	srv := &http.Server{
		Addr:    httpServerPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shuting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}

	log.Println("server exited.")
}

func getCustomerByID(c *gin.Context) {
	customerID := c.Param(getCustomerByIdRouteParameter)

	result, err := entClient.Customer.
		Query().
		Where(customer.ID(customerID)).
		Only(c)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.String(http.StatusNotFound, getCustomerByIdNotFoundError)
		} else {
			fmt.Printf("system error -> %v\n", err)
			c.String(http.StatusForbidden, getCustomerByIdSystemError)
		}
	} else {
		c.JSON(http.StatusOK, result)
	}
}
