package main

import (
	"fmt"
	//"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
	"flag"
	"os"
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

var (
	hbaseHost string = "http://hadoop-1.jcloud.local"
	hbasePort string
	hbaseNamespace string
)

func init() {
	hbasePort = os.Getenv("BSI_HBASE_HBASEDEMO_PORT")
	hbaseNamespace = os.Getenv("d479d1429ab811e6b845fa163d0e0615")
}

func main() {
	fmt.Println("Let's go!!!")

	createTable := flag.String("create", "default", "create table.")
	lsTables := flag.String("ls", "", "list all tables.")
	flag.Parse()

	if *createTable != "" {
		columns := make([]column, 0)
		column := column{Name: "cf", Bloomfilter: true, Time_to_live: 10, In_memory: false, Max_versions: 2, Compression: "", Max_value_length: 50, Block_cache_enabled: true}
		columns = append(columns, column)

		tableName := hbaseNamespace+":"+*createTable
		tableInfo := tableInfo{Name: tableName, Column_families: columns}

		body, err := json.Marshal(tableInfo)
		if err != nil {
			fmt.Println("Marshal err:", err)
			return
		}

		url := hbaseHost+":"+hbasePort+"/"+hbaseNamespace+":"+*createTable+"/schema"
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
	}else if *lsTables != "" {
		url := hbaseHost+":"+hbasePort
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("NewRequest err:", err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("http.DefaultClient.Do err:", err)
			return
		}

		if resp.StatusCode == http.StatusOK {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("ReadAll err:", err)
				return
			}
			fmt.Println(string(respBody))
		} else {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("ReadAll err:", err)
				return
			}
			fmt.Println(string(respBody))
			return
		}
	} else {
		fmt.Println("End...")
	}

	return
}
