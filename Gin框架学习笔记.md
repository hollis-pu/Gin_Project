# Gin框架学习笔记

[toc]

2023.09.13

学习教程：[七米-基于gin框架和gorm的web开发实战](https://www.bilibili.com/video/BV1gJ411p7xC)

参考博客：[李文周的博客](https://www.liwenzhou.com/categories/Golang/)

## HTTP通信

Go语言是一门强大的编程语言，内置了用于HTTP通信的标准库。使用Go语言进行HTTP通信通常涉及以下步骤：

1. 导入必要的包：首先，您需要导入Go语言标准库中的`net/http`包，以便使用HTTP相关功能。

2. 创建HTTP路由器：您可以使用`http.NewServeMux()`创建一个HTTP路由器，以处理不同的HTTP请求。路由器将根据请求的URL路径分发请求给不同的处理函数。

3. 创建处理函数：您需要编写处理HTTP请求的函数。这些函数通常具有两个参数：`http.ResponseWriter`用于发送HTTP响应，以及`*http.Request`用于接收HTTP请求。

4. 注册处理函数：将处理函数与路由器中的特定路径相关联，以确保正确的处理程序用于特定请求。

5. 启动HTTP服务器：最后，您需要启动HTTP服务器，监听指定的端口，并开始接受和处理HTTP请求。

以下是一个简单的示例，演示了如何使用Go语言进行HTTP通信：

```go
package main

import (
    "fmt"
    "net/http"
)

// 处理HTTP请求的处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // 向客户端发送响应
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // 创建HTTP路由器
    mux := http.NewServeMux()

    // 注册处理函数
    mux.HandleFunc("/", helloHandler)

    // 启动HTTP服务器并监听端口
    port := 9090
    serverAddr := fmt.Sprintf(":%d", port)
    fmt.Printf("Server listening on port %d...\n", port)
    err := http.ListenAndServe(serverAddr, mux)
    if err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}
```

在这个示例中，我们创建了一个简单的HTTP服务器，它监听端口8080，并将所有请求路由到`helloHandler`函数。`helloHandler`函数会向客户端发送"Hello, World!"的响应。您可以在浏览器中访问`http://localhost:9090`来测试这个简单的HTTP服务器。

请注意，这只是一个基本示例。在实际应用中，您可能需要处理不同类型的HTTP请求，解析请求参数，设置响应头等等。 Go语言的`net/http`包提供了丰富的功能来满足各种HTTP通信需求。

## Gin框架入门

### 1.介绍

Gin是一个流行的Go语言Web框架，用于构建高性能的Web应用程序和API。它提供了一组强大的工具和功能，使开发者能够轻松地构建可伸缩、高效和可维护的Web应用程序。以下是Gin框架的详细介绍：

1. **轻量级和高性能**：Gin是一个轻量级的框架，它在性能上表现出色。它使用了快速的HTTP路由引擎，因此非常适合构建高性能的Web应用程序。

2. **HTTP路由**：Gin提供了强大而灵活的HTTP路由功能，允许您定义各种路由规则和参数。您可以轻松地定义RESTful API端点，并支持路径参数、查询参数和请求体参数的处理。

3. **中间件**：中间件是Gin的一个强大特性，它允许您在请求进入处理程序之前或响应离开处理程序之后执行某些操作。这使得实现身份验证、日志记录、错误处理等功能非常容易。

4. **错误处理**：Gin具有内置的错误处理机制，允许您定义全局的错误处理程序或为特定路由定义自定义错误处理。这有助于提供清晰的错误信息和响应。

5. **请求和响应处理**：Gin提供了方便的方法来解析请求参数、读取请求体、设置响应头和发送JSON、XML等响应。它支持自动绑定请求数据到Go结构体，使得处理请求变得非常简单。

6. **模板引擎**：虽然Gin本身不包含模板引擎，但它很容易与第三方模板引擎集成，例如HTML模板或JSON模板。

7. **WebSocket支持**：Gin也支持WebSocket，使您能够构建实时应用程序，如聊天应用或实时通知系统。

8. **热重载**：Gin具有热重载功能，当您修改代码时，可以自动重新加载应用程序，无需手动停止和启动服务器。

以下是一个简单的Gin示例，展示了如何创建一个Hello World的Web应用程序：

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // 创建Gin引擎
    r := gin.Default()

    // 定义路由处理函数
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    // 启动Web服务器
    r.Run(":8080")
}
```

要开始使用Gin框架，您需要首先安装它，可以使用以下命令：

```
go get -u github.com/gin-gonic/gin
```

然后，您可以根据您的需求构建更复杂的Web应用程序和API，利用Gin提供的丰富功能和中间件来实现各种功能。 Gin的文档和社区资源丰富，可帮助您更深入地了解如何使用它构建强大的Web应用程序。

### 2.官方文档和学习资料

要深入学习Gin框架，您可以查阅以下官方文档和其他学习资料：

1. **官方文档**：
   - [Gin框架官方文档](https://gin-gonic.com/docs/): 官方文档包含了Gin框架的详细说明、示例和API文档。这是学习Gin框架的最佳起点。
2. **GitHub仓库**：
   - [Gin GitHub仓库](https://github.com/gin-gonic/gin): 您可以在GitHub上查看Gin框架的源代码，并参与讨论或报告问题。
3. **Gin Web框架示例代码**（英文）：
   - [Gin Web框架示例代码 - GitHub](https://github.com/gin-gonic/examples): 在Gin框架的官方GitHub仓库中，您可以找到一些示例代码，包括REST API和WebSocket示例。
4. **Go 技术论坛**：
   - [Go 技术论坛](https://learnku.com/go/)：这个网站提供了关于Go Web开发的系列教程，其中包括了有关Gin框架的学习资源。

## RESTful API

REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

推荐阅读[阮一峰 理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)

简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

- `GET`用来获取资源
- `POST`用来新建资源
- `PUT`用来更新资源
- `DELETE`用来删除资源。

只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

例如，我们现在要编写一个管理书籍的系统，我们可以查询对一本书进行查询、创建、更新和删除等操作，我们在编写程序的时候就要设计客户端浏览器与我们Web服务端交互的方式和路径。按照经验我们通常会设计成如下模式：

| 请求方法 |     URL      |     含义     |
| :------: | :----------: | :----------: |
|   GET    |    /book     | 查询书籍信息 |
|   POST   | /create_book | 创建书籍记录 |
|   POST   | /update_book | 更新书籍信息 |
|   POST   | /delete_book | 删除书籍信息 |

同样的需求我们按照RESTful API设计如下：

| 请求方法 |  URL  |     含义     |
| :------: | :---: | :----------: |
|   GET    | /book | 查询书籍信息 |
|   POST   | /book | 创建书籍记录 |
|   PUT    | /book | 更新书籍信息 |
|  DELETE  | /book | 删除书籍信息 |

Gin框架支持开发RESTful API的开发。

```go
func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
}
```

开发RESTful API的时候我们通常使用[Postman](https://www.getpostman.com/)来作为客户端的测试工具。

## Gin渲染

参考博客：[Gin框架介绍及使用](https://www.liwenzhou.com/posts/Go/gin/)

### HTML渲染

我们首先定义一个存放模板文件的`templates`文件夹，然后在其内部按照业务分别定义一个`posts`文件夹和一个`users`文件夹。 `posts/index.html`文件的内容如下：

```html
{{define "posts/index.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}
```

`users/index.html`文件的内容如下：

```html
{{define "users/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>users/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}
```

Gin框架中使用`LoadHTMLGlob()`或者`LoadHTMLFiles()`方法进行HTML模板渲染。

```go
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
```

### 自定义模板函数

定义一个不转义相应内容的`safe`模板函数如下：

```go
func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})

	router.Run(":8080")
}
```

在`index.tmpl`中使用定义好的`safe`模板函数：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>修改模板引擎的标识符</title>
</head>
<body>
<div>{{ . | safe }}</div>
</body>
</html>
```

### 静态文件处理

当我们渲染的HTML文件中引用了静态文件时，我们只需要按照以下方式在渲染页面前调用`gin.Static`方法即可。

```go
func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
   // ...
	r.Run(":8080")
}
```

### 使用模板继承

Gin框架默认都是使用单模板，如果需要使用`block template`功能，可以通过`"github.com/gin-contrib/multitemplate"`库实现，具体示例如下：

首先，假设我们项目目录下的templates文件夹下有以下模板文件，其中`home.tmpl`和`index.tmpl`继承了`base.tmpl`：

```bash
templates
├── includes
│   ├── home.tmpl
│   └── index.tmpl
├── layouts
│   └── base.tmpl
└── scripts.tmpl
```

然后我们定义一个`loadTemplates`函数如下：

```go
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
```

我们在`main`函数中

```go
func indexFunc(c *gin.Context){
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context){
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main(){
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}
```

### 补充文件路径处理

关于模板文件和静态文件的路径，我们需要根据公司/项目的要求进行设置。可以使用下面的函数获取当前执行程序的路径。

```go
func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}
```

目前学习到第9集。

### 返回Json

2023.09.14

```go
func main() {
    r := gin.Default()

    // 方法1：使用map
    r.GET("/json01", func(c *gin.Context) {
       data := gin.H{ // gin.H is a shortcut for map[string]any
          "name":   "tom",
          "age":    26,
          "gender": "male",
       }
       c.JSON(http.StatusOK, data)
    })

    // 方法2：使用struct，使用tag来定制json字段的名称。
    r.GET("/json02", func(c *gin.Context) {
       var person struct {
          Name   string `json:"name"` // 这里一定要是可导出的，否则json获取不到数据。如果想要指定json字段的名称，可以使用Tag。
          Age    int    `json:"age"`
          Gender string `json:"gender"`
       }
       person.Name = "Jane"
       person.Age = 24
       person.Gender = "female"
       c.JSON(http.StatusOK, person) // 传入结构体类型的数据
    })

    r.Run(":9003")
}
```

## context组件

在Gin框架中，`context`是一个非常重要的组件，它用于处理HTTP请求和响应，以及在请求处理过程中传递数据和控制流。以下是对Gin框架中`context`的详细介绍：

1. **HTTP请求和响应管理**：
   `context`对象封装了HTTP请求和响应的所有信息，包括HTTP方法、请求路径、请求头、请求参数、响应状态码、响应头等。通过`context`，您可以轻松地访问和操作这些信息。

2. **参数解析**：
   Gin的`context`提供了方便的方法来解析HTTP请求中的参数，包括查询参数、表单数据、JSON请求体等。您可以使用`Bind`、`ShouldBind`等方法来将请求参数绑定到Go结构体中，从而方便地进行验证和处理。

3. **路由参数**：
   `context`还允许您从URL路径中提取参数。例如，如果您定义了一个路由`/user/:id`，则可以使用`context.Param("id")`来获取`:id`占位符的值。

4. **中间件支持**：
   Gin的中间件是一种机制，可以在请求处理过程中添加预处理逻辑，如身份验证、日志记录、错误处理等。`context`对象可以在中间件之间传递，以便在请求处理过程中共享数据和控制流。中间件可以通过`context`的方法来操作请求和响应。

5. **错误处理**：
   在处理请求期间，如果发生错误，您可以使用`context`的`Abort`、`JSON`等方法来处理错误并生成适当的响应。例如，您可以使用`context.AbortWithStatus`来中止请求处理并返回指定的HTTP状态码。

6. **响应生成**：
   使用`context`，您可以方便地生成HTTP响应，包括设置响应状态码、响应头和响应体。例如，您可以使用`context.JSON`、`context.String`等方法来生成JSON响应或纯文本响应。

7. **上下文数据传递**：
   您可以在`context`对象中存储自定义数据，这些数据在请求处理过程中可以跨中间件和处理函数传递。这对于在整个请求周期中共享信息非常有用。

8. **请求上下文的超时和取消**：
   Gin的`context`支持上下文的超时和取消。这意味着您可以设置一个超时，以确保长时间运行的请求不会无限期地等待响应。

总之，Gin框架的`context`是一个非常重要的组件，它为处理HTTP请求和响应提供了丰富的功能和工具，使得开发Web应用变得更加便捷和灵活。通过熟练使用`context`，您可以更好地控制和管理您的Web应用程序的行为。

## 获取参数

### 1.query string参数

在 Gin 框架中，你可以使用上下文对象（`c`）来获取查询字符串（query string）参数。查询字符串参数通常是在 URL 中以 `?` 后面的键值对形式传递的，如 `http://example.com/path?param1=value1&param2=value2`。

以下是如何在 Gin 框架中获取查询字符串参数的示例：

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // 定义一个路由处理程序，用于获取查询字符串参数
    router.GET("/query", func(c *gin.Context) {
        // 使用 c.DefaultQuery 方法获取查询字符串参数，并设置默认值（可选）
        param1 := c.DefaultQuery("param1", "default_value")

        // 使用 c.Query 方法获取查询字符串参数，如果参数不存在则返回空字符串
        param2 := c.Query("param2")

        // 打印获取到的参数值
        c.JSON(200, gin.H{
            "param1": param1,
            "param2": param2,
        })
    })

    router.Run(":8080")
}
```

在上述示例中，我们创建了一个路由 `/query`，在这个路由下定义了一个处理程序函数。在处理程序函数内部，我们使用 `c.DefaultQuery` 方法来获取查询字符串参数 `param1`，并设置了一个默认值。同时，使用 `c.Query` 方法来获取查询字符串参数 `param2`。然后，我们将获取到的参数值返回给客户端。

运行这个示例，你可以通过访问 `http://localhost:8080/query?param1=value1&param2=value2` 来测试获取查询字符串参数的功能。

### 2.Form表单参数

在 Gin 框架中，你可以使用上下文对象（`c`）来获取通过 POST 请求提交的表单参数。以下是如何获取表单参数的示例：

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // 定义一个路由处理程序，用于处理 POST 请求提交的表单参数
    router.POST("/submit", func(c *gin.Context) {
        // 使用 c.PostForm 方法获取表单参数
        param1 := c.PostForm("param1")
        param2 := c.PostForm("param2")

        // 打印获取到的参数值
        c.JSON(200, gin.H{
            "param1": param1,
            "param2": param2,
        })
    })

    router.Run(":8080")
}
```

在上述示例中，我们创建了一个路由 `/submit`，在这个路由下定义了一个处理程序函数。在处理程序函数内部，我们使用 `c.PostForm` 方法来获取表单参数 `param1` 和 `param2` 的值。然后，我们将获取到的参数值返回给客户端。

要测试这个示例，你可以使用一个工具（例如，Postman 或 cURL）来发送 POST 请求，将表单参数发送到 `http://localhost:8080/submit`。

注意：在实际应用中，你可能需要先检查参数是否存在，以及进行错误处理等操作，以确保应用的安全性和稳定性。

### 3.路径参数

在 Gin 框架中，你可以使用路由参数来获取路径中的参数。路径参数是路由中的一部分，可以通过路由模式中的占位符来捕获。以下是如何获取路径参数的示例：

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // 定义一个路由，包含路径参数
    router.GET("/user/:id", func(c *gin.Context) {
        // 使用 c.Param 方法来获取路径参数
        userID := c.Param("id")

        // 打印获取到的路径参数值
        c.JSON(200, gin.H{
            "userID": userID,
        })
    })

    router.Run(":8080")
}
```

在上述示例中，我们创建了一个路由 `/user/:id`，其中 `:id` 是路径参数的占位符。在路由处理程序函数内部，我们使用 `c.Param` 方法来获取路径参数的值，并将其存储在 `userID` 变量中。然后，我们将获取到的路径参数值返回给客户端。

你可以通过访问类似 `http://localhost:8080/user/123` 的 URL 来测试路径参数的功能，其中 `123` 是路径参数的值。

