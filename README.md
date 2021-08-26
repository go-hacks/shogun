# shogun
Go Shotgun

Proof of concept to send data across a network without actually making any TCP connections or transmitting any data in the packets.

Naturally, it is rather slow, about 100bytes/s on my local machine at the default 10ms delay. If set to 1ms it's around 770bytes/s but I doubt that would work anywhere besides the same local machine.

It could be faster with more ports but the math would be very weird as even 2 bytes would require all existing ports and generally wouldn't work or you'd have to use multiple file readers but then positional and file size data would have to be communicated as well (probably down the first set of 256). Then you could have several 256 port byte streams and theoretically multiply the speed to upwards of 20-100KB/s on a local network depending on how low you could get the sleep delay and still keep data integrity.

Note: Sleep delay may need to be increased if on a slower network so that packets don't arrive out of order.
This will decrease transfer speeds but necessary if your network has high jitter (inconsistent ping).
May be almost unusable on bad cell data. You'd have to slow it down so far it would be impractical.

The test sends 2KB (2048 bytes) of random data.
Obviously you could replace this with anything, like a file reader. 

P.S. I am not responsible for what you do with this. I just made it for fun :) 
