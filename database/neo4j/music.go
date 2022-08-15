package neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func HelloWorld(uri, username, password string) (string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	greeting, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n) RETURN n", nil)
		if err != nil {
			return nil, err
		}
		if result.Next() {
			fmt.Println(result.Record().Values)
			return "result.Record().Values", nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}
	return greeting.(string), nil
}
