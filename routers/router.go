package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/notify/controllers"
)

func Setup(poxy resins.Epoxi) {
	//Subscribe
	subCtrl := &controllers.Subscribe{}
	subGroup := routing.NewRouteGroup("subscribe", mix.JSON)
	subGroup.AddRoute("Subscribe", "", "POST", roletype.User, subCtrl.Post)
	poxy.AddGroup(subGroup)

	//Publish
	pubCtrl := &controllers.Publish{}
	pubGroup := routing.NewRouteGroup("publish", mix.JSON)
	pubGroup.AddRoute("Publish", "", "POST", roletype.User, pubCtrl.Post)
	poxy.AddGroup(pubGroup)
}
