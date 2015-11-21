package main

import (
	"log"
	"net"
)

// This struct handles:
// - users joining
// - users disconnecting
// - receiving individual messages from users and broadcasting them
// to other users
type ChatRoom struct {
}

// NewChatRoom will create a ChatRoom
func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

// Listen for messages in the ChatRoom
func (cr *ChatRoom) ListenForMessages() {
}

// Logout a user from the ChatRoom
func (cr *ChatRoom) Logout(username string) {
}

// Allows a user to join the ChatRoom
func (cr *ChatRoom) Join(conn net.Conn) {
}

// Broadcast a message
func (cr *ChatRoom) Broadcast(msg string) {
}

// This struct handles:
// - reading lines of data from user socket and notifying the chatroom
// there is a new message
// - writing data back to the socket (eg messages from other users)
type ChatUser struct {
}

// NewChatUser creates a new ChatUser
func NewChatUser(conn net.Conn) *ChatUser {
	return &ChatUser{}
}

// Read incoming messages in a loop
func (cu *ChatUser) ReadIncomingMessages(chatroom *ChatRoom) {
}

// Wait for outgoing messages in a loop, and write them
func (cu *ChatUser) WriteOutgoingMessages(chatroom *ChatRoom) {
}

// Login the user
func (cu *ChatUser) Login(chatroom *ChatRoom) error {
	return nil
}

// Read a line from the socket
func (cu *ChatUser) ReadLine() (string, error) {
	return "", nil
}

// Write a line from the socket
func (cu *ChatUser) WriteString(msg string) error {
	return nil
}

// Put a message on the outgoing message queue
func (cu *ChatUser) Send(msg string) {
}

// Close the socket
func (cu *ChatUser) Close() {
}

// Main function to create a socket, bind to port 6677
// and loop while waiting for connections
//
// When it receives a connection, it will pass it to
// `chatroom.Join()`
//
func main() {
	log.Println("Chat server starting")

	// Create a TCP listener on port 6677

	// Create a new instance of chatroom using NewChatRoom()
	// and call chatroom.ListenForMessages()

	// Loop and listen for accepted connections on port 6677

	// Print out remote address of connection using log
}
