# SWEB
The Simple WEB Server.

I wrote this very small web server to replace the fatal flaw in the `python2 -m SimpleHTTPServer 8080`
web server that could only handle a single request at a time. If you have a use case to download 
files from the web server, or have several people download files at the same time, the python 
implementation becomes useless... This little server is designed to be small, portable, and
generally the most useful, small webserver out there.

### Example Usage
* `sweb`
	1. Serves the directory you called it from
	2. Binds to port 8080
	3. Binds to the 0.0.0.0 address

* `sweb -d ~ -p 9000 -a localhost`
	1. Serves your home directory (~)
	2. Binds to port 9000
	3. Binds to 127.0.0.1 (localhost) only
