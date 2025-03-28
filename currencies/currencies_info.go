package currencies

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/eosorio/personalf-services/databaseInfo"
)

// GetCurrencies returns info from all the records in the database table currencies. If currencyID is 0 it will return all the records
func GetCurrencies(currencyID int64, currencyType int64) ([]CurrenciesRecordSet, error) {
	dbConnectInfo := databaseInfo.GetConfigFromEnv()

	connectString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		dbConnectInfo.Hostname, dbConnectInfo.Name, dbConnectInfo.User, dbConnectInfo.Pass)
	query := "SELECT id, code, name, 0 AS type_id FROM currencies "
	if currencyID != 0 || currencyType != 0 {
		query = query + "WHERE "
	}
	if currencyID != 0 {
		query = fmt.Sprintf("%s id=%d", query, currencyID)
		if currencyType != 0 {
			query = query + " AND "
		}
	}
	if currencyType != 0 {
		query = fmt.Sprintf("%s type_id=%d", query, currencyType)
	}

	query = fmt.Sprintf("%s ORDER BY name ASC", query)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query(query)
	log.Printf("query: #%s#\n", query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	currenciesList := make([]CurrenciesRecordSet, 0)

	for rows.Next() {
		var currencyItem CurrenciesRecordSet
		if err := rows.Scan(&currencyItem.ID,
			&currencyItem.Code,
			&currencyItem.Name,
			&currencyItem.TypeID); err != nil {
			log.Fatal(err)
			return nil, err
		}
		currenciesList = append(currenciesList, currencyItem)
	}

	return currenciesList, nil
}
