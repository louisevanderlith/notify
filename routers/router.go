package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/notify/controllers"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.User, mix.JSON, &controllers.Subscribe{}, &controllers.Publish{})
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
}
