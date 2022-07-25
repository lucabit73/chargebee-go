package card

import (
	"fmt"
	"net/url"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/card"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/cards/%v", url.PathEscape(id)), nil)
}
func UpdateCardForCustomer(id string, params *card.UpdateCardForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/credit_card", url.PathEscape(id)), params)
}
func SwitchGatewayForCustomer(id string, params *card.SwitchGatewayForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/switch_gateway", url.PathEscape(id)), params)
}
func CopyCardForCustomer(id string, params *card.CopyCardForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/copy_card", url.PathEscape(id)), params)
}
func DeleteCardForCustomer(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_card", url.PathEscape(id)), nil)
}