如果你需要在路由中使用多个路径参数，只需在路由模式中添加更多的占位符，然后在处理程序函数中使用 `c.Param` 方法来获取它们的值。例如，`/user/:id/:name` 可以捕获两个路径参数 `id` 和 `name` 的值。

### 4.参数绑定

在 Gin 框架中，参数绑定是一种非常方便的方式，用于将 HTTP 请求中的参数（例如查询字符串参数、表单参数、路径参数）绑定到 Go 结构体的字段上，以便于处理请求和验证输入数据。Gin 框架支持多种参数绑定方式，包括 Query 参数绑定、表单参数绑定、JSON 参数绑定和路径参数绑定等。

以下是 Gin 框架中参数绑定的示例：

假设我们有一个简单的 Go 结构体定义如下：

```go
type User struct {
    ID   int    `form:"id" json:"id" binding:"required"`
    Name string `form:"name" json:"name" binding:"required"`
}
```

然后，我们可以通过不同的方式绑定参数到这个结构体。

1. **Query 参数绑定**：从查询字符串参数中绑定数据。

```go
func GetUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindQuery(&user); err == nil {
        // 在这里可以使用 user 结构体中的字段
        c.JSON(http.StatusOK, user)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
```

2. **表单参数绑定**：从 POST 请求的表单参数中绑定数据。

```go
func CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBind(&user); err == nil {
        // 在这里可以使用 user 结构体中的字段
        c.JSON(http.StatusCreated, user)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
```

3. **JSON 参数绑定**：从 JSON 请求体中绑定数据。

```go
func UpdateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err == nil {
        // 在这里可以使用 user 结构体中的字段
        c.JSON(http.StatusOK, user)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
```

4. **路径参数绑定**：从路径参数中绑定数据。

```go
func GetUserByID(c *gin.Context) {
    var user User
    if err := c.ShouldBindUri(&user); err == nil {
        // 在这里可以使用 user 结构体中的字段
        c.JSON(http.StatusOK, user)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
```

这些示例展示了不同的参数绑定方式。在每个示例中，我们使用 `ShouldBind`、`ShouldBindQuery`、`ShouldBindJSON` 和 `ShouldBindUri` 方法来将请求参数绑定到 `User` 结构体中，并进行错误处理。在实际应用中，你可以根据请求的内容选择合适的参数绑定方式来处理数据。同时，你可以使用 `binding` 标签来定义参数的验证规则，以确保数据的完整性和正确性。

## 文件上传

在 Gin 框架中进行文件上传是一项常见的任务，它允许你接收客户端上传的文件并对其进行处理。以下是如何在 Gin 框架中执行文件上传的示例：

首先，确保你已经导入了 Gin 框架和 `multipart/form-data` 包：

```go
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
```

然后，你可以创建一个路由来处理文件上传请求。下面是一个示例：

```go
func main() {
    r := gin.Default()

    // 设置上传文件的最大大小（可选）
    r.MaxMultipartMemory = 8 << 20 // 8 MB

    // 创建一个路由，用于处理文件上传
    r.POST("/upload", func(c *gin.Context) {
        // 从请求中获取上传的文件
        file, err := c.FormFile("file")

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        // 你可以在这里执行对上传文件的操作，比如保存到服务器或处理文件内容

        // 示例：将上传的文件保存到服务器
        filePath := "uploads/" + file.Filename
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "File uploaded successfully",
            "filename": file.Filename,
        })
    })

    r.Run(":8080")
}
```

在上述示例中，我们首先创建了一个路由 `/upload`，用于处理文件上传请求。在处理程序函数内部，我们使用 `c.FormFile` 方法来获取上传的文件。然后，我们可以选择保存文件到服务器或进行其他操作。

在示例中，我们使用 `c.SaveUploadedFile` 方法将上传的文件保存到服务器。你需要指定一个目标文件路径，这里我们将文件保存在 `uploads/` 目录下。

确保你在服务器上创建了 `uploads/` 目录，以便保存上传的文件。

最后，我们返回一个 JSON 响应，通知客户端文件上传成功。

要测试这个示例，你可以使用表单提交一个包含文件上传字段的 POST 请求。如果你使用 Postman 等工具，确保将请求类型设置为 `multipart/form-data`，并添加一个文件字段（通常命名为 `file`）。

这是一个简单的文件上传示例，你可以根据实际需求对文件进行更复杂的处理。同时，你也可以添加文件类型验证、文件大小限制等更多的安全措施。

## 重定向

在 Gin 框架中，你可以实现 HTTP 重定向和路由重定向，它们有一些区别：

### 1.HTTP 重定向

HTTP 重定向是通过发送特定的 HTTP 响应来实现的，通常使用 3xx 状态码来指示重定向。在 Gin 中，你可以使用 `c.Redirect` 方法来执行 HTTP 重定向。

```go
func RedirectHandler(c *gin.Context) {
    c.Redirect(http.StatusFound, "/new-location")
}
```

- `c` 是 Gin 上下文对象。
- `http.StatusFound` 是 HTTP 302 Found 状态码，用于指示重定向。
- `"/new-location"` 是重定向目标 URL。

**HTTP 重定向是在客户端与服务器之间进行的，客户端收到重定向响应后，会根据响应头中的新地址重新发起请求。**

> 用户侧的URL地址会发生改变。

### 2.路由重定向

路由重定向是指在应用程序的路由层级进行的重定向，它是通过修改路由规则来实现的。在 Gin 中，你可以使用 `c.Request.URL.Path` 来进行路由重定向。

```go
func RouteRedirectHandler(c *gin.Context) {
    // 使用 c.Request.URL.Path 进行路由重定向
    c.Request.URL.Path = "/new-route"
    router.HandleContext(c)
}
```

**路由重定向是在服务器端进行的，客户端不需要知道关于重定向的任何信息**，它只需向原始路径发出请求，服务器会根据路由规则将其重定向到新的路径。

> 用户侧的URL地址不会发生变化。

总的来说，HTTP 重定向是通过发送 HTTP 响应告知客户端进行重定向，而路由重定向是在服务器端通过修改路由规则将请求重定向到新的路由处理程序函数。在 Gin 中，你可以根据具体需求选择使用哪种方式来实现重定向。

## 路由

在 Gin 框架中，路由可以分为以下几种类别：

1. **基本路由**：

   基本路由是 Gin 中最简单的路由形式，它指定了请求的 HTTP 方法（GET、POST、PUT、DELETE 等）和路径。基本路由由 `GET`、`POST`、`PUT`、`DELETE` 等方法调用创建。

   ```go
   router.GET("/get", func(c *gin.Context) {
       // 处理 GET 请求
   })
   
   router.POST("/post", func(c *gin.Context) {
       // 处理 POST 请求
   })
   
   // 同样的方式创建其他 HTTP 方法的路由
   ```

   > `Any` 和 `NoRoute` 是两个用于处理路由的特殊方法，它们具有不同的用途：
   >
   > 1. **Any 方法**：
   >
   >    `Any` 方法用于注册一个路由处理器，该处理器会匹配所有 HTTP 请求方法（GET、POST、PUT、DELETE 等）。这意味着你可以使用 `Any` 方法来处理多个不同 HTTP 方法的请求。
   >
   >    ```go
   >    router.Any("/route", func(c *gin.Context) {
   >        // 这个处理程序将匹配所有 HTTP 方法的请求
   >    })
   >    ```
   >
   >    这对于需要在多个不同的请求方法下执行相同操作的路由非常有用。然而，需要注意的是，在某些情况下，你可能需要谨慎使用 `Any` 方法，因为它可能导致代码不够清晰。
   >
   > 2. **NoRoute 方法**：
   >
   >    `NoRoute` 方法用于设置一个处理程序，该处理程序会在请求的路径没有匹配到任何路由时执行。换句话说，当客户端请求的路径在路由中没有匹配项时，将触发 `NoRoute` 处理程序。
   >
   >    ```go
   >    router.NoRoute(func(c *gin.Context) {
   >        // 这个处理程序将在没有匹配路由的情况下执行
   >        c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
   >    })
   >    ```
   >
   >    `NoRoute` 处理程序通常用于自定义 404 错误页面或返回 JSON 404 响应。它允许你更精确地控制在请求未匹配到任何路由时客户端会收到什么响应。
   >

2. **路由参数**：

   路由参数允许你在路径中包含变量，这些变量可以捕获 URL 中的特定部分。例如，`:id` 可以用于捕获数字或字符串等。

   ```go
   router.GET("/user/:id", func(c *gin.Context) {
       userID := c.Param("id")
       // 使用 userID 处理请求
   })
   ```

3. **查询字符串参数**：

   查询字符串参数是通过 URL 的查询字符串传递的参数，可以通过 `c.Query` 或 `c.DefaultQuery` 方法来获取。

   ```go
   router.GET("/search", func(c *gin.Context) {
       query := c.Query("q")
       // 使用 query 处理请求
   })
   ```

4. **分组路由**：

   分组路由允许你将多个路由分组在一起，以便在它们之间应用相同的中间件或共享相同的路由前缀。这有助于组织和管理路由。

   ```go
   v1 := router.Group("/v1")
   {
       v1.GET("/user", func(c *gin.Context) {
           // 处理 v1 版本的用户路由
       })
   
       v1.GET("/post", func(c *gin.Context) {
           // 处理 v1 版本的文章路由
       })
   }
   ```

5. **中间件**：

   Gin 框架支持中间件，它们可以在路由处理之前或之后执行特定的逻辑。中间件可以用于验证身份、日志记录、权限控制等。

   ```go
   // 中间件示例
   router.GET("/admin", AuthMiddleware(), AdminHandler)
   ```

6. **分组中的中间件**：

   你可以在分组中为一组路由应用中间件，以便将相同的中间件应用于多个路由。

   ```go
   // 在分组中使用中间件
   v1 := router.Group("/v1")
   v1.Use(AuthMiddleware())
   {
       v1.GET("/user", UserHandler)
       v1.GET("/post", PostHandler)
   }
   ```

7. **静态文件服务**：

   Gin 框架允许你为静态文件（如图像、样式表、脚本等）提供服务，以便可以通过 HTTP 访问它们。这可以通过 `gin.Static` 方法来实现。

   ```go
   router.Static("/static", "/path/to/static/files")
   ```

这些是 Gin 框架中常见的路由类别。你可以根据应用程序的需求组合和使用这些不同类型的路由。 Gin 的灵活性使得可以创建复杂的路由结构，以满足各种需求。

## 中间件

在 Gin 框架中，中间件是一种机制，它允许你在请求到达路由处理程序之前或之后执行一些逻辑。中间件通常用于执行一些公共的任务，例如身份验证、日志记录、数据验证、权限控制等。Gin 中间件是一组函数，可以按照指定的顺序应用到路由上。

以下是 Gin 框架中间件的详细介绍：

1. **创建中间件**：

   中间件是一个函数，其签名如下：

   ```go
   func(c *gin.Context)
   ```

   中间件函数接受一个 Gin 上下文对象 `c`，它允许你访问请求信息、响应信息以及进行处理。

2. **应用中间件**：

   在 Gin 中，可以使用 `Use` 方法来将中间件应用到路由组或单个路由上。例如：

   ```go
   // 将中间件应用到所有路由组上
   router.Use(LoggerMiddleware(), AuthMiddleware())
   
   // 将中间件应用到路由分组中
   v1 := router.Group("/v1")
   v1.Use(LoggerMiddleware(),AuthMiddleware())
   
   // 将中间件应用到单个路由上
   router.GET("/admin", AuthMiddleware(), AdminHandler)
   ```

   注意，中间件的执行顺序与它们添加到路由上的顺序有关，按照添加的顺序依次执行。

   > 你可以将中间件应用到全局（对所有路由有效）或仅应用到特定的路由组或路由上。使用 `Use` 方法添加到全局中间件，使用 `Group` 方法创建路由组并在路由组上应用中间件，或者在单个路由上使用 `Use` 方法应用中间件。

3. **中间件示例**：

   下面是一些常见的中间件示例：

   - **日志记录中间件**：用于记录请求和响应信息，以便调试和监控。

   - **身份验证中间件**：用于验证用户身份，通常检查令牌或会话是否有效。

   - **权限控制中间件**：用于检查用户是否有权限执行特定操作。

   - **CORS 中间件**：用于处理跨域资源共享 (CORS) 请求，以确保安全的跨域通信。

   - **缓存中间件**：用于缓存响应，以减少服务器负载和提高性能。

   - **数据验证中间件**：用于验证请求数据的有效性，防止恶意输入。

4. **中间件的执行流程**：

   Gin 中间件的执行流程是串行的，按照添加的顺序依次执行。在每个中间件中，你可以执行一些前置操作，然后使用 `c.Next()` 将控制权传递给下一个中间件或路由处理程序。如果不调用 `c.Next()`，后续中间件或路由将不会执行，请求会被终止。

   ```go
   func MyMiddleware(c *gin.Context) {
       // 前置操作
       c.Next() // 传递控制权给下一个中间件或路由处理程序
       // 后置操作
   }
   ```

   如果需要在中间件中执行一些操作后终止请求流程，可以使用 `c.Abort()`，它会阻止后续中间件和路由的执行。

5. **中间件的参数**：

   你可以将参数传递给中间件函数，以便在中间件中访问它们。例如，如果你需要在中间件中传递配置信息，可以这样做：

   ```go
   func CustomMiddleware(config Config) gin.HandlerFunc {
       return func(c *gin.Context) {
           // 使用 config 参数执行中间件操作
           c.Next()
       }
   }
   
   // 在路由中应用带有参数的中间件
   router.Use(CustomMiddleware(myConfig))
   ```

6. **内置中间件**：

   Gin 框架提供了一些内置的中间件，如 `gin.Logger()` 用于记录请求日志、`gin.Recovery()` 用于处理恐慌和崩溃等。你可以使用这些内置中间件来快速添加常见功能。

7. **中间件中的错误处理**：

   如果中间件中发生了错误，你可以使用 `c.AbortWithError` 来中止请求并返回错误响应。这可以用于处理身份验证失败、权限不足等情况。

   ```go
   func AuthMiddleware(c *gin.Context) {
       if !isAuthenticated(c) {
           c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
           return
       }
       c.Next()
   }
   ```

总的来说，Gin 框架中间件提供了一种灵活和强大的方式来管理和处理请求。你可以根据应用程序的需求创建自定义中间件，并将它们应用到不同的路由或路由组上，以实现各种功能和控制逻辑。

## Gorm

### 1.ORM介绍

ORM 是对象关系映射（Object-Relational Mapping）的缩写，它是一种用于处理关系型数据库（如MySQL、PostgreSQL、SQLite等）和面向对象编程语言之间数据转换的技术。

关系型数据库使用表来存储数据，每个表都有特定的列和行。面向对象编程语言（如Java、Python、Go等）使用对象来表示数据和行为。ORM 技术的主要目标是解决关系数据库和面向对象编程语言之间的数据不匹配问题，将数据库中的数据映射到对象上，以便更轻松地进行数据操作。

