-- *********************************************
-- * SQL MySQL generation                      
-- *--------------------------------------------
-- * DB-MAIN version: 11.0.2              
-- * Generator date: Sep 20 2021              
-- * Generation date: Sat Feb  1 19:17:10 2025 
-- * LUN file: db/TRACK_PAYMENTS.lun 
-- * Schema: track_payments_logico/1 
-- ********************************************* 


-- Database Section
-- ________________ 

create database track_payments;
use track_payments;


-- Tables Section
-- _____________ 

create table CITY (
     name varchar(32) not null,
     constraint IDCITY primary key (name));

create table DETAIL_ORDER (
     quantity int not null,
     unit_price int not null,
    -- foreign keys:
     nameItem varchar(32) not null,
     paymentId int not null,
     constraint IDDETAIL_ORDER primary key (paymentId, nameItem));

create table ITEM (
     name varchar(32) not null,
     constraint IDITEM primary key (name));

create table PAYMENT (
     paymentId int not null auto_increment,
     date datetime not null,
     total_price int not null,
    -- foreign keys:
     city varchar(32) not null,
     shop varchar(32) not null,
     payment_method varchar(32) not null,
     constraint IDPAYMENT_ID primary key (paymentId));

create table PAYMENT_METHOD (
     method varchar(32) not null,
     constraint IDPAYMENT_METHOD primary key (method));

create table SHOP (
     name varchar(32) not null,
     constraint IDSHOP primary key (name));


-- Constraints Section
-- ___________________ 

alter table DETAIL_ORDER add constraint FKcontains
     foreign key (paymentId)
     references PAYMENT (paymentId);

alter table DETAIL_ORDER add constraint FKpart_of
     foreign key (nameItem)
     references ITEM (name);

alter table PAYMENT add constraint FKin
     foreign key (city)
     references CITY (name);

alter table PAYMENT add constraint FKinside
     foreign key (shop)
     references SHOP (name);

alter table PAYMENT add constraint FKpayed_via
     foreign key (payment_method)
     references PAYMENT_METHOD (method);


-- Index Section
-- _____________ 

