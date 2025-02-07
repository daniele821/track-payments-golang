#!/bin/python3

import server
import configs

server.run("payments", "localhost", configs.FLAGS.port, configs.FLAGS.noserver, configs.FLAGS.gui)
