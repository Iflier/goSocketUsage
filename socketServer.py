# -*- coding:utf-8 -*-
"""
Dec: server端，可以用Go语言的client接收，也可以用Python语言的client接收
Created on: 2018.10.11
Author: Iflier
"""

import time
import socket
import argparse
from datetime import datetime


ap = argparse.ArgumentParser()
ap.add_argument("-p", "--port", type=int, default=60000, help="Specify server open which one port")
args = vars(ap.parse_args())

ADDRESS = ("127.0.0.1", args["port"])

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.bind(ADDRESS)
sock.listen()
conn, addr = sock.accept()
print("addr: {0}".format(addr))

for _ in range(10):
    numBytesWritten = conn.send("Yes.".encode())
    print("On {0}, written {1:^5,d} bytes.".format(datetime.now().strftime("%c"), numBytesWritten))
    time.sleep(1.0)
    receivedBytes = conn.recv(128)
    print("On {0}, received: {1}".format(datetime.now().strftime("%c"), receivedBytes.decode()))

numBytesWritten = conn.send("exit".encode())
print("On {0}, written {1:^5,d} bytes.".format(datetime.now().strftime("%c"), numBytesWritten))
conn.close()
sock.close()
print("Done.")
