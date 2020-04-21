package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"strings"
)

func URLFor(routeinfos []gin.RouteInfo) func(string, ...interface{}) string {
	httpMethod := "GET"
	controllerName := "controller"
	return func(endpoint string, values ...interface{}) string {
		if len(endpoint) == 0 {
			return "/"
		}
		paths := strings.Split(endpoint, ".")
		//if len(paths) <= 1 {
		//	fmt.Println("urlfor endpoint must like path.controller.method")
		//	return "/"
		//}
		if len(values)%2 != 0 {
			fmt.Println("urlfor params must key-value pair")
			return "/"
		}
		params := make(map[string]string)
		if len(values) > 0 {
			key := ""
			for k, v := range values {
				if k%2 == 0 {
					key = fmt.Sprint(v)
				} else {
					params[key] = fmt.Sprint(v)
				}
			}
		}
		//controllerName := strings.Join(paths[:len(paths)-1], "/")
		methodName := paths[len(paths)-1]
		for _, route := range routeinfos {
			url, ok := getURL(route, controllerName, methodName, params, httpMethod)
			if ok {
				return url
			}
		}
		return "/"
	}
}

func getURL(r gin.RouteInfo, controllerName, methodName string, params map[string]string, httpMethod string) (string, bool) {
	handlers := strings.Split(r.Handler, "/")
	url := r.Path
	if r.Method == httpMethod && handlers[len(handlers)-1] == controllerName+"."+methodName {
		//url += toURL(params)
		for k, v := range params {
			if k[:1] == ":" {
				url = url + "/" + k + "=" + v
			}
		}
		for k, _ := range params {
			if k[:1] == ":" {
				url = strings.Replace(url, k+"=", "", -1)
				delete(params, k)
			}
		}
		return url + toURL(params), true
	}
	return "/", false
}

func toURL(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	u := "?"
	for k, v := range params {
		u += k + "=" + v + "&"
	}
	return strings.TrimRight(u, "&")
}
