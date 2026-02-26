package http

import (
	"mime"
	"net/http"
	"strings"
)

// IsRequestPrintable attempts to determine if the the entity of the provided
// request can be printed for human consumption without modification.
func IsRequestPrintable(req *http.Request) bool {
	return isEntityPrintable(req.Header)
}

// IsResponsePrintable attempts to determine if the the entity of the provided
// response can be printed for human consumption without modification.
func IsResponsePrintable(rsp *http.Response) bool {
	return isEntityPrintable(rsp.Header)
}

// isEntityPrintable attempts to determine if the the entity of the request or
// response associated with the provied header can be printed for human
// consumption without modification.
func isEntityPrintable(hdr http.Header) bool {
	return hdr.Get("Content-Encoding") == "" && IsMimetypePrintable(hdr.Get("Content-Type"))
}

// IsMimetypePrintable attempts to determine if the provided mimetype can be
// printed for human consumption without modification.
func IsMimetypePrintable(t string) bool {
	m, p, err := mime.ParseMediaType(t)
	if err != nil {
		return false // if the mimetype is invalid, we assume it's not printable
	}
	switch {
	case m == "application/json":
		return true // this is a special case
	case m == "application/x-www-form-urlencoded":
		return true // this is a special case
	case strings.HasPrefix(m, "text/"):
		return true // if it's text, it's printable
	case p["charset"] != "":
		return true // if a charset is defined, it's printable
	default:
		return false // otherwise, we must assume it's not printable
	}
}
