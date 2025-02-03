package serverfx

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type ServerFX struct {
	opts options
	*http.Server
	name string
}

func New(appName string, opts ...Option) *ServerFX {
	o := options{
		name:         "http",
		readTimeout:  5 * time.Second,
		writeTimeout: 40 * time.Second,
		idleTimeout:  90 * time.Second,
		addr:         ":8081",
		handler:      http.HandlerFunc(defaultHandler),
	}

	for _, opt := range opts {
		opt(&o)
	}

	srv := &ServerFX{
		opts: o,
		Server: &http.Server{
			Addr:         o.addr,
			Handler:      o.handler,
			ReadTimeout:  o.readTimeout,
			WriteTimeout: o.writeTimeout,
			IdleTimeout:  o.idleTimeout,
			ConnState:    o.connState,
		},
		name: o.name,
	}
	if srv.Server.ConnState == nil {
		httpCount := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: strings.ReplaceAll(appName, "-", "_"),
				Subsystem: strings.ReplaceAll(o.name, "-", "_"),
				Name:      "count_of_tcp_connections",
				Help:      "Number of tcp connections",
			},
		)
		prometheus.MustRegister(httpCount)

		srv.Server.ConnState = func(c net.Conn, state http.ConnState) {
			// nolint
			switch state {
			case http.StateNew:
				httpCount.Inc()
			case http.StateHijacked, http.StateClosed:
				httpCount.Dec()
			}
		}
	}

	return srv
}

func (s *ServerFX) Start() error {
	fmt.Printf("http server started at port %s\n", s.opts.addr)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *ServerFX) Stop(ctx context.Context) error {
	fmt.Println("http server stopping")
	return s.Server.Shutdown(ctx)
}

func (s *ServerFX) Name() string {
	return s.name
}

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("Hello, you need a routing!\n")); err != nil {
		fmt.Println(err)
	}
}
