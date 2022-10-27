// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/internal/client"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: configure,
		Schema:               zschemas.ProviderSchema,
		DataSourcesMap: map[string]*schema.Resource{
			"zedcloud_datastore":        datasources.NewDataStore(),
			"zedcloud_edgeapp":          datasources.NewEdgeApp(),
			"zedcloud_edgeapp_instance": getAppInstDataSourceSchema(),
			"zedcloud_edgenode":         getEdgeNodeDataSourceSchema(),
			"zedcloud_image":            getImageDataSourceSchema(),
			"zedcloud_network":          getNetworkDataSourceSchema(),
			"zedcloud_network_instance": getNetInstDataSourceSchema(),
			"zedcloud_volume_instance":  getVolumeInstanceDataSourceSchema(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgeapp":          resources.NewEdgeApp(),
			"zedcloud_edgeapp_instance": getAppInstResourceSchema(),
			"zedcloud_edgenode":         getEdgeNodeResourceSchema(),
			"zedcloud_image":            getImageResourceSchema(),
			"zedcloud_network":          getNetworkResourceSchema(),
			"zedcloud_network_instance": getNetInstResourceSchema(),
			"zedcloud_volume_instance":  getVolumeInstanceResourceSchema(),
		},
	}
}

// configure parses the config into the Terraform provider meta object
func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := resourcedata.GetStr(d, "token")
	username := resourcedata.GetStr(d, "username")
	password := resourcedata.GetStr(d, "password")
	zedcloudUrl := resourcedata.GetStr(d, "zedcloud_url")

	var diags diag.Diagnostics
	client, err := client.NewClient(zedcloudUrl, token, username, password)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("could not configure provider: %w", err))
	}

	log.Println("successfully configured provider")
	return client, diags
}
