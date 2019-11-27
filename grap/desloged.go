package grap

import (
	"log"
	"time"
	DDBB "github.com/siulfe/gql/Database"
)


func DesLoged(){

	for{
		time.Sleep(5 * time.Minute)

		for x := range tokens {
			if tokens[x] < int(time.Now().Unix()){

				id := UserIDFromToken(x)

				_,err := DDBB.GetDB().Exec(DDBB.DESLOGED, id)

				if err != nil{
					log.Println("A ocurrido un error al deslogear un usuario. ID: ", id, "\n error: ",err)
				}else{
					delete(tokens,x)	
				}
			}
		}

	}
}