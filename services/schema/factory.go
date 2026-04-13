package schema

import "fmt"

func GetSchemaProvider(driver string) (SchemaProvider, error) {
	switch driver {
	case "postgres":
		return &PostgresProvider{}, nil
	case "mysql":
		return &MySQLProvider{}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}
}
