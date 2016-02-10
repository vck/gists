import socket
import sys 


sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
host = socket.gethostname()
port = int(sys.argv[1])
sock.bind(('localhost', port))
sock.listen(5)
print 'server listening on 192.168.43.219:%s'% port
while True:
   client_connection, address = sock.accept()
   print 'connection from ', address
   client_connection.sendall("""\
   	HTTP/1.1 200 OK

   	Hello, World
	""")
   client_connection.close()

   


