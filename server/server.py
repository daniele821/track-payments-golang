#!/bin/python3

from database import Db

db = Db("test")
print(db.getCity())
print(db.getItem())
print(db.getShop())
print(db.getMethod())
print(db.getPayment())
print(db.getDetail())
print()
for line in db.getFullDetail():
    print(line)
