Trivial File Transport Protocol (TFTP )
--
--

This is a golang implementation of the TFTP protocol. This repo currently supports only the server. It is possible to use the tftp linux utility (http://www.tutorialspoint.com/unix_commands/tftp.htm) to work with the server. The client is on its way.

<h1>Features</h1>
<lo>
<li>putting file on server</li>
<li>getting file from server</li>
<li>concurrent requests</li>
<li>deduplicated storage</li>
<li>handles large files</li>
<li>built with extendibility in mind</li>
<li>modularized</li>


</lo>





<h1>How to run server</h1>
go to tftp repo and run:  
go run ./tftp_server/server.go


<h1>How to run tests</h1>
go to tftp repo and run:  
go test ./tests/

