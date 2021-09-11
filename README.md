# shogun
Go Shotgun

Proof of concept to send data across a network without actually making any TCP connections or transmitting any data in the packets.

Build with chmod +x build && ./build

Start server with ./server fileName

Begin file transfer with ./client

Naturally, it is rather slow, about 12KB/s on my local machine.

It could be faster with more ports but the math would be very weird as even 2 bytes would require all existing ports and generally wouldn't work or you'd have to use multiple file readers but then positional and file size data would have to be communicated as well (probably down the first set of 256). Then you could have several 256 port byte streams and theoretically multiply the speed. 

Note: File is loaded entirely into memory. Don't use it on files bigger than your RAM. 

P.S. I am not responsible for what you do with this. I just made it for fun :) 
