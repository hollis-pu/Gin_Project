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

```template
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

```template
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

### 2.路由重定向

路由重定向是指在应用程序的路由层级进行的重定向，它是通过修改路由规则来实现的。在 Gin 中，你可以使用 `c.Redirect` 方法或 `c.Request.URL.Path` 来进行路由重定向。

```go
func RouteRedirectHandler(c *gin.Context) {
    // 使用 c.Redirect 进行路由重定向
    c.Redirect(http.StatusFound, "/new-route")
}
```

或者

```go
func RouteRedirectHandler(c *gin.Context) {
    // 使用 c.Request.URL.Path 进行路由重定向
    c.Request.URL.Path = "/new-route"
    router.HandleContext(c)
}
```

**路由重定向是在服务器端进行的，客户端不需要知道关于重定向的任何信息**，它只需向原始路径发出请求，服务器会根据路由规则将其重定向到新的路径。

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
