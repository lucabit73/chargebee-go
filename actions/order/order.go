package order

import (
	"fmt"
	"net/url"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/order"
)

func Create(params *order.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders"), params)
}
func Update(id string, params *order.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v", url.PathEscape(id)), params)
}
func ImportOrder(params *order.ImportOrderRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/import_order"), params)
}
func AssignOrderNumber(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/assign_order_number", url.PathEscape(id)), nil)
}
func Cancel(id string, params *order.CancelRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/cancel", url.PathEscape(id)), params)
}
func CreateRefundableCreditNote(id string, params *order.CreateRefundableCreditNoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/create_refundable_credit_note", url.PathEscape(id)), params)
}
func Reopen(id string, params *order.ReopenRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/reopen", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/orders/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/delete", url.PathEscape(id)), nil)
}
func List(params *order.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/orders"), params)
}
func OrdersForInvoice(id string, params *order.OrdersForInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/orders", url.PathEscape(id)), params)
}
func Resend(id string, params *order.ResendRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/orders/%v/resend", url.PathEscape(id)), params)
}
