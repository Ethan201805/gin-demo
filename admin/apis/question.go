package apis

import (
	"gin-web-demo/Test/gin-demo/admin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//api首页
func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddQuestions(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBind(&question); err != nil {
		log.Fatal(err.Error())
	}
	res, err := question.AddQuestion()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": res,
	})

}
