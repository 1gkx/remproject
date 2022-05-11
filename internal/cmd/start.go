package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"remproject/internal/transport"

	"github.com/go-kit/kit/log"
	"github.com/urfave/cli/v2"
)

var Start = func(c *cli.Context) error {

	_, cancel := context.WithCancel(c.Context)
	defer cancel()

	logger := log.NewLogfmtLogger(os.Stdout)

	httpAddr := fmt.Sprintf(":%s", "3000")

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%v", <-c)
	}()

	go func() {
		logger.Log("event", "Server starting...", "address", httpAddr)
		server := &http.Server{
			Addr:    httpAddr,
			Handler: transport.New(logger),
		}
		errs <- server.ListenAndServe()
	}()

	logger.Log("exit: %v\n", <-errs)

	return nil
}
