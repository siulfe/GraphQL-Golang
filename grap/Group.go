package grap

import (
	DDBB "github.com/siulfe/gql/Database"
	"errors"
)
func (m *Group)Create() error{

	err :=DDBB.GetDB().QueryRow(DDBB.CREATE_GROUP,m.UserID,m.Title).Scan(&m.ID)

	return err
}

func (m *Group)Delete() error{

	rows, err := DDBB.GetDB().Exec(DDBB.DELETE_GROUP,m.ID)

	if err != nil{
		return err
	}

	count,err := rows.RowsAffected()

	if count == 0{
		return errors.New("No se ha eliminado ningun registro")
	}

	return err
}

func (m *Group)Update() error{
	rows, err := DDBB.GetDB().Exec(DDBB.UPDATE_GROUP,m.ID,m.Title)

	if err != nil{
		return err
	}

	count,err := rows.RowsAffected()

	if count == 0{
		return errors.New("No se ha actualizado ningun registro")
	}

	return err
}

func GetAllGroups()([]*Group,error){
	var groups []*Group

	rows, err := DDBB.GetDB().Query(DDBB.SELECT_GROUPS)

	if err != nil{
		return nil,err
	}

	for rows.Next(){
		var group Group

		err := rows.Scan(&group.ID,&group.Title,&group.UserID,&group.CreateAt,&group.CloseAt)

		if err != nil{
			return nil, err
		}

		groups = append(groups,&group)
	}

	return groups,nil
}