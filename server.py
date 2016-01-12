import socket
import sys 


sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

host = socket.gethostname()
port = int(sys.argv[1])
sock.bind(('192.168.43.219', port))
sock.listen(5)
print 'server listening on 192.168.43.219:%s'% port
while True:
   c, add = sock.accept()
   print 'connection from ', add
   c.send(""" 
	<html>
		<h1>cieee connect ke server gw ciee </h1>
	</html>
	""")
   


