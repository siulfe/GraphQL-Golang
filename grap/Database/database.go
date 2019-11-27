package database

import (
	"log"
	"database/sql"
	"os"
	"github.com/siulfe/gql/crypto"
	_ "github.com/lib/pq"
)

var DB *sql.DB


func ConnectDatabase(resp map[string]string) error{

	db, err := sql.Open("postgres",resp["url"])

	if err != nil{
		return err
	}

	DB = db

	err = crearBD(resp)

	return err
}

func GetDB() *sql.DB{
	return DB
}

func crearBD(resp map[string]string) error{

	_,err :=DB.Exec(VERIFICAR)

	if err == nil{

		DB.Exec(RESET_STATUS)

		return nil
	}
 	
	DB.Close()

	err = abrirConexion(resp["url-create-DDBB"])

	if err != nil{
		log.Println("Error al abrir la conexion para crear la base de datos", err)
		return err
	}

 	_,err = DB.Exec(CREATE_DATABASE)
	
	if err != nil{
		log.Println("Error al crear base de datos",err)
		return err
	}

	
	DB.Close()

	err = abrirConexion(os.Getenv("DB_URL"))

	if err != nil{
		log.Println("Error al abrir la conexion a la base de datos del chat", err)
		return err
	}

	_,err = DB.Exec(CREATE_DIRECCION_TABLE)

	if err != nil{
		log.Println("Error al crear table direccion",err)
		return err
	}
	_,err = DB.Exec(CREATE_USERG_TABLE)
	
	if err != nil{
		log.Println("Error al crear table userg",err)
		return err
	}

	_,err = DB.Exec(CREATE_GROUP_TABLE)
	
	if err != nil{
		log.Println("Error al crear table group",err)
		return err
	}

	_,err = DB.Exec(CREATE_MESSAGE_TABLE)

	if err != nil{
		log.Println("Error al crear table message ",err)
		return err
	}


	_,err = DB.Exec(FIRST_DIRECCION)

	if err != nil{
		log.Println("Error al crear direccion admin ",err)
		return err
	}


	password,err := crypto.HashPassword("adminChatG")

	if err != nil{
		log.Println("Error al crear password para usuario admin ",err)
		return err
	}

	_,err = DB.Exec(FIRST_USER,password)

	if err != nil{
		log.Println("Error al crear usuario admin ",err)
		return err
	}

	return nil
}


func abrirConexion(url string) error{
	db, err := sql.Open("postgres",url)

	if err != nil{
		log.Println("Error al abrir la conexion", err)
		return err
	}

	DB = db

	return nil
}