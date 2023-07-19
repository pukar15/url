package router

import (
	"net/http"

	constant "github.com/pukarlamichhane/url-shorter.git/constant"
	controller "github.com/pukarlamichhane/url-shorter.git/controller"
)

var urlShortner = Routes{
	Route{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},
	Route{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectURL},
}