以下是 ORM 技术的一些主要特点和好处：

1. **对象表示数据**：ORM 允许开发人员使用对象来表示数据库中的数据，使数据操作更直观、自然。

2. **抽象数据库差异**：ORM 库通常提供了一种抽象层，使开发人员可以编写与数据库无关的代码。这意味着你可以在不更改代码的情况下切换不同的数据库系统。

3. **自动生成 SQL**：ORM 库通常能够自动生成 SQL 查询和操作语句，减少了手动编写 SQL 的工作，同时也有助于防止 SQL 注入攻击。

4. **关系管理**：ORM 提供了方便的方式来管理数据库表之间的关系，包括一对一、一对多、多对多等关系。

5. **数据验证**：ORM 库通常提供了数据验证功能，帮助开发人员确保数据的完整性和正确性。

6. **性能优化**：一些高级 ORM 库具有性能优化功能，可以生成高效的 SQL 查询，并使用缓存等技术来提高查询速度。

7. **跨平台和多语言支持**：ORM 库通常支持多种编程语言和不同的操作系统，使其适用于多种开发环境。

8. **减少重复代码**：ORM 可以减少编写重复数据库访问代码的工作，提高了开发效率。

一些流行的 ORM 框架包括 Hibernate（Java）、Entity Framework（C#）、Django ORM（Python）、SQLAlchemy（Python）、GORM（Go）等。

### 2.GORM介绍

GORM（Golang Object Relational Mapping）是一个流行的 Go 语言 ORM 框架，用于处理关系型数据库的数据操作。它提供了强大的工具和功能，使开发人员能够更轻松地在 Go 语言中进行数据库操作，同时减少了与数据库交互的重复工作。

以下是 GORM 的一些主要特点和功能：

1. **数据库无关性**：GORM 支持多种关系型数据库，包括 MySQL、PostgreSQL、SQLite、SQL Server 等。这意味着你可以使用相同的 GORM 代码与不同类型的数据库交互。

2. **模型定义**：你可以创建 Go 结构体来定义数据库表的模型，并使用 GORM 的标签来指定字段名称、类型、约束等信息。这使得模型的定义非常直观。

   ```go
   type User struct {
       gorm.Model
       Name  string
       Email string `gorm:"unique"`
   }
   ```

3. **CRUD 操作**：GORM 提供了一组强大的方法，用于执行数据库的增、删、改、查操作，包括创建记录、查询记录、更新记录和删除记录。

   ```go
   // 创建记录
   db.Create(&user)

   // 查询记录
   db.First(&user, 1)

   // 更新记录
   db.Model(&user).Update("Name", "NewName")

   // 删除记录
   db.Delete(&user)
   ```

4. **事务支持**：GORM 允许你使用事务来确保一组操作要么全部成功，要么全部失败。这对于保持数据库的一致性非常重要。

   ```go
   tx := db.Begin()
   // 执行数据库操作
   if err := tx.Commit().Error; err != nil {
       tx.Rollback()
   }
   ```

5. **自动迁移**：GORM 支持自动迁移功能，可以根据模型定义自动创建数据库表，这减少了手动创建表的工作。

   ```go
   db.AutoMigrate(&User{})
   ```

6. **关联关系**：GORM 支持定义和管理数据库表之间的关联关系，包括一对一、一对多、多对多等。

   ```go
   type User struct {
       gorm.Model
       Name     string
       Email    string `gorm:"unique"`
       Articles []Article
   }

   type Article struct {
       gorm.Model
       Title   string
       Content string
       UserID  uint
   }
   ```

7. **钩子函数**：你可以在模型中定义钩子函数，以在记录创建、更新、删除等事件发生时执行自定义逻辑。

   ```go
   func (u *User) BeforeCreate(tx *gorm.DB) error {
       // 在创建记录之前执行自定义逻辑
   }
   ```

8. **复杂查询**：GORM 提供了丰富的查询方法，支持复杂的条件查询、排序、分页、预加载等。

   ```go
   db.Where("age > ?", 18).Order("created_at desc").Limit(10).Find(&users)
   ```

9. **批量操作**：你可以使用 GORM 执行批量插入、更新、删除等操作，从而提高了数据库操作的效率。

   ```go
   db.Create(&users)
   db.Model(&users).Update("Status", "inactive")
   db.Delete(&users)
   ```

总的来说，GORM 是一个功能强大且易于使用的 Go 语言 ORM 框架，它使开发人员能够轻松地与关系型数据库进行交互，同时提供了丰富的功能和灵活的查询语言。如果你在使用 Go 语言开发应用程序，并需要与数据库交互，GORM 是一个很好的选择。它的活跃社区和广泛的文档也使得学习和使用 GORM 更加容易。

官方文档：https://gorm.io/zh_CN/

GitHub：https://github.com/go-gorm/gorm

目前学习到第18集。

### 3.入门使用

2023.09.15

