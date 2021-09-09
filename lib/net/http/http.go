package http

import (
	"net/http"

	qlang "github.com/topxeq/qlang/spec"
)

// -----------------------------------------------------------------------------

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name":            "net/http",
	"request":          http.NewRequest,
	"readRequest":      http.ReadRequest,
	"readResponse":     http.ReadResponse,
	"parseHTTPVersion": http.ParseHTTPVersion,
	"parseTime":        http.ParseTime,

	"NewRequest":       http.NewRequest,
	"ReadRequest":      http.ReadRequest,
	"ReadResponse":     http.ReadResponse,
	"ParseHTTPVersion": http.ParseHTTPVersion,
	"ParseTime":        http.ParseTime,

	"get":      http.Get,
	"post":     http.Post,
	"postForm": http.PostForm,
	"head":     http.Head,

	"Get":      http.Get,
	"Post":     http.Post,
	"PostForm": http.PostForm,
	"Head":     http.Head,

	"handle":            http.Handle,
	"handleFunc":        http.HandleFunc,
	"serveMux":          http.NewServeMux,
	"serve":             http.Serve,
	"listenAndServe":    http.ListenAndServe,
	"listenAndServeTLS": http.ListenAndServeTLS,

	"Handle":            http.Handle,
	"HandleFunc":        http.HandleFunc,
	"NewServeMux":       http.NewServeMux,
	"Serve":             http.Serve,
	"ListenAndServe":    http.ListenAndServe,
	"ListenAndServeTLS": http.ListenAndServeTLS,

	"error":           http.Error,
	"notFound":        http.NotFound,
	"notFoundHandler": http.NotFoundHandler,
	"redirect":        http.Redirect,
	"redirectHandler": http.RedirectHandler,
	"fileServer":      http.FileServer,
	"fileTransport":   http.NewFileTransport,
	"serveContent":    http.ServeContent,
	"serveFile":       http.ServeFile,
	"setCookie":       http.SetCookie,
	"stripPrefix":     http.StripPrefix,
	"timeoutHandler":  http.TimeoutHandler,

	"Error":            http.Error,
	"NotFound":         http.NotFound,
	"NotFoundHandler":  http.NotFoundHandler,
	"Redirect":         http.Redirect,
	"RedirectHandler":  http.RedirectHandler,
	"FileServer":       http.FileServer,
	"NewFileTransport": http.NewFileTransport,
	"ServeContent":     http.ServeContent,
	"ServeFile":        http.ServeFile,
	"SetCookie":        http.SetCookie,
	"StripPrefix":      http.StripPrefix,
	"TimeoutHandler":   http.TimeoutHandler,

	"statusText":           http.StatusText,
	"canonicalHeaderKey":   http.CanonicalHeaderKey,
	"detectContentType":    http.DetectContentType,
	"maxBytesReader":       http.MaxBytesReader,
	"proxyFromEnvironment": http.ProxyFromEnvironment,
	"proxyURL":             http.ProxyURL,

	"StatusText":           http.StatusText,
	"CanonicalHeaderKey":   http.CanonicalHeaderKey,
	"DetectContentType":    http.DetectContentType,
	"MaxBytesReader":       http.MaxBytesReader,
	"ProxyFromEnvironment": http.ProxyFromEnvironment,
	"ProxyURL":             http.ProxyURL,

	"DefaultTransport": http.DefaultTransport,
	"DefaultClient":    http.DefaultClient,
	"DefaultServeMux":  http.DefaultServeMux,

	"Dir":      qlang.StructOf((*http.Dir)(nil)),
	"Client":   qlang.StructOf((*http.Client)(nil)),
	"Cookie":   qlang.StructOf((*http.Cookie)(nil)),
	"Header":   qlang.StructOf((*http.Header)(nil)),
	"Request":  qlang.StructOf((*http.Request)(nil)),
	"Response": qlang.StructOf((*http.Response)(nil)),
	"Server":   qlang.StructOf((*http.Server)(nil)),

	"StatusContinue":           http.StatusContinue,           // RFC 7231, 6.2.1
	"StatusSwitchingProtocols": http.StatusSwitchingProtocols, // RFC 7231, 6.2.2
	"StatusProcessing":         http.StatusProcessing,         // RFC 2518, 10.1
	"StatusEarlyHints":         http.StatusEarlyHints,         // RFC 8297

	"StatusOK":                   http.StatusOK,                   // RFC 7231, 6.3.1
	"StatusCreated":              http.StatusCreated,              // RFC 7231, 6.3.2
	"StatusAccepted":             http.StatusAccepted,             // RFC 7231, 6.3.3
	"StatusNonAuthoritativeInfo": http.StatusNonAuthoritativeInfo, // RFC 7231, 6.3.4
	"StatusNoContent":            http.StatusNoContent,            // RFC 7231, 6.3.5
	"StatusResetContent":         http.StatusResetContent,         // RFC 7231, 6.3.6
	"StatusPartialContent":       http.StatusPartialContent,       // RFC 7233, 4.1
	"StatusMultiStatus":          http.StatusMultiStatus,          // RFC 4918, 11.1
	"StatusAlreadyReported":      http.StatusAlreadyReported,      // RFC 5842, 7.1
	"StatusIMUsed":               http.StatusIMUsed,               // RFC 3229, 10.4.1

	"StatusMultipleChoices":   http.StatusMultipleChoices,   // RFC 7231, 6.4.1
	"StatusMovedPermanently":  http.StatusMovedPermanently,  // RFC 7231, 6.4.2
	"StatusFound":             http.StatusFound,             // RFC 7231, 6.4.3
	"StatusSeeOther":          http.StatusSeeOther,          // RFC 7231, 6.4.4
	"StatusNotModified":       http.StatusNotModified,       // RFC 7232, 4.1
	"StatusUseProxy":          http.StatusUseProxy,          // RFC 7231, 6.4.5
	"StatusTemporaryRedirect": http.StatusTemporaryRedirect, // RFC 7231, 6.4.7
	"StatusPermanentRedirect": http.StatusPermanentRedirect, // RFC 7538, 3

	"StatusBadRequest":                   http.StatusBadRequest,                   // RFC 7231, 6.5.1
	"StatusUnauthorized":                 http.StatusUnauthorized,                 // RFC 7235, 3.1
	"StatusPaymentRequired":              http.StatusPaymentRequired,              // RFC 7231, 6.5.2
	"StatusForbidden":                    http.StatusForbidden,                    // RFC 7231, 6.5.3
	"StatusNotFound":                     http.StatusNotFound,                     // RFC 7231, 6.5.4
	"StatusMethodNotAllowed":             http.StatusMethodNotAllowed,             // RFC 7231, 6.5.5
	"StatusNotAcceptable":                http.StatusNotAcceptable,                // RFC 7231, 6.5.6
	"StatusProxyAuthRequired":            http.StatusProxyAuthRequired,            // RFC 7235, 3.2
	"StatusRequestTimeout":               http.StatusRequestTimeout,               // RFC 7231, 6.5.7
	"StatusConflict":                     http.StatusConflict,                     // RFC 7231, 6.5.8
	"StatusGone":                         http.StatusGone,                         // RFC 7231, 6.5.9
	"StatusLengthRequired":               http.StatusLengthRequired,               // RFC 7231, 6.5.10
	"StatusPreconditionFailed":           http.StatusPreconditionFailed,           // RFC 7232, 4.2
	"StatusRequestEntityTooLarge":        http.StatusRequestEntityTooLarge,        // RFC 7231, 6.5.11
	"StatusRequestURITooLong":            http.StatusRequestURITooLong,            // RFC 7231, 6.5.12
	"StatusUnsupportedMediaType":         http.StatusUnsupportedMediaType,         // RFC 7231, 6.5.13
	"StatusRequestedRangeNotSatisfiable": http.StatusRequestedRangeNotSatisfiable, // RFC 7233, 4.4
	"StatusExpectationFailed":            http.StatusExpectationFailed,            // RFC 7231, 6.5.14
	"StatusTeapot":                       http.StatusTeapot,                       // RFC 7168, 2.3.3
	"StatusMisdirectedRequest":           http.StatusMisdirectedRequest,           // RFC 7540, 9.1.2
	"StatusUnprocessableEntity":          http.StatusUnprocessableEntity,          // RFC 4918, 11.2
	"StatusLocked":                       http.StatusLocked,                       // RFC 4918, 11.3
	"StatusFailedDependency":             http.StatusFailedDependency,             // RFC 4918, 11.4
	"StatusTooEarly":                     http.StatusTooEarly,                     // RFC 8470, 5.2.
	"StatusUpgradeRequired":              http.StatusUpgradeRequired,              // RFC 7231, 6.5.15
	"StatusPreconditionRequired":         http.StatusPreconditionRequired,         // RFC 6585, 3
	"StatusTooManyRequests":              http.StatusTooManyRequests,              // RFC 6585, 4
	"StatusRequestHeaderFieldsTooLarge":  http.StatusRequestHeaderFieldsTooLarge,  // RFC 6585, 5
	"StatusUnavailableForLegalReasons":   http.StatusUnavailableForLegalReasons,   // RFC 7725, 3

	"StatusInternalServerError":           http.StatusInternalServerError,           // RFC 7231, 6.6.1
	"StatusNotImplemented":                http.StatusNotImplemented,                // RFC 7231, 6.6.2
	"StatusBadGateway":                    http.StatusBadGateway,                    // RFC 7231, 6.6.3
	"StatusServiceUnavailable":            http.StatusServiceUnavailable,            // RFC 7231, 6.6.4
	"StatusGatewayTimeout":                http.StatusGatewayTimeout,                // RFC 7231, 6.6.5
	"StatusHTTPVersionNotSupported":       http.StatusHTTPVersionNotSupported,       // RFC 7231, 6.6.6
	"StatusVariantAlsoNegotiates":         http.StatusVariantAlsoNegotiates,         // RFC 2295, 8.1
	"StatusInsufficientStorage":           http.StatusInsufficientStorage,           // RFC 4918, 11.5
	"StatusLoopDetected":                  http.StatusLoopDetected,                  // RFC 5842, 7.2
	"StatusNotExtended":                   http.StatusNotExtended,                   // RFC 2774, 7
	"StatusNetworkAuthenticationRequired": http.StatusNetworkAuthenticationRequired, // RFC 6585, 6

}

// -----------------------------------------------------------------------------
