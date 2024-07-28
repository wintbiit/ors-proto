package proto

/*
public enum e_Role_Location
{
  ecl_null,
  ecl_Gate,
  ecl_Hall,
  ecl_Room,
  ecl_Match,
  ecl_max,
}

*/

type e_Role_Location int32

const (
	e_Role_Location_ecl_null  e_Role_Location = 0
	e_Role_Location_ecl_Gate  e_Role_Location = 1
	e_Role_Location_ecl_Hall  e_Role_Location = 2
	e_Role_Location_ecl_Room  e_Role_Location = 3
	e_Role_Location_ecl_Match e_Role_Location = 4
	e_Role_Location_ecl_max   e_Role_Location = 5
)
