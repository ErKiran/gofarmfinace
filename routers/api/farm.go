package api

import "github.com/gin-gonic/gin"

type FarmForm struct{
	Name string `json:"name" binding:"required"`
	Location string `json:"location"`
	FarmSize float64 `json:"farmSize"`
	FarmType string `json:"farmType"`
}

func CreateFarm(c *gin.Context){
	
}