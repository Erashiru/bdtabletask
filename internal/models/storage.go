package models

import (
	"database/sql"
	"log"
)

type Product struct {
	ID   int
	Name string
}

type Shelf struct {
	ID                int
	Name              string
	IsPrimary         bool
	Additional        []string
	AdditionalShelfID int
	ProductID         int
	ProductName       string
	Quantity          int
	OrderNumber       int
}

type StorageDB struct {
	DB *sql.DB
}

func (m *StorageDB) Insert() error {
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM Products").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {

		_, err = m.DB.Exec(`
		INSERT OR IGNORE INTO Products (Name, Type) VALUES
		('Ноутбук', 'Ноутбуки'),
		('Телевизор', 'Телевизоры'),
		('Телефон', 'Телефоны'),
		('Системный блок', 'Системные блоки'),
		('Часы', 'Часы'),
		('Микрофон', 'Микрофоны');
	`)
		if err != nil {
			log.Fatal(err)
		}

		// Добавляем данные в таблицу Shelves
		_, err = m.DB.Exec(`
		INSERT OR IGNORE INTO Shelves (Name, IsPrimary) VALUES
		('А', 1),
		('А', 0),
		('Б', 1),
		('В', 0),
		('Ж', 1),
		('З', 0);	
	`)
		if err != nil {
			log.Fatal(err)
		}

		// Добавляем данные в таблицу ShelfProducts
		_, err = m.DB.Exec(`
		INSERT OR IGNORE INTO ShelfProducts (ShelfID, ProductID, IsPrimary, AdditionalShelfID) VALUES
		(1, 1, 1, 0), -- Ноутбук на Стеллаже А (главный)
		(1, 2, 1, 0), -- Телевизор на Стеллаже А (главный)
		(3, 3, 1, 6), -- Телефон на Стеллаже Б (главный)
		(5, 4, 1, 0),
		(5, 5, 1, 1),
		(5, 6, 1, 0);
	`)
		if err != nil {
			log.Fatal(err)
		}

		// Добавляем данные в таблицу Orders
		_, err = m.DB.Exec(`
		INSERT OR IGNORE INTO Orders (OrderNumber, Status) VALUES
		(10, 'в обработке'),
		(11, 'доставлен'),
		(14, 'в обработке'),
		(15, 'выполнен');
	`)
		if err != nil {
			log.Fatal(err)
		}

		// Добавляем данные в таблицу OrderProducts
		_, err = m.DB.Exec(`
		INSERT OR IGNORE INTO OrderProducts (OrderID, ProductID, Quantity) VALUES
		(1, 1, 2), -- Ноутбук заказа 10, 2 шт
		(2, 2, 3), -- TV заказа 11, 3 шт
		(3, 1, 3), -- Ноутбук заказа 14, 3 шт
		(1, 3, 1), -- Телефон заказа 10, 1 шт
		(1, 6, 1), -- Микрофон заказа 10, 1 шт
		(3, 4, 4), -- Системный блок заказа 14, 4 шт
		(4, 5, 1); -- Часы заказа 15, 1 шт 
	`)
		if err != nil {
			log.Fatal(err)
		}

	} 

	return nil
}

func (m *StorageDB) Get(order string) ([]Shelf, error) {
	stmt := `
   SELECT S.ID, S.Name, S.IsPrimary, S2.Name AS AdditionalShelf, SP.AdditionalShelfID, P.ID AS ProductID, P.Name AS ProductName, OP.Quantity, O.OrderNumber
   FROM OrderProducts OP
   JOIN Orders O ON OP.OrderID = O.ID
   JOIN Products P ON OP.ProductID = P.ID
   JOIN ShelfProducts SP ON P.ID = SP.ProductID
   JOIN Shelves S ON SP.ShelfID = S.ID
   LEFT JOIN Shelves S2 ON SP.AdditionalShelfID = S2.ID
   WHERE O.OrderNumber = ?;
`

	rows, err := m.DB.Query(stmt, order)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelves []Shelf

	for rows.Next() {
		var shelf Shelf
		var additionalShelf sql.NullString

		err := rows.Scan(&shelf.ID, &shelf.Name, &shelf.IsPrimary, &additionalShelf, &shelf.AdditionalShelfID, &shelf.ProductID, &shelf.ProductName, &shelf.Quantity, &shelf.OrderNumber)
		if err != nil {
			return nil, err
		}

		if additionalShelf.Valid {
			shelf.Additional = append(shelf.Additional, additionalShelf.String)
		}

		shelves = append(shelves, shelf)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shelves, nil
}
