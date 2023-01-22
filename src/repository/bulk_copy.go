package repository

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"packform-test/config"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func BulkCopy(table string, columns []string, reader *io.Reader, batchSize int) int64 {
	lines := 0
	loc, _ := time.LoadLocation("Melbourne/Australia")
	copyTotal := int64(0)
	batch := make([][]string, 0, batchSize)
	csvReader := csv.NewReader(*reader)

	database_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.Close(context.Background())

	// Read line at a time
	csvReader.Read()

	for {
		buffer, err := csvReader.Read()
		// Cap the buffer based on the actual length read
		if err != nil {
			// When the error is EOF, there are no longer any bytes to read
			// meaning the request is completed
			if err == io.EOF {
				break
			}

			// If the error is an unexpected EOF, the requested size to read
			// was larger than what was available. This is not an issue for
			// as long as the length returned above is used, or the buffer
			// is capped as it is above. Any error that is not an unexpected
			// EOF is an actual error, which we handle accordingly
			if err != io.ErrUnexpectedEOF {
				panic(err)
			}
		} else {
			batch = append(batch, buffer)
			lines += 1
			if lines >= batchSize {
				// You may now use the batch to bulk insert into db
				copyCount, _ := conn.CopyFrom(
					context.Background(),
					pgx.Identifier{table},
					append(columns, "created_at", "updated_at"),
					pgx.CopyFromSlice(len(batch), func(i int) ([]interface{}, error) {
						old := batch[i]
						new := make([]interface{}, len(old))
						for k, v := range old {
							if v == "" {
								new[k] = nil
							} else if strings.HasPrefix(v, "{") {
								v = strings.TrimLeft(v, "{")
								v = strings.TrimRight(v, "}")
								new[k] = strings.Split(v, ",")
							} else {
								new_val, err := time.ParseInLocation(time.RFC3339, v, loc)
								if err == nil {
									new[k] = new_val
								} else {
									new[k] = v
								}
							}
						}
						return append(new, time.Now(), time.Now()), nil
					}),
				)
				copyTotal += copyCount
				lines = 0
				batch = batch[:0] // Keep initialized values and make length 0 so no performance hit of emptying all elements
			}
		}
	}
	if len(batch) > 0 {
		copyCount, err := conn.CopyFrom(
			context.Background(),
			pgx.Identifier{table},
			append(columns, "created_at", "updated_at"),
			pgx.CopyFromSlice(len(batch), func(i int) ([]interface{}, error) {
				old := batch[i]
				new := make([]interface{}, len(old))
				for k, v := range old {
					if v == "" {
						new[k] = nil
					} else if strings.HasPrefix(v, "{") {
						v = strings.TrimLeft(v, "{")
						v = strings.TrimRight(v, "}")
						new[k] = strings.Split(v, ",")
					} else {
						new_val, err := time.ParseInLocation(time.RFC3339, v, loc)
						if err == nil {
							new[k] = new_val
						} else {
							new[k] = v
						}
					}
				}
				return append(new, time.Now(), time.Now()), nil
			}),
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to copy from csv: %v\n", err)
			panic(err)
		}
		batch = batch[:0]
		copyTotal += copyCount
	}
	return copyTotal
}
