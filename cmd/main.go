package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/handlers"
)

func main() {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Chat room and message routes
	router.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	router.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")
	router.HandleFunc("/rooms/{room_id}/messages", handlers.GetMessageHistory).Methods("GET")

	// Reactions and read receipts
	router.HandleFunc("/messages/{message_id}/reactions", handlers.GetReactions).Methods("GET")
	router.HandleFunc("/reactions", handlers.AddReaction).Methods("POST")
	router.HandleFunc("/messages/{message_id}/read_receipts", handlers.GetReadReceipts).Methods("GET")
	router.HandleFunc("/read_receipts", handlers.SaveReadReceipt).Methods("POST")

	// WebSocket route
	router.HandleFunc("/ws", websockets.JoinRoom).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
