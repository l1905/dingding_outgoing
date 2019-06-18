package outquery

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"outgoing/conf"
	"strconv"
	"time"
)

//mysql配置
//var dbConfig string = "xxxxx"

//mysql实例
var dbCon *sql.DB

/**
初始化mysql连接
*/
func NewMysql() {
	fmt.Println("init执行开始")
	//https://github.com/go-sql-driver/mysql/wiki/Examples
	//https://blog.csdn.net/rambo_huang/article/details/60604924
	var err error
	dbCon, err = sql.Open("mysql", conf.Conf.MySQL.Host)
	if err != nil {
		fmt.Println(err)
	}

	err = dbCon.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

func QueryAction(keyword string) ([]map[string]string, error){
	var queryObject []map[string]string
	sqlTpl := "SELECT id, tag, url, item_desc FROM outgoing WHERE  tag like ? "
	fmt.Println(sqlTpl)

	tag := fmt.Sprintf("%%%s%%", keyword )
	fmt.Println(tag)

	rows, err := dbCon.Query(sqlTpl,  tag)
	if err != nil {
		fmt.Println(err)
		return queryObject, err
	}
	for rows.Next() {
		var id int
		var tag, url, item_desc string
		err = rows.Scan(&id, &tag, &url, &item_desc)
		if err != nil {
			fmt.Println(err)
			return queryObject, err
		}

		row := make(map[string]string)

		row["id"] = strconv.Itoa(id)
		row["tag"] = tag
		row["url"] = url
		row["item_desc"] = item_desc

		queryObject = append(queryObject, row)

	}
	fmt.Println(queryObject)

	return queryObject, nil
}

func InsertAction(tag string, url string, item_desc string) (int, error)   {
	t := time.Now().Local()
	creationTime := t.Format("2006-01-02 15:04:05")

	stmt, _ := dbCon.Prepare("INSERT INTO outgoing(tag, url, item_desc, creation_time) VALUES (?, ?, ?, ?)" )
	res, _ := stmt.Exec(tag, url, item_desc, creationTime)

	id, err := res.LastInsertId()

	return int(id), err

}

func DelAction(id string) (int, error)  {

	stmt, _ := dbCon.Prepare("DELETE FROM outgoing WHERE id = ?" )
	res, _ := stmt.Exec(id)

	deleteID, err := res.RowsAffected()

	return int(deleteID), err

}

func UpdateAction() {
	//todo
}