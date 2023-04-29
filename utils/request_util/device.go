package request_util

import (
	"net/http"
	"strings"
)

type DeviceType string

const (
	MOBILE DeviceType = "Mobile"
	TABLET DeviceType = "Tab"
	WEB    DeviceType = "Web"
	TV     DeviceType = "TV"
)

func GetDeviceType(r *http.Request) string {

	if isUserAgent(r, "Android", "webOS", "iPhone", "BlackBerry", "Windows Phone") {
		return string(MOBILE)
	}
	if isUserAgent(r, "iPad", "iPod", "tablet", "RX-34", "FOLIO") ||
		(isUserAgent(r, "Kindle", "Mac OS") && isUserAgent(r, "Silk")) ||
		(isUserAgent(r, "AppleWebKit") && isUserAgent(r, "Silk")) {
		return string(TABLET)
	}
	if isUserAgent(r, "TV", "NetCast", "boxee", "Kylo", "Roku", "DLNADOC") {
		return string(TV)
	}

	return string(WEB)
}

func isUserAgent(r *http.Request, userAgents ...string) bool {
	userAgent := r.Header.Get("User-Agent")
	for _, v := range userAgents {
		if strings.Contains(userAgent, v) {
			return true
		}
	}
	return false
}
