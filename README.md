#-\\ simple-chatroom-single-conversational //-
This is the backend part of the single conversational chatroom using go lang...
\\project details//
The websocket is used in the code(gorilla websocket)
>>>rather than using the http connexioniin which the client connects with the server(request and response) and the connection gets terminated,
using websocket(gorilla) the connection remains stable bidirectional connection....
>>> Four channels are created forward,join,leave,client..
>>>forward:holds all the incoming messages,join:join is a channel wishing to join the server(room),leave:wishing to leave the server(room),client:holds all the current client...
>>>




