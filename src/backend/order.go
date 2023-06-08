package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"costomerName"`
	Total        int    `json:"total"`
	Status       string `json:"status"`
	// we're missing the connection between order and orderItem
	// so adding a slice of orderItem to the order struct
	Items []orderItem `json:"items"`
}

type orderItem struct {
	ID        int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func getOrders(db *sql.DB) ([]order, error) {
	// get all orders in JSON format
	row, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	orders := []order{}

	for row.Next() {
		var o order
		if err := row.Scan(&o.ID, &o.CustomerName, &o.Total, &o.Status); err != nil {
			return nil, err
		}

		err = o.getOrderItems(db)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, nil
}

func (o *order) getOrder(db *sql.DB) error {
	// get a specific order in JSON format
	err := db.QueryRow("SELECT customerName, total, status FROM orders WHERE id = ?", o.ID).Scan(&o.CustomerName, &o.Total, &o.Status)
	if err != nil {
		return err
	}

	err = o.getOrderItems(db) // NB! =, not :=, b/c there is err above declared
	if err != nil {
		return err
	}
	return nil
}

// Helper function to get order items
func (o *order) getOrderItems(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", o.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	orderItems := []orderItem{}

	for rows.Next() {
		var oi orderItem
		if err := rows.Scan(&oi.ID, &oi.ProductID, &oi.Quantity); err != nil {
			return err
		}
		orderItems = append(orderItems, oi)
	}

	o.Items = orderItems
	return nil
}

func (o *order) createOrder(db *sql.DB) error {
	resp, err := db.Exec("INSERT INTO orders(customerName, total, status) VALUES(?, ?, ?)", o.CustomerName, o.Total, o.Status)
	if err != nil {
		return err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return err
	}
	o.ID = int(id)

	return nil
}

func (oi *orderItem) createOrderItem(db *sql.DB) error {
	// we don't need to return the orderItem ID hence the _ because do not care about response
	_, err := db.Exec("INSERT INTO order_items(order_id, product_id, quantity) VALUES(?, ?, ?)", oi.ID, oi.ProductID, oi.Quantity)
	if err != nil {
		return err
	}

	return nil
}
