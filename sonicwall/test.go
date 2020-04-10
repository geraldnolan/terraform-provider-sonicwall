package sonicwall

//import "github.com/geraldnolan/terraform-provider-sonicwall/sonicwall"

/*type SonicWallClient struct {
	username string
	password string
	hostname string
	port     string
}*/

type SonicWallSchema struct {
	name            string
	command_type    string
	zone_assignment string
	sw_type         string
	ipv4address     string
}

func test() {
	/*plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sonicwall.Provider,
	})*/

	client := SonicWallClient{
		username: "test",
		password: "test",
		hostname: "192.168.68.1",
		port:     "22",
	}

	sonicwallSchema := SonicWallSchema{
		name:            "test (private)",
		command_type:    "address-object",
		zone_assignment: "host",
		sw_type:         "LAN",
		ipv4address:     "192.168.68.55",
	}

	command := sonicwallSchema.command_type + " ipv4 \"" + sonicwallSchema.name + "\" " + sonicwallSchema.sw_type + " " + sonicwallSchema.ipv4address + " zone " + sonicwallSchema.zone_assignment

	shellCommand := []string{
		"config",
		command,
		"commit",
		"end",
		"exit",
	}

	RunSSHCommand(client.username, client.password, client.hostname, client.port, shellCommand)
}
