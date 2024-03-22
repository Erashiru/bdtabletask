CREATE TABLE IF NOT EXISTS Products (
    ID INTEGER PRIMARY KEY,
    Name TEXT,
    Type TEXT
);

CREATE TABLE IF NOT EXISTS Shelves (
    ID INTEGER PRIMARY KEY,
    Name TEXT,
    IsPrimary INTEGER -- 1 для главного стеллажа, 0 для второстепенного
);

CREATE TABLE  IF NOT EXISTS ShelfProducts (
    ID INTEGER PRIMARY KEY,
    ShelfID INTEGER,
    ProductID INTEGER,
    IsPrimary INTEGER,
    AdditionalShelfID INTEGER,
    FOREIGN KEY (ShelfID) REFERENCES Shelves(ID),
    FOREIGN KEY (ProductID) REFERENCES Products(ID)
);

CREATE TABLE IF NOT EXISTS Orders (
    ID INTEGER PRIMARY KEY,
    OrderNumber INTEGER,
    Status TEXT
);

CREATE TABLE IF NOT EXISTS OrderProducts (
    ID INTEGER PRIMARY KEY,
    OrderID INTEGER,
    ProductID INTEGER,
    Quantity INTEGER,
    FOREIGN KEY (OrderID) REFERENCES Orders(ID),
    FOREIGN KEY (ProductID) REFERENCES Products(ID)
);
