package dbutils

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectDB() (neo4j.DriverWithContext, context.Context, error) {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", "")) // TODO use env variables
	if err != nil {
		fmt.Println(err)
		return nil, ctx, err
	}

	ctx = context.WithValue(ctx, "driver", driver)
	return driver, ctx, nil
}

//func CloseDB(driver neo4j.DriverWithContext, ctx context.Context) {
//	err := driver.Close(ctx)
//	if err != nil {
//		return
//	}
//}
