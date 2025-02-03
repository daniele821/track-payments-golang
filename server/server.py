#!/bin/python3

import sys
import os

def open_link(link):
    if sys.platform == "linux":
        if "TERMUX_VERSION" in os.environ:
            os.system(f"termux-open {link}")
        else:
            os.system(f"xdg-open {link}")
    elif sys.platform == "win32":
        os.system(f"start {link}")
    elif sys.platform == "darwin":
        os.system(f"open {link}")
    else:
        raise Exception("unable to open link: unknown platform!")


        
