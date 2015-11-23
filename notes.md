## Setup the chat project

Nothing much to do over here other than set up the scaffolding code.

## Create the TCP socket

- Create a TCP listener on port 6677. To do this we will use net.Listen function.  
The net.Listen function takes a network type and a network address, and return a listener interface along with any error encountered.  
A listener is a generic interface for stream oriented portocols.

- We create a new chatroom using the NewChatRoom function.

- Now we run an infiinte loop and listen for accepted connections on port 6677.   
To do that we use the Accept() function defined on the listener interface. The Accept() function accepts, waits for and returns the next connection on listener. It returns a net.Conn connection and any error that may occur.  
So what we are doing is that, first we create a listener of type Listener interface for our TCP network on port 6677. Then in an infinite loop, we wait for any incoming connection, which we return and store in the variable `c`. 

- Next we print the remote address of our connection. We can do that by using the RemoteAddr() function defined on the connection interface.
 
 ## Populate `ChatRoom` and `ChatUser`
 
 Design of datastructure and requirements:
 - There's one `ChatRoom` in the app.
 - The `ChatRoom` must know about all active connections.
 - Each connection is related to a user connecting to the `ChatRoom` and hence is tracked in a `ChatUser` object.
 - The `ChatRoom` must be able to recive messages from a single connection, and broadcast to all other connections.
 - When a new connection is established, the `ChatRoom` must be notified of these new connections.
 
 Our `ChatUser` struct contains a private member reader, of type *bufio.Reader, and a private member writer, of type *bufio.Writer.
 Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
 
 We create the `reader` and `writer` variables, using `bufio.NewReader` and `bufio.NewWriter` functions. We will be passing the net.Conn to these functions to get the `reader` and `writer` variables. Interesting thing is that net.Conn interface has the function `Read` and `Write`, hence it satisfies both the `io.Reader` and `io.Writer` interface. So we can pass net.Conn to `bufio.NewReader` and `bufio.NewWriter` functions.
 
 In our main function, inside our infinite for loop, we add each accepted connection to our ChatRoom using the `chatroom.Join()` method. We use the `go` keyword there and make sure that we create a separate goroutine for each accepted connection. 
 
 ## Login to the Chat Server

Now we will:  
- Print a banner every time a user connects.
- Implement a chatuser.Login() method to be able to read a username from the user.
 
 - We are creating a new goroutine for each connection that is accepted by the server and is joined to the `ChatRoom`.
 - In this goroutine when we add a connection to our `ChatRoom`, we will create a new `ChatUser` object for each connection.
 - Then we use this `ChatUser` object to login by invoking its Login() method.
 - We notify the addition of a new user by putting the newly created `ChatUser` object on the `joins` channel of the `ChatRoom`.

 The `chatuser.Login` method is called everytime we accept a new connection, add it to our `ChatRoom` and create a new `ChatUser` object. Once the user  Login, the first thing we do is display our banner. We write the message of the banner to our socket connection by calling the `WriteString` method of `ChatUser`.  
 Inside the `WriteString` method of `ChatUser`, we call the `WriteString` method of our buffered writer. After that we need to call the Flush method of our buffered writer so that it writes any buffered data to the underlying `io.Writer`, which in our case is the socket connection, `net.Conn`.  
 
 Now we are going to read from the socket connection. We want to ask the user's `username` and store it in `ChatUser.username` field.
 
 To read the string from the socket connection we will be using the ReadLine() method of `ChatUser`. In the ReadLine() method, we will call the `ReadString('\n')` method of buffered reader instead of the `ReadLine()` method as mentioned in the docs.
 
 
 
 
 
 
 