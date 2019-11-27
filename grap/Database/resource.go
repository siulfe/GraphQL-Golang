package database

import (
	"strings"
	"bufio"
	"os"
)

var VERIFICAR string ="select 1 from userg"

var RESET_STATUS string ="update userg set status_logger = false"

var LOGGED string = "update userg set status_logger = true where id = $1"

var DESLOGED string ="update userg set status_logger = false where id = $1"

var CREATE_USER string = `insert into userg (name,lastname,identification,age,direccion_id,password,rol,createat) 
					values($1,$2,$3,$4,$5,$6,$7,now()) returning id`

var DELETE_USER string = "delete from userg where id = $1"

var UPDATE_USER string ="update userg set name = $2, lastname = $3, identification= $4, age =$5, updateAt = now() where id= $1"

var SELECT_USERS string = `select userg.id,userg.name,userg.lastname,userg.identification,userg.age,
							userg.createAt,userg.updateAt,userg.password 
							,direccion.id, direccion.casa from userg
							left join direccion on direccion.id = userg.direccion_id
							where rol != 'admin'`

var SELECT_USER_NAME string="select password,rol,id from userg where name =$1"

var SELECT_USER_NAME_ROL string ="select 1 from userg where name =$1 and rol =$2 and status_logger = true"

var CREATE_DIRECCION string ="insert into direccion (casa) values($1) returning id"

var DELETE_DIRECCION string ="delete from direccion where id = $1"

var UPDATE_DIRECCION string = "update direccion set casa = $2 where id = $1"

var SELECT_DIRECCIONS string ="select * from direccion where casa != ''"

//********************************************************************************************************************************

var SELECT_MESSAGES string ="select id,message,user_id,group_id from message where group_id = $1"

var CREATE_MESSAGE string ="insert into message (message,user_id,group_id,createat) values($1,$2,$3,now()) returning id"

var SELECT_GROUPS string ="select id,title,user_id,createat,updateat from groups"

var CREATE_GROUP string ="insert into groups (user_id,title,createat) values ($1,$2,now()) returning id"

var DELETE_GROUP string ="delete from groups where id=$1"

var UPDATE_GROUP string ="update groups set title = $2 where id = $1"



//***********************************************************************************************************************************

var CREATE_DATABASE string ="create database chatg"

var CREATE_USERG_TABLE string = `create table if not exists userg(
							id serial not null,
							name varchar(255) not null unique,
							rol varchar(255) not null,
							lastname varchar(255),
							password varchar(255) not null,
							createat timestamp not null default now(),
							updateat timestamp,
							age int,
							identification varchar(255) not null,
							direccion_id integer not null,
							status_logger bool not null default false,
							primary key(id),
							foreign key(direccion_id) references direccion(id)
						)`
var CREATE_DIRECCION_TABLE string = `create table if not exists direccion(
							id serial not null,
							casa varchar(255) not null unique,
							primary key(id)
						)`
var CREATE_GROUP_TABLE string = `create table if not exists groups(
							id serial not null,
							user_id integer not null,
							createat timestamp not null,
							updateat timestamp,
							title varchar(255) not null,
							primary key(id),
							foreign key(user_id) references userg(id)
						)`
var CREATE_MESSAGE_TABLE string = `create table if not exists message(
							id serial not null,
							user_id integer not null,
							group_id integer not null,
							createat timestamp not null,
							message varchar(255) not null,
							primary key(id),
							foreign key(user_id) references userg(id),
							foreign key(group_id) references groups(id)
						)`

//***********************************************************************************************************************************
var FIRST_DIRECCION string = "insert into direccion (casa) values('')"

var FIRST_USER string = "insert into userg (name,identification,rol,password,direccion_id) values('admin','00000000','admin',$1,1)"

//var PASSWORD_FIRST_USER string ="adminChatG"


func LeerArchivo() (map[string]string, error){
	resp := make(map[string]string)
	archivo,err := os.Open("config.txt")
	if err != nil{
		return nil,err
	}

	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan(){
		fila := strings.Split(scanner.Text(),":")

		if fila[0][0] == '#'{
			continue
		}

		resp[fila[0]] = fila[1]
	}

	return resp,nil
}

