package main

import (
	"bufio"
	"fmt"
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

	// Listen in a loop for any messages on the channels in the
	// ChatRoom object and act accordingly
	//
	// We will run a separate goroutine for this
	go func() {
		for {
			select {
			case msg := <-cr.incoming:
				cr.Broadcast(msg)

			case user := <-cr.joins:
				cr.users[user.username] = user
				cr.Broadcast("*** " + user.username + " just joined the channel")
			}
		}
	}()
}

// Logout a user from the ChatRoom
func (cr *ChatRoom) Logout(username string) {
}

// Allows a user to join the ChatRoom
// we run a separate goroutine for each user.
func (cr *ChatRoom) Join(conn net.Conn) {

	// Create a new ChatUser object using NewChatUser
	cu := NewChatUser(conn)

	// call chatuser.Login on this object and verify there is no error
	err := cu.Login(cr)

	if err != nil {
		log.Println("Error while logging in using newly created user")
		log.Println("Error: (%s)", err)
		os.Exit(1)
	}

	// Notifies of a new user by putting the newly created ChatUser
	// on the ChatRoom.joins channel
	cr.joins <- cu
}

// Broadcast a message
func (cr *ChatRoom) Broadcast(msg string) {
	// Broadcast the message to each of the user
	for _, user := range cr.users {
		user.Send(msg)
	}
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
	go func() {
		for {
			msg, _ := cu.ReadLine()
			msg = "[" + cu.username + "]" + msg
			chatroom.incoming <- msg
		}
	}()
}

// Wait for outgoing messages in a loop, and write them
func (cu *ChatUser) WriteOutgoingMessages(chatroom *ChatRoom) {
	// Constantly read a message from the chatuser.outgoing channel
	// and write the message to the socket connection.
	go func() {
		for {
			msg := <-cu.outgoing
			msg += "\n"
			cu.WriteString(msg)
		}
	}()
}

// Login the user
// This method is called everytime a new connection is added to our
// ChatRoom and we create a new ChatUser object
func (cu *ChatUser) Login(chatroom *ChatRoom) error {

	// Create a helpful banner for the user to see when it logs in
	var msg, owner string
	owner = "Mukul"
	msg = fmt.Sprintf("Welcome to %s's Ultimate chat server!\n", owner)
	cu.WriteString(msg)

	// Ask for username
	cu.WriteString("Please Enter your username: ")

	// Read the username from the socket connection
	// we ignore the error, because our functions are written in a way that they
	// will halt execution if an error is encountered.
	cu.username, _ = cu.ReadLine()

	log.Println("User logged in : ", cu.username)

	// Welcome the user by writing out its name in the connection
	cu.WriteString("Welcome " + cu.username + "\n")

	// Start listening on outgoing channel using WriteOutgoingMessages method
	cu.WriteOutgoingMessages(chatroom)

	// Start the goroutine to read an incoming message using ReadIncomingMessages
	cu.ReadIncomingMessages(chatroom)

	return nil
}

// Read a line from the socket
func (cu *ChatUser) ReadLine() (string, error) {

	// We have a ReadLine() function on the Reader interface
	// but according to the docs we should prefer ReadString('\n')
	// method.
	s, err := cu.reader.ReadString('\n')
	s = s[:len(s)-2]
	if err != nil {
		log.Println("Error while reading from socket connection")
		log.Println("Error (%s)", err)
		os.Exit(1)
	}
	return s, nil
}

// Write a line to the socket
func (cu *ChatUser) WriteString(msg string) error {

	// Write the string message on our socket connection
	// which is an io.Writer
	_, err := cu.writer.WriteString(msg)
	if err != nil {
		log.Println("Error while writing the message to socket connection")
		log.Println("Error (%s)", err)
		os.Exit(1)
	}

	// Flush the writer so it writes any buffered data to the underlying
	// io.Writer, which is net.Conn in this case
	err = cu.writer.Flush()
	if err != nil {
		log.Println("Error while flushing the message")
		log.Println("Error (%s)", err)
		os.Exit(1)
	}

	return nil
}

// Put a message on the outgoing message queue
func (cu *ChatUser) Send(msg string) {
	// Place the message on the outgoing channel
	cu.outgoing <- msg
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

		// add this connection to our ChatRoom
		// Create a separate goroutine for each connection
		go chatroom.Join(c)

	}

}
