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
 
 