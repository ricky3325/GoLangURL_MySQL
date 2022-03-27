package main

import (
    "net/http"
    "encoding/json"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "time"
    "io/ioutil"
    "math"
    "strings"
    "errors"
    "github.com/gorilla/mux"
)

/*func main()  {

    http.HandleFunc("/login1", login1)
    http.HandleFunc("/login2", login2)
    http.ListenAndServe("0.0.0.0:8080", nil)
}*/

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/login1/{id}", login1)
    r.HandleFunc("/login1", login1)
    r.HandleFunc("/login2", login2)
    http.ListenAndServe("0.0.0.0:8080", r)
}

type FullData struct {
	ID  int
	Url string
	ExpireAt string
}

type Cumstomer struct {
	ID  int
	Username string
	Password string
}

type CretUrsho struct {
    Url         string `json:"url"`
    ExpireAt    string `json:"expireAt"`
}

type Resp struct {
    Code    string `json:"code"`
    Msg     string `json:"msg"`
}

type  Auth struct {
    Username string `json:"username"`
    Pwd      string   `json:"password"`
}

const (
	UserName     string = "root"
	Password     string = "12345"
	Addr         string = "mysql"
	Port         int    = 3306
	Database     string = "mydb"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

/*type App struct {
	MyDB    *sql.DB
}*/
var (
	MyDB         *sql.DB
)

const (
    alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    length   = uint64(len(alphabet))
)

//post接口接收json數據
func login1(writer http.ResponseWriter,  request *http.Request)  {
    /*var auth Auth
    if err := json.NewDecoder(request.Body).Decode(&auth); err != nil {
        request.Body.Close()
        log.Fatal(err)
    }
    var result  Resp
    if auth.Username == "admin" && auth.Pwd == "123456" {
        result.Code = "200"
        result.Msg = "登錄成功"
    } else {
        result.Code = "401"
        result.Msg = "賬戶名或密碼錯誤"
    }
    if err := json.NewEncoder(writer).Encode(result); err != nil {
        log.Fatal(err)
    }*/

    /*var result  Resp
    result.Code = "401"
    result.Msg = "登錄失敗"

    if err := json.NewEncoder(writer).Encode(result); err != nil {
        log.Fatal(err)
    }*/
    //var results []string
    if request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Error reading request body",
				http.StatusInternalServerError)
		}
		//results = append(results, string(body))

        var postData CretUrsho
        if err := json.Unmarshal(body, &postData); err != nil {   // Parse []byte to go struct pointer
            fmt.Println("Can not unmarshal JSON", err)
        }
        fmt.Println(postData.Url)

        x, y := autoAdd(postData.Url, postData.ExpireAt)
        z := x + y + "POST done"
		fmt.Fprint(writer, z)
	} else if request.Method == "GET"{
        vars := mux.Vars(request)
        id, ok := vars["id"]
        if !ok {
            fmt.Println("id is missing in parameters")
        }
        fmt.Println(`id := `, id)
        fmt.Fprint(writer, "GET done")
    }else {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//接收x-www-form-urlencoded類型的post請求或者普通get請求
func login2(writer http.ResponseWriter,  request *http.Request)  {
    request.ParseForm()
    username, uError :=  request.Form["username"]
    pwd, pError :=  request.Form["password"]

    var result  Resp
    if !uError || !pError {
        result.Code = "401"
        result.Msg = "登錄失敗"
    } else if username[0] == "admin" && pwd[0] == "0" {
        result.Code = "200"
        result.Msg = "0Connect"
        DbConnectSQL()
    }else if username[0] == "admin" && pwd[0] == "1" {
        result.Code = "200"
        result.Msg = "1CreateTable"
        CreateTable()
    }else if username[0] == "3"{
        result.Code = "200"
        result.Msg = "3ReadFullData"
        ReadFullData(pwd[0])
    }else if username[0] == "admin" && pwd[0] == "4" {
        result.Code = "200"
        result.Msg = "4SHOW_TABLES"
        SHOW_TABLES()
    }else {
        result.Code = "203"
        result.Msg = "賬戶名或密碼錯誤"
    }
    if err := json.NewEncoder(writer).Encode(result); err != nil {
        log.Fatal(err)
    }
}

func DbConnectSQL(){
    //組合sql連線字串
    conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
    //連接MySQL
    //DB, err := sql.Open("mysql", conn)
    //MyDB = DB
    err := error(nil)
    MyDB, err = sql.Open("mysql", conn)
    if err != nil {
        fmt.Println("connection to mysql failed:", err)
        return
    }else {
        fmt.Println("connected to mysql")
    }
    MyDB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
    MyDB.SetMaxOpenConns(MaxOpenConns)
    MyDB.SetMaxIdleConns(MaxIdleConns)
    //fmt.Println("connected to mysql")
}

/*func CreateTable() {
	sql := `CREATE TABLE cumstomer(
	id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	username VARCHAR(64),
	password VARCHAR(64),
	status INT(4),
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	); `

	if _, err := MyDB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}*/

func CreateTable() {
	sql := `CREATE TABLE urshoner(
	id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	url NVARCHAR(2084) NOT NULL,
	ExpireAt timestamp NOT NULL
	); `

	if _, err := MyDB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}

/*func AddCumstomer() {
	//sql := `insert INTO cumstomer(username,password) values('test','123456'); `
    result, err := MyDB.Exec("insert INTO cumstomer(username,password) values(?,?)", "test", "123456")
    //result, err := MyDB.Exec(sql)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
    //sql.Result 的LastInsertId()可取得AUTO_INCREMENT的值
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return
	}
	fmt.Println("Insert data id:", lastInsertID)

    //RowsAffected() 影響的資料筆數，如果很嚴謹的寫法會判斷RowsAffected()是否與新增的資料筆數一致
	rowsaffected, err := result.RowsAffected() 
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}*/

func ReadFullData(Num string) {
    var fullData FullData
    //單筆資料
	row := MyDB.QueryRow("select id,url,ExpireAt from urshoner where id=?", Num)
    //Scan對應的欄位與select語法的欄位順序一致
	if err := row.Scan(&fullData.ID, &fullData.Url, &fullData.ExpireAt); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
    fmt.Println("fullData.ID:", fullData.ID)
    fmt.Println("fullData.Url:", fullData.Url)
    fmt.Println("fullData.ExpireAt:", fullData.ExpireAt)
    A := Encode(uint64(fullData.ID))
    fmt.Println("Encode(fullData.ID):", A)
    B, C := Decode(A)
    fmt.Println("Decode(A)B:", B)
    fmt.Println("Decode(A)C:", C)
	fmt.Println("fullData:%+v:", fullData)
} 

func SHOW_TABLES() {
    sql := `SHOW TABLES;`
    
        if _, err := MyDB.Exec(sql); err != nil {
            fmt.Println("SHOW TABLES failed:", err)
            return
        }
        fmt.Println("SHOW TABLES successd")
} 

func autoAdd(Url string, ExpireAt string) (string, string) {
	//sql := `insert INTO cumstomer(username,password) values('test','123456'); `
    result, err := MyDB.Exec("insert INTO urshoner(url,ExpireAt) values(?,?)", Url, ExpireAt)
    //result, err := MyDB.Exec(sql)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return "fail", "fail"
	}
    //sql.Result 的LastInsertId()可取得AUTO_INCREMENT的值
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return "fail", "fail"
	}
	fmt.Println("Insert data id:", lastInsertID)

    //RowsAffected() 影響的資料筆數，如果很嚴謹的寫法會判斷RowsAffected()是否與新增的資料筆數一致
	rowsaffected, err := result.RowsAffected() 
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return "fail", "fail"
	}
	fmt.Println("Affected rows:", rowsaffected)
    x := Encode(uint64(lastInsertID))
    y := "http://localhost/login1/" + x
    return x, y
}

func Encode(number uint64) string {
    var encodedBuilder strings.Builder
    encodedBuilder.Grow(11)
  
    for ; number > 0; number = number / length {
       encodedBuilder.WriteByte(alphabet[(number % length)])
    }
  
    return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
    var number uint64
  
    for i, symbol := range encoded {
       alphabeticPosition := strings.IndexRune(alphabet, symbol)
  
       if alphabeticPosition == -1 {
          return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
       }
       number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
    }
  
    return number, nil
}