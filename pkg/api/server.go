package api

import (
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Serve(api *Api, address string) {

	// configure app routes
	handler := register(api)

	// don't let a panic crash the server.
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	// log all requests using logrus logger
	handler = handlers.LoggingHandler(
		logrus.WithField("prefix", "httpd").WriterLevel(logrus.InfoLevel),
		handler)

	server := &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadHeaderTimeout: 1 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	logrus.Infof("Start http server on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	logrus.Info("Server shutdown completed.")
}
