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
 
 
 
 
 