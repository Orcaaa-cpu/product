package helper

import "database/sql"

func MigrateUsers(db *sql.DB) {
	sql := `
    CREATE TABLE users (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		username varchar(255) NOT NULL,
		password varchar(255) NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;
    `

	_, err := db.Exec(sql)
	PanicError(err)
}

func MigrateProduct(db *sql.DB) {
	sql := `
    CREATE TABLE items (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		username varchar(255) NOT NULL,
		password varchar(255) NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;
    `

	_, err := db.Exec(sql)
	PanicError(err)

}
