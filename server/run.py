#!/bin/python3

import server
import database
import utils
import configs

if not configs.FLAGS.noserver:
    server_address = server.run_server()
    print("started server on " + server_address)
    if configs.FLAGS.gui:
        utils.open_link(server_address)
        print("launching browser to connect to the server...")
    else:
        print("not launching browser!")
