package main

import (
    "fmt"
    "net/http"
	"net/url"
	"github.com/tidwall/gjson"
	"io/ioutil"
	//"time"
	"reflect"
	"strings"
	"strconv"
)

// 获取token
func getTocken() string {

	// 判断处于有效期的token是否存在
	// ... ToDo

	// 不存在，则重新获取
	query := url.Values{
		"app_id": {"*******"},
		"app_secret": {"*******************************"},
	}
	resp, err := http.PostForm("https://open.youjiuhealth.com/mch/v3/session", query)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	strBody := string(body)
	access_token := gjson.Get(strBody, "access_token")
	
	return access_token.String()
}

func typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}

func getData(path string,query url.Values) string {

	token := getTocken()
	req_data := query.Encode()
	url := "http://open.cc/mch" + path
	
	//fmt.Println(url)
	
	req, err := http.NewRequest("GET", url, strings.NewReader(req_data))
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	if err != nil {
        panic(err)
    }

	client := &http.Client{}
    resp, error := client.Do(req)
    if error != nil {
        panic(error)
    }
    defer resp.Body.Close()

    result, _ := ioutil.ReadAll(resp.Body)
    body := string(result)	
	strBody := string(body)
	
	return strBody
}

// 获取报告列表
func getReportsList(query url.Values) string {
	ret := getData("/reports",query)
	return ret
}

// 获取报告详情
func getReportsDetail(id int) string {
	ret := getData("/reports/" + strconv.Itoa(id),url.Values{} )
	return ret
}

// 获取小程序码
func getMiniProgramCode(id int) string {
	ret := getData("/reports/" + strconv.Itoa(id) + "/miniProgramCode",url.Values{} )
	return ret
}

// 获取商家列表
func getClients(query url.Values) string {
	ret := getData("/clients",query)
	return ret
}

// 获取商家设备列表
func getClientDevices(client_id int) string {
	ret := getData("/clients/" + strconv.Itoa(client_id) + "/devices",url.Values{})
	return ret
}



func main() {
	// 获取token
	//+---------------------------------------
	//var ret = getTocken()
	
	// 获取报告列表
	//+---------------------------------------
	/*
	var query = url.Values{}
	query.Add("page", "2")
	var ret = getReportsList(query)
	*/
	
	
	
	// 获取报告详情
	//+---------------------------------------
	//measurementId := 39335628
	//var ret = getReportsDetail(measurementId)
	
	// 获取小程序码
	//+---------------------------------------
	//measurementId := 39335628
	//var ret = getMiniProgramCode(measurementId)
	
	// 获取商家列表
	//+---------------------------------------	
	var query = url.Values{}
	query.Add("page", "1")
	var ret = getClients(query)

	
	// 获取商家设备列表
	//+---------------------------------------
	// client_id := 3236199
	// var ret = getClientDevices(client_id)
	
	fmt.Println( ret )
	
}








