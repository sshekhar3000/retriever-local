package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"
	"encoding/json"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool

func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    databaseUrl := os.Getenv("DATABASE_URL")
    if databaseUrl == "" {
        log.Fatal("DATABASE_URL environment variable is required")
    }

    dbPool, err = pgxpool.New(context.Background(), databaseUrl)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    fmt.Println("Connected to the database successfully")
}

func GetDB() *pgxpool.Pool {
    return dbPool
}

func Close() {
    if dbPool != nil {
        dbPool.Close()
        fmt.Println("Database connection pool closed")
    }
}

func getRows(artist string) ([]Event, error) {
    query := `SELECT id, title, category, alternate_titles, labels, country, entities, start_local, end_local, timezone, location, geo, artist_name 
              FROM events WHERE artist_name = $1`
	rows, err := dbPool.Query(context.Background(), query, artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var title, category, country, startLocal, endLocal, timezone, artistName sql.NullString
		var alternateTitles, labels, entities, location, geo sql.NullString

		err := rows.Scan(&event.ID, &title, &category, &alternateTitles, &labels, &country, &entities, &startLocal, &endLocal, &timezone, &location, &geo, &artistName)
		if err != nil {
			return nil, err
		}

		event.Title = nullStringToString(title)
		event.Category = nullStringToString(category)
		event.AlternateTitles = nullStringToJSON(alternateTitles)
		event.Labels = nullStringToJSON(labels)
		event.Country = nullStringToString(country)
		event.Entities = nullStringToJSON(entities)
		event.StartLocal = nullStringToString(startLocal)
		event.EndLocal = nullStringToString(endLocal)
		event.Timezone = nullStringToString(timezone)
		event.Location = nullStringToJSON(location)
		event.Geo = nullStringToJSON(geo)
		event.ArtistName = nullStringToString(artistName)

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func getRowsByCountryDate(artist, country, start, end string) ([]Event, error) {
    query := `SELECT id, title, category, alternate_titles, labels, country, entities, start_local, end_local, timezone, location, geo, artist_name 
              FROM events 
              WHERE artist_name = $1 AND country = $2 AND start_local >= $3 AND end_local <= $4`
    rows, err := dbPool.Query(context.Background(), query, artist, country, start, end)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var events []Event
    for rows.Next() {
        var event Event
        var title, category, country, startLocal, endLocal, timezone, artistName sql.NullString
        var alternateTitles, labels, entities, location, geo sql.NullString

        err := rows.Scan(&event.ID, &title, &category, &alternateTitles, &labels, &country, &entities, &startLocal, &endLocal, &timezone, &location, &geo, &artistName)
        if err != nil {
            return nil, err
        }

        event.Title = nullStringToString(title)
        event.Category = nullStringToString(category)
        event.AlternateTitles = nullStringToJSON(alternateTitles)
        event.Labels = nullStringToJSON(labels)
        event.Country = nullStringToString(country)
        event.Entities = nullStringToJSON(entities)
        event.StartLocal = nullStringToString(startLocal)
        event.EndLocal = nullStringToString(endLocal)
        event.Timezone = nullStringToString(timezone)
        event.Location = nullStringToJSON(location)
        event.Geo = nullStringToJSON(geo)
        event.ArtistName = nullStringToString(artistName)

        events = append(events, event)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return events, nil
}
func getRowsByCountry(artist, country string)([]Event, error){
    query := `SELECT id, title, category, alternate_titles, labels, country, entities, start_local, end_local, timezone, location, geo, artist_name 
              FROM events 
              WHERE artist_name = $1 AND country = $2`
    rows, err := dbPool.Query(context.Background(), query, artist, country)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var events []Event
    for rows.Next() {
        var event Event
        var title, category, country, startLocal, endLocal, timezone, artistName sql.NullString
        var alternateTitles, labels, entities, location, geo sql.NullString

        err := rows.Scan(&event.ID, &title, &category, &alternateTitles, &labels, &country, &entities, &startLocal, &endLocal, &timezone, &location, &geo, &artistName)
        if err != nil {
            return nil, err
        }

        event.Title = nullStringToString(title)
        event.Category = nullStringToString(category)
        event.AlternateTitles = nullStringToJSON(alternateTitles)
        event.Labels = nullStringToJSON(labels)
        event.Country = nullStringToString(country)
        event.Entities = nullStringToJSON(entities)
        event.StartLocal = nullStringToString(startLocal)
        event.EndLocal = nullStringToString(endLocal)
        event.Timezone = nullStringToString(timezone)
        event.Location = nullStringToJSON(location)
        event.Geo = nullStringToJSON(geo)
        event.ArtistName = nullStringToString(artistName)

        events = append(events, event)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return events, nil
}

func nullStringToString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return "null"
}

func nullStringToJSON(ns sql.NullString) json.RawMessage {
    if ns.Valid {
        return json.RawMessage(ns.String)
    }
    return json.RawMessage("{}")
}