package controllers

import (
	"github.com/revel/revel"
	"github.com/shubhamp-zeotap/Upflow/app"
)

type App struct {
	*revel.Controller
}

func prepareHeaders(a *App) {
	a.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")
	a.Response.Out.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	a.Response.Out.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (c App) Index() revel.Result {
	prepareHeaders(&c)
	return c.RenderJSON("")
}

func (c App) Plugins() revel.Result {
	prepareHeaders(&c)
	return c.RenderJSON(app.SendPlugins())
}

func (c App) Workflows() revel.Result {
	prepareHeaders(&c)
	return c.RenderJSON(app.SendWorkflows())
}

func (c App) UpdateWorkflow(channel string) revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	return c.RenderJSON(app.UpdateWorkflow(channel, jsonData))
}

func (c App) RunWorkflow(channel string) revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	//fmt.Println(jsonData)
	return c.RenderJSON(app.RunWorkflow(channel, jsonData))
}

func (c App) SystemConfig() revel.Result {
	prepareHeaders(&c)
	return c.RenderJSON(app.SendSystemInputs())
}
