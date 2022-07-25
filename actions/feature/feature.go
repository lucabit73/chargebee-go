package feature

import (
	"fmt"
	"net/url"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/feature"
)

func List(params *feature.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/features"), params)
}
func Create(params *feature.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features"), params)
}
func Update(id string, params *feature.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/features/%v", url.PathEscape(id)), nil)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/delete", url.PathEscape(id)), nil)
}
func Activate(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/activate_command", url.PathEscape(id)), nil)
}
func Archive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/archive_command", url.PathEscape(id)), nil)
}
func Reactivate(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/reactivate_command", url.PathEscape(id)), nil)
}
