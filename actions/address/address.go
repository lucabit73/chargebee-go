package address

import (
	"fmt"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/address"
)

func Retrieve(params *address.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/addresses"), params)
}
func Update(params *address.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/addresses"), params)
}
