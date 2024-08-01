package AllTest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	type PostReqType struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	router := gin.New()

	router.GET("/getReq", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"rsp":  fmt.Sprintf("更换专辑:resources/songList.txt"),
		})
	})

	router.Run(":8080") // 8080端口，底层调用的是net/http包，也是单独启协程进行监听

}
