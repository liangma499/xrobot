package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       7,
		Password: "crot568Ze#mTeQC",
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		fmt.Println("err", err)
	}
	//fmt.Println(in)
	file, err := os.Open("./config/game_result_crash.csv")
	if err != nil {
		return
	}
	defer file.Close()
	//读取csv文件
	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}
	list := toStruct(data)
	table := make([]*redis.Z, 0)
	mastK := 1
	for k, val := range list {
		mastK = k + 1
		table = append(table, &redis.Z{Score: float64(mastK), Member: val.Hash})
	}
	err = rdb.ZAdd(context.Background(), "hashResultCrash", table...).Err()
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("success", mastK)

	}
}

// 转换为结构体切片
func toStruct(data [][]string) []Info {
	var shoppingList []Info
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Info
			for j, field := range line {
				if j == 0 {
					rec.Id = field
				} else if j == 1 {
					rec.Hash = field
				} else if j == 2 {
					decimal, _ := strconv.ParseFloat(field, 64)
					rec.Multi = decimal
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

type Info struct {
	Id    string
	Hash  string
	Multi float64
}
