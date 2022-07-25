package resourcemigration

import (
	"fmt"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/resourcemigration"
)

func RetrieveLatest(params *resourcemigration.RetrieveLatestRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/resource_migrations/retrieve_latest"), params)
}
