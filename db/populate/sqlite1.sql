PRAGMA foreign_keys=ON;
BEGIN TRANSACTION;
CREATE TABLE CITY (
     name           TEXT    not null,
     constraint IDCITY primary key (name)
);
INSERT INTO CITY VALUES('Cesena');
CREATE TABLE ITEM (
     name           TEXT    not null,
     constraint IDITEM primary key (name)
);
INSERT INTO ITEM VALUES('Pane');
CREATE TABLE PAYMENT_METHOD (
     method         TEXT    not null,
     constraint IDPAYMENT_METHOD primary key (method)
);
INSERT INTO PAYMENT_METHOD VALUES('Contante');
CREATE TABLE SHOP (
     name           TEXT    not null,
     constraint IDSHOP primary key (name)
);
INSERT INTO SHOP VALUES('Coop');
CREATE TABLE PAYMENT (
     paymentId      INTEGER not null,
     date           TEXT    not null,
     total_price    INTEGER not null,
     city           TEXT    not null,
     shop           TEXT    not null,
     payment_method TEXT    not null,
     constraint IDPAYMENT primary key (paymentId),
     foreign key (city) references CITY(name),
     foreign key (shop) references SHOP(name),
     foreign key (payment_method) references PAYMENT_METHOD(method)
);
INSERT INTO PAYMENT VALUES(1,'2025-02-02 03:27:42',1000,'Cesena','Coop','Contante');
INSERT INTO PAYMENT VALUES(2,'2025-02-02 03:27:57',1000,'Cesena','Coop','Contante');
INSERT INTO PAYMENT VALUES(3,'2025-02-02 03:28:05',130,'Cesena','Coop','Contante');
CREATE TABLE DETAIL_ORDER (
     nameItem       TEXT    not null,
     paymentId      INTEGER not null,
     quantity       INTEGER not null,
     unit_price     INTEGER not null,
     constraint IDDETAIL_ORDER primary key (paymentId, nameItem),
     foreign key (paymentId) references PAYMENT(paymentId),
     foreign key (nameItem) references ITEM(name)
);
INSERT INTO DETAIL_ORDER VALUES('Pane',2,12,100);
INSERT INTO DETAIL_ORDER VALUES('Pane',3,12,100);
COMMIT;

