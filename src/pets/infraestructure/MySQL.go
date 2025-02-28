package infraestructure

import (
	"API-HEX-GO/src/pets/domain"
	"API-HEX-GO/src/core"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.Pet) error {
	query := "INSERT INTO Pet (Nombre, Edad) VALUES (?, ?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Edad)
	return err
}

func (r *MySQLRepository) Delete(nombre string) error {
	query := "DELETE FROM Pet WHERE Nombre = ?"
	_, err := r.conn.DB.Exec(query, nombre)
	return err
}

func (r *MySQLRepository) Update(nombre string, p *domain.Pet) error {
	query := "UPDATE Pet SET Nombre = ?, Edad = ? WHERE nombre = ?"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Edad, nombre)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) GetAll() ([]domain.Pet, error) {
	query := "SELECT Nombre, Edad FROM Pet"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []domain.Pet
	for rows.Next() {
		var pet domain.Pet
		if err := rows.Scan(&pet.Nombre, &pet.Edad); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return pets, nil
}
