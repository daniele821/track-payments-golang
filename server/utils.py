#!/bin/python3

import sys
import os
import subprocess


def open_link(link):
    if sys.platform == "linux":
        if "TERMUX_VERSION" in os.environ:
            subprocess.run(["termux-open", link])
        else:
            subprocess.run(["xdg-open", link])
    elif sys.platform == "win32":
        # `shell=True` is required on Windows!
        subprocess.run(["start", link], shell=True)
    elif sys.platform == "darwin":
        subprocess.run(["open", link])
    else:
        raise Exception("unable to open link: unknown platform!")
