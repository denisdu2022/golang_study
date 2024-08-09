export default { //注意,对象要抛出后,其他文件中才能引入使用
    host: 'http://127.0.0.1:8080/api',  //后台项目将通过这个域名和端口来启动
    wsHost: 'ws://127.0.0.1:8080/api', //websocket服务端的api地址
}