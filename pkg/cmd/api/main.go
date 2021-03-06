package api

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/btwiuse/k8cc/pkg/apiserver/service"
	"github.com/btwiuse/k8cc/pkg/controller"
)

func RunApiServer(args []string) {
	var (
		fset          = flag.NewFlagSet("api", flag.ExitOnError)
		httpAddr      = fset.String("http.addr", ":8080", "HTTP listen address")
		kubeConfig    = fset.String("kube.config", "", "Kubeconfig path")
		kubeMasterURL = fset.String("kube.master-url", "", "Kubernetes master URL")
	)
	fset.Parse(args)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	sharedClient, err := controller.NewSharedClient(*kubeMasterURL, *kubeConfig)
	if err != nil {
		/* #nosec */
		_ = logger.Log("err", err)
		os.Exit(1)
	}

	var s service.Service
	{
		s = service.NewService(sharedClient, log.With(logger, "component", "Service"))
		s = service.LoggingMiddleware(logger)(s)
	}

	var h http.Handler
	{
		h = service.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error, 1)
	defer close(errs)
	stopCh := make(chan struct{})

	go func() {
		// handle signals
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		close(stopCh)
		errs <- fmt.Errorf("%s", sig)
	}()

	go func() {
		errs <- sharedClient.Run(stopCh)
	}()

	go func() {
		/* #nosec */
		_ = logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	// this last one takes ownership of the main goroutine
	//if err = operator.Run(2, stopCh); err != nil {
	//	errs <- err
	//}

	/* #nosec */
	_ = logger.Log("exit", <-errs)
}
