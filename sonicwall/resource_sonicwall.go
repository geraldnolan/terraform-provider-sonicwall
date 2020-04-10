package sonicwall

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSonicWallRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceSonicWallRecordCreate,
		Read:   resourceSonicWallRecordRead,
		Delete: resourceSonicWallRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_assignment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"command_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipv4address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSonicWallRecordCreate(d *schema.ResourceData, m interface{}) error {
	//convert the interface so we can use the variables like username, etc
	client := m.(*SonicWallClient)

	name := d.Get("name").(string)
	zone_assignment := d.Get("zone_assignment").(string)
	command_type := d.Get("command_type").(string)
	sw_type := d.Get("type").(string)
	ipv4address := d.Get("ipv4address").(string)

	var id string = name + "_" + zone_assignment + "_" + command_type + "_" + sw_type + "_" + ipv4address

	var shellCommand []string

	switch command_type {
	case "address-object":
		if ipv4address == "" {
			return errors.New("Must provide ipv4address if command_type is 'address-object'")
		}
		//psCommand = "Add-DNSServerResourceRecord -ZoneName " + zone_name + " -" + record_type + " -Name " + record_name + " -IPv4Address " + ipv4address
		//"address-object ipv4 \"AnyObject (Chuck)\" host 192.168.68.95 zone LAN",
		command := command_type + " ipv4 \"" + name + "\" " + sw_type + " " + ipv4address + " zone " + zone_assignment
		shellCommand = []string{
			"config",
			command,
			"commit",
			"end",
			"exit",
		}
	case "group-objects":
		/*if hostnamealias == "" {
			return errors.New("Must provide hostnamealias if record_type is 'CNAME'")
		}*/
		//psCommand = "Add-DNSServerResourceRecord -ZoneName " + zone_name + " -" + record_type + " -Name " + record_name + " -HostNameAlias " + hostnamealias
	default:
		return errors.New("Unknown command_type. This provider currently only supports 'address-object' and 'group-objects' records")
	}

	_, err := RunSSHCommand(client.username, client.password, client.hostname, client.port, shellCommand)

	if err != nil {
		if err.Error() == "wait: remote command exited without exit status or exit signal" {
			// ExitMissingError is returned if a session is torn down cleanly, but
			// the server sends no confirmation of the exit status.
			//set err to nill
			//err = nil
			//return b.String(), nil

		} else {
			//something bad happened
			return err
		}
	}

	d.SetId(id)

	return nil
}

func resourceSonicWallRecordRead(d *schema.ResourceData, m interface{}) error {
	//convert the interface so we can use the variables like username, etc
	/*client := m.(*SonicWallClient)*/

	name := d.Get("name").(string)
	zone_assignment := d.Get("zone_assignment").(string)
	command_type := d.Get("command_type").(string)
	sw_type := d.Get("type").(string)
	//ipv4address := d.Get("ipv4address").(string)

	//var psCommand string = "try { $record = Get-DnsServerResourceRecord -ZoneName " + zone_name + " -RRType " + record_type + " -Name " + record_name + " -ErrorAction Stop } catch { $record = '''' }; if ($record) { write-host 'RECORD_FOUND' }"
	/*var psCommand []string = {
		"show address-object ipv4 \"CV-Exchange (Private)\""
	}*/

	/*command := "show " + command_type + " ipv4 \"" + name + "\""
	commands := []string{
		command,
	}*/

	/*_, err := RunSSHCommand(client.username, client.password, client.hostname, client.port, commands)
	if err != nil {
		if !strings.Contains(err.Error(), "ObjectNotFound") {
			//something bad happened
			return err
		} else {
			//not able to find the record - this is an error but ok
			d.SetId("")
			return nil
		}
	}*/

	var id string = name + "_" + zone_assignment + "_" + command_type + "_" + sw_type
	d.SetId(id)

	return nil
}

func resourceSonicWallRecordDelete(d *schema.ResourceData, m interface{}) error {
	//convert the interface so we can use the variables like username, etc
	/*client := m.(*SonicWallClient)*/

	//zone_name := d.Get("zone_name").(string)
	//record_type := d.Get("record_type").(string)
	//record_name := d.Get("record_name").(string)

	/*name := d.Get("name").(string)
	zone_assignment := d.Get("zone_assignment").(string)
	command_type := d.Get("command_type").(string)
	sw_type := d.Get("type").(string)
	ipv4address := d.Get("ipv4address").(string)*/

	//Remove-DnsServerResourceRecord -ZoneName "contoso.com" -RRType "A" -Name "Host01"
	//var psCommand string = "Remove-DNSServerResourceRecord -ZoneName " + zone_name + " -RRType " + record_type + " -Name " + record_name + " -Confirm:$false -Force"
	//var psCommand string = "show address-object ipv4 \"CV-Exchange (Private)\""
	/*
		command := "no " + command_type + " ipv4 " + name + " " + sw_type + " " + ipv4address + " zone " + zone_assignment
		commands := []string{
			"config",
			command,
			"exit",
		}*/

	/*_, err := RunSSHCommand(client.username, client.password, client.hostname, client.port, commands)
	if err != nil {
	//something bad happened
		return err
	}*/

	// d.SetId("") is automatically called assuming delete returns no errors, but it is added here for explicitness.
	d.SetId("")

	return nil
}
