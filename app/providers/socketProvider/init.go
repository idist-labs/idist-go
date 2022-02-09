package socketProvider

func Init() {
	//fmt.Println("------------------------------------------------------------")
	//router := gin.New()
	//config := configProvider.GetConfig()
	//
	//opts := engineio.Options{
	//	Transports: []transport.Transport{
	//		websocket.Default},
	//}
	//server := socketio.NewServer(&opts)
	//if _, err := server.Adapter(&socketio.RedisAdapterOptions{
	//	Host:   config.GetString("redis.host"),
	//	Port:   config.GetString("redis.port"),
	//	Prefix: "socket.io",
	//}); err != nil {
	//	log.Fatal("error:", err)
	//} else {
	//	fmt.Println("Setup Redis broadcast adapter done!")
	//}
	//server.OnConnect("/", sockets.Onconnect)
	//server.OnDisconnect("/", sockets.OnDisconnect)
	////routes.SocketRoutes(server)
	//router.Any("/api/socket/*any", gin.WrapH(server))
	//go func() {
	//	if err := server.Serve(); err != nil {
	//		loggerProvider.Logger.Error("Start socket server has err", zap.Error(err))
	//	} else {
	//		loggerProvider.Logger.Info("Socket server started")
	//	}
	//}()
	//defer server.Close()
}
