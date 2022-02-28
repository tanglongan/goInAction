package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// GO 提供了database/sql包用于对SQL数据库的访问
// 作为操作数据库的入口对象sql.DB，它主要有2个功能：通过数据库驱动为我们提供管理底层数据库连接的打开和关闭、为我们管理数据库连接池

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

type Cate struct {
	cid     int
	cname   string
	addtime int
	scope   int
}

func main() {
	dbw := DbWorker{Dsn: "root:nYhKT@2687zsmB@tcp(9.135.132.31:3306)/learn?charset=utf8mb4&&parseTime=True"}

	dbtemp, err := sql.Open("mysql", dbw.Dsn)
	dbw.Db = dbtemp

	if err != nil {
		panic(err)
	}

	defer dbw.Db.Close()

	dbw.insertData()

	dbw.deleteData()

	dbw.editData()

	dbw.queryData()
}

// insertData 添加数据
func (dbw *DbWorker) insertData() {
	stmt, _ := dbw.Db.Prepare(`INSERT INTO t_article_cate(cname,addtime,scope) VALUES(?,?,?)`)
	defer stmt.Close()

	for i := 1; i <= 10; i++ {
		ret, err := stmt.Exec(fmt.Sprintf("%s%d", "栏目", i), time.Now().Unix(), 10)

		if err != nil {
			fmt.Printf("insert data error: %v\n", err)
		}

		if LastInsertId, err := ret.LastInsertId(); err != nil {
			fmt.Println("LastInsertId: ", LastInsertId)
		}

		if RowsAffected, err := ret.RowsAffected(); err != nil {
			fmt.Println("RowsAffected: ", RowsAffected)
		}
	}

}

// deleteData 删除数据
func (dbw *DbWorker) deleteData() {
	stmt, _ := dbw.Db.Prepare(`DELETE FROM t_article_cate WHERE cid=?`)
	defer stmt.Close()
	ret, err := stmt.Exec(20)

	if err != nil {
		fmt.Printf("delete data error: %v\n", err)
		return
	}

	if RowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected: ", RowsAffected)
	}

}

// editData 更新数据
func (dbw *DbWorker) editData() {
	stmt, _ := dbw.Db.Prepare(`UPDATE t_article_cate SET scope=? WHERE cid=?`)
	defer stmt.Close()

	ret, err := stmt.Exec(123, 11)
	if err != nil {
		fmt.Printf("update data error:%v\n", err)
		return
	}

	if RowsAffected, err := ret.RowsAffected(); err != nil {
		fmt.Println("RowsAffected: ", RowsAffected)
	}

}

func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT cid, cname, addtime, scope From t_article_cate where status=?`)

	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query Data error: %v\n", err)
	}

	columns, _ := rows.Columns()
	fmt.Println(columns)

	for rows.Next() {
		var c Cate
		rows.Scan(&c.cid, &c.cname, &c.scope, &c.addtime)
		fmt.Println(c)
	}

}
