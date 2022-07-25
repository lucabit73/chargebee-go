package entitlementoverride

import (
	"fmt"
	"net/url"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/entitlementoverride"
)

func AddEntitlementOverrideForSubscription(id string, params *entitlementoverride.AddEntitlementOverrideForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/entitlement_overrides", url.PathEscape(id)), params)
}
func ListEntitlementOverrideForSubscription(id string, params *entitlementoverride.ListEntitlementOverrideForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/entitlement_overrides", url.PathEscape(id)), params)
}
