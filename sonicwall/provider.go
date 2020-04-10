package sonicwall

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("USERNAME", nil),
				Description: "Username to connect.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD", nil),
				Description: "The password to connect.",
			},
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HOSTNAME", nil),
				Description: "The Sonicwall Server to connect to.",
			},
			"port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PORT", false),
				Description: "The TCP port used",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sonicwall": resourceSonicWallRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	if username == "" {
		return nil, fmt.Errorf("The 'username' property was not specified.")
	}

	password := d.Get("password").(string)
	if password == "" {
		return nil, fmt.Errorf("The 'password' property was not specified and usessh was false.")
	}

	hostname := d.Get("hostname").(string)
	if hostname == "" {
		return nil, fmt.Errorf("The 'hostname' property was not specified.")
	}

	port := d.Get("port").(string)
	if port == "" {
		return nil, fmt.Errorf("The 'port' property was not specified.")
	}

	client := SonicWallClient{
		username: username,
		password: password,
		hostname: hostname,
		port:     port,
	}

	return &client, nil
}

type SonicWallClient struct {
	username string
	password string
	hostname string
	port     string
}
