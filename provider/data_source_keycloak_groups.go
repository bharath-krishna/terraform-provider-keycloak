package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func dataSourceKeycloakGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKeycloakGroupsRead,
		Schema: map[string]*schema.Schema{
			"realm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceKeycloakGroupsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	keycloakClient := meta.(*keycloak.KeycloakClient)

	realmId := data.Get("realm_id").(string)
	fmt.Println("group details are here:")
	fmt.Println(realmId)

	group_list, err := keycloakClient.GetGroups(ctx, realmId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(realmId)

	mapFromGroupsToData(data, group_list)

	return nil
}
