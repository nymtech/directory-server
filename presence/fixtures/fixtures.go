package fixtures

import "github.com/nymtech/nym-directory/models"

func GoodCocoHost() models.CocoHostInfo {
	good := models.CocoHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host",
			PubKey: "pubkey",
		},
		Type: "type",
	}
	return good
}

func GoodHost() models.MixHostInfo {
	good := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host",
			PubKey: "pubkey",
		},
	}
	return good
}

func GoodMixProviderHost() models.MixProviderHostInfo {
	client1 := models.RegisteredClient{PubKey: "client1"}
	client2 := models.RegisteredClient{PubKey: "client2"}
	clients := []models.RegisteredClient{client1, client2}
	good := models.MixProviderHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host",
			PubKey: "pubkey",
		},
		RegisteredClients: clients,
	}
	return good
}

func XssCocoHost() models.CocoHostInfo {
	xss := models.CocoHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')",
			PubKey: "pubkey<script>alert('gotcha')",
		},
		Type: "type<script>alert('gotcha')",
	}
	return xss
}

func XssMixHost() models.MixHostInfo {
	xss := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')</script>",
			PubKey: "pubkey<script>alert('gotcha')</script>",
		},
	}
	return xss
}

func XssMixProviderHost() models.MixProviderHostInfo {
	client1 := models.RegisteredClient{PubKey: "client1<script>alert('gotcha')</script>"}
	client2 := models.RegisteredClient{PubKey: "client2<script>alert('gotcha')</script>"}
	clients := []models.RegisteredClient{client1, client2}
	xss := models.MixProviderHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')</script>",
			PubKey: "pubkey<script>alert('gotcha')</script>",
		},
		RegisteredClients: clients,
	}
	return xss
}