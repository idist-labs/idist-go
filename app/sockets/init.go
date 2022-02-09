package sockets

//
//import (
//	"fmt"
//	"log"
//	"time"
//)
//
//func Onconnect(s socketio.Conn) error {
//	s.SetContext("/")
//	s.Join("chat")
//	fmt.Println("Connected:", s.ID())
//	s.Emit("user_connect", s.ID())
//	return nil
//}
//
//func OnDisconnect(s socketio.Conn, msg string) {
//	fmt.Println("Somebody just close the connection ")
//	s.LeaveAll()
//}
//
//func OnError(s socketio.Conn, e error) {
//	log.Println("meet error:", e)
//}
//
//func OnNotice(s socketio.Conn, msg string) {
//	log.Println("notice:", msg)
//	time.Sleep(10 * time.Second)
//	s.Emit("reply", "received: "+msg)
//}
//
//func OnChat(s socketio.Conn, data interface{}) {
//	time.Sleep(10 * time.Second)
//	s.Emit("notice", data)
//	for i := 0; i > 0; i++ {
//		s.Emit("internal", i)
//	}
//}
