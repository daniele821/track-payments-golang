#!/bin/python3

from database import Db

db = Db("test")
try:
    db.insertCity("Asti")
    db.insertCity("Cesena")
    db.insertItem("Pane")
    db.insertItem("Briosche")
    db.insertShop("Coop")
    db.insertShop("Conad")
    db.insertShop("Paninaro")
    db.insertMethod("Contante")
    db.insertMethod("Postepay")
    db.insertMethod("SanPaolo")
    db.insertPayment("2025-01-01 12:34:32", "Asti", "Coop", "Contante")
    db.insertDetail("Pane", 1, 3, 100)
    db.insertDetail("Briosche", 1, 3, 100)
except Exception as e:
    print(f'\x1b[1;31mFAILURE ---> {e}\x1b[m')
print(db.deletePayment(1))
print(db.getCity())
print(db.getItem())
print(db.getShop())
print(db.getMethod())
print(db.getPayment())
print(db.getDetail())
print()
for line in db.getFullDetail():
    print(line)
