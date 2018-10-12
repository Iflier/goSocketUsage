# -*- coding:utf-8 -*-
"""
Dec: socker客户端，接受Go语言server端发送的消息
Created on: 2018.10.11
Author: Iflier
"""

import time
import socket
import argparse
from datetime import datetime


ap = argparse.ArgumentParser()
ap.add_argument("-p", "--port", type=int, default=60000, help="Specify server port to connect")
args = vars(ap.parse_args())

ADDRESS = ("127.0.0.1", args["port"])
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(ADDRESS)

while True:
    receivedBytes = sock.recv(128)
    print("On {0}, received: {1}".format(datetime.now().strftime("%c"), receivedBytes.decode()))
    if receivedBytes.decode().lower() in ["exit", "quit"]:
        break
    time.sleep(1.0)
    sendBytesNum = sock.send("Happy.".encode())
    print("On {0}, written {1:^5,d} bytes.".format(datetime.now().strftime("%c"), sendBytesNum))

sock.close()
print("Done.")
