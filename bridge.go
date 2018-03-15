package api

type bridge struct {
	alert                 *AlertAPI
	apRules               *APRulesAPI
	applications          *ApplicationsAPI
	aps                   *APsAPI
	avc                   *AVCAPI
	blockClient           *BlockClientAPI
	certstore             *CertStoreAPI
	clients               *ClientsAPI
	cluster               *ClusterAPI
	configuration         *ConfigurationAPI
	configurationSettings *ConfigurationSettingsAPI
	controlPlanes         *ControlPlanesAPI
	controller            *ControllerAPI
	dhcpData              *DHCPDataAPI
	ftps                  *FTPSAPI
	globalSettings        *GlobalSettingsAPI
	identity              *IdentityAPI
	lwapp2scg             *LWAPP2SCGAPI
	maps                  *MapsAPI
	planes                *PlanesAPI
	precedence            *PrecedenceAPI
	profiles              *ProfilesAPI
	query                 *QueryAPI
	restart               *RestartAPI
	rkszones              *RuckusZonesAPI
	rogue                 *RogueAPI
	rogueaps              *RogueAPsAPI
	sci                   *SCIAPI
	services              *ServicesAPI
	session               *SessionAPI
	shutdown              *ShutdownAPI
	smsGateway            *SMSGatewayAPI
	system                *SystemAPI
	tool                  *ToolAPI
	upgrade               *UpgradeAPI
	userGroups            *UserGroupsAPI
	users                 *UsersAPI
	vlanpoolings          *VLANPoolingsAPI
}

func newBridge(c *Client) *bridge {
	return &bridge{
		alert:                 &AlertAPI{c},
		apRules:               &APRulesAPI{c},
		applications:          &ApplicationsAPI{c},
		aps:                   &APsAPI{c},
		avc:                   &AVCAPI{c},
		blockClient:           &BlockClientAPI{c},
		certstore:             &CertStoreAPI{c},
		clients:               &ClientsAPI{c},
		cluster:               &ClusterAPI{c},
		configuration:         &ConfigurationAPI{c},
		configurationSettings: &ConfigurationSettingsAPI{c},
		controlPlanes:         &ControlPlanesAPI{c},
		controller:            &ControllerAPI{c},
		dhcpData:              &DHCPDataAPI{c},
		ftps:                  &FTPSAPI{c},
		globalSettings:        &GlobalSettingsAPI{c},
		identity:              &IdentityAPI{c},
		lwapp2scg:             &LWAPP2SCGAPI{c},
		maps:                  &MapsAPI{c},
		planes:                &PlanesAPI{c},
		precedence:            &PrecedenceAPI{c},
		profiles:              &ProfilesAPI{c},
		query:                 &QueryAPI{c},
		restart:               &RestartAPI{c},
		rkszones:              &RuckusZonesAPI{c},
		rogue:                 &RogueAPI{c},
		rogueaps:              &RogueAPsAPI{c},
		sci:                   &SCIAPI{c},
		services:              &ServicesAPI{c},
		session:               &SessionAPI{c},
		shutdown:              &ShutdownAPI{c},
		smsGateway:            &SMSGatewayAPI{c},
		system:                &SystemAPI{c},
		tool:                  &ToolAPI{c},
		upgrade:               &UpgradeAPI{c},
		userGroups:            &UserGroupsAPI{c},
		users:                 &UsersAPI{c},
		vlanpoolings:          &VLANPoolingsAPI{c},
	}
}

func (b *bridge) Alert() *AlertAPI {
	return b.alert
}
func (b *bridge) APRules() *APRulesAPI {
	return b.apRules
}
func (b *bridge) Applications() *ApplicationsAPI {
	return b.applications
}
func (b *bridge) APs() *APsAPI {
	return b.aps
}
func (b *bridge) AVC() *AVCAPI {
	return b.avc
}
func (b *bridge) BlockClient() *BlockClientAPI {
	return b.blockClient
}
func (b *bridge) CertStore() *CertStoreAPI {
	return b.certstore
}
func (b *bridge) Clients() *ClientsAPI {
	return b.clients
}
func (b *bridge) Cluster() *ClusterAPI {
	return b.cluster
}
func (b *bridge) Configuration() *ConfigurationAPI {
	return b.configuration
}
func (b *bridge) ConfigurationSettings() *ConfigurationSettingsAPI {
	return b.configurationSettings
}
func (b *bridge) ControlPlanes() *ControlPlanesAPI {
	return b.controlPlanes
}
func (b *bridge) Controller() *ControllerAPI {
	return b.controller
}
func (b *bridge) DHCPData() *DHCPDataAPI {
	return b.dhcpData
}
func (b *bridge) FTPS() *FTPSAPI {
	return b.ftps
}
func (b *bridge) GlobalSettings() *GlobalSettingsAPI {
	return b.globalSettings
}
func (b *bridge) Identity() *IdentityAPI {
	return b.identity
}
func (b *bridge) LWAPP2SCG() *LWAPP2SCGAPI {
	return b.lwapp2scg
}
func (b *bridge) Maps() *MapsAPI {
	return b.maps
}
func (b *bridge) Planes() *PlanesAPI {
	return b.planes
}
func (b *bridge) Precedence() *PrecedenceAPI {
	return b.precedence
}
func (b *bridge) Profiles() *ProfilesAPI {
	return b.profiles
}
func (b *bridge) Query() *QueryAPI {
	return b.query
}
func (b *bridge) Restart() *RestartAPI {
	return b.restart
}
func (b *bridge) RuckusZones() *RuckusZonesAPI {
	return b.rkszones
}
func (b *bridge) Rogue() *RogueAPI {
	return b.rogue
}
func (b *bridge) RogueAPs() *RogueAPsAPI {
	return b.rogueaps
}
func (b *bridge) SCI() *SCIAPI {
	return b.sci
}
func (b *bridge) Services() *ServicesAPI {
	return b.services
}
func (b *bridge) Session() *SessionAPI {
	return b.session
}
func (b *bridge) Shutdown() *ShutdownAPI {
	return b.shutdown
}
func (b *bridge) SMSGateway() *SMSGatewayAPI {
	return b.smsGateway
}
func (b *bridge) System() *SystemAPI {
	return b.system
}
func (b *bridge) Tool() *ToolAPI {
	return b.tool
}
func (b *bridge) Upgrade() *UpgradeAPI {
	return b.upgrade
}
func (b *bridge) UserGroups() *UserGroupsAPI {
	return b.userGroups
}
func (b *bridge) Users() *UsersAPI {
	return b.users
}
func (b *bridge) VLANPoolings() *VLANPoolingsAPI {
	return b.vlanpoolings
}
