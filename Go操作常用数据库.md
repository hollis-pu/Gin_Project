# Go操作常用数据库

2023.09.27

学习教程：[Go Web开发进阶实战（gin框架）(共23:04:35)](https://study.163.com/course/introduction.htm?courseId=1210171207)

## 入门介绍

### database/sql包

`database/sql` 包是 Go 语言标准库中用于数据库操作的核心包之一。它提供了一种通用的接口，使你可以连接和操作不同类型的关系型数据库，包括 MySQL、PostgreSQL、SQLite、Microsoft SQL Server 等。以下是 `database/sql` 包的一些关键概念和功能：

1. **数据库连接池管理**：
   - `sql.Open(driverName, dataSourceName)`：用于打开与数据库的连接，其中 `driverName` 是数据库驱动程序的名称，`dataSourceName` 包含数据库连接的信息（如用户名、密码、主机地址、端口号和数据库名称）。这个函数返回一个 `*DB` 对象，表示与数据库的连接。
   - `*DB`：表示一个数据库连接池，它负责管理多个数据库连接，以便在需要时将连接分配给不同的数据库操作。

2. **执行 SQL 查询和操作**：
   - `(*DB).Query(query string, args ...interface{})`：用于执行查询操作，并返回一个 `*Rows` 对象，以便逐行读取结果。
   - `(*DB).Exec(query string, args ...interface{})`：用于执行非查询操作，如插入、更新和删除，并返回一个 `Result` 对象，表示执行的结果。

3. **处理查询结果**：
   - `*Rows`：用于处理查询结果集。你可以使用 `Next()` 方法逐行遍历结果，然后使用 `Scan()` 方法将结果映射到 Go 变量中。
   - `*Rows` 和 `*Result` 对象都应该在使用后关闭，通常使用 `defer` 语句来确保关闭。

4. **处理错误**：
   - 大多数函数都会返回一个错误，你应该始终检查这些错误以处理潜在的问题，如数据库连接失败、查询错误等。

5. **数据库驱动程序**：
   - `database/sql` 包本身只提供了通用的数据库接口，具体的数据库支持需要使用相应的数据库驱动程序。每个驱动程序都实现了 `database/sql/driver` 接口，以便与 `database/sql` 包协同工作。
   - 针对不同数据库的驱动程序通常需要单独导入。例如，要连接到 MySQL，你需要导入 `"github.com/go-sql-driver/mysql"`。

6. **事务管理**：
   - `(*DB).Begin()`：用于启动一个数据库事务。事务允许你执行一系列操作，然后要么全部提交，要么全部回滚。

7. **连接池管理**：
   - `(*DB).SetMaxOpenConns(n int)` 和 `(*DB).SetMaxIdleConns(n int)`：用于设置最大打开连接数和最大空闲连接数，以控制连接池的大小。

8. **预处理语句**：
   - `(*DB).Prepare(query string)`：用于创建预处理语句，可以多次执行，提高性能和安全性。

总之，`database/sql` 包提供了一个通用的、数据库无关的方式来连接和操作关系型数据库。你可以根据需要选择合适的数据库驱动程序，并使用标准的 Go 代码来执行数据库操作。这使得 Go 语言在处理数据库时非常灵活和强大。

### 安装Mysql驱动

在终端中运行以下命令来安装MySQL驱动程序包：

```go
go get -u github.com/go-sql-driver/mysql
```

这将下载并安装 "github.com/go-sql-driver/mysql" 包及其依赖项。

### Open()方法

`database/sql` 包中的 `Open` 函数用于打开与数据库的连接。它是连接到关系型数据库的第一步，返回一个 `*DB` 对象，代表了与数据库的连接。以下是关于 `Open` 函数的详细介绍：

```go
func Open(driverName, dataSourceName string) (*DB, error)
```

- `driverName`：字符串，表示要使用的数据库驱动程序的名称。不同的数据库类型使用不同的驱动程序。例如，如果要连接到 MySQL 数据库，可以使用 `"mysql"` 作为驱动程序名称。
- `dataSourceName`：字符串，包含连接数据库的详细信息，通常包括用户名、密码、主机地址、端口号和数据库名称。这个字符串的格式取决于所使用的数据库驱动程序和数据库类型。

**返回值**：

- `*DB`：表示一个数据库连接池，可以用于执行 SQL 查询和操作。它是 `database/sql` 包中的核心类型。它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用。
- `error`：如果连接过程中发生错误，则返回一个非 `nil` 的错误。否则，返回 `nil`。

**示例用法**：

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func main() {
    // 打开与 MySQL 数据库的连接
    db, err := sql.Open("mysql", "用户名:密码@tcp(localhost:3306)/数据库名")
    if err != nil {
        fmt.Println("数据库连接失败:", err)
        return
    }
    defer db.Close() // 确保在函数退出时关闭数据库连接

    // 现在可以使用 db 执行 SQL 查询和操作
}
```

在打开连接时，`Open` 函数只是验证其参数格式是否正确，不会实际连接到数据库，它会为后续的数据库操作创建了一个连接池对象。返回的DB对象可以安全地被多个goroutine并发使用，并且维护其自己的空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。

如果要检查数据源的名称是否真实有效，应该调用Ping方法。

### Ping()方法

`database/sql` 包中的 `Ping` 函数用于测试与数据库的连接是否仍然有效。它是一个在连接到数据库之后用于验证连接状态的有用工具。以下是关于 `Ping` 函数的详细介绍：

```go
func (db *DB) Ping() error
```

- `db`：一个 `*DB` 对象，表示与数据库的连接。你可以使用 `sql.Open` 函数打开连接并得到这个对象。

**返回值**：

- `error`：如果连接正常，返回 `nil`。如果连接失败或不可用，则返回一个非 `nil` 的错误。

**示例用法**：

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func main() {
    // 打开与 MySQL 数据库的连接
    db, err := sql.Open("mysql", "用户名:密码@tcp(localhost:3306)/数据库名")
    if err != nil {
        fmt.Println("数据库连接失败:", err)
        return
    }
    defer db.Close()

    // 使用Ping函数检查连接状态
    err = db.Ping()
    if err != nil {
        fmt.Println("连接不可用:", err)
        return
    }

    fmt.Println("数据库连接正常")
    // 现在可以继续执行其他数据库操作
}
```

在上面的示例中，`Ping` 函数用于检查数据库连接是否可用。如果连接正常，它将返回 `nil`，否则返回一个描述连接问题的错误。

### 连接参数设置

使用 `database/sql` 包连接数据库时，你可以通过配置来控制连接池和数据库操作的行为。以下是一些常用的配置选项：

1. **设置最大打开连接数（`SetMaxOpenConns`）**：
   - 使用 `(*DB).SetMaxOpenConns(n int)` 方法可以限制连接池中的最大打开连接数。这是为了防止过多的数据库连接占用系统资源。
   - 示例：`db.SetMaxOpenConns(10)`，表示最多同时打开 10 个数据库连接。
2. **设置最大空闲连接数（`SetMaxIdleConns`）**：
   - 使用 `(*DB).SetMaxIdleConns(n int)` 方法可以设置连接池中的最大空闲连接数。空闲连接是在连接池中保留但当前未使用的连接。
   - 示例：`db.SetMaxIdleConns(5)`，表示连接池中最多保留 5 个空闲连接。
3. **设置连接的最大生存时间（`SetConnMaxLifetime`）**：
   - 使用 `(*DB).SetConnMaxLifetime(d time.Duration)` 方法可以设置数据库连接的最大生存时间。连接在达到该生存时间后将被关闭，并在需要时重新创建。
   - 示例：`db.SetConnMaxLifetime(1 * time.Hour)`，表示数据库连接的最大生存时间为 1 小时。

### 匿名导入

```go
import (
    _ "github.com/go-sql-driver/mysql"
)
```

在Go语言中，匿名导入是通过使用下划线 "_" 来导入一个包，这意味着你导入了该包，但不直接引用它的标识符。匿名导入的主要目的是执行包的初始化代码。

具体到上面的例子，`_ "github.com/go-sql-driver/mysql"` 表示你导入了名为 "github.com/go-sql-driver/mysql" 的包，但不会直接使用该包中的标识符。这通常用于导入数据库驱动程序等需要在初始化时注册自身的包。

在这种情况下，当你执行程序时，Go 编译器会确保导入的包（在这里是 "github.com/go-sql-driver/mysql"）的初始化代码会被执行。初始化代码通常用于注册数据库驱动程序，以便在程序中使用 `database/sql` 包时可以选择该驱动程序。

`mysql`包`init()`函数内容如下：

```go
func init() {
	sql.Register("mysql", &MySQLDriver{})  // 注册mysql驱动
}
```

`Register`函数内容如下：

```go
func Register(name string, driver driver.Driver) {
    driversMu.Lock()  // 加锁
    defer driversMu.Unlock()
    if driver == nil {
       panic("sql: Register driver is nil")
    }
    if _, dup := drivers[name]; dup {
       panic("sql: Register called twice for driver " + name)
    }
    drivers[name] = driver
}
```

## mysql的CRUD操作

### 1.单行查询

```go
// 定义一个全局对象db
var db *sql.DB

// User 创建一个结构体和数据库表进行映射
type User struct {
	id   int
	name string
	age  int
}

// 1.单行查询：使用db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。
func queryRowData() {
	sqlStr := "select id,name,age from user where id=?"
	var user User
	row := db.QueryRow(sqlStr, 1) // 得到一行的记录
	// 要确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := row.Scan(&user.Id, &user.Name, &user.Age) // 将查询结果的列和结构体字段进行映射赋值（根据查询出来的列的顺序进行依次赋值）
	// 也可以直接链式调用: err := db.QueryRow(sqlStr, 1).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("user:", user) // user: {1 Tom 25}
}
```

### 2.多行查询

```go
// 2.多行查询：使用db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
func queryMultiRowData() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}
```

### 3.插入

```go
// 3.插入数据：插入、更新和删除操作都使用Exec方法。
func insertRowData() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38) // 第一个返回值类型为Result，保存了执行结果的信息（自增ID、影响行数）
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
```

### 4.更新

```go
// 4.更新数据
func updateRowData() {
	sqlStr := "update user set age=? where name = ?"
	ret, err := db.Exec(sqlStr, 30, "Tom")
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}
```

### 5.删除

```go
// 5.删除数据
func deleteRowData() {
	sqlStr := "delete from user where name = ?"
	ret, err := db.Exec(sqlStr, "王五")
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
```

## sql预处理

可以使用`database/sql`包来执行SQL预处理操作。预处理允许你提前定义SQL语句，并在需要时将参数传递给它，以安全地执行SQL查询和操作，同时防止SQL注入攻击。

**什么是预处理？**

普通SQL语句执行过程：

1. 客户端对SQL语句进行占位符替换得到完整的SQL语句。
2. 客户端发送完整SQL语句到MySQL服务端
3. MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：

1. 把SQL语句分成两部分，命令部分与数据部分。
2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
3. 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
4. MySQL服务端执行完整的SQL语句并将结果返回给客户端。

**为什么要预处理？**

1. 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
2. 避免SQL注入问题。

示例：

```go
// database/sql中使用下面的Prepare方法来实现预处理操作
func prepareQueryDemo() {
    sqlStr := "select id, name, age from user where id > ?"  // ? 是占位符，表示稍后会用真实的参数值替换它们
    stmt, err := db.Prepare(sqlStr)
    if err != nil {
       fmt.Printf("prepare failed, err:%v\n", err)
       return
    }
    defer stmt.Close()
    rows, err := stmt.Query(0)  // 使用真实值来替换占位符
    if err != nil {
       fmt.Printf("query failed, err:%v\n", err)
       return
    }
    defer rows.Close()
    // 循环读取结果集中的数据
    for rows.Next() {
       var u User
       err := rows.Scan(&u.Id, &u.Name, &u.Age)
       if err != nil {
          fmt.Printf("scan failed, err:%v\n", err)
          return
       }
       fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
    }
}

// 插入、更新和删除操作的预处理十分类似，这里以插入操作的预处理为例
// 预处理插入示例
func prepareInsertDemo() {
    sqlStr := "insert into user(name, age) values (?,?)"
    stmt, err := db.Prepare(sqlStr)
    if err != nil {
       fmt.Printf("prepare failed, err:%v\n", err)
       return
    }
    defer stmt.Close()
    _, err = stmt.Exec("Tony", 23)
    if err != nil {
       fmt.Printf("insert failed, err:%v\n", err)
       return
    }
    _, err = stmt.Exec("Alice", 19)
    if err != nil {
       fmt.Printf("insert failed, err:%v\n", err)
       return
    }
    fmt.Println("insert success.")
}

func main() {
    prepareQueryDemo()
    prepareInsertDemo()
}
```

## 事务

### 介绍

事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。A转账给B，这里面就需要执行两次update操作。

在MySQL中只有使用了`Innodb`数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。

### 事务的ACID

通常事务必须满足4个条件（ACID）：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。

|  条件  |                             解释                             |
| :----: | :----------------------------------------------------------: |
| 原子性 | 一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。 |
| 一致性 | 在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。 |
| 隔离性 | 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。 |
| 持久性 | 事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。 |

### 事务相关方法

Go语言中使用以下三个方法实现MySQL中的事务操作。 开始事务

```go
func (db *DB) Begin() (*Tx, error)
```

提交事务

```go
func (tx *Tx) Commit() error
```

回滚事务

```go
func (tx *Tx) Rollback() error
```

### 事务示例

下面的代码演示了一个简单的事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。

```go
// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
```

## sqlx库

2023.09.28

在项目中我们通常可能会使用`database/sql`连接MySQL数据库。本文借助使用`sqlx`实现批量插入数据的例子，介绍了`sqlx`中可能被你忽视了的`sqlx.In`和`DB.NamedExec`方法。

### sqlx介绍

在项目中我们通常可能会使用`database/sql`连接MySQL数据库。`sqlx`可以认为是Go语言内置`database/sql`的超集，它在优秀的内置`database/sql`基础上提供了一组扩展。这些扩展中除了大家常用来查询的`Get(dest interface{}, ...) error`和`Select(dest interface{}, ...) error`外还有很多其他强大的功能。

### 安装sqlx

```go
go get github.com/jmoiron/sqlx
```

### 基本使用

#### 连接数据库

```go
var db *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
```

#### 查询

查询单行数据示例代码如下：

```go
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}
```

查询多行数据示例代码如下：

```go
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
```

#### 插入、更新和删除

sqlx中的exec方法与原生sql中的exec使用基本一致：

```go
// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
```

#### NamedExec

`DB.NamedExec`方法用来绑定SQL语句与结构体或map中的同名字段。

```go
func insertUserDemo()(err error){
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "七米",
			"age": 28,
		})
	return
}
```

#### NamedQuery

与`DB.NamedExec`同理，这里是支持查询。

```go
func namedQuery(){
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "七米",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}
```

### 事务操作

对于事务操作，我们可以使用`sqlx`中提供的`db.Beginx()`和`tx.Exec()`方法。示例代码如下：

```go
func transactionDemo2()(err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err!= nil{
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err!=nil{
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}
```

### sqlx.In

`sqlx.In`是`sqlx`提供的一个非常方便的函数。

#### sqlx.In的批量插入示例

**表结构**

为了方便演示插入数据操作，这里创建一个`user`表，表结构如下：

```sql
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

**结构体**

定义一个`user`结构体，字段通过tag与数据库中user表的列一致。

```go
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}
```

**bindvars**（绑定变量）

查询占位符`?`在内部称为**bindvars（查询占位符）**,它非常重要。你应该始终使用它们向数据库发送值，因为它们可以防止SQL注入攻击。`database/sql`不尝试对查询文本进行任何验证；它与编码的参数一起按原样发送到服务器。除非驱动程序实现一个特殊的接口，否则在执行之前，查询是在服务器上准备的。因此`bindvars`是特定于数据库的:

- MySQL中使用`?`
- PostgreSQL使用枚举的`$1`、`$2`等bindvar语法
- SQLite中`?`和`$1`的语法都支持
- Oracle中使用`:name`的语法

`bindvars`的一个常见误解是，它们用来在sql语句中插入值。它们其实仅用于参数化，不允许更改SQL语句的结构。例如，使用`bindvars`尝试参数化列或表名将不起作用：

```go
// ？不能用来插入表名（做SQL语句中表名的占位符）
db.Query("SELECT * FROM ?", "mytable")
 
