package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {

	//数据库连接信息
	connStr := "postgres://postgres:password@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	//延迟关闭
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// db.Ping : 检测数据库是否连接成功
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	//调用创建表函数
	createProductTable(db)

	//向表中插入数据

	//product := Product{
	//	Name:      "book",
	//	Price:     15.55,
	//	Available: true,
	//}
	//pk := insertProduct(db, product)
	//fmt.Printf("ID = %d\n", pk)

	//数据库查询,返回一行

	//var name string
	//var price float64
	//var available bool
	//
	//query := "SELECT name,available,price FROM product WHERE id = $1"
	//err = db.QueryRow(query, 4).Scan(&name, &available, &price)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Fatalf("No rows found with ID %d", 4) //如果错误信息等于sql行错误,则返回: No rows found with ID 4
	//
	//	}
	//	log.Fatal(err)
	//}
	//fmt.Printf("Name: %s\n", name)
	//fmt.Printf("Available: %t\n", available) //%t 打印bool值 true or false
	//fmt.Printf("Price: %f\n", price)

	//数据库查询, 返回多行
	data := []Product{}

	rows, err := db.Query("SELECT name,available,price FROM product")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	//to scan DB vals
	var name string
	var available bool
	var price float64
	//循环rows
	for rows.Next() {
		err := rows.Scan(&name, &available, &price)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}
	fmt.Println("data: ", data)

}

func createProductTable(db *sql.DB) {
	/*Product Table
	- ID
	- Name
	- Price
	- Available
	- Date Created
	*/
	qyery := `CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
	price NUMERIC(6,2) NOT NULL,
    available BOOLEAN,
    cretaed timestamp DEFAULT NOW()
)`
	//db.Exec 执行后不会返回,执行返回需要使用db.Query
	//只需要接収err来判断SQL语句是否执行成功
	_, err := db.Exec(qyery)
	if err != nil {
		log.Fatal(err)
	}

}

//向表中插入数据

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name,price,available) 
		VALUES ($1,$2,$3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
