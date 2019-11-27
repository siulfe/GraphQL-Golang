package grap

import (
	DDBB "github.com/siulfe/gql/Database"
	"log"
)


type chanMessages struct{
	GroupID int
	chanel chan []*Message
}

func (m *Message)Create() error{

	err:= DDBB.GetDB().QueryRow(DDBB.CREATE_MESSAGE,m.Message,m.UserID,m.GroupID).Scan(&m.ID)

	return err
}

func GetAllMessages(group int)([]*Message,error){
	var messages []*Message
	rows,err :=  DDBB.GetDB().Query(DDBB.SELECT_MESSAGES,group)

	if err != nil{
		return nil,err
	}

	for rows.Next(){
		var message Message

		err :=rows.Scan(&message.ID,&message.Message,&message.UserID,&message.GroupID)

		if err != nil{
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages,nil
}


func EnviarMessage(chanels []chanMessages,group int) error {
	
	msg, err := GetAllMessages(group)

	if err != nil{
		return err
	}

	for i :=0;i <len(chanels);i++{
		if chanels[i].GroupID == group{
			chanels[i].chanel <- msg
		}
	}

	return nil 
}

func EnviarMessageOneChanel(chanel chan<- []*Message,group int) {
		
	mensajes,err := GetAllMessages(group)

	if err != nil {
		log.Println("Error al enviar mensaje a un canal: ",err)
		return 
	}

	chanel <- mensajes
}