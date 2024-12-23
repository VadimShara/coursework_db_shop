package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/VadimShara/coursework_DB_shop/internal/config"
	"github.com/VadimShara/coursework_DB_shop/internal/models"
)

var (
	ErrSessionNotFound = errors.New("user session not found")
)

func PGConnectionStr(config config.PostgresConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
}

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(config config.PostgresConfig) *DB {
	dbUrl := PGConnectionStr(config)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatal("unable to parse db url")
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal("unable to connect to db")
	}

	return &DB{Pool: pool}
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) AddProduct(product models.Product) error {
	const query = `INSERT INTO products (name_product, product_init, purchase_price, selling_price) VALUES ($1, $2, $3, $4)`

	_, err := db.Pool.Exec(context.Background(), query, product.Name_product, product.Product_init, product.Purchase_price, product.Selling_price)
	if err != nil {
		fmt.Errorf("failed to add product")
	}

	return nil
}

func (db *DB) DeleteProduct(id int) error {
	const query = `DELETE FROM products WHERE id = $1`

	_, err := db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Errorf("failed to delete product")
	}

	return nil
}

func (db *DB) ListProducts() ([]models.Product, error) {
	const query = `SELECT * FROM products`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to list products")
	}

	products := make([]models.Product, 0)

	for rows.Next() {
		var p models.Product 

		if err := rows.Scan(&p); err != nil {
			return nil, fmt.Errorf("failed to list products")
		}

		products = append(products, p)
	}

	return products, nil
}

func (db *DB) SearchProduct(name string) ([]models.Product, error) {
	const query = `SELECT * FROM products WHERE name_product = $1`

	rows, err := db.Pool.Query(context.Background(), query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to list products")
	}

	products := make([]models.Product, 0)

	for rows.Next() {
		var p models.Product 

		if err := rows.Scan(&p); err != nil {
			return nil, fmt.Errorf("failed to list products")
		}

		products = append(products, p)
	}

	return products, nil
}

func (db *DB) AddSeller(seller models.Seller) error {
	const query = `INSERT INTO sellers (name_seller, commission_rate) VALUES ($1, $2)`

	_, err := db.Pool.Exec(context.Background(), query, seller.Name_seller, seller.Commission_rate)
	if err != nil {
		return fmt.Errorf("failed to add seller")
	}

	return nil
}

func (db *DB) DeleteSeller(id int) error {
	const query = `DELETE FROM sellers WHERE id = $1`

	_, err := db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete seller")
	}

	return nil
}

func (db *DB) ListSellers() ([]models.Seller, error) {
	const query = `SELECT * FROM sellers`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to list sellers")
	}

	sellers := make([]models.Seller, 0)

	for rows.Next() {
		var s models.Seller

		if err := rows.Scan(&s); err != nil {
			return nil, fmt.Errorf("failed to list sellers")
		}

		sellers = append(sellers, s)
	}

	return sellers, nil
}

func (db *DB) SearchSeller(name string) ([]models.Seller, error) {
	const query = `SELECT * FROM sellers WHERE name_seller = $1`

	rows, err := db.Pool.Query(context.Background(), query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to list sellers")
	}

	sellers := make([]models.Seller, 0)

	for rows.Next() {
		var s models.Seller

		if err := rows.Scan(&s); err != nil {
			return nil, fmt.Errorf("failed to list sellers")
		}

		sellers = append(sellers, s)
	}

	return sellers, nil
}

func (db *DB) AddSale(sale models.Sale) error {
	const query = `INSERT INTO sales (product_id, seller_id, number, saleDateTime) VALUES ($1, $2, $3, $4)`

	_, err := db.Pool.Exec(context.Background(), query, sale.Product_id, sale.Seller_id, sale.Number, time.Now())
	if err != nil {
		return fmt.Errorf("failed to add sale")
	}

	return nil
}

func (db *DB) DeleteSale(id int) error {
	const query = `DELETE FROM sale WHERE id = $1`

	_, err := db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete sale")
	}

	return nil
}

func (db *DB) ListSales() ([]models.Seller, error) {
	const query = `SELECT * FROM sales`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to list sales")
	}

	sales := make([]models.Sale, 0)

	for rows.Next() {
		var s models.Sale

		if err := rows.Scan(&s); err != nil {
			return nil, fmt.Errorf("failed to list sales")
		}

		sales = append(sales, s)
	}

	return sales, nil
}

func (db *DB) SearchSale(name string) ([]models.Seller, error) {
	const query = `SELECT * FROM sales LEFT JOIN products WHERE name_product = $1`

	rows, err := db.Pool.Query(context.Background(), query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to list sales")
	}

	sales := make([]models.Seller, 0)

	for rows.Next() {
		var s models.Sale

		if err := rows.Scan(&s); err != nil {
			return nil, fmt.Errorf("failed to list sales")
		}

		sales = append(sales, s)
	}

	return sales, nil
}