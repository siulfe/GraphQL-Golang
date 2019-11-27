package grap

import (
	 DDBB "github.com/siulfe/gql/Database"
	 "errors"
)

func (d *Direccion) Create() error{

	err := DDBB.DB.QueryRow(DDBB.CREATE_DIRECCION,d.Casa).Scan(&d.ID)

	return err
}

func (d *Direccion) Delete() error{

	row,err := DDBB.GetDB().Exec(DDBB.DELETE_DIRECCION,d.ID)

	if err != nil{
		return err
	}

	count,err := row.RowsAffected()

	if count == 0{
		return errors.New("No se ha podido eliminar ninguna direccion")
	}

	return err
}

func (d *Direccion) Update() error{

	row,err := DDBB.GetDB().Exec(DDBB.UPDATE_DIRECCION,d.ID,d.Casa)

	if err != nil{
		return err
	}

	count,err := row.RowsAffected()

	if count == 0{
		return errors.New("No se ha podido actualizar ninguna direccion")
	}

	return err
}

func  GetAllDireccions() ([]*Direccion,error){
	var direcciones []*Direccion

	rows,err := DDBB.DB.Query(DDBB.SELECT_DIRECCIONS)

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	for rows.Next(){
		var d Direccion
		err :=rows.Scan(&d.ID,&d.Casa)

		if err != nil{
			continue
		}
		direcciones = append(direcciones,&d)
	}


	return direcciones,nil
}

func enviarDireccion(chanels []chan []*Direccion){

	direcciones,err := GetAllDireccions()

	if err != nil{
		return
	}

	for i := 0; i < len(chanels); i++{
		chanels[i]  <- direcciones
	}
}

func enviarDireccionOneChanel(w chan []*Direccion){

	direcciones,err := GetAllDireccions()

	if err != nil{
		return 
	}

	w <- direcciones
}