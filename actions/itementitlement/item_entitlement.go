package itementitlement

import (
	"fmt"
	"net/url"

	"github.com/lucabit73/chargebee-go"
	"github.com/lucabit73/chargebee-go/models/itementitlement"
)

func ItemEntitlementsForItem(id string, params *itementitlement.ItemEntitlementsForItemRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id)), params)
}
func ItemEntitlementsForFeature(id string, params *itementitlement.ItemEntitlementsForFeatureRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id)), params)
}
func AddItemEntitlements(id string, params *itementitlement.AddItemEntitlementsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id)), params)
}
func UpsertOrRemoveItemEntitlementsForItem(id string, params *itementitlement.UpsertOrRemoveItemEntitlementsForItemRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id)), params)
}
