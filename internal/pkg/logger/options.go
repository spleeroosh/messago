package logger

import (
	"io"
	"time"
)

type Option func(*options)

type options struct {
	level             Level
	env               string
	buildCommit       string
	buildTime         time.Time
	noTimestamp       bool
	writer            io.Writer
	prettify          bool
	overrideStdLogOut bool
}

func WithLevel(lvl Level) Option {
	return func(o *options) {
		o.level = lvl
	}
}

func WithEnv(env string) Option {
	return func(o *options) {
		o.env = env
	}
}

func WithBuildCommit(commit string) Option {
	return func(o *options) {
		o.buildCommit = commit
	}
}

func WithBuildTime(t time.Time) Option {
	return func(o *options) {
		o.buildTime = t
	}
}

func WithNoTimestamp(no bool) Option {
	return func(o *options) {
		o.noTimestamp = no
	}
}

func WithWriter(w io.Writer) Option {
	return func(o *options) {
		o.writer = w
	}
}

func WithPrettify(prettify bool) Option {
	return func(o *options) {
		o.prettify = prettify
	}
}

func WithOverrideStdLogOut(override bool) Option {
	return func(o *options) {
		o.overrideStdLogOut = override
	}
}
