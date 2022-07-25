package portalsession

import (
	"fmt"
	"net/url"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/portalsession"
)

func Create(params *portalsession.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/portal_sessions/%v", url.PathEscape(id)), nil)
}
func Logout(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/logout", url.PathEscape(id)), nil)
}
func Activate(id string, params *portalsession.ActivateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/activate", url.PathEscape(id)), params)
}
