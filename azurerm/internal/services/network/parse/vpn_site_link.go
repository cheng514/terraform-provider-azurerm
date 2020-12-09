package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type VpnSiteLinkId struct {
	SubscriptionId string
	ResourceGroup  string
	VpnSiteName    string
	Name           string
}

func NewVpnSiteLinkID(subscriptionId, resourceGroup, vpnSiteName, name string) VpnSiteLinkId {
	return VpnSiteLinkId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		VpnSiteName:    vpnSiteName,
		Name:           name,
	}
}

func (id VpnSiteLinkId) String() string {
	segments := []string{
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
		fmt.Sprintf("Vpn Site Name %q", id.VpnSiteName),
		fmt.Sprintf("Name %q", id.Name),
	}
	return strings.Join(segments, " / ")
}

func (id VpnSiteLinkId) ID(_ string) string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/vpnSites/%s/vpnSiteLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.VpnSiteName, id.Name)
}

// VpnSiteLinkID parses a VpnSiteLink ID into an VpnSiteLinkId struct
func VpnSiteLinkID(input string) (*VpnSiteLinkId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := VpnSiteLinkId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.VpnSiteName, err = id.PopSegment("vpnSites"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("vpnSiteLinks"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}