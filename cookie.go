package session

import (
	"time"
)

type Cookie struct {
	Name string

	Value string

	Opts *cookieOpts
}

type cookieOpts struct {
	Secure bool

	HTTPOnly bool

	Domain string

	Path string

	SameSite SameSite

	MaxAge time.Duration

	Expiry time.Time
}

type Options func(*cookieOpts)

type SameSite string

const (
	SameSiteLax    SameSite = "lax"
	SameSiteNone   SameSite = "none"
	SameSiteStrict SameSite = "strict"
)

var defaultCookieOpts = &cookieOpts{
	Secure: false,

	HTTPOnly: false,

	Domain: "",

	Path: "/",

	SameSite: SameSiteLax,

	MaxAge: 0,

	Expiry: time.Time{},
}

func NewCookie(name, value string, opts ...Options) *Cookie {
	defaultOpts := defaultCookieOpts
	cookie := &Cookie{
		Name:  name,
		Value: value,
		Opts:  defaultOpts,
	}

	for _, opt := range opts {
		opt(defaultOpts)
	}

	return cookie
}

func SetSecure() Options {
	return func(co *cookieOpts) {
		co.Secure = true
	}
}

func SetHTTPOnly() Options {
	return func(co *cookieOpts) {
		co.HTTPOnly = true
	}
}

func SetDomain(domain string) Options {
	return func(co *cookieOpts) {
		co.Domain = domain
	}
}

func SetPath(path string) Options {
	return func(co *cookieOpts) {
		co.Path = path
	}
}

func SetSameSite(samesite SameSite) Options {
	return func(co *cookieOpts) {
		co.SameSite = samesite
	}
}

func SetMaxAge(max_age time.Duration) Options {
	return func(co *cookieOpts) {
		co.MaxAge = max_age
	}
}

func SetExpiry(expiry time.Time) Options {
	return func(co *cookieOpts) {
		co.Expiry = expiry
	}
}
