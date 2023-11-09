package Section_7___Golang_Database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('Masil', 'Masil')"
	//script := "UPDATE customer SET name = 'Chahyo' WHERE id = 'aji')"
	//script := "DELETE FROM customer WHERE id = 'aji';)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}
