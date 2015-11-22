package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

// This struct handles:
// - users joining
// - users disconnecting
// - receiving individual messages from users and broadcasting them
// to other users
//
// `users` contain ChatUser connections.
// `incoming` receives incoming messages from ChatUser connections.
// `joins` receives incoming new ChatUser connections.
// `disconnects` receives disconnect notifications.
//
type ChatRoom struct {
	users       map[string]*ChatUser
	incoming    chan string
	joins       chan *ChatUser
	disconnects chan string
}

// NewChatRoom will create a ChatRoom
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users:       make(map[string]*ChatUser),
		incoming:    make(chan string),
		joins:       make(chan *ChatUser),
		disconnects: make(chan string),
	}
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
//
// `conn` is the socket
// `disconnect` indicated whether or not the socket is disconnected
// `username` is the chat username
// `outgoing` is a channel with all pending outgoing messages
//  to be written to the socket.
// `reader` is the buffered socket read stream
//  `writer` is the buffered socket write stream
//
type ChatUser struct {
	conn       net.Conn
	disconnect bool
	username   string
	outgoing   chan string
	reader     *bufio.Reader
	writer     *bufio.Writer
}

// NewChatUser creates a new ChatUser
func NewChatUser(conn net.Conn) *ChatUser {

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	return &ChatUser{
		conn:       conn,
		disconnect: false,
		reader:     reader,
		writer:     writer,
		outgoing:   make(chan string),
		username:   "",
	}
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
	listener, err := net.Listen("tcp", ":6677")

	if err != nil {
		log.Println("Error whole listening on port 6677")
		log.Println("Error: (%s)", err)
		os.Exit(1)
	}

	// Create a new instance of chatroom using NewChatRoom()
	chatroom := NewChatRoom()

	// and call chatroom.ListenForMessages()
	chatroom.ListenForMessages()

	// Loop and listen for accepted connections on port 6677
	for {

		// accept, wait for and return the next connection on listener
		c, err := listener.Accept()

		if err != nil {
			log.Println("Error while accepting network connections")
			log.Println("Error : (%s)", err)
			os.Exit(1)
		}

		// Print out the remote address of the connection
		addr := c.RemoteAddr()

		log.Println("The remote address of connection is %s", addr)

	}

}
