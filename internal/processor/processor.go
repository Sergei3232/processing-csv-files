package processor

import "github.com/Sergei3232/processing-csv-files/internal/app/queue/postgres"

type csvProcessor struct {
	ImageLoader *postgres.PostgresSQL
	FileStorage *postgres.PostgresSQL
	ListFile    []string
}