// ？也不能用来插入列名（做SQL语句中列名的占位符）
db.Query("SELECT ?, ? FROM people", "name", "location")
```

##### 自己拼接语句实现批量插入

比较笨，但是很好理解。就是有多少个User就拼接多少个`(?, ?)`。

```go
// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users) * 2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := DB.Exec(stmt, valueArgs...)
	return err
}
```

##### 使用sqlx.In实现批量插入

前提是需要我们的结构体实现`driver.Valuer`接口：

```go
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}
```

使用`sqlx.In`实现批量插入代码如下：

```go
// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := DB.Exec(query, args...)
	return err
}
```

##### 使用NamedExec实现批量插入

**注意** ：该功能需1.3.1版本以上，并且1.3.1版本目前还有点问题，sql语句最后不能有空格和`;`，详见[issues/690](https://github.com/jmoiron/sqlx/issues/690)。

使用`NamedExec`实现批量插入的代码如下：

```go
// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := DB.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}
```

把上面三种方法综合起来试一下：

```go
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	u1 := User{Name: "七米", Age: 18}
	u2 := User{Name: "q1mi", Age: 28}
	u3 := User{Name: "小王子", Age: 38}

	// 方法1
	users := []*User{&u1, &u2, &u3}
	err = BatchInsertUsers(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	}

	// 方法2
	users2 := []interface{}{u1, u2, u3}
	err = BatchInsertUsers2(users2)
	if err != nil {
		fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	}

	// 方法3
	users3 := []*User{&u1, &u2, &u3}
	err = BatchInsertUsers3(users3)
	if err != nil {
		fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	}
}
```

#### sqlx.In的查询示例

关于`sqlx.In`这里再补充一个用法，在`sqlx`查询语句中实现In查询和FIND_IN_SET函数。即实现`SELECT * FROM user WHERE id in (3, 2, 1);`和`SELECT * FROM user WHERE id in (3, 2, 1) ORDER BY FIND_IN_SET(id, '3,2,1');`。

##### in查询

查询id在给定id集合中的数据。

```go
// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int)(users []User, err error){
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = DB.Rebind(query)

	err = DB.Select(&users, query, args...)
	return
}
```

##### in查询和FIND_IN_SET函数

查询id在给定id集合的数据并维持给定id集合的顺序。

```go
// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int)(users []User, err error){
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = DB.Rebind(query)

	err = DB.Select(&users, query, args...)
	return
}
```

当然，在这个例子里面你也可以先使用`IN`查询，然后通过代码按给定的ids对查询结果进行排序。

## go-redis库

### 安装

Go 社区中目前有很多成熟的 redis client 库，比如 https://github.com/gomodule/redigo 和https://github.com/go-redis/redis，读者可以自行选择适合自己的库。这里使用 go-redis 这个库来操作 Redis 数据库。

使用以下命令下安装 go-redis 库。

```bash
go get github.com/go-redis/redis/v8
```

### 连接

#### 普通连接模式

go-redis 库中使用 redis.NewClient 函数连接 Redis 服务器。

```go
rdb := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // 密码
	DB:       0,  // 数据库
	PoolSize: 20, // 连接池大小
})
```

除此之外，还可以使用 redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。

```go
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

rdb := redis.NewClient(opt)
```

#### TLS连接模式

如果使用的是 TLS 连接方式，则需要使用 tls.Config 配置。

```go
rdb := redis.NewClient(&redis.Options{
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
		// Certificates: []tls.Certificate{cert},
    // ServerName: "your.domain.com",
	},
})
```

#### Redis Sentinel模式

使用下面的命令连接到由 Redis Sentinel 管理的 Redis 服务器。

```go
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},
})
```

#### Redis Cluster模式

使用下面的命令连接到 Redis Cluster，go-redis 支持按延迟或随机路由命令。

```go
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

    // 若要根据延迟或随机路由命令，请启用以下命令之一
    // RouteByLatency: true,
    // RouteRandomly: true,
})
```