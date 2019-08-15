package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/notify/controllers"
)

func Setup(poxy *droxolite.Epoxy) {
	//Subscribe
	subCtrl := &controllers.Subscribe{}
	subGroup := droxolite.NewRouteGroup("subscribe", subCtrl)
	subGroup.AddRoute("Subscribe", "/", "POST", roletype.User, subCtrl.Post)
	poxy.AddGroup(subGroup)

	//Publish
	pubCtrl := &controllers.Publish{}
	pubGroup := droxolite.NewRouteGroup("publish", pubCtrl)
	pubGroup.AddRoute("Publish", "/", "POST", roletype.User, pubCtrl.Post)
	poxy.AddGroup(pubGroup)
}
