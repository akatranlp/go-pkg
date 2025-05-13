package middleware

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"slices"
	"strings"
)

type Middleware = func(next http.Handler) http.Handler

type hostCtxKey struct{}
type protoCtxKey struct{}

// Forwarded looks at the headers of the requests and populates the X-Forwarded-Host and X-Forwarded-Proto values
// in the context and puts the first trusted X-Forwarded-For value as r.RemoteAddr
func Forwarded(trustedProxies int) Middleware {
	return func(next http.Handler) http.Handler {
		if trustedProxies < 1 {
			return next
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			var remoteAddr string
			forwardedFor := collectHeaderValues(r, "X-Forwarded-For")
			if len(forwardedFor) > 0 && len(forwardedFor) < trustedProxies {
				remoteAddr = forwardedFor[0]
			} else if len(forwardedFor) > 0 {
				remoteAddr = forwardedFor[len(forwardedFor)-trustedProxies]
			}
			forwardedHost := collectHeaderValues(r, "X-Forwarded-Host")
			if len(forwardedHost) > 0 {
				ctx = context.WithValue(ctx, hostCtxKey{}, forwardedHost[0])
			}
			forwardedProto := collectHeaderValues(r, "X-Forwarded-Proto")
			if len(forwardedProto) > 0 {
				ctx = context.WithValue(ctx, protoCtxKey{}, forwardedProto[0])
			}

			r = r.WithContext(ctx)
			if remoteAddr != "" {
				r.RemoteAddr = remoteAddr
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Proto gets the first X-Forwarded-Proto header from the request
// Use this to look if the original request was https or http.
//
// In Combination with r.Host You can get the origin of the request like $proto://$r.Host
func Proto(r *http.Request) string {
	if proto, ok := r.Context().Value(protoCtxKey{}).(string); ok {
		return proto
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

// Origin uses the Proto function and r.Host to give you the origin of the original request
func Origin(r *http.Request) string {
	proto := Proto(r)
	return fmt.Sprintf("%s://%s", proto, r.Host)
}

// Host gets the first X-Forwarded-Host header from the request
//
// Please use r.Host directly because this value can easily be spoofed
// and not all reverseProxies set these field
func Host(r *http.Request) string {
	if host, ok := r.Context().Value(hostCtxKey{}).(string); ok {
		return host
	}
	return r.Host
}

func mapSeq[T, U any](seq iter.Seq[T], mapFn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(mapFn(v)) {
				return
			}
		}
	}
}

func collectHeaderValues(r *http.Request, header string) []string {
	return slices.Collect(mapSeq(strings.SplitSeq(strings.Join(r.Header[header], ","), ","), func(s string) string {
		return strings.TrimSpace(s)
	}))
}
