# PaillierVoteSystemRemoteRouter

这个项目是一个使用 Paillier 密码体系实现的远程投票系统的服务器端实现。它使用 Go 语言编写，支持多用户投票、投票计数和加密数据存储等功能。
# 项目结构

该项目包含以下文件夹和文件：

    main 文件夹：包含了服务器端的主要代码实现，包括路由器、投票逻辑、加密和解密功能等。

    frontend 文件夹：包含了客户端的代码实现，使用 HTML、CSS 和 JavaScript 实现。

    crypto 文件夹：包含了 Paillier 加密算法的实现，以及其他密码学算法的实现。

    database 文件夹：包含了数据库模型的定义和实现。

    tests 文件夹：包含了一些单元测试和集成测试的代码。

    其他文件：如 .gitignore 文件和 README.md 文件等。

在 main 文件夹中，主要文件有：

    router.go：定义了一个 HTTP 服务器的路由器，它处理客户端发来的请求，将请求路由到正确的处理程序。

    vote.go：定义了投票逻辑的实现，包括投票计数、加密和解密等。

    crypto.go：包含了 Paillier 加密算法的实现。

    db.go：定义了数据库模型的实现，包括创建数据库表、插入、查询等。

    server.go：定义了服务器的入口函数，它负责初始化路由器、连接数据库、启动 HTTP 服务器等。

此外，还有其他一些辅助文件和配置文件，如 config.go、utils.go、docker-compose.yml 等。
# 如何运行

要运行该项目，您需要使用以下步骤：

    安装 Go 和 Docker。

    克隆该项目的仓库。

    进入项目根目录，运行以下命令来构建 Docker 镜像：

docker-compose build

运行以下命令来启动 Docker 容器：

    docker-compose up

    在浏览器中打开 http://localhost:8080 来访问该项目的主页。

# 如何使用

在使用该项目之前，您需要了解以下几点：

    该项目需要使用 Paillier 密码体系进行加密和解密。

    您需要在 config.go 文件中设置数据库的用户名、密码和主机名等参数。

    您需要在 config.go 文件中设置 Paillier 密钥的长度和公钥、私钥的文件路径等参数。

使用该项目的步骤如下：
    在浏览器中打开 http://localhost:8080/login 页面，使用您的用户名和密码登录。

    在浏览器中打开 http://localhost:8080/vote 页面，选择您要投票的候选人，并提交投票。

    在浏览器中打开 http://localhost:8080/result 页面，查看投票结果。

# 贡献者

该项目的贡献者包括：

    Microsoft-tele

    其他参与贡献的开发者


# 结语

该项目是一个使用 Paillier 密码体系实现的远程投票系统的服务器端实现。它具有多用户投票、投票计数和加密数据存储等功能，可以用于各种需要保证安全性的投票场景。如果您有任何问题或建议，请随时与我们联系。
