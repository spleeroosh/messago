package serverfx

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

type Option func(o *options)

type options struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	connState    func(net.Conn, http.ConnState)
	handler      http.Handler
	addr         string
	name         string
}

func Name(name string) Option {
	return func(s *options) {
		s.name = name
	}
}

func ConnectionState(f func(net.Conn, http.ConnState)) Option {
	return func(o *options) { o.connState = f }
}

func Handler(h http.Handler) Option {
	return func(o *options) { o.handler = h }
}

func ReadTimeout(t time.Duration) Option {
	return func(o *options) { o.readTimeout = t }
}

func WriteTimeout(t time.Duration) Option {
	return func(o *options) { o.writeTimeout = t }
}

func IdleTimeout(t time.Duration) Option {
	return func(o *options) { o.idleTimeout = t }
}

func Port(p int) Option {
	return func(o *options) { o.addr = ":" + strconv.Itoa(p) }
}
