package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
	"github.com/olivere/elastic"
)
import "github.com/notifications/db"

type Data struct {
	Id  int		`json:"id"`
	Data  map[string]string `json:"data"`
}

var (
	dao  db.Api
	route *gin.Engine
	events   chan string
	esClient    *elastic.Client
)

func Initialize() {
	dao = db.GetUrlDao()
	route = gin.Default()
	events = make(chan string)
	elasticUrl := "192.168.12.90:9200"
	esClient, _ = elastic.NewClient(
		elastic.SetURL(elasticUrl),
		elastic.SetSniff(false),
	)
	/*if e != nil {
		panic(e)
	}*/
}

func runWorkers(count int) {
	for i:=0;i<count;i++ {
		go worker(events)
	}
}

func setupRoutes()  {
	route.POST("/trigger", TriggerNotification)
}

func Run() {
	runWorkers(1)
	setupRoutes()
	route.Run(":" + viper.GetString("server.port"))
}


func TriggerNotification(c *gin.Context) {
	var payload SegmentPayload
	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	segmentId,_ := strconv.Atoi(c.PostForm("segment_id"))
	notificationId,_ := strconv.Atoi(c.PostForm("notification_id"))
	templateId,_ := strconv.Atoi(c.PostForm("template_id"))
	go readSegment(segmentId, notificationId, templateId)
}


func readSegment(segmentId int, notificationId int, templateId int) {
	now := time.Now()
	data,err := readFromDo("https://kreate-test-data.sgp1.digitaloceanspaces.com/data_10000.json")
	read_time := time.Since(now)

	fmt.Println("Read time taken", read_time)

	if err != nil {
		return
	}
	for _, item := range data {
		events <- fmt.Sprintf("%s", item.Data["name"])
	}
}

type Index struct {
	Message string
}

func sendNotification(message string) {
	m1 := Index{Message:message}
	ctx := context.Background()
	put, err := esClient.
		Index().
		Index("notifications").
		BodyJson(m1).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(put, message)
}

func worker(ch chan string) {
	for message := range ch {
		sendNotification(message)
	}
}


func readFromDo(url string) ([]Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data []Data
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &data); err != nil {
		return nil, err
	}
	return data, nil
}


/*
func ExpandUrl(c *gin.Context) {
	uid := c.Param("uid")
	redirectUrl, err := dao.GetUrl(uid)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	} else {
		c.Redirect(301, redirectUrl)
	}
}

func ShrinkUrl(c *gin.Context) {
	var payload Payload
	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	url := c.PostForm("url")
	uid, error := dao.InsertUrl(url)
	if error != nil {
		c.JSON(500, gin.H{"message": error.Error()})
		return
	} else {
		c.JSON(201, gin.H{"url": uid})
	}
}
*/