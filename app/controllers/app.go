package controllers

import (
	"github.com/revel/revel"
	"github.com/shubhamp-zeotap/Upflow/app"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Plugins() revel.Result {
	return c.RenderJSON(app.SendPlugins()) 
}

func (c App) Workflows() revel.Result {
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
	return c.RenderJSON(app.SendSystemInputs())
}

