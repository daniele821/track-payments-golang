#!/bin/python3

import server
import database
import utils
import configs

FLAGS = configs.FLAGS

server_address = server.run_server()
print("started server on " + server_address)
if FLAGS.gui:
    utils.open_link(server_address)
    print("launching browser to connect to the server...")
else:
    print("SKIP: launching browser")
