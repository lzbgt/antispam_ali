package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"greensdksample"
	"net/http"
	"strconv"
	"time"
	//"encoding/json"
	"uuid"
)

const apiPath = "/green/text/scan"
const apiPathImageSync = "/green/image/scan"
const apiPathImageAsync = "/green/image/asyncscan"
const apiPathVideoSync = "/green/text/scan"
const apiPathVideoAsync = "/green/text/scan"
const bizType = ""
var clientInfo = &greensdksample.ClinetInfo{Ip: "127.0.0.1"}
var profile = &greensdksample.Profile{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret}

func HandlePostMedia(c *gin.Context) {
	media := c.Param("media")
	rpc := c.Param("rpc")

	if media == "" {
		media = "text"
	}
	if rpc == "" {
		rpc = "sync"
	}

	fmt.Println(media, rpc)
	var post greensdksample.PostContents
	if c.BindJSON(&post) == nil {
		if len(post.Contents) > 0 {
			//fmt.Printf("%q\n", post.Contents)
			scenes := []string{"antispam"}
			tasks := make([]greensdksample.Task, len(post.Contents))
			for key := range post.Contents {
				tasks[key].DataId = uuid.Rand().Hex()
				tasks[key].Content = string(post.Contents[key])
			}
	
			bizData := greensdksample.BizData{BizType: bizType, Scenes: scenes, Tasks: tasks}
			var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile: *profile}
			tstart := time.Now()
			result := client.GetResponse(apiPath, *clientInfo, bizData)
			elapse := strconv.FormatFloat(time.Since(tstart).Seconds(), 'f', 5, 64)
			c.Header("x-elapse", elapse)
			c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(result))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "no contents found in request", "code": "400"})
			//fmt.Println("no access")
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid request", "code":"400"})
	}
}