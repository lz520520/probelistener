花半小时写的一个简单小工具，凑合用。

监听目标服务器出网请求包，判断哪个端口可出网，使用-h查看。

主要好处是同端口的多个连接请求不阻塞，会过滤请求头部，减少误报。