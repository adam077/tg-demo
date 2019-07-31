package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tg-demo/src/data"
	"tg-demo/src/utils"
)

type Screen struct {
	Id      string      `json:"screenId"`
	Name    string      `json:"name"`
	Content interface{} `json:"content"`
}

type Component struct {
	Id      string      `json:"componentId"`
	Name    string      `json:"name"`
	Path    string      `json:"path"`
	Content interface{} `json:"content"`
}

func GetScreens(c *gin.Context) {
	userId := c.GetHeader("UserId")
	users := data.GetUserWithScreens(userId)
	result := make([]Screen, 0)
	for _, user := range users {
		for _, UserScreen := range user.UserScreens {
			if UserScreen.Screen != nil {
				toAdd := Screen{
					Name: UserScreen.Screen.Name,
				}
				json.Unmarshal(UserScreen.Screen.Content, &toAdd.Content)
				result = append(result, toAdd)
			}
		}
	}
	utils.SuccessResp(c, "", result)
}

func AddScreen(c *gin.Context) {
	body, _ := c.GetRawData()
	var one Screen
	json.Unmarshal(body, &one)
	oneToAdd := data.Screen{}
	oneToAdd.ID = utils.GetUUID()
	oneToAdd.Name = one.Name
	oneToAdd.Content, _ = json.Marshal(one.Content)
	result := data.AddOne(&oneToAdd)
	utils.SuccessResp(c, "", result)
}

func DeleteScreen(c *gin.Context) {
	body, _ := c.GetRawData()
	var one Screen
	json.Unmarshal(body, &one)
	oneToAdd := data.Screen{}
	oneToAdd.ID = one.Id
	result := data.DeleteOne(&oneToAdd)
	utils.SuccessResp(c, "", result)
}

func PatchScreen(c *gin.Context) {
	body, _ := c.GetRawData()
	var one Screen
	json.Unmarshal(body, &one)
	oneToAdd := data.Screen{}
	oneToAdd.ID = one.Id
	content, _ := json.Marshal(one.Content)
	result := data.UpdateOne(&oneToAdd, map[string]interface{}{"content": content})
	utils.SuccessResp(c, "", result)
}

func GetComponents(c *gin.Context) {
	components := data.GetComponents()
	result := make([]Component, 0)
	for _, component := range components {
		toAdd := Component{
			Id:   component.ID,
			Name: component.Name,
			Path: component.Path,
		}
		json.Unmarshal(component.Content, &toAdd.Content)
		result = append(result, toAdd)
	}
	utils.SuccessResp(c, "", result)
}
