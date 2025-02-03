#!/bin/python3

import server
import database
import utils

server_address = server.run_server()
print("started server on " + server_address)
utils.open_link(server_address)
print("launching browser to connect to the server...")