学习教程：[枫枫知道-golang最简单的gorm教程](https://www.bilibili.com/video/BV1xg411t7RZ)

参考博客：[枫枫知道-Gorm文档](https://docs.fengfengzhidao.com/#/docs/gorm%E6%96%87%E6%A1%A3/1.%E8%BF%9E%E6%8E%A5)

#### 1.下载Gorm和数据库驱动

```go
go get gorm.io/driver/mysql
go get gorm.io/gorm
```

#### 2.简单连接

```go
username := "root"  //账号
password := "root"  //密码
host := "127.0.0.1" //数据库地址，可以是Ip或者域名
port := 3306        //数据库端口
Dbname := "gorm"   //数据库名
timeout := "10s"    //连接超时，10秒

// root:root@tcp(127.0.0.1:3306)/gorm?
dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
db, err := gorm.Open(mysql.Open(dsn))
if err != nil {
  panic("连接数据库失败, error=" + err.Error())
}
// 连接成功
fmt.Println(db)
```

### 4.高级配置

#### 跳过默认事务

为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这样可以获得60%的性能提升。

```go
db, err := gorm.Open(mysql.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})
```

#### 命名策略

gorm采用的命名策略是，表名是蛇形复数，字段名是蛇形单数。默认情况下，名为 `ID` 的字段会作为表的主键。

例如：

```go
type Student struct {
  Name      string
  Age       int
  MyStudent string
}
```

gorm会为我们这样生成表结构：

```go
CREATE TABLE `students` (`name` longtext,`age` bigint,`my_student` longtext)
```

我们也可以修改这些策略：

```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
  NamingStrategy: schema.NamingStrategy{
    TablePrefix:   "f_",  // 表名前缀
    SingularTable: false, // 单数表名
    NoLowerCase:   false, // 关闭小写转换
  },
})
```

> 也可以指定需要绑定的表的名称
>
> ```go
> func (Student) TableName() string { // 指定表名
>   return "student"
> }
> ```

#### 显示日志

gorm的默认日志是只打印错误和慢SQL

我们可以自己设置

```go
var mysqlLogger logger.Interface
// 要显示的日志等级
mysqlLogger = logger.Default.LogMode(logger.Info)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
  Logger: mysqlLogger,
})
```

如果你想自定义日志的显示，那么可以使用如下代码：

```go
newLogger := logger.New(
  log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
  logger.Config{
    SlowThreshold:             time.Second, // 慢 SQL 阈值
    LogLevel:                  logger.Info, // 日志级别
    IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
    Colorful:                  true,        // 使用彩色打印
  },
)

db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
  Logger: newLogger,
})
```

部分展示日志

```go
var model Student
session := DB.Session(&gorm.Session{Logger: newLogger})
session.First(&model)
// SELECT * FROM `students` ORDER BY `students`.`name` LIMIT 1
```

如果只想某些语句显示日志

```go
DB.Debug().First(&model)
```

### 5.模型

模型是标准的 struct，由 Go 的基本数据类型、实现了 `Scanner` 和 `Valuer` 接口的自定义类型及其指针或别名组成

#### 定义一张表

```go
type Student struct {
  ID    uint // 默认使用ID作为主键
  Name  string
  Email *string // 使用指针是为了存空值
}
```

常识：小写属性是不会生成字段的。

#### 自动生成表结构

```go
// 可以放多个
DB.AutoMigrate(&Student{})
```

`AutoMigrate`的逻辑是**只新增，不删除，不修改（大小会修改）**。

例如：将Name修改为Name1，进行迁移，会多出一个name1的字段。

生成的表结构如下：

```SQL
CREATE TABLE `f_students` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,`email` longtext,PRIMARY KEY (`id`))
```

默认的类型太大了。

#### 修改大小

有两种方式：

```go
Name  string  `gorm:"type:varchar(12)"`
Name  string  `gorm:"size:2"`
```

字段标签

> `type` 定义字段类型
> `size` 定义字段大小
> `column` 自定义列名
> `primaryKey` 将列定义为主键
> `unique` 将列定义为唯一键
> `default` 定义列的默认值
> `not null` 不可为空
> `embedded` 嵌套字段
> `embeddedPrefix` 嵌套字段前缀
> `comment` 注释

多个标签之前用 `;` 连接

```go
type StudentInfo struct {
  Email  *string `gorm:"size:32"` // 使用指针是为了存空值
  Addr   string  `gorm:"column:y_addr;size:16"`
  Gender bool    `gorm:"default:true"`
}
type Student struct {
  Name string      `gorm:"type:varchar(12);not null;comment:用户名"`
  UUID string      `gorm:"primaryKey;unique;comment:主键"`
  Info StudentInfo `gorm:"embedded;embeddedPrefix:s_"`
}

// 建表语句
CREATE TABLE `students` (
    `name` varchar(12) NOT NULL COMMENT '用户名',
    `uuid` varchar(191) UNIQUE COMMENT '主键',
    `s_email` varchar(32),
    `s_y_addr` varchar(16),
    `s_gender` boolean DEFAULT true,
    PRIMARY KEY (`uuid`)
)

```

### 6.单表CRUD

先使用gorm对单张表进行增删改查

表结构

```go
type Student struct {
  ID     uint   `gorm:"size:3"`
  Name   string `gorm:"size:8"`
  Age    int    `gorm:"size:3"`
  Gender bool
  Email  *string `gorm:"size:32"`
}
```

#### 1.插入操作

##### 1.添加记录

```go
email := "xxx@qq.com"
// 创建记录
student := Student{
  Name:   "枫枫",
  Age:    21,
  Gender: true,
  Email:  &email,
}
DB.Create(&student)
```

有两个地方需要注意

1. 指针类型是为了更好的存null类型，但是传值的时候，也记得传指针
2. Create接收的是一个指针，而不是值

由于我们传递的是一个指针，调用完Create之后，student这个对象上面就有该记录的信息了，如创建的id

```go
DB.Create(&student)
fmt.Printf("%#v\n", student)  
// main.Student{ID:0x2, Name:"zhangsan", Age:23, Gender:false, Email:(*string)(0x11d40980)}
```

##### 2.批量插入

Create方法还可以用于插入多条记录

```go
var studentList []Student
for i := 0; i < 100; i++ {
  studentList = append(studentList, Student{
    Name:   fmt.Sprintf("机器人%d号", i+1),
    Age:    21,
    Gender: true,
    Email:  &email,
  })
}
DB.Create(&studentList)  // 使用切片进行批量插入
```

#### 2.查询操作

##### 1.查询单条记录

```go
var student Student
DB.Take(&student)
fmt.Println(student)
```

获取单条记录的方法很多，我们对比sql就很直观了

```go
DB = DB.Session(&gorm.Session{Logger: Log})
var student Student
DB.Take(&student)  
// SELECT * FROM `students` LIMIT 1
DB.First(&student) 
// SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
DB.Last(&student)  
// SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
```

##### 2.根据主键查询

```go
var student Student
DB.Take(&student, 2)
fmt.Println(student)

student = Student{} // 重新赋值
DB.Take(&student, "4")
fmt.Println(student)
```

Take的第二个参数，默认会根据主键查询，可以是字符串，可以是数字

##### 3.根据其他条件查询

```go
var student Student
DB.Take(&student, "name = ?", "机器人27号")
fmt.Println(student)
```

使用？作为占位符，将查询的内容放入?

```go
SELECT * FROM `students` WHERE name = '机器人27号' LIMIT 1
```

这样可以有效的防止sql注入，它的原理就是将参数全部转义，如

```go
DB.Take(&student, "name = ?", "机器人27号' or 1=1;#")
// SELECT * FROM `students` WHERE name = '机器人27号\' or 1=1;#' LIMIT 1
```

##### 4.根据struct查询

```go
var student Student
// 只能有一个主要值
student.ID = 2
//student.Name = "枫枫"
DB.Take(&student)
fmt.Println(student)
```

##### 5.获取查询结果

获取查询的记录数

```go
count := DB.Find(&studentList).RowsAffected
```

是否查询失败

```go
err := DB.Find(&studentList).Error
```

查询失败有查询为空，查询条件错误，sql语法错误

可以使用判断

```go
var student Student
err := DB.Take(&student, "xx").Error
switch err {
case gorm.ErrRecordNotFound:
  fmt.Println("没有找到")
default:
  fmt.Println("sql错误")
}
```

##### 6.查询多条记录

```go
var studentList []Student
DB.Find(&studentList)
for _, student := range studentList {
  fmt.Println(student)
}

// 由于email是指针类型，所以看不到实际的内容
// 但是序列化之后，会转换为我们可以看得懂的方式
var studentList []Student
DB.Find(&studentList)
for _, student := range studentList {

  data, _ := json.Marshal(student)
  fmt.Println(string(data))
}
```

##### 7.根据主键列表查询

```go
var studentList []Student
DB.Find(&studentList, []int{1, 3, 5, 7})
DB.Find(&studentList, 1, 3, 5, 7)  // 一样的
fmt.Println(studentList)
```

##### 8.根据其他条件查询

```go
DB.Find(&studentList, "name in ?", []string{"枫枫", "zhangsan"})
```

#### 3.更新

更新的前提的先查询到记录

##### 1.Save保存所有字段

用于单个记录的全字段更新

它会保存所有字段，即使零值也会保存

```go
var student Student
DB.Take(&student)
student.Age = 23
// 全字段更新
DB.Save(&student)
// UPDATE `students` SET `name`='枫枫',`age`=23,`gender`=true,`email`='xxx@qq.com' WHERE `id` = 1
```

零值也会更新

```go
var student Student
DB.Take(&student)
student.Age = 0
// 全字段更新
DB.Save(&student)
// UPDATE `students` SET `name`='枫枫',`age`=0,`gender`=true,`email`='xxx@qq.com' WHERE `id` = 1
```

##### 2.更新指定字段

可以使用select选择要更新的字段

```go
var student Student
DB.Take(&student)
student.Age = 21
// 全字段更新
DB.Select("age").Save(&student)
// UPDATE `students` SET `age`=21 WHERE `id` = 1
```

##### 3.批量更新

例如我想给年龄21的学生，都更新一下邮箱

```go
var studentList []Student
DB.Find(&studentList, "age = ?", 21).Update("email", "is21@qq.com")
```

还有一种更简单的方式

```go
DB.Model(&Student{}).Where("age = ?", 21).Update("email", "is21@qq.com")
// UPDATE `students` SET `email`='is22@qq.com' WHERE age = 21
```

这样的更新方式也是可以更新零值的

##### 4.更新多列

如果是结构体，它默认不会更新零值

```go
email := "xxx@qq.com"
DB.Model(&Student{}).Where("age = ?", 21).Updates(Student{
  Email:  &email,
  Gender: false,  // 这个不会更新
})

// UPDATE `students` SET `email`='xxx@qq.com' WHERE age = 21
```

如果想让他更新零值，用select就好

```go
email := "xxx1@qq.com"
DB.Model(&Student{}).Where("age = ?", 21).Select("gender", "email").Updates(Student{
  Email:  &email,
  Gender: false,
})
// UPDATE `students` SET `gender`=false,`email`='xxx1@qq.com' WHERE age = 21
```

如果不想多写几行代码，则推荐使用map（这样也可以更新零值）

```go
DB.Model(&Student{}).Where("age = ?", 21).Updates(map[string]any{
  "email":  &email,
  "gender": false,
})
```

更新选定字段

- Select选定字段
- Omit忽略字段

#### 4.删除

##### 1.根据结构体删除

```go
// student 的 ID 是 `10`
db.Delete(&student)
// DELETE from students where id = 10;
```

##### 2.删除多个

```go
db.Delete(&Student{}, []int{1,2,3})

// 查询到的切片列表
db.Delete(&studentList)
```

### 7.Hook

参考：https://gorm.io/zh_CN/docs/hooks.html

Hook 是在创建、查询、更新、删除等操作之前、之后调用的函数。

如果您已经为模型定义了指定的方法，它会在创建、更新、查询、删除时自动被调用。如果任何回调返回错误，GORM 将停止后续的操作并回滚事务。

钩子方法的函数签名应该是 `func(*gorm.DB) error`。

Gorm 提供了多个 Hook 点，例如 `BeforeCreate`、`AfterCreate`、`BeforeUpdate`、`AfterUpdate` 等，让你可以在模型创建、更新等事件发生时执行自定义代码。以下是如何在 Gorm 中创建 Hook 的示例：

```go
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}
// 也可以这样创建BeforeCreate的hook，会在创建对象前，自动调用BeforeCreate中的代码
// func (user *User) BeforeCreate(tx *gorm.DB) (err error){
//     fmt.Println("Before creating a new record")
//     return nil
// }

func main() {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 自动迁移模型
	db.AutoMigrate(&User{})

	// 创建 Hook，示例是在创建记录之前打印一条消息
	db.BeforeCreate(func(tx *gorm.DB) (err error) {
		fmt.Println("Before creating a new record")
		return nil
	})

	// 创建用户记录
	user := User{Name: "John Doe"}
	db.Create(&user)

	// 查询用户记录
	var retrievedUser User
	db.First(&retrievedUser, user.ID)
	fmt.Printf("Retrieved user: %v\n", retrievedUser)
}
```

上述示例中，我们创建了一个名为 `User` 的模型，并在模型的 `BeforeCreate` Hook 中打印了一条消息。在创建用户记录时，Hook 会在创建之前执行，然后继续执行正常的数据库操作。

你可以根据需要在其他 Hook 点执行自定义操作，例如 `BeforeUpdate`、`AfterUpdate`、`BeforeDelete`、`AfterDelete` 等，以满足你的应用程序需求。在 Hook 中，你可以执行任何 Go 代码，从而实现更复杂的逻辑。请注意，Hook 函数必须具有特定的签名，以匹配 Gorm 的要求。

### 8.高级查询

重新构造一些数据用于查询

```go
func main(){
  var studentList []Student
  DB.Find(&studentList).Delete(&studentList)
  studentList = []Student{
    {ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
    {ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
    {ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
    {ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
    {ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
    {ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
    {ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
    {ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
    {ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
  }
  DB.Create(&studentList)
}

func PtrString(email string) *string {
  return &email
}

```

#### 1.Where

等价于sql语句中的where

```go
var users []Student
// 查询用户名是枫枫的
DB.Where("name = ?", "枫枫").Find(&users)
fmt.Println(users)
// 查询用户名不是枫枫的
DB.Where("name <> ?", "枫枫").Find(&users)
fmt.Println(users)
// 查询用户名包含 如燕，李元芳的
DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
fmt.Println(users)
// 查询姓李的
DB.Where("name like ?", "李%").Find(&users)
fmt.Println(users)
// 查询年龄大于23，是qq邮箱的
DB.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
fmt.Println(users)
// 查询是qq邮箱的，或者是女的
DB.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
fmt.Println(users)
```

#### 2.使用结构体查询

使用结构体查询，会过滤零值，并且结构体中的条件都是`and`关系。

```go
// 会过滤零值
DB.Where(&Student{Name: "李元芳", Age: 0}).Find(&users)
fmt.Println(users)
```

#### 3.使用map查询

不会过滤零值

```go
DB.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)
// SELECT * FROM `students` WHERE `age` = 0 AND `name` = '李元芳'
fmt.Println(users)
```

#### 4.Not条件

和where中的not等价

```go
// 排除年龄大于23的
DB.Not("age > 23").Find(&users)
fmt.Println(users)
```

#### 5.Or条件

和where中的or等价

```go
DB.Or("gender = ?", false).Or(" email like ?", "%@qq.com").Find(&users)
fmt.Println(users)
```

#### 6.Select 选择字段

```go
DB.Select("name", "age").Find(&users)
fmt.Println(users)
// 没有被选中，会被赋零值

```

可以使用扫描Scan，将选择的字段存入另一个结构体中

```go
type User struct {
  Name string
  Age  int
}
var students []Student
var users []User
DB.Select("name", "age").Find(&students).Scan(&users)
fmt.Println(users)

```

这样写也是可以的，不过最终会查询两次，还是不这样写

```go
SELECT `name`,`age` FROM `students`
SELECT `name`,`age` FROM `students`
```

这样写就只查询一次了

```go
type User struct {
  Name string
  Age  int
}
var users []User
DB.Model(&Student{}).Select("name", "age").Scan(&users)
fmt.Println(users)
```

还可以这样

```go
var users []User
DB.Table("students").Select("name", "age").Scan(&users)
fmt.Println(users)
```

Scan是根据column列名进行扫描的

```go
type User struct {
  Name123 string `gorm:"column:name"`
  Age     int
}
var users []User
DB.Table("students").Select("name", "age").Scan(&users)
fmt.Println(users)
```

#### 7.排序

根据年龄倒序

```go
var users []Student
DB.Order("age desc").Find(&users)
fmt.Println(users)
// desc    降序
// asc     升序
```

注意order的顺序

#### 8.分页查询

```go
var users []Student
// 一页两条，第1页
DB.Limit(2).Offset(0).Find(&users)
fmt.Println(users)
// 第2页
DB.Limit(2).Offset(2).Find(&users)
fmt.Println(users)
// 第3页
DB.Limit(2).Offset(4).Find(&users)
fmt.Println(users)

```

通用写法

```go
var users []Student
// 一页多少条
limit := 2
// 第几页
page := 1
offset := (page - 1) * limit
DB.Limit(limit).Offset(offset).Find(&users)
fmt.Println(users)
```

#### 9.去重

```go
var ageList []int
DB.Table("students").Select("age").Distinct("age").Scan(&ageList)
fmt.Println(ageList)
```

或者

```go
DB.Table("students").Select("distinct age").Scan(&ageList)
```

#### 10.分组查询

```go
var ageList []int
// 查询男生的个数和女生的个数
DB.Table("students").Select("count(id)").Group("gender").Scan(&ageList)
fmt.Println(ageList)
```

有个问题，哪一个是男生个数，那个是女生个数

所以我们应该精确一点

```go
type AggeGroup struct {
  Gender int
  Count  int `gorm:"column:count(id)"`
}

var agge []AggeGroup
// 查询男生的个数和女生的个数
DB.Table("students").Select("count(id)", "gender").Group("gender").Scan(&agge)
fmt.Println(agge)
```

如何再精确一点，具体的男生名字，女生名字

```go
type AggeGroup struct {
  Gender int
  Count  int    `gorm:"column:count(id)"`
  Name   string `gorm:"column:group_concat(name)"`
}

var agge []AggeGroup
// 查询男生的个数和女生的个数
DB.Table("students").Select("count(id)", "gender", "group_concat(name)").Group("gender").Scan(&agge)
fmt.Println(agge)
```

#### 11.执行原生sql

```go
type AggeGroup struct {
  Gender int
  Count  int    `gorm:"column:count(id)"`
  Name   string `gorm:"column:group_concat(name)"`
}

var agge []AggeGroup
DB.Raw(`SELECT count(id), gender, group_concat(name) FROM students GROUP BY gender`).Scan(&agge)

fmt.Println(agge)
```

#### 12.子查询

查询大于平均年龄的用户

```go
# 原生sql
select * from students where age > (select avg(age) from students);
```

使用gorm编写

```go
var users []Student
DB.Model(Student{}).Where("age > (?)", DB.Model(Student{}).Select("avg(age)")).Find(&users)
fmt.Println(users)
```

#### 13.命名参数

```go
var users []Student

DB.Where("name = @name and age = @age", sql.Named("name", "枫枫"), sql.Named("age", 23)).Find(&users)
DB.Where("name = @name and age = @age", map[string]any{"name": "枫枫", "age": 23}).Find(&users)
fmt.Println(users)
```

#### 14.find到map

```go
var res []map[string]any
DB.Table("students").Find(&res)
fmt.Println(res)
```

#### 15.查询引用Scope

可以再model层写一些通用的查询方式，这样外界就可以直接调用方法即可

```go
func Age23(db *gorm.DB) *gorm.DB {
  return db.Where("age > ?", 23)
}

func main(){
  var users []Student
  DB.Scopes(Age23).Find(&users)
  fmt.Println(users)
}
```

### 9.表的关联

#### 1.一对多

一个用户可以发布多篇文章，一篇文章属于一个用户

```go
type User struct {
    ID       uint      `gorm:"size:4"`
    Name     string    `gorm:"size:8"`
    Articles []Article `gorm:"foreignKey:UID"` // 用户拥有的文章列表
}

type Article struct {
    ID    uint   `gorm:"size:4"`
    Title string `gorm:"size:120"`
    UID   uint   `gorm:"size:4"`         // 属于   这里的类型要和引用的外键类型一致，包括大小
    User  User   `gorm:"foreignKey:UID"` // 属于   这里重写了外键关联
}
```

关于外键命名，外键名称就是关联表名+ID，类型是uint。

也可以使用`gorm:"foreignKey`来重写外键，将外键指定为其他的字段。

##### 1.添加

###### 1.创建

创建用户，并且创建文章

```go
a1 := Article{Title: "python"}
a2 := Article{Title: "golang"}
user := User{Name: "枫枫", Articles: []Article{a1, a2}}
DB.Create(&user)
```

gorm自动创建了两篇文章，以及创建了一个用户，还将他们的关系给关联上了

创建文章，关联已有用户

```go
a1 := Article{Title: "golang零基础入门", UserID: 1}
DB.Create(&a1)
var user User
DB.Take(&user, 1)
DB.Create(&Article{Title: "python零基础入门", User: user})
```

###### 2.外键添加

给现有用户绑定文章

```go
var user User
DB.Take(&user, 2)

var article Article
DB.Take(&article, 5)

user.Articles = []Article{article}
DB.Save(&user)
```

也可以用Append方法

```go
var user User
DB.Take(&user, 2)

var article Article
DB.Take(&article, 5)

//user.Articles = []Article{article}
//DB.Save(&user)

DB.Model(&user).Association("Articles").Append(&article)
```

给现有文章关联用户

```go
var article Article
DB.Take(&article, 5)

article.UserID = 2
DB.Save(&article)
```

也可用Append方法

```go
var user User
DB.Take(&user, 2)

var article Article
DB.Take(&article, 5)

DB.Model(&article).Association("User").Append(&user)
```

##### 2.查询

查询用户，显示用户的文章列表

```go
var user User
DB.Take(&user, 1)
fmt.Println(user)
```

直接这样，是显示不出文章列表

###### 1.预加载

我们必须要使用预加载来加载文章列表

```go
var user User
DB.Preload("Articles").Take(&user, 1)
fmt.Println(user)
```

预加载的名字就是外键关联的属性名

查询文章，显示文章用户的信息

同样的，使用预加载

```go
var article Article
DB.Preload("User").Take(&article, 1)
fmt.Println(article)
```

###### 2.嵌套预加载

查询文章，显示用户，并且显示用户关联的所有文章，这就得用到嵌套预加载了

```go
var article Article
DB.Preload("User.Articles").Take(&article, 1)
fmt.Println(article)
```

###### 3.带条件的预加载

查询用户下的所有文章列表，过滤某些文章

```go
var user User
DB.Preload("Articles", "id = ?", 1).Take(&user, 1)
fmt.Println(user)
```

这样，就只有id为1的文章被预加载出来了

###### 4.自定义预加载

```go
var user User
DB.Preload("Articles", func(db *gorm.DB) *gorm.DB {
  return db.Where("id in ?", []int{1, 2})
}).Take(&user, 1)
fmt.Println(user)
```

##### 3.删除

###### 1.级联删除

删除用户，与用户关联的文章也会删除

```go
var user User
DB.Take(&user, 1)
DB.Select("Articles").Delete(&user)
```

###### 2.清除外键关系

删除用户，与将与用户关联的文章，外键设置为null

```go
var user User
DB.Preload("Articles").Take(&user, 2)
DB.Model(&user).Association("Articles").Delete(&user.Articles)
```

学习到第18集。（https://www.bilibili.com/video/BV1xg411t7RZ?p=18）

#### 2.一对一

2023.09.16

一对一关系比较少，一般用于表的扩展

例如一张用户表，有很多字段

那么就可以把它拆分为两张表，常用的字段放主表，不常用的字段放详情表

##### 1.表结构搭建

```go
type User struct {
  ID       uint
  Name     string
  Age      int
  Gender   bool
  UserInfo UserInfo // 通过UserInfo可以拿到用户详情信息
}

type UserInfo struct {
  UserID uint // 外键
  ID     uint
  Addr   string
  Like   string
}
```

##### 2.添加记录

添加用户，自动添加用户详情

```go
DB.Create(&User{
  Name:   "枫枫",
  Age:    21,
  Gender: true,
  UserInfo: UserInfo{
    Addr: "湖南省",
    Like: "写代码",
  },
})
```

添加用户详情，关联已有用户

这个场景特别适合网站的注册，以及后续信息完善

刚开始注册的时候，只需要填写很基本的信息，这就是添加主表的一条记录

注册进去之后，去个人中心，添加头像，修改地址...

这就是添加附表

```go
DB.Create(&UserInfo{
  UserID: 2,
  Addr:   "南京市",
  Like:   "吃饭",
})
```

当然，也可以直接把用户对象传递进来

我们需要改一下表结构

```go
type User struct {
  ID       uint
  Name     string
  Age      int
  Gender   bool
  UserInfo UserInfo // 通过UserInfo可以拿到用户详情信息
}

type UserInfo struct {
  User *User  // 要改成指针，不然就嵌套引用了
  UserID uint // 外键
  ID     uint
  Addr   string
  Like   string
}
```

不限于重新迁移，直接添加即可

```go
var user User
DB.Take(&user, 2)
DB.Create(&UserInfo{
  User: &user,
  Addr: "南京市",
  Like: "吃饭",
})
```

##### 3.查询

一般是通过主表查副表

```go
var user User
DB.Preload("UserInfo").Take(&user)
fmt.Println(user)
```

#### 3.多对多

多对多关系，需要用第三张表存储两张表的关系

##### 1.表结构搭建

```go
type Tag struct {
  ID       uint
  Name     string
  Articles []Article `gorm:"many2many:article_tags;"` // 用于反向引用
}

type Article struct {
  ID    uint
  Title string
  Tags  []Tag `gorm:"many2many:article_tags;"`
}
```

##### 2.多对多添加

添加文章，并创建标签

```go
DB.Create(&Article{
  Title: "python基础课程",
  Tags: []Tag{
    {Name: "python"},
    {Name: "基础课程"},
  },
})
```

添加文章，选择标签

```go
var tags []Tag
DB.Find(&tags, "name = ?", "基础课程")
DB.Create(&Article{
  Title: "golang基础",
  Tags:  tags,
})
```

##### 3.多对多查询

查询文章，显示文章的标签列表

```go
var article Article
DB.Preload("Tags").Take(&article, 1)
fmt.Println(article)
```

查询标签，显示文章列表

```go
var tag Tag
DB.Preload("Articles").Take(&tag, 2)
fmt.Println(tag)
```

##### 4.多对多更新

移除文章的标签

```go
var article Article
DB.Preload("Tags").Take(&article, 1)
DB.Model(&article).Association("Tags").Delete(article.Tags)
fmt.Println(article)
```

更新文章的标签

```go
var article Article
var tags []Tag
DB.Find(&tags, []int{2, 6, 7})

DB.Preload("Tags").Take(&article, 2)
DB.Model(&article).Association("Tags").Replace(tags)
fmt.Println(article)
```

##### 5.自定义连接表

默认的连接表，只有双方的主键id，展示不了更多信息了

这是官方的例子，我修改了一下

```go
type Article struct {
  ID    uint
  Title string
  Tags  []Tag `gorm:"many2many:article_tags"`
}

type Tag struct {
  ID   uint
  Name string
}

type ArticleTag struct {
  ArticleID uint `gorm:"primaryKey"`
  TagID     uint `gorm:"primaryKey"`
  CreatedAt time.Time
}

```

##### 6.生成表结构

```go
// 设置Article的Tags表为ArticleTag
DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
// 如果tag要反向应用Article，那么也得加上
// DB.SetupJoinTable(&Tag{}, "Articles", &ArticleTag{})
err := DB.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{})
fmt.Println(err)
```

##### 7.操作案例

举一些简单的例子

1. 添加文章并添加标签，并自动关联
2. 添加文章，关联已有标签
3. 给已有文章关联标签
4. 替换已有文章的标签
5. 添加文章并添加标签，并自动关联

```go
DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})  // 要设置这个，才能走到我们自定义的连接表
DB.Create(&Article{
  Title: "flask零基础入门",
  Tags: []Tag{
    {Name: "python"},
    {Name: "后端"}, 
    {Name: "web"},
  },
})
// CreatedAt time.Time 由于我们设置的是CreatedAt，gorm会自动填充当前时间，
// 如果是其他的字段，需要使用到ArticleTag 的添加钩子 BeforeCreate
```

**1.添加文章，关联已有标签**

```go
DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
var tags []Tag
DB.Find(&tags, "name in ?", []string{"python", "web"})
DB.Create(&Article{
  Title: "flask请求对象",
  Tags:  tags,
})
```

**2.给已有文章关联标签**

```go
DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
article := Article{
  Title: "django基础",
}
DB.Create(&article)
var at Article
var tags []Tag
DB.Find(&tags, "name in ?", []string{"python", "web"})
DB.Take(&at, article.ID).Association("Tags").Append(tags)
```

**3.替换已有文章的标签**

```go
var article Article
var tags []Tag
DB.Find(&tags, "name in ?", []string{"后端"})
DB.Take(&article, "title = ?", "django基础")
DB.Model(&article).Association("Tags").Replace(tags)
```

**4.查询文章列表，显示标签**

```go
var articles []Article
DB.Preload("Tags").Find(&articles)
fmt.Println(articles)
```

**5.SetupJoinTable**

添加和更新的时候得用这个

这样才能走自定义的连接表，以及走它的钩子函数

查询则不需要这个

**6.自定义连接表主键**

这个功能还是很有用的，例如你的文章表 可能叫ArticleModel，你的标签表可能叫TagModel

那么按照gorm默认的主键名，那就分别是ArticleModelID，TagModelID，太长了，根本就不实用

这个地方，官网给的例子看着也比较迷，不过我已经跑通了

主要是要修改这两项

joinForeignKey 连接的主键id

JoinReferences 关联的主键id

```go
type ArticleModel struct {
  ID    uint
  Title string
  Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
}

type TagModel struct {
  ID       uint
  Name     string
  Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
}

type ArticleTagModel struct {
  ArticleID uint `gorm:"primaryKey"` // article_id
  TagID     uint `gorm:"primaryKey"` // tag_id
  CreatedAt time.Time
}
```

**7.生成表结构**

```go
DB.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
DB.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
err := DB.AutoMigrate(&ArticleModel{}, &TagModel{}, &ArticleTagModel{})
fmt.Println(err)
```

添加，更新，查询操作和上面的都是一样

**8.操作连接表**

如果通过一张表去操作连接表，这样会比较麻烦

比如查询某篇文章关联了哪些标签

或者是举个更通用的例子，用户和文章，某个用户在什么时候收藏了哪篇文章

无论是通过用户关联文章，还是文章关联用户都不太好查

最简单的就是直接查连接表

```go
type UserModel struct {
  ID       uint
  Name     string
  Collects []ArticleModel `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}

type ArticleModel struct {
  ID    uint
  Title string
  // 这里也可以反向引用，根据文章查哪些用户收藏了
}

// UserCollectModel 用户收藏文章表
type UserCollectModel struct {
  UserID    uint `gorm:"primaryKey"` // article_id
  ArticleID uint `gorm:"primaryKey"` // tag_id
  CreatedAt time.Time
}

func main() {
  DB.SetupJoinTable(&UserModel{}, "Collects", &UserCollectModel{})
  err := DB.AutoMigrate(&UserModel{}, &ArticleModel{}, &UserCollectModel{})
  fmt.Println(err)
}
```

常用的操作就是根据用户查收藏的文章列表

```go
var user UserModel
DB.Preload("Collects").Take(&user, "name = ?", "枫枫")
fmt.Println(user)
```

但是这样不太好做分页，并且也拿不到收藏文章的时间

```go
var collects []UserCollectModel
DB.Find(&collects, "user_id = ?", 2)
fmt.Println(collects)
```

这样虽然可以查到用户id，文章id，收藏的时间，但是搜索只能根据用户id搜，返回也拿不到用户名，文章标题等

我们需要改一下表结构，不需要重新迁移，加一些字段

```go
type UserModel struct {
  ID       uint
  Name     string
  Collects []ArticleModel `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}

type ArticleModel struct {
  ID    uint
  Title string
}

// UserCollectModel 用户收藏文章表
type UserCollectModel struct {
  UserID       uint         `gorm:"primaryKey"` // article_id
  UserModel    UserModel    `gorm:"foreignKey:UserID"`
  ArticleID    uint         `gorm:"primaryKey"` // tag_id
  ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
  CreatedAt    time.Time
}
```

**9.查询**

```go
var collects []UserCollectModel

var user UserModel
DB.Take(&user, "name = ?", "枫枫")
// 这里用map的原因是如果没查到，那就会查0值，如果是struct，则会忽略零值，全部查询
DB.Debug().Preload("UserModel").Preload("ArticleModel").Where(map[string]any{"user_id": user.ID}).Find(&collects)

for _, collect := range collects {
  fmt.Println(collect)
}
```

### 10.自定义数据类型

在Go语言中使用Gorm框架自定义数据类型时，你可以创建一个自定义类型并实现Gorm的`Scanner`和`Valuer`接口，以便将其映射到数据库字段，包括JSON类型字段。下面是一个示例，演示如何自定义数据类型并将其映射到JSON字段。

#### 1.Json

假设你要创建一个自定义数据类型`MyJSON`，它表示一个可以存储为JSON的数据结构，并将其映射到数据库中的JSON字段。

1. 创建自定义数据类型：

```go
package models

import (
    "database/sql/driver"
    "encoding/json"
    "errors"
)

// MyJSON 自定义JSON类型
type MyJSON map[string]interface{}

// 实现Valuer接口，将MyJSON类型转换为数据库字段值
func (mj MyJSON) Value() (driver.Value, error) {
    return json.Marshal(mj)
}

// 实现Scanner接口，从数据库字段值扫描到MyJSON类型
func (mj *MyJSON) Scan(value interface{}) error {
    if value == nil {
        *mj = nil
        return nil
    }

    // 将数据库中的JSON字符串解析为MyJSON类型
    var jsonData []byte
    switch val := value.(type) {
    case []byte:
        jsonData = val
    case string:
        jsonData = []byte(val)
    default:
        return errors.New("Failed to scan MyJSON from database")
    }

    return json.Unmarshal(jsonData, mj)
}
```

2. 在模型中使用自定义数据类型：

```go
package models

import "github.com/jinzhu/gorm"

type Item struct {
    ID   uint
    Data MyJSON // 使用自定义JSON数据类型
}
```

3. 在数据库迁移中使用自定义数据类型：

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "your_project/models"
)

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("Failed to connect database")
    }
    defer db.Close()

    // 自动迁移数据库结构
    db.AutoMigrate(&models.Item{})
}
```

现在，你已经成功地创建了一个自定义数据类型`MyJSON`，并在Gorm模型中使用它。这使你能够将包含JSON数据的结构存储为数据库中的JSON字段，并在需要时自动转换为`MyJSON`类型。你可以在模型中使用`MyJSON`字段来存储和检索JSON数据。

#### 2.Enum

```go
type Status int

