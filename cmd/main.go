package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"outgoing/conf"
	"outgoing/parse"
	"reflect"
)

/*{
"conversationId": "cidRKA2a+fKIUeTmUj8rYScUQ==",
"atUsers": [{
"dingtalkId": "$:LWCP_v1:$CnBeWRt3EMDJYzUJQ3g9tpwvRAEveAAR"
}],
"chatbotUserId": "$:LWCP_v1:$CnBeWRt3EMDJYzUJQ3g9tpwvRAEveAAR",
"msgId": "msgomRlBpPoT4046yWeyTWmRg==",
"senderNick": "dddd",
"isAdmin": false,
"sessionWebhookExpiredTime": 1560698502162,
"createAt": 1560697302117,
"conversationType": "2",
"senderId": "$:LWCP_v1:$ECf4QkQBPxA6U1111iAlPvCcA==",
"conversationTitle": "dddd-私有文件传输",
"isInAtList": true,
"sessionWebhook": "https://oapi.dingtalk.com/robot/sendBySession?session=18651ed4d21900356a7f2348791775b91122",
"text": {
"content": " /add  你好"
},
"msgtype": "text"
}*/

//json定义
type params struct {
	MsgType string `json:"msgtype"`
	Text params_text `json:"text"`
	CreateAt int64 `json:"createAt"`
	ConversationType string `json:"conversationType"`
	ConversationId string `json:"conversationId"`

	ConversationTitle string `json:"conversationTitle"`
	SenderId string `json:"senderId"`
	SenderNick string `json:"senderNick"`
	SenderStaffId string `json:"senderStaffId"`
	IsAdmin bool `json:"isAdmin"`
	Context string `json:"context"`
	ChatbotCorpId string `json:"chatbotCorpId"`
	ChatbotUserId string `json:"chatbotUserId"`

}
type params_text struct {
	Content string `json:"content"`
}

//resp
type resp struct {
	Msgtype string `json:"msgtype"`
	Text text `json:"text"`
}

type text struct {
	Content string `json:"content"`
}


func dingtalk(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("Hello, world mutex..."))
	token := req.Header.Get("token")
	arrStrings := []string{"121212122323233434", "121212122323233433", "112233"}
	stringExist, _ := InArray(token, arrStrings)

	var parseStr string
	if stringExist != true {
		parseStr = "token校验错误"

	} else {
		fmt.Println(token)

		//token鉴权处理
		body, _ := ioutil.ReadAll(req.Body)


		log.Println(string(body))
		param := &params{}
		_ = json.Unmarshal(body, param)

		log.Println(param)

		//调用解析方法
		parseStr = parse.ParseParam(param.Text.Content)
	}

	respData := &resp{}
	respData.Msgtype = "text"
	respData.Text.Content = parseStr

	resp, _ := json.Marshal(respData)

	w.Write([]byte(resp))
}

func InArray(needle interface{}, haystack interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func main() {
	fmt.Println("hello outgoing")

	flag.Parse()
	conf.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", dingtalk)
	http.ListenAndServe(":8000", mux)
}
