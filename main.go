package main

import (
	"fmt"
	//"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
	"flag"
)

//{"name":"test5", "column_families":[{
//#>              "name":"columnfam1",
//#>              "bloomfilter":true,
//#>              "time_to_live":10,
//#>              "in_memory":false,
//#>              "max_versions":2,
//#>              "compression":"",
//#>              "max_value_length":50,
//#>              "block_cache_enabled":true
//#>           }
//#> ]}

type tableInfo struct {
	Name            string   `json:"name"`
	Column_families []column `json:"column_families"`
}

type column struct {
	Name                string `json:"name"`
	Bloomfilter         bool   `json:"bloomfilter"`
	Time_to_live        int    `json:"time_to_live"`
	In_memory           bool   `json:"in_memory"`
	Max_versions        int    `json:"max_versions"`
	Compression         string `json:"compression,omitempty"`
	Max_value_length    int    `json:"max_value_length"`
	Block_cache_enabled bool   `json:"block_cache_enabled"`
}

func init() {
	//os.Getenv("BSI_HDFS_HBASEFORDEMO_HOST")
}

func main() {
	fmt.Println("Let's go!!!")

	table := flag.String("ct", "", "create table")
	flag.Parse()

	if *table == "" {
		return
	} else {
		fmt.Println(*table)
	}

	columns := make([]column, 0)
	column := column{Name: "cf", Bloomfilter: true, Time_to_live: 10, In_memory: false, Max_versions: 2, Compression: "", Max_value_length: 50, Block_cache_enabled: true}
	columns = append(columns, column)

	tableInfo := tableInfo{Name: "59e695e786ba11e68852fa163d0e0615:t2", Column_families: columns}

	body, err := json.Marshal(tableInfo)
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}

	url := "http://36.110.132.55:8070/59e695e786ba11e68852fa163d0e0615:t1/schema"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("NewRequest err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient.Do err:", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK.")
		return
	} else {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ReadAll err:", err)
			return
		}
		fmt.Println(string(respBody))
	}

}