const (
    Running Status = iota // 使用iota来递增常量
    Except
    OffLine
)

type Host struct {
    IP     string `json:"ip"`
    Status Status `json:"status"`
}

// MarshalJSON 实现MarshalJSON()方法，会在json.Marshal()时被自动调用
func (status Status) MarshalJSON() ([]byte, error) { // 作Status数字和字符串之间的映射
    var str string
    switch status {
    case Running:
       str = "Running"
    case Except:
       str = "Except"
    case OffLine:
       str = "OffLine"
    }
    return json.Marshal(str)
}

func main() {
    DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
    _ = DB.AutoMigrate(&Host{}) // 自动迁移数据库结构

    // 添加数据
    DB.Create(&Host{
       IP:     "192.168.12.56",
       Status: Running,
    })
    DB.Create(&Host{
       IP:     "192.168.12.88",
       Status: OffLine,
    })

    // 查询数据
    var host Host
    DB.Take(&host, "ip = ?", "192.168.12.56")
    fmt.Println(host) // {192.168.12.56 1}
    // 这样打印出来的值是数据库中实际存放的值，如果我们想要给前端返回Status值而不是数字，就需要实现MarshalJSON()方法。
    // 在通过json.Marshal()对host进行序列化时，会自动调用MarshalJSON()函数，将数字转换为其对应的字符串。

    // 序列化为JSON格式的字符串
    data, _ := json.Marshal(host) // 自动调用MarshalJSON()函数
    fmt.Println(string(data))     // {"ip":"192.168.12.56","status":"Running"}
}
```

MarshalJSON 方法：

> 在Gorm中，有时你可能需要自定义如何将数据库模型（model）的字段值转换为JSON格式。为了实现这个目标，你可以在Gorm的模型上定义一个 `MarshalJSON` 方法，该方法会在将模型序列化为JSON时被自动调用。

### 11.事务

事务就是用户定义的一系列数据库操作，这些操作可以视为一个完成的逻辑处理工作单元，要么全部执行，要么全部不执行，是不可分割的工作单元。

很形象的一个例子，张三给李四转账100元，在程序里面，张三的余额就要-100，李四的余额就要+100 整个事件是一个整体，哪一步错了，整个事件都是失败的

gorm事务默认是开启的。为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。

如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升。

一般不推荐禁用

```go
// 全局禁用
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})

```

本节课表结构

```go
type User struct {
  ID    uint   `json:"id"`
  Name  string `json:"name"`
  Money int    `json:"money"`
}

// InnoDB引擎才支持事务，MyISAM不支持事务
// DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

```

##### 1.普通事务

以张三给李四转账为例，不使用事务的后果

```go
var zhangsan, lisi User
DB.Take(&zhangsan, "name = ?", "张三")
DB.Take(&lisi, "name = ?", "李四")
// 张三给李四转账100元
// 先给张三-100
zhangsan.Money -= 100
DB.Model(&zhangsan).Update("money", zhangsan.Money)
// 模拟失败的情况

// 再给李四+100
lisi.Money += 100
DB.Model(&lisi).Update("money", lisi.Money)
```

在失败的情况下，要么张三白白损失了100，要么李四凭空拿到100元

这显然是不合逻辑的，并且不合法的

那么，使用事务是怎样的

```go
var zhangsan, lisi User
DB.Take(&zhangsan, "name = ?", "张三")
DB.Take(&lisi, "name = ?", "李四")
// 张三给李四转账100元
DB.Transaction(func(tx *gorm.DB) error {

  // 先给张三-100
  zhangsan.Money -= 100
  err := tx.Model(&zhangsan).Update("money", zhangsan.Money).Error
  if err != nil {
    fmt.Println(err)
    return err
  }

  // 再给李四+100
  lisi.Money += 100
  err = tx.Model(&lisi).Update("money", lisi.Money).Error
  if err != nil {
    fmt.Println(err)
    return err
  }
  // 提交事务
  return nil
})
```

使用事务之后，他们就是一体，一起成功，一起失败

**db.Transaction 方法**：

> 签名：`func (db *DB) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error`
>
> `fc func(tx *gorm.DB) error`: 这是一个接受 `*gorm.DB` 参数并返回 `error` 的函数，用于定义你要在事务中执行的操作。在这个函数内部，你可以执行数据库的查询、插入、更新和删除操作，如果发生错误，应该返回一个非空的错误值，以便事务能够回滚。
>
> `opts ...*sql.TxOptions`: 这是一个可选参数，用于指定事务的选项，例如事务的隔离级别、只读事务等。`sql.TxOptions` 是一个结构体，你可以在其中设置事务的属性。默认情况下，如果不提供这个参数，Gorm会使用数据库的默认选项。
>
> 如果在事务中的任何操作中发生错误，事务将自动回滚，并且错误将被传递给调用方。在这个方法内部，如果发生错误，应该返回一个非空的错误值，以便事务能够回滚。返回`nil`表示事务操作成功。

##### 2.手动事务

```go
// 开始事务
tx := db.Begin()

// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
tx.Create(...)

// ...

// 遇到错误时回滚事务
tx.Rollback()

// 否则，提交事务
tx.Commit()
```

刚才的代码也可以这样实现

```go
var zhangsan, lisi User
DB.Take(&zhangsan, "name = ?", "张三")
DB.Take(&lisi, "name = ?", "李四")

// 张三给李四转账100元
tx := DB.Begin()

// 先给张三-100
zhangsan.Money -= 100
err := tx.Model(&zhangsan).Update("money", zhangsan.Money).Error
if err != nil {
  tx.Rollback()
}

