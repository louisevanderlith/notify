package handles

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(secret, securityUrl string) http.Handler {
	//Subscribe
	/*subCtrl := &controllers.Subscribe{}
	subGroup := routing.NewRouteGroup("subscribe", mix.JSON)
	subGroup.AddRoute("Subscribe", "", "POST", roletype.User, subCtrl.Post)
	e.AddBundle(subGroup)

	//Publish
	pubCtrl := &controllers.Publish{}
	pubGroup := routing.NewRouteGroup("publish", mix.JSON)
	pubGroup.AddRoute("Publish", "", "POST", roletype.User, pubCtrl.Post)
	e.AddBundle(pubGroup)*/
	return mux.NewRouter()
}
