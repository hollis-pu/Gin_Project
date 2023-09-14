# Gin框架学习笔记

2023.09.13

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

## Gin框架

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

```template
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