// 再给李四+100
lisi.Money += 100
err = tx.Model(&lisi).Update("money", lisi.Money).Error
if err != nil {
  tx.Rollback()
}
// 提交事务
tx.Commit()
```

已学完：[枫枫知道-golang最简单的gorm教程](https://www.bilibili.com/video/BV1xg411t7RZ)

已学完：[【最新Go Web开发教程】基于gin框架和gorm的web开发实战 (七米出品)](https://www.bilibili.com/video/BV1gJ411p7xC)

目前学习内容：[七米-Go Web开发进阶实战（gin框架）（共23小时）](https://study.163.com/course/introduction.htm?courseId=1210171207)

## Gin框架源码

可以参考的教程：[小徐先生1212-gin框架底层技术原理剖析](https://www.bilibili.com/video/BV1zm4y177mb)

### 前缀树

```go
// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    c := engine.pool.Get().(*Context)
    c.writermem.reset(w)  // 取出来之后再做初始化
    c.Request = req
    c.reset()

    engine.handleHTTPRequest(c)  // 处理http请求

    engine.pool.Put(c)
}
```

通过对象池，减少垃圾回收的次数和内存申请的消耗。

```go
// 这样可以在使用len(t)的时候，只计算一次len(t)，减少计算的次数。
for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {  // 把应该立即返回的条件写在前面
			continue  
		}
```

```go
type Engine struct {
	...
}

var _ IRouter = (*Engine)(nil)
```

这样写是为了确保Engine结构体实现了IRouter接口，把问题暴露在编译阶段（如果在写某个结构体的时候，需要让改结构体实现某个接口，就可以采用这样的写法）

```go
type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
    for _, tree := range trees {
       if tree.method == method {
          return tree.root
       }
    }
    return nil
}
```

每种请求方法都对应了一棵单独的树（共9种请求方法），methodTrees切片中存放了所有的请求方法，通过遍历methodTrees，就可以得到该方法对应的树结构。

```go
func Default() *Engine {
    ...
	engine := New()
	...
}

func New() *Engine {
    ...
	trees:make(methodTrees, 0, 9),
    ...
}
```

在初始化结构体时，一次性把容量申请到位，防止频繁地去扩容。

```go
type node struct {
   // 节点路径，比如上面的s，earch，和upport
	path      string
	// 和children字段对应, 保存的是分裂的分支的第一个字符
	// 例如search和support, 那么s节点的indices对应的"eu"
	// 代表有两个分支, 分支的首字母分别是e和u
	indices   string
	// 儿子节点
	children  []*node
	// 处理函数链条（切片）
	handlers  HandlersChain
	// 优先级，子节点、子子节点等注册的handler数量
	priority  uint32
	// 节点类型，包括static, root, param, catchAll
	// static: 静态节点（默认），比如上面的s，earch等节点
	// root: 树的根节点
	// catchAll: 有*匹配的节点
	// param: 参数节点
	nType     nodeType
	// 路径上最大参数个数
	maxParams uint8
	// 节点是否是参数节点，比如上面的:post
	wildChild bool
	// 完整路径
	fullPath  string
}
```

这个node就是构造路由前缀树的结构体。

### 路由注册

```go
func Default() *Engine {
    debugPrintWARNINGDefault()
    engine := New()
    engine.Use(Logger(), Recovery())
    return engine
}
```

```go
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodGet, relativePath, handlers)
}
```

在使用gin.Default()获取一个路由实例的时候，返回的是一个Engine类型的变量，但是在route.Get()时，Get()方法绑定的类型却是RouterGroup，为什么可以这样呢？

```go
type Engine struct {
    RouterGroup
	...
}
```

因为Engine结构体中嵌套了RouterGroup。RouterGroup实现了Get()方法。用Engine类型的变量来调用Get()方法，是多态的一种体现。

## Zap

2023.09.28

### 介绍

Zap是一个高性能的、结构化的日志库，专门为Go语言设计和优化。它是由Uber开发并维护的，旨在提供高性能的日志记录，同时保持简单和易用。以下是关于Zap库的详细介绍：

1. **高性能**：Zap被设计为一个极其高性能的日志库。它采用了零分配（zero-allocation）的设计，意味着在记录日志时不会分配不必要的内存，从而减少了垃圾回收的压力。这对于高并发和低延迟的应用程序特别重要。

2. **结构化日志**：Zap鼓励使用结构化日志记录方式，而不是传统的文本日志。结构化日志可以更轻松地进行查询和分析，因为每个日志消息都包含了一组字段和值。这有助于构建更可读和可搜索的日志。

3. **多日志级别**：Zap支持多种日志级别，包括Debug、Info、Warn、Error和DPanic（处理恶劣情况但不终止应用程序）。你可以根据需要选择适当的级别来记录不同严重程度的信息。

4. **自定义日志输出**：Zap允许你将日志输出到不同的目标，包括控制台、文件、网络和其他自定义目标。你可以根据应用程序的需求轻松定制日志输出。

5. **性能测量**：Zap内置了性能测量工具，可以用于测量日志记录操作的性能，以便进行优化。

6. **丰富的上下文**：Zap允许你添加上下文信息，例如请求ID、用户ID等，以便更容易跟踪问题和分析日志。

7. **延迟初始化**：Zap支持延迟初始化，这意味着你可以在应用程序启动时设置全局的日志配置，然后在需要时延迟初始化具体的日志记录器，以提高性能。

8. **社区支持**：Zap是一个活跃的开源项目，拥有大量的社区支持和贡献，因此你可以期望在文档、示例和问题解决方案方面找到丰富的资源。

下面是一个简单的示例，展示了如何在Go中使用Zap库：

```go
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 配置日志记录器
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	// 记录日志消息
	logger.Info("This is an info message",
		zap.String("key", "value"),
		zap.Int("count", 42),
	)

	// 关闭日志记录器
	defer logger.Sync()
}
```

在实际应用程序中，你可以根据需要自定义日志记录器的配置，包括输出目标、日志级别和其他设置。Zap库的灵活性和性能使其成为许多Go项目的首选日志库之一。你可以查看Zap的官方文档以获取更多信息和示例：https://pkg.go.dev/go.uber.org/zap

### 默认的Go Logger

在介绍Uber-go的zap包之前，让我们先看看Go语言提供的基本日志功能。Go语言提供的默认日志包是https://golang.org/pkg/log/。

**实现Go Logger**

实现一个Go语言中的日志记录器非常简单——创建一个新的日志文件，然后设置它为日志的输出位置。

**设置Logger**

我们可以像下面的代码一样设置日志记录器

```go
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/Users/q1mi/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}
```

**使用Logger**

让我们来写一些虚拟的代码来使用这个日志记录器。

在当前的示例中，我们将建立一个到URL的HTTP连接，并将状态代码/错误记录到日志文件中。

```go
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}
```

**Logger的运行**

现在让我们执行上面的代码并查看日志记录器的运行情况。

```go
func main() {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
```

当我们执行上面的代码，我们能看到一个`test.log`文件被创建，下面的内容会被添加到这个日志文件中。

```bash
2019/05/24 01:14:13 Error fetching url www.google.com : Get www.google.com: unsupported protocol scheme ""
2019/05/24 01:14:14 Status Code for http://www.google.com : 200 OK
```

**Go Logger的优势和劣势**

**优势**

它最大的优点是使用非常简单。我们可以设置任何`io.Writer`作为日志记录输出并向其发送要写入的日志。

**劣势**

- 仅限基本的日志级别

  - 只有一个`Print`选项。不支持`INFO`/`DEBUG`等多个级别。

- 对于错误日志，它有

  ```
  Fatal
  ```

  和

  ```
  Panic
  ```

  - Fatal日志通过调用`os.Exit(1)`来结束程序
  - Panic日志在写入日志消息之后抛出一个panic
  - 但是它缺少一个ERROR日志级别，这个级别可以在不抛出panic或退出程序的情况下记录错误

- 缺乏日志格式化的能力——例如记录调用者的函数名和行号，格式化日期和时间格式，等等。

- 不提供日志切割的能力。

### Uber-go Zap

[Zap](https://github.com/uber-go/zap)是非常快的、结构化的，分日志级别的Go日志库。

**为什么选择Uber-go zap**

- 它同时提供了结构化日志记录和printf风格的日志记录
- 它非常的快

根据Uber-go Zap的文档，它的性能比类似的结构化日志包更好——也比标准库更快。 以下是Zap发布的基准测试信息

记录一条消息和10个字段:

|     Package     |    Time     | Time % to zap | Objects Allocated |
| :-------------: | :---------: | :-----------: | :---------------: |
|      ⚡️ zap      |  862 ns/op  |      +0%      |    5 allocs/op    |
| ⚡️ zap (sugared) | 1250 ns/op  |     +45%      |   11 allocs/op    |
|     zerolog     | 4021 ns/op  |     +366%     |   76 allocs/op    |
|     go-kit      | 4542 ns/op  |     +427%     |   105 allocs/op   |
|    apex/log     | 26785 ns/op |    +3007%     |   115 allocs/op   |
|     logrus      | 29501 ns/op |    +3322%     |   125 allocs/op   |
|      log15      | 29906 ns/op |    +3369%     |   122 allocs/op   |

记录一个静态字符串，没有任何上下文或printf风格的模板：

|     Package      |    Time    | Time % to zap | Objects Allocated |
| :--------------: | :--------: | :-----------: | :---------------: |
|      ⚡️ zap       | 118 ns/op  |      +0%      |    0 allocs/op    |
| ⚡️ zap (sugared)  | 191 ns/op  |     +62%      |    2 allocs/op    |
|     zerolog      |  93 ns/op  |     -21%      |    0 allocs/op    |
|      go-kit      | 280 ns/op  |     +137%     |   11 allocs/op    |
| standard library | 499 ns/op  |     +323%     |    2 allocs/op    |
|     apex/log     | 1990 ns/op |    +1586%     |   10 allocs/op    |
|      logrus      | 3129 ns/op |    +2552%     |   24 allocs/op    |
|      log15       | 3887 ns/op |    +3194%     |   23 allocs/op    |

### 安装

运行下面的命令安装zap

```bash
go get -u go.uber.org/zap
```

### 配置Zap Logger

Zap提供了两种类型的日志记录器—`Sugared Logger`和`Logger`。

在性能很好但不是很关键的上下文中，使用`SugaredLogger`。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。

在每一微秒和每一次内存分配都很重要的上下文中，使用`Logger`。它甚至比`SugaredLogger`更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。

#### Logger

- 通过调用`zap.NewProduction()`/`zap.NewDevelopment()`或者`zap.Example()`创建一个Logger。
- 上面的每一个函数都将创建一个logger。唯一的区别在于它将记录的信息不同。例如production logger默认记录调用函数信息、日期和时间等。
- 通过Logger调用Info/Error等。
- 默认情况下日志都会打印到应用程序的console界面。

```go
var logger *zap.Logger

func main() {
	InitLogger()
  defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
```

在上面的代码中，我们首先创建了一个Logger，然后使用Info/ Error等Logger方法记录消息。

日志记录器方法的语法是这样的：

```go
func (log *Logger) MethodXXX(msg string, fields ...Field) 
```

其中`MethodXXX`是一个可变参数函数，可以是Info / Error/ Debug / Panic等。每个方法都接受一个消息字符串和任意数量的`zapcore.Field`场参数。

每个`zapcore.Field`其实就是一组键值对参数。

我们执行上面的代码会得到如下输出结果：

```bash
{"level":"error","ts":1572159218.912792,"caller":"zap_demo/temp.go:25","msg":"Error fetching url..","url":"www.sogo.com","error":"Get www.sogo.com: unsupported protocol scheme \"\"","stacktrace":"main.simpleHttpGet\n\t/Users/q1mi/zap_demo/temp.go:25\nmain.main\n\t/Users/q1mi/zap_demo/temp.go:14\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
{"level":"info","ts":1572159219.1227388,"caller":"zap_demo/temp.go:30","msg":"Success..","statusCode":"200 OK","url":"http://www.sogo.com"}
```

#### Sugared Logger

现在让我们使用Sugared Logger来实现相同的功能。

- 大部分的实现基本都相同。
- 惟一的区别是，我们通过调用主logger的`. Sugar()`方法来获取一个`SugaredLogger`。
- 然后使用`SugaredLogger`以`printf`格式记录语句

下面是修改过后使用`SugaredLogger`代替`Logger`的代码：

```go
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
  logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
```

当你执行上面的代码会得到如下输出：

```bash
{"level":"error","ts":1572159149.923002,"caller":"logic/temp2.go:27","msg":"Error fetching URL www.sogo.com : Error = Get www.sogo.com: unsupported protocol scheme \"\"","stacktrace":"main.simpleHttpGet\n\t/Users/q1mi/zap_demo/logic/temp2.go:27\nmain.main\n\t/Users/q1mi/zap_demo/logic/temp2.go:14\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
{"level":"info","ts":1572159150.192585,"caller":"logic/temp2.go:29","msg":"Success! statusCode = 200 OK for URL http://www.sogo.com"}
```

你应该注意到的了，到目前为止这两个logger都打印输出JSON结构格式。

在本博客的后面部分，我们将更详细地讨论SugaredLogger，并了解如何进一步配置它。

### 定制logger

#### 将日志写入文件而不是终端

我们要做的第一个更改是把日志写入文件，而不是打印到应用程序控制台。

- 我们将使用`zap.New(…)`方法来手动传递所有配置，而不是使用像`zap.NewProduction()`这样的预置方法来创建logger。

```go
func New(core zapcore.Core, options ...Option) *Logger
```

`zapcore.Core`需要三个配置——`Encoder`，`WriteSyncer`，`LogLevel`。

1.**Encoder**:编码器(如何写入日志)。我们将使用开箱即用的`NewJSONEncoder()`，并使用预先设置的`ProductionEncoderConfig()`。

```go
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
```

2.**WriterSyncer** ：指定日志将写到哪里去。我们使用`zapcore.AddSync()`函数并且将打开的文件句柄传进去。

```go
file, _ := os.Create("./test.log")
writeSyncer := zapcore.AddSync(file)
```

3.**Log Level**：哪种级别的日志将被写入。

我们将修改上述部分中的Logger代码，并重写`InitLogger()`方法。其余方法—`main()` /`SimpleHttpGet()`保持不变。

```go
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
```

当使用这些修改过的logger配置调用上述部分的`main()`函数时，以下输出将打印在文件——`test.log`中。

```bash
{"level":"debug","ts":1572160754.994731,"msg":"Trying to hit GET request for www.sogo.com"}
{"level":"error","ts":1572160754.994982,"msg":"Error fetching URL www.sogo.com : Error = Get www.sogo.com: unsupported protocol scheme \"\""}
{"level":"debug","ts":1572160754.994996,"msg":"Trying to hit GET request for http://www.sogo.com"}
{"level":"info","ts":1572160757.3755069,"msg":"Success! statusCode = 200 OK for URL http://www.sogo.com"}
```

#### 将JSON Encoder更改为普通的Log Encoder

现在，我们希望将编码器从JSON Encoder更改为普通Encoder。为此，我们需要将`NewJSONEncoder()`更改为`NewConsoleEncoder()`。

```go
return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
```

当使用这些修改过的logger配置调用上述部分的`main()`函数时，以下输出将打印在文件——`test.log`中。

```bash
1.572161051846623e+09	debug	Trying to hit GET request for www.sogo.com
1.572161051846828e+09	error	Error fetching URL www.sogo.com : Error = Get www.sogo.com: unsupported protocol scheme ""
1.5721610518468401e+09	debug	Trying to hit GET request for http://www.sogo.com
1.572161052068744e+09	info	Success! statusCode = 200 OK for URL http://www.sogo.com
```

#### 更改时间编码并添加调用者详细信息

鉴于我们对配置所做的更改，有下面两个问题：

- 时间是以非人类可读的方式展示，例如1.572161051846623e+09
- 调用方函数的详细信息没有显示在日志中

我们要做的第一件事是覆盖默认的`ProductionConfig()`，并进行以下更改:

- 修改时间编码器
- 在日志文件中使用大写字母记录日志级别

```go
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
```

接下来，我们将修改zap logger代码，添加将调用函数信息记录到日志中的功能。为此，我们将在`zap.New(..)`函数中添加一个`Option`。

```go
logger := zap.New(core, zap.AddCaller())
```

当使用这些修改过的logger配置调用上述部分的`main()`函数时，以下输出将打印在文件——`test.log`中。

```bash
2019-10-27T15:33:29.855+0800	DEBUG	logic/temp2.go:47	Trying to hit GET request for www.sogo.com
2019-10-27T15:33:29.855+0800	ERROR	logic/temp2.go:50	Error fetching URL www.sogo.com : Error = Get www.sogo.com: unsupported protocol scheme ""
2019-10-27T15:33:29.856+0800	DEBUG	logic/temp2.go:47	Trying to hit GET request for http://www.sogo.com
2019-10-27T15:33:30.125+0800	INFO	logic/temp2.go:52	Success! statusCode = 200 OK for URL http://www.sogo.com
```

#### AddCallerSkip

当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，想要获得准确的调用信息就需要通过`AddCallerSkip`函数来跳过。

```go
logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
```

#### 将日志输出到多个位置

我们可以将日志同时输出到文件和终端。

```go
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	// 利用io.MultiWriter支持文件和终端两个输出目标
	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}
```

#### 将err日志单独输出到文件

有时候我们除了将全量日志输出到`xx.log`文件中之外，还希望将`ERROR`级别的日志单独输出到一个名为`xx.err.log`的日志文件中。我们可以通过以下方式实现。

```go
func InitLogger() {
	encoder := getEncoder()
	// test.log记录全量日志
	logF, _ := os.Create("./test.log")
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
	// test.err.log记录ERROR级别的日志
	errF, _ := os.Create("./test.err.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)
	logger = zap.New(core, zap.AddCaller())
}
```

### 使用Lumberjack进行日志切割归档

这个日志程序中唯一缺少的就是日志切割归档功能。

> *Zap本身不支持切割归档日志文件*

官方的说法是为了添加日志切割归档功能，我们将使用第三方库[Lumberjack](https://github.com/natefinch/lumberjack)来实现。

目前只支持按文件大小切割，原因是按时间切割效率低且不能保证日志数据不被破坏。详情戳https://github.com/natefinch/lumberjack/issues/54。

想按日期切割可以使用[github.com/lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)这个库，虽然目前不维护了，但也够用了。

```go
// 使用file-rotatelogs按天切割日志

import rotatelogs "github.com/lestrrat-go/file-rotatelogs"

l, _ := rotatelogs.New(
	filename+".%Y%m%d%H%M",
	rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
	rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
)
zapcore.AddSync(l)
```

#### 安装

执行下面的命令安装 Lumberjack v2 版本。

```bash
go get gopkg.in/natefinch/lumberjack.v2
```

#### zap logger中加入Lumberjack

要在zap中加入Lumberjack支持，我们需要修改`WriteSyncer`代码。我们将按照下面的代码修改`getLogWriter()`函数：

```go
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
```

Lumberjack Logger采用以下属性作为输入:

- Filename: 日志文件的位置
- MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
- MaxBackups：保留旧文件的最大个数
- MaxAges：保留旧文件的最大天数
- Compress：是否压缩/归档旧文件

### 测试所有功能

最终，使用Zap/Lumberjack logger的完整示例代码如下：

```go
package main

import (
	"net/http"

	"gopkg.in/natefinch/lumberjack.v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
```

执行上述代码，下面的内容会输出到文件——test.log中。

```go
2019-10-27T15:50:32.944+0800	DEBUG	logic/temp2.go:48	Trying to hit GET request for www.sogo.com
2019-10-27T15:50:32.944+0800	ERROR	logic/temp2.go:51	Error fetching URL www.sogo.com : Error = Get www.sogo.com: unsupported protocol scheme ""
2019-10-27T15:50:32.944+0800	DEBUG	logic/temp2.go:48	Trying to hit GET request for http://www.sogo.com
2019-10-27T15:50:33.165+0800	INFO	logic/temp2.go:53	Success! statusCode = 200 OK for URL http://www.sogo.com
```

同时，可以在`main`函数中循环记录日志，测试日志文件是否会自动切割和归档（日志文件每1MB会切割并且在当前目录下最多保存5个备份）。

至此，我们总结了如何将Zap日志程序集成到Go应用程序项目中。

## Gin框架中使用Zap

使用zap接收gin框架默认的日志并配置日志归档

我们在基于gin框架开发项目时通常都会选择使用专业的日志库来记录项目中的日志，go语言常用的日志库有`zap`、`logrus`等。网上也有很多类似的教程，我之前也翻译过一篇[《在Go语言项目中使用Zap日志库》](https://www.liwenzhou.com/posts/Go/zap/)。

但是我们该如何在日志中记录gin框架本身输出的那些日志呢？

### gin默认的中间件

首先我们来看一个最简单的gin项目：

```go
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String("hello liwenzhou.com!")
	})
	r.Run(
}
```

接下来我们看一下`gin.Default()`的源码：

```go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```

也就是我们在使用`gin.Default()`的同时是用到了gin框架内的两个默认中间件`Logger()`和`Recovery()`。

其中`Logger()`是把gin框架本身的日志输出到标准输出（我们本地开发调试时在终端输出的那些日志就是它的功劳），而`Recovery()`是在程序出现panic的时候恢复现场并写入500响应的。

### 基于zap的中间件

我们可以模仿`Logger()`和`Recovery()`的实现，使用我们的日志库来接收gin框架默认输出的日志。

这里以zap为例，我们实现两个中间件如下：

```go
// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
```

*如果不想自己实现，可以使用github上有别人封装好的https://github.com/gin-contrib/zap。*

这样我们就可以在gin框架中使用我们上面定义好的两个中间件来代替gin框架默认的`Logger()`和`Recovery()`了。

```go
r := gin.New()
r.Use(GinLogger(), GinRecovery())
```

### 在gin项目中使用zap

最后我们再加入我们项目中常用的日志切割，完整版的`logger.go`代码如下：

```go
package logger

import (
	"gin_zap_demo/config"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

// InitLogger 初始化Logger
func InitLogger(cfg *config.LogConfig) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
```

然后定义日志相关配置：

```go
type LogConfig struct {
	Level string `json:"level"`
	Filename string `json:"filename"`
	MaxSize int `json:"maxsize"`
	MaxAge int `json:"max_age"`
	MaxBackups int `json:"max_backups"`
}
```

在项目中先从配置文件加载配置信息，再调用`logger.InitLogger(config.Conf.LogConfig)`即可完成logger实例的初识化。其中，通过`r.Use(logger.GinLogger(), logger.GinRecovery(true))`注册我们的中间件来使用zap接收gin框架自身的日志，在项目中需要的地方通过使用`zap.L().Xxx()`方法来记录自定义日志信息。

```go
package main

import (
	"fmt"
	"gin_zap_demo/config"
	"gin_zap_demo/logger"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config from config.json
	if len(os.Args) < 1 {
		return
	}

	if err := config.Init(os.Args[1]); err != nil {
		panic(err)
	}
	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	gin.SetMode(config.Conf.Mode)

	r := gin.Default()
	// 注册zap相关中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		// 假设你有一些数据需要记录到日志中
		var (
			name = "q1mi"
			age  = 18
		)
		// 记录日志并使用zap.Xxx(key, val)记录相关字段
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))

		c.String(http.StatusOK, "hello liwenzhou.com!")
	})

	addr := fmt.Sprintf(":%v", config.Conf.Port)
	r.Run(addr)
}
```

## Viper

[Viper](https://github.com/spf13/viper)是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

鉴于`viper`库本身的README已经写得十分详细，这里就将其翻译成中文，并在最后附上两个项目中使用`viper`的示例代码以供参考。

### 安装

```bash
go get github.com/spf13/viper
```

### 什么是Viper？

Viper是适用于Go应用程序（包括`Twelve-Factor App`）的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。它支持以下特性：

- 设置默认值
- 从`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件读取配置信息
- 实时监控和重新读取配置文件（可选）
- 从环境变量中读取
- 从远程配置系统（etcd或Consul）读取并监控配置变化
- 从命令行参数读取配置
- 从buffer读取配置
- 显式配置值

