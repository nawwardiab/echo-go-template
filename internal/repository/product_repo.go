package repository

import (
	"echo-server/internal/model"
	"fmt"

	"github.com/jackc/pgx"
)

type ProductRepo struct {
	db *pgx.Conn
}

func NewProductRepo(db *pgx.Conn) *ProductRepo {
	return &ProductRepo{db: db}
}

// GeAllProducts – queries db products table and returns products details
func (r *ProductRepo) GetAllProducts() ([]model.Product, error) {
  rows, queryErr := r.db.Query(`SELECT * FROM products`)
  if queryErr != nil {
    return nil, fmt.Errorf("postgres: ListAllProducts products: %w", queryErr)
  }
  defer rows.Close()

  // slice to hold data from returned rows
  var list []model.Product

  // Loop through rows, using Scan to assign column data to struct fields.
  for rows.Next() {
    var p model.Product
    scanErr := rows.Scan(
      &p.ID,
      &p.Title,
      &p.Year,
      &p.Artist,
      &p.Img,
      &p.Price,
      &p.Genre,
      )
    if scanErr != nil {
      return nil, fmt.Errorf("postgres: scan product: %w", scanErr)
    } else {
      list = append(list, p)
    }
  }
  return list, rows.Err()
}

// GetProductDetails – queries db products table and returns queried product's details
func (r *ProductRepo) GetProductDetails(id int) (*model.Product, error) {
  query := `SELECT * FROM products WHERE id=$1`
  row := r.db.QueryRow(query, id)

  var p model.Product
  scanErr := row.Scan(&p.ID, &p.Title, &p.Year, &p.Artist, &p.Img, &p.Price, &p.Genre,)
  
  if scanErr != nil {
    return nil, fmt.Errorf("postgres: scan product %d: %w", id, scanErr)
  } else {
    return &p, nil
  } 
}