package purchase

import (
	"fmt"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/purchase"
)

func Create(params *purchase.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases"), params)
}
func Estimate(params *purchase.EstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/purchases/estimate"), params)
}
