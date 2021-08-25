# shogun
Go Shotgun

Proof of concept to send data across a network without actually making any TCP connections or transmitting any data in the UDP packets.
Naturally, it is rather slow, about 100bytes/s on my local machine. 

Note: Sleep delay may need to be increased if on a slower network so that packets don't arrive out of order. 
The test sends 2048 bytes of random data.
Obviously you could replace this with anything, like a file reader. 

P.S. I am not responsible for what you do with this.
I just made it for fun :) 
