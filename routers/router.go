package routers

import (
	"github.com/louisevanderlith/droxolite"
)

func Setup(poxy *droxolite.Epoxy) {
	//History
	/*histCtrl := &controllers.HistoryController{}
	histGroup := droxolite.NewRouteGroup("history", histCtrl)
	histGroup.AddRoute("/", "POST", roletype.Owner, histCtrl.Post)
	histGroup.AddRoute("/{vehicleKey:[0-9]+\x60[0-9]+}", "GET", roletype.User, histCtrl.GetByVehicle)
	poxy.AddGroup(histGroup)*/

	//ctrlmap := EnableFilters(s, host)

	/*profCtrl := controllers.NewProfileCtrl(ctrlmap)

	beego.Router("/v1/profile", profCtrl, "post:Post;put:Put")
	beego.Router("/v1/profile/:site", profCtrl, "get:GetOne")
	beego.Router("/v1/profile/all/:pagesize", profCtrl, "get:Get")


	Sal nog die een moet beplan. (Push to Client...)
	*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner
	emptyMap["PUT"] = roletype.Owner

	ctrlmap.Add("/v1/profile", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}*/