### 为什么选择Viper?

在构建现代应用程序时，你无需担心配置文件格式；你想要专注于构建出色的软件。Viper的出现就是为了在这方面帮助你的。

Viper能够为你执行下列操作：

1. 查找、加载和反序列化`JSON`、`TOML`、`YAML`、`HCL`、`INI`、`envfile`和`Java properties`格式的配置文件。
2. 提供一种机制为你的不同配置选项设置默认值。
3. 提供一种机制来通过命令行参数覆盖指定选项的值。
4. 提供别名系统，以便在不破坏现有代码的情况下轻松重命名参数。
5. 当用户提供了与默认值相同的命令行或配置文件时，可以很容易地分辨出它们之间的区别。

Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目:

- 显示调用`Set`设置值
- 命令行参数（flag）
- 环境变量
- 配置文件
- key/value存储
- 默认值

**重要：** 目前Viper配置的键（Key）是大小写不敏感的。目前正在讨论是否将这一选项设为可选。

### 把值存入Viper

#### 建立默认值

一个好的配置系统应该支持默认值。键不需要默认值，但如果没有通过配置文件、环境变量、远程配置或命令行标志（flag）设置键，则默认值非常有用。

例如：

```go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

#### 读取配置文件

Viper需要最少知道在哪里查找配置文件的配置。Viper支持`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件。Viper可以搜索多个路径，但目前单个Viper实例只支持单个配置文件。Viper不默认任何配置搜索路径，将默认决策留给应用程序。

下面是一个如何使用Viper搜索和读取配置文件的示例。不需要任何特定的路径，但是至少应该提供一个配置文件预期出现的路径。

```go
viper.SetConfigFile("./config.yaml") // 指定配置文件路径
viper.SetConfigName("config") // 配置文件名称(无扩展名)
viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件的错误
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

在加载配置文件出错时，你可以像下面这样处理找不到配置文件的特定情况：

```go
if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // 配置文件未找到错误；如果需要可以忽略
    } else {
        // 配置文件被找到，但产生了另外的错误
    }
}

// 配置文件找到并成功解析
```

*注意[自1.6起]：* 你也可以有不带扩展名的文件，并以编程方式指定其格式。对于位于用户`$HOME`目录中的配置文件没有任何扩展名，如`.bashrc`。

**这里补充两个问题供读者解答并自行验证**

当你使用如下方式读取配置时，viper会从`./conf`目录下查找任何以`config`为文件名的配置文件，如果同时存在`./conf/config.json`和`./conf/config.yaml`两个配置文件的话，`viper`会从哪个配置文件加载配置呢？

```go
viper.SetConfigName("config")
viper.AddConfigPath("./conf")
```

在上面两个语句下搭配使用`viper.SetConfigType("yaml")`指定配置文件类型可以实现预期的效果吗？

#### 写入配置文件

从配置文件中读取配置文件是有用的，但是有时你想要存储在运行时所做的所有修改。为此，可以使用下面一组命令，每个命令都有自己的用途:

- WriteConfig - 将当前的`viper`配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
- SafeWriteConfig - 将当前的`viper`配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
- WriteConfigAs - 将当前的`viper`配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
- SafeWriteConfigAs - 将当前的`viper`配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。

根据经验，标记为`safe`的所有方法都不会覆盖任何文件，而是直接创建（如果不存在），而默认行为是创建或截断。

一个小示例：

```go
viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 因为该配置文件写入过，所以会报错
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

#### 监控并重新读取配置文件

Viper支持在运行时实时读取配置文件的功能。

需要重新启动服务器以使配置生效的日子已经一去不复返了，viper驱动的应用程序可以在运行时读取配置文件的更新，而不会错过任何消息。

只需告诉viper实例watchConfig。可选地，你可以为Viper提供一个回调函数，以便在每次发生更改时运行。

**确保在调用`WatchConfig()`之前添加了所有的配置路径。**

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  // 配置文件发生变更之后会调用的回调函数
	fmt.Println("Config file changed:", e.Name)
})
```

#### 从io.Reader读取配置

Viper预先定义了许多配置源，如文件、环境变量、标志和远程K/V存储，但你不受其约束。你还可以实现自己所需的配置源并将其提供给viper。

```go
viper.SetConfigType("yaml") // 或者 viper.SetConfigType("YAML")

// 任何需要将此配置添加到程序中的方法。
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name") // 这里会得到 "steve"
```

#### 覆盖设置

这些可能来自命令行标志，也可能来自你自己的应用程序逻辑。

```go
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)
```

#### 注册和使用别名

别名允许多个键引用单个值

```go
viper.RegisterAlias("loud", "Verbose")  // 注册别名（此处loud和Verbose建立了别名）

viper.Set("verbose", true) // 结果与下一行相同
viper.Set("loud", true)   // 结果与前一行相同

viper.GetBool("loud") // true
viper.GetBool("verbose") // true
```

#### 使用环境变量

Viper完全支持环境变量。这使`Twelve-Factor App`开箱即用。有五种方法可以帮助与ENV协作:

- `AutomaticEnv()`
- `BindEnv(string...) : error`
- `SetEnvPrefix(string)`
- `SetEnvKeyReplacer(string...) *strings.Replacer`
- `AllowEmptyEnv(bool)`

*使用ENV变量时，务必要意识到Viper将ENV变量视为区分大小写。*

Viper提供了一种机制来确保ENV变量是惟一的。通过使用`SetEnvPrefix`，你可以告诉Viper在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用这个前缀。

`BindEnv`使用一个或两个参数。第一个参数是键名称，第二个是环境变量的名称。环境变量的名称区分大小写。如果没有提供ENV变量名，那么Viper将自动假设ENV变量与以下格式匹配：前缀+ “_” +键名全部大写。当你显式提供ENV变量名（第二个参数）时，它 **不会** 自动添加前缀。例如，如果第二个参数是“id”，Viper将查找环境变量“ID”。

在使用ENV变量时，需要注意的一件重要事情是，每次访问该值时都将读取它。Viper在调用`BindEnv`时不固定该值。

`AutomaticEnv`是一个强大的助手，尤其是与`SetEnvPrefix`结合使用时。调用时，Viper会在发出`viper.Get`请求时随时检查环境变量。它将应用以下规则。它将检查环境变量的名称是否与键匹配（如果设置了`EnvPrefix`）。

`SetEnvKeyReplacer`允许你使用`strings.Replacer`对象在一定程度上重写 Env 键。如果你希望在`Get()`调用中使用`-`或者其他什么符号，但是环境变量里使用`_`分隔符，那么这个功能是非常有用的。可以在`viper_test.go`中找到它的使用示例。

或者，你可以使用带有`NewWithOptions`工厂函数的`EnvKeyReplacer`。与`SetEnvKeyReplacer`不同，它接受`StringReplacer`接口，允许你编写自定义字符串替换逻辑。

默认情况下，空环境变量被认为是未设置的，并将返回到下一个配置源。若要将空环境变量视为已设置，请使用`AllowEmptyEnv`方法。

#### Env 示例：

```go
SetEnvPrefix("spf") // 将自动转为大写
BindEnv("id")

