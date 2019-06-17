package parse

import (
	"flag"
	"fmt"
	"outgoing/outquery"
	"strconv"
	"strings"
)

func ParseParam(myString string) (string) {
	var helper, tag, url, item_desc, id string

	var add = flag.NewFlagSet("add", flag.ContinueOnError)
	add.StringVar(&helper,"h","","帮助使用")
	add.StringVar(&tag,"tag","","标签")
	add.StringVar(&url,"url","","跳转URL")
	add.StringVar(&item_desc,"item_desc","","文档具体描述")
	add.StringVar(&id,"id","","文档ID")

	myString = strings.TrimSpace(myString)
	cmd := strings.Fields(myString)

	//处理字符串
	newList := HandleString(cmd[1:])
	add.Parse(newList)

	outquery.NewMysql()

	var resultStr string
	if cmd[0] == "/add" || cmd[0] == "/update" || cmd[0] == "/del" {
		//处理空格
		if cmd[0] == "/add" {
			fmt.Println("新增")
			if len(helper) > 0 {
				fmt.Println("请输出/add帮助文档")
				resultStr = "@机器人 /add -tag {标题} -url {跳转URL} -item_desc {具体描述}"
				goto RESULT
			}
			if len(tag) <= 0  {
				fmt.Println("tag不能为空")
				resultStr = "tag不能为空"
				goto RESULT
			}
			if len(url) <= 0 {
				fmt.Println("url不能为空")
				resultStr = "url不能为空"
				goto RESULT
			}
			insertID, err := outquery.InsertAction(tag, url, item_desc)
			if err != nil {
				fmt.Println("插入失败")
				resultStr = "插入失败"
				goto RESULT
			}
			resultStr = "插入成功:" + strconv.Itoa(insertID)


		}
		if cmd[0] == "/update" {

			fmt.Println("更新")
		}

		if cmd[0] == "/del" {
			fmt.Println("删除")
			if len(helper) > 0 {
				fmt.Println("请输出/del 帮助文档")
				resultStr = "@机器人 /del -id {ID}"
				goto RESULT
			}
			if len(id) <= 0  {
				fmt.Println("id不能为空")
				fmt.Println(id)
				resultStr = "id不能为空"
				goto RESULT
			}

			deleteID, err := outquery.DelAction(id)
			if err != nil {
				fmt.Println("删除失败")
				resultStr = "删除失败"
				goto RESULT
			}
			resultStr = "删除成功:" + strconv.Itoa(deleteID)
		}
	} else {
		//trim字符串， 做查询处理
		rows, _ := outquery.QueryAction(myString)

		for _, row := range rows {
			txt := fmt.Sprintf("%s== %s == %s(%s)\n", row["id"], row["url"], row["tag"], row["item_desc"])
			resultStr = resultStr + txt
		}
		if len(resultStr) <= 0 {
			resultStr = "无数据"
		}
	}

	//其他都是查询， 走查询操作


	//QueryAction("nihao")


RESULT:
	return resultStr
	//其他，都等于查询
}


func HandleString(stringList []string) ([]string){
	var newList []string

	checkType := 1
	cache_str := ""
	for _, currentStr := range(stringList) {
		checkTag := currentStr == "-tag"
		checkDesc := currentStr == "-item_desc"
		checkUrl := currentStr == "-url"
		checkH := currentStr == "-h"
		checkID := currentStr == "-id"

		fmt.Println(currentStr)
		fmt.Println("=====")

		if len(currentStr) <= 0 {
			continue
		}
		if checkType == 1 {
			//-tag -desc, -url, sdfsdf fdsfsdf

			if checkTag || checkDesc || checkUrl || checkH || checkID {
				newList = append(newList, currentStr)
				checkType = 3;
				cache_str = ""
			}
		} else if checkType == 2 || checkType == 3 {
			if checkTag || checkDesc || checkUrl || checkH || checkID{
				newList = append(newList, cache_str)
				newList = append(newList, currentStr)
				checkType = 3
				cache_str = ""
			} else {
				if len(cache_str) <= 0 {
					cache_str = currentStr
				} else {
					cache_str = cache_str + "=" + currentStr
				}

				checkType = 2
			}
		}
		// -232323 3434 5434 4545 232 54545 323 45454 32323

	}
	if checkType == 3 {
		newList = append(newList, "")
	} else if checkType == 2 {
		newList = append(newList, cache_str)
	}
	fmt.Println(newList)
	return newList
}
