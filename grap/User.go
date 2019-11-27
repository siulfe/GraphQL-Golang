package grap

import (
	DDBB "github.com/siulfe/gql/Database"
	"github.com/siulfe/gql/crypto"
	"errors"
)


func (u *User) Create() error{

	err := DDBB.GetDB().QueryRow(DDBB.CREATE_USER, u.Name,u.LastName,u.Identification,u.Age,u.Direccion.ID,u.Password,u.Rol).Scan(&u.ID)

	return err
}

func (u *User) Delete() error{
	row,err := DDBB.GetDB().Exec(DDBB.DELETE_USER,u.ID)

	if err != nil{
		return err
	}

	count,err := row.RowsAffected()

	if count == 0{
		return errors.New("No se ha eliminado ningun usuario")
	}

	return err
}

func (u *User) Update() error{

	row,err := DDBB.GetDB().Exec(DDBB.UPDATE_USER,u.ID,u.Name,u.LastName,u.Identification,u.Age)

	if err != nil{
		return err
	}

	count, err:= row.RowsAffected()

	if count == 0{
		return errors.New("No se ha actualizado ningun usuario")
	}

	return err
}

func GetAllUsers() ([]*User,error){
	var users []*User
	rows, err := DDBB.GetDB().Query(DDBB.SELECT_USERS)

	if err != nil{
		return nil,err
	}

	for rows.Next(){
		var user User = User{
			Direccion: &Direccion{},
		}
		
		err :=rows.Scan(&user.ID,&user.Name,&user.LastName,&user.Identification,&user.Age,
							&user.CreateAt,&user.UpdateAt,&user.Password,&user.Direccion.ID,&user.Direccion.Casa)

		if err != nil{
			return nil, err
		}

		users = append(users, &user)

	}

	return users,nil
}


func (u *User) Login() error{
	var user User

	err := DDBB.GetDB().QueryRow(DDBB.SELECT_USER_NAME, u.Name).Scan(&user.Password,&u.Rol,&u.ID)

	if err != nil{
		return err
	}

	if !crypto.ComparePassword(*u.Password, *user.Password){
		return errors.New("Usuario invalido")
	}

	err = loged(u.ID)

	return err
}


func loged(w int) error{
	_,err := DDBB.GetDB().Exec(DDBB.LOGGED,w)
	return err
}

func logoutUser(w int) error{
	_,err := DDBB.GetDB().Exec(DDBB.DESLOGED,w)
	return err
}