package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func main() {

	ctx := context.Background()
	// 设置日志输出
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)

	host := "http://127.0.0.1:5601"
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	// 使用client进行查询
	q := elastic.NewMatchQuery("user", "olivere")
	rsp, err := client.Search().Index("twitter").Query(q).Do(ctx)
	if err != nil {
		panic(err)
	}

	total := rsp.Hits.TotalHits.Value
	fmt.Printf("Found %d tweets\n", total)

	for _, hit := range rsp.Hits.Hits {
		jsonData, err := hit.Source.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Printf(string((jsonData)))
	}

	// 插入数据
	tweet1 := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		BodyJson(tweet1).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	// 新建索引，mapping
	createIndex, err := client.CreateIndex("mygoods").BodyString(mapping).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}

}

const mapping = `
{
	"mappings":{
		"properties":{
			"name":{
				"type":"keyword",
				"analyzer": "ik_max_word",
			},
			"id":{
				"type":"integer",
			}
		}
	}
}`

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