os.Setenv("SPF_ID", "13") // 通常是在应用程序之外完成的

id := Get("id") // 13
```

#### 使用Flags

Viper 具有绑定到标志的能力。具体来说，Viper支持[Cobra](https://github.com/spf13/cobra)库中使用的`Pflag`。

与`BindEnv`类似，该值不是在调用绑定方法时设置的，而是在访问该方法时设置的。这意味着你可以根据需要尽早进行绑定，即使在`init()`函数中也是如此。

对于单个标志，`BindPFlag()`方法提供此功能。

例如：

```go
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

你还可以绑定一组现有的pflags （pflag.FlagSet）：

举个例子：

```go
pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") // 从viper而不是从pflag检索值
```

在 Viper 中使用 pflag 并不阻碍其他包中使用标准库中的 flag 包。pflag 包可以通过导入这些 flags 来处理flag包定义的flags。这是通过调用pflag包提供的便利函数`AddGoFlagSet()`来实现的。

例如：

```go
package main

import (
	"flag"
	"github.com/spf13/pflag"
)

func main() {

	// 使用标准库 "flag" 包
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // 从 viper 检索值

	...
}
```

#### flag接口

如果你不使用`Pflag`，Viper 提供了两个Go接口来绑定其他 flag 系统。

`FlagValue`表示单个flag。这是一个关于如何实现这个接口的非常简单的例子：

```go
type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }
```

一旦你的 flag 实现了这个接口，你可以很方便地告诉Viper绑定它：

```go
viper.BindFlagValue("my-flag-name", myFlag{})
```

`FlagValueSet`代表一组 flags 。这是一个关于如何实现这个接口的非常简单的例子:

```go
type myFlagSet struct {
	flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}
```

一旦你的flag set实现了这个接口，你就可以很方便地告诉Viper绑定它：

```go
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)
```

#### 远程Key/Value存储支持

在Viper中启用远程支持，需要在代码中匿名导入`viper/remote`这个包。

```
import _ "github.com/spf13/viper/remote"
```

Viper将读取从Key/Value存储（例如etcd或Consul）中的路径检索到的配置字符串（如`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式）。这些值的优先级高于默认值，但是会被从磁盘、flag或环境变量检索到的配置值覆盖。（译注：也就是说Viper加载配置值的优先级为：磁盘上的配置文件>命令行标志位>环境变量>远程Key/Value存储>默认值。）

Viper使用[crypt](https://github.com/bketelsen/crypt)从K/V存储中检索配置，这意味着如果你有正确的gpg密匙，你可以将配置值加密存储并自动解密。加密是可选的。

你可以将远程配置与本地配置结合使用，也可以独立使用。

`crypt`有一个命令行助手，你可以使用它将配置放入K/V存储中。`crypt`默认使用在[http://127.0.0.1:4001](http://127.0.0.1:4001/)的etcd。

```bash
$ go get github.com/bketelsen/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
```

确认值已经设置：

```bash
$ crypt get -plaintext /config/hugo.json
```

有关如何设置加密值或如何使用Consul的示例，请参见`crypt`文档。

#### 远程Key/Value存储示例-未加密

##### etcd

```go
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

##### Consul

你需要 Consul Key/Value存储中设置一个Key保存包含所需配置的JSON值。例如，创建一个key`MY_CONSUL_KEY`将下面的值存入Consul key/value 存储：

```json
{
    "port": 8080,
    "hostname": "liwenzhou.com"
}
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json") // 需要显示设置成json
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port")) // 8080
fmt.Println(viper.Get("hostname")) // liwenzhou.com
```

##### Firestore

```go
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") // 配置的格式: "json", "toml", "yaml", "yml"
err := viper.ReadRemoteConfig()
```

当然，你也可以使用`SecureRemoteProvider`。

#### 远程Key/Value存储示例-加密

```go
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

#### 监控etcd中的更改-未加密

```go
// 或者你可以创建一个新的viper实例
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

// 第一次从远程读取配置
err := runtime_viper.ReadRemoteConfig()

// 反序列化
runtime_viper.Unmarshal(&runtime_conf)

// 开启一个单独的goroutine一直监控远端的变更
go func(){
	for {
	    time.Sleep(time.Second * 5) // 每次请求后延迟一下

	    // 目前只测试了etcd支持
	    err := runtime_viper.WatchRemoteConfig()
	    if err != nil {
	        log.Errorf("unable to read remote config: %v", err)
	        continue
	    }

	    // 将新配置反序列化到我们运行时的配置结构体中。你还可以借助channel实现一个通知系统更改的信号
	    runtime_viper.Unmarshal(&runtime_conf)
	}
}()
```

### 从Viper获取值

在Viper中，有几种方法可以根据值的类型获取值。存在以下功能和方法:

- `Get(key string) : interface{}`
- `GetBool(key string) : bool`
- `GetFloat64(key string) : float64`
- `GetInt(key string) : int`
- `GetIntSlice(key string) : []int`
- `GetString(key string) : string`
- `GetStringMap(key string) : map[string]interface{}`
- `GetStringMapString(key string) : map[string]string`
- `GetStringSlice(key string) : []string`
- `GetTime(key string) : time.Time`
- `GetDuration(key string) : time.Duration`
- `IsSet(key string) : bool`
- `AllSettings() : map[string]interface{}`

需要认识到的一件重要事情是，每一个Get方法在找不到值的时候都会返回零值。为了检查给定的键是否存在，提供了`IsSet()`方法。

例如：

```go
viper.GetString("logfile") // 不区分大小写的设置和获取
if viper.GetBool("verbose") {
    fmt.Println("verbose enabled")
}
```

#### 访问嵌套的键

访问器方法也接受深度嵌套键的格式化路径。例如，如果加载下面的JSON文件：

```json
{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}
```

Viper可以通过传入`.`分隔的路径来访问嵌套字段：

```go
GetString("datastore.metric.host") // (返回 "127.0.0.1")
```

这遵守上面建立的优先规则；搜索路径将遍历其余配置注册表，直到找到为止。(译注：因为Viper支持从多种配置来源，例如磁盘上的配置文件>命令行标志位>环境变量>远程Key/Value存储>默认值，我们在查找一个配置的时候如果在当前配置源中没找到，就会继续从后续的配置源查找，直到找到为止。)

例如，在给定此配置文件的情况下，`datastore.metric.host`和`datastore.metric.port`均已定义（并且可以被覆盖）。如果另外在默认值中定义了`datastore.metric.protocol`，Viper也会找到它。

然而，如果`datastore.metric`被直接赋值覆盖（被flag，环境变量，`set()`方法等等…），那么`datastore.metric`的所有子键都将变为未定义状态，它们被高优先级配置级别“遮蔽”（shadowed）了。

最后，如果存在与分隔的键路径匹配的键，则返回其值。例如：

```go
{
    "datastore.metric.host": "0.0.0.0",
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

GetString("datastore.metric.host") // 返回 "0.0.0.0"
```

#### 提取子树

从Viper中提取子树。

例如，`viper`实例现在代表了以下配置：

```yaml
app:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80
```

执行后：

```go
subv := viper.Sub("app.cache1")
```

`subv`现在就代表：

```yaml
max-items: 100
item-size: 64
```

假设我们现在有这么一个函数：

```go
func NewCache(cfg *Viper) *Cache {...}
```

它基于`subv`格式的配置信息创建缓存。现在，可以轻松地分别创建这两个缓存，如下所示：

```go
cfg1 := viper.Sub("app.cache1")
cache1 := NewCache(cfg1)

cfg2 := viper.Sub("app.cache2")
cache2 := NewCache(cfg2)
```

#### 反序列化

你还可以选择将所有或特定的值解析到结构体、map等。

有两种方法可以做到这一点：

- `Unmarshal(rawVal interface{}) : error`
- `UnmarshalKey(key string, rawVal interface{}) : error`

举个例子：

```go
type config struct {
	Port int
	Name string
	PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
```

如果你想要解析那些键本身就包含`.`(默认的键分隔符）的配置，你需要修改分隔符：

```go
v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]interface{}{
    "ingress": map[string]interface{}{
        "annotations": map[string]interface{}{
            "traefik.frontend.rule.type":"PathPrefix",
            "traefik.ingress.kubernetes.io/ssl-redirect": "true",
        },
    },
})

type config struct {
	Chart struct{
        Values map[string]interface{}
    }
}

var C config

v.Unmarshal(&C)
```

Viper还支持解析到嵌入的结构体：

```go
/*
Example config:

module:
    enabled: true
    token: 89h3f98hbwf987h3f98wenf89ehf
*/
type config struct {
	Module struct {
		Enabled bool

		moduleConfig `mapstructure:",squash"`
	}
}

// moduleConfig could be in a module specific package
type moduleConfig struct {
	Token string
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
```

Viper在后台使用[github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)来解析值，其默认情况下使用`mapstructure`tag。

**注意** 当我们需要将viper读取的配置反序列到我们定义的结构体变量中时，一定要使用`mapstructure`tag哦！

#### 序列化成字符串

你可能需要将viper中保存的所有设置序列化到一个字符串中，而不是将它们写入到一个文件中。你可以将自己喜欢的格式的序列化器与`AllSettings()`返回的配置一起使用。

```go
import (
    yaml "gopkg.in/yaml.v2"
    // ...
)

func yamlStringSettings() string {
    c := viper.AllSettings()
    bs, err := yaml.Marshal(c)
    if err != nil {
        log.Fatalf("unable to marshal config to YAML: %v", err)
    }
    return string(bs)
}
```

### 使用单个还是多个Viper实例?

Viper是开箱即用的。你不需要配置或初始化即可开始使用Viper。由于大多数应用程序都希望使用单个中央存储库管理它们的配置信息，所以viper包提供了这个功能。它类似于单例模式。

在上面的所有示例中，它们都以其单例风格的方法演示了如何使用viper。

#### 使用多个viper实例

你还可以在应用程序中创建许多不同的viper实例。每个都有自己独特的一组配置和值。每个人都可以从不同的配置文件，key value存储区等读取数据。每个都可以从不同的配置文件、键值存储等中读取。viper包支持的所有功能都被镜像为viper实例的方法。

例如：

```go
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")

//...
```

当使用多个viper实例时，由用户来管理不同的viper实例。

### 使用Viper示例

假设我们的项目现在有一个`./conf/config.yaml`配置文件，内容如下：

```yaml
port: 8123
version: "v1.2.3"
```

接下来通过示例代码演示两种在项目中使用`viper`管理项目配置信息的方式。

#### 直接使用viper管理配置

这里用一个demo演示如何在gin框架搭建的web项目中使用`viper`，使用viper加载配置文件中的信息，并在代码中直接使用`viper.GetXXX()`方法获取对应的配置值。

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()

	r := gin.Default()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	if err := r.Run(
		fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}
```

#### 使用结构体变量保存配置信息

除了上面的用法外，我们还可以在项目中定义与配置文件对应的结构体，`viper`加载完配置信息后使用结构体变量保存配置信息。

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

var Conf = new(Config)

func main() {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()               // 读取配置信息
	if err != nil {                           // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	r := gin.Default()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, Conf.Version)
	})

	if err := r.Run(fmt.Sprintf(":%d", Conf.Port)); err != nil {
		panic(err)
	}
}
```

**参考链接：**https://github.com/spf13/viper/blob/master/README.md

## 优雅关机和重启

我们编写的Web项目部署之后，经常会因为需要进行配置变更或功能迭代而重启服务，单纯的`kill -9 pid`的方式会强制关闭进程，这样就会导致服务端当前正在处理的请求失败，那有没有更优雅的方式来实现关机或重启呢？

> 阅读本文需要了解一些UNIX系统中`信号`的概念，请提前查阅资料预习。

### 优雅地关机

#### 什么是优雅关机？

优雅关机就是服务端关机命令发出后不是立即关机，而是等待当前还在处理的请求全部处理完毕后再退出程序，是一种对客户端友好的关机方式。而执行`Ctrl+C`关闭服务端时，会强制结束进程导致正在访问的请求出现问题。

#### 如何实现优雅关机？

Go 1.8版本之后， http.Server 内置的 [Shutdown()](https://golang.org/pkg/net/http/#Server.Shutdown) 方法就支持优雅地关机，具体示例如下：

```go
// +build go1.8

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  // 此处不会阻塞
	<-quit  // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
```

如何验证优雅关机的效果呢？

上面的代码运行后会在本地的`8080`端口开启一个web服务，它只注册了一条路由`/`，后端服务会先sleep 5秒钟然后才返回响应信息。

我们按下`Ctrl+C`时会发送`syscall.SIGINT`来通知程序优雅关机，具体做法如下：

1. 打开终端，编译并执行上面的代码
2. 打开一个浏览器，访问`127.0.0.1:8080/`，此时浏览器白屏等待服务端返回响应。
3. 在终端**迅速**执行`Ctrl+C`命令给程序发送`syscall.SIGINT`信号
4. 此时程序并不立即退出而是等我们第2步的响应返回之后再退出，从而实现优雅关机。

### 优雅地重启

优雅关机实现了，那么该如何实现优雅重启呢？

我们可以使用 [fvbock/endless](https://github.com/fvbock/endless) 来替换默认的 `ListenAndServe`启动服务来实现， 示例代码如下：

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "hello gin!")
	})
	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	if err := endless.ListenAndServe(":8080", router); err!=nil{
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")
}
```

如何验证优雅重启的效果呢？

我们通过执行`kill -1 pid`命令发送`syscall.SIGINT`来通知程序优雅重启，具体做法如下：

1. 打开终端，`go build -o graceful_restart`编译并执行`./graceful_restart`,终端输出当前pid(假设为43682)
2. 将代码中处理请求函数返回的`hello gin!`修改为`hello q1mi!`，再次编译`go build -o graceful_restart`
3. 打开一个浏览器，访问`127.0.0.1:8080/`，此时浏览器白屏等待服务端返回响应。
4. 在终端**迅速**执行`kill -1 43682`命令给程序发送`syscall.SIGHUP`信号
5. 等第3步浏览器收到响应信息`hello gin!`后再次访问`127.0.0.1:8080/`会收到`hello q1mi!`的响应。
6. 在不影响当前未处理完请求的同时完成了程序代码的替换，实现了优雅重启。

但是需要注意的是，此时程序的PID变化了，因为`endless` 是通过`fork`子进程处理新请求，待原进程处理完当前请求后再退出的方式实现优雅重启的。所以当你的项目是使用类似`supervisor`的软件管理进程时就**不适用**这种方式了。

### 总结

无论是优雅关机还是优雅重启归根结底都是通过监听特定系统信号，然后执行一定的逻辑处理保障当前系统正在处理的请求被正常处理后再关闭当前进程。使用优雅关机还是使用优雅重启以及怎么实现，这就需要根据项目实际情况来决定了。
