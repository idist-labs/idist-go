package socketProvider

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"idist-go/app/sockets"
	"idist-go/routes"
)

func Init(router *gin.Engine) {
	fmt.Println("------------------------------------------------------------")
	server := socketio.NewServer(nil)
	server.OnConnect("/", sockets.Onconnect)
	server.OnDisconnect("/", sockets.OnDisconnect)
	routes.SocketRoutes(server)
	go server.Serve()
	defer server.Close()
	router.GET("/socket/*any", gin.WrapH(server))
}
