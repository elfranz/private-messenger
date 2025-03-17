# private-messenger

how am I going to communicate between 2 clients?
investigate what methods of communication there are.

SUMMARIZE THIS AND INCLUDE IMAGES TO EXPLAIN NETWORK APPLICATION
https://www.youtube.com/watch?v=1bRfeKQfDbU

we are going to use TCP protocol since it guarantees reliable data transfer and built in security (TCP with SSL).
also speed in our use case is not priority since its very important to send and receive messages without missing any information
the abstraction from the transport layer that permits us to send packets between applications are sockets.

question raised: do I really need a server to handle communication between clients?
Network Application Architecture: - Client-Server  - Peer to Peer (P2P)
Im more inclined to go P2P since the project will prioritize security and privacy in communication.
Could try both approaches? experiment a bit with them
Not having a server removes risk of it being breached and communications intercepted.
Still gotta investigate more pros and cons in case im missing something on p2p.

/// EVALUATE ARCHITECTURE PROS CONS HERE ///

WHAT IS A SOCKET