package infraestructure

import (
	"API-HEX-GO/src/Products/domain"
	"API-HEX-GO/src/core"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.Product) error {
	query := "INSERT INTO Product (Nombre, Precio) VALUES (?, ?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Precio)
	return err
}

func (r *MySQLRepository) Delete(p string)error{
	nombre :=p
	query := "DELETE FROM Product WHERE Nombre = ?"
	_,err :=r.conn.DB.Exec(query,nombre)
	return err
}

func (r *MySQLRepository) Update(id int,p *domain.Product)error{
	query := "UPDATE Product SET nombre = ?, precio = ? WHERE id = ?"
    _, err := r.conn.DB.Exec(query, p.Nombre, p.Precio,id)
    if err != nil {
        return err
    }
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Product, error) {
	query := "SELECT id, nombre, precio FROM Product"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Precio); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}


func (r *MySQLRepository) GetById(id int) (*domain.Product, error) {
	query := "SELECT nombre, precio FROM Product WHERE id = ?"
	row := r.conn.DB.QueryRow(query, id)

	var product domain.Product
	if err := row.Scan(&product.Nombre, &product.Precio); err != nil {
		return nil, err
	}

	return &product, nil
}


