package api

type bridge struct {
	alert                 *Alert
	apRules               *APRules
	applications          *Applications
	aps                   *APs
	avc                   *AVC
	blockClient           *BlockClient
	certstore             *CertStore
	clients               *Clients
	cluster               *Cluster
	configuration         *Configuration
	configurationSettings *ConfigurationSettings
	controlPlanes         *ControlPlanes
	controller            *Controller
	dhcpData              *DHCPData
	ftps                  *FTPS
	globalSettings        *GlobalSettings
	identity              *Identity
	lwapp2scg             *LWAPP2SCG
	maps                  *Maps
	planes                *Planes
	precedence            *Precedence
	profiles              *Profiles
	query                 *Query
	restart               *Restart
	rkszones              *RuckusZones
	rogue                 *Rogue
	rogueaps              *RogueAPs
	sci                   *SCI
	services              *Services
	session               *Session
	shutdown              *Shutdown
	smsGateway            *SMSGateway
	system                *System
	tool                  *Tool
	upgrade               *Upgrade
	userGroups            *UserGroups
	users                 *Users
	vlanpoolings          *VLANPoolings
}

func newBridge(c *Client) *bridge {
	return &bridge{
		alert:                 &Alert{c},
		apRules:               &APRules{c},
		applications:          &Applications{c},
		aps:                   &APs{c},
		avc:                   &AVC{c},
		blockClient:           &BlockClient{c},
		certstore:             &CertStore{c},
		clients:               &Clients{c},
		cluster:               &Cluster{c},
		configuration:         &Configuration{c},
		configurationSettings: &ConfigurationSettings{c},
		controlPlanes:         &ControlPlanes{c},
		controller:            &Controller{c},
		dhcpData:              &DHCPData{c},
		ftps:                  &FTPS{c},
		globalSettings:        &GlobalSettings{c},
		identity:              &Identity{c},
		lwapp2scg:             &LWAPP2SCG{c},
		maps:                  &Maps{c},
		planes:                &Planes{c},
		precedence:            &Precedence{c},
		profiles:              &Profiles{c},
		query:                 &Query{c},
		restart:               &Restart{c},
		rkszones:              &RuckusZones{c},
		rogue:                 &Rogue{c},
		rogueaps:              &RogueAPs{c},
		sci:                   &SCI{c},
		services:              &Services{c},
		session:               &Session{c},
		shutdown:              &Shutdown{c},
		smsGateway:            &SMSGateway{c},
		system:                &System{c},
		tool:                  &Tool{c},
		upgrade:               &Upgrade{c},
		userGroups:            &UserGroups{c},
		users:                 &Users{c},
		vlanpoolings:          &VLANPoolings{c},
	}
}

func (b *bridge) Alert() *Alert {
	return b.alert
}
func (b *bridge) APRules() *APRules {
	return b.apRules
}
func (b *bridge) Applications() *Applications {
	return b.applications
}
func (b *bridge) APs() *APs {
	return b.aps
}
func (b *bridge) AVC() *AVC {
	return b.avc
}
func (b *bridge) BlockClient() *BlockClient {
	return b.blockClient
}
func (b *bridge) CertStore() *CertStore {
	return b.certstore
}
func (b *bridge) Clients() *Clients {
	return b.clients
}
func (b *bridge) Cluster() *Cluster {
	return b.cluster
}
func (b *bridge) Configuration() *Configuration {
	return b.configuration
}
func (b *bridge) ConfigurationSettings() *ConfigurationSettings {
	return b.configurationSettings
}
func (b *bridge) ControlPlanes() *ControlPlanes {
	return b.controlPlanes
}
func (b *bridge) Controller() *Controller {
	return b.controller
}
func (b *bridge) DHCPData() *DHCPData {
	return b.dhcpData
}
func (b *bridge) FTPS() *FTPS {
	return b.ftps
}
func (b *bridge) GlobalSettings() *GlobalSettings {
	return b.globalSettings
}
func (b *bridge) Identity() *Identity {
	return b.identity
}
func (b *bridge) LWAPP2SCG() *LWAPP2SCG {
	return b.lwapp2scg
}
func (b *bridge) Maps() *Maps {
	return b.maps
}
func (b *bridge) Planes() *Planes {
	return b.planes
}
func (b *bridge) Precedence() *Precedence {
	return b.precedence
}
func (b *bridge) Profiles() *Profiles {
	return b.profiles
}
func (b *bridge) Query() *Query {
	return b.query
}
func (b *bridge) Restart() *Restart {
	return b.restart
}
func (b *bridge) RuckusZones() *RuckusZones {
	return b.rkszones
}
func (b *bridge) Rogue() *Rogue {
	return b.rogue
}
func (b *bridge) RogueAPs() *RogueAPs {
	return b.rogueaps
}
func (b *bridge) SCI() *SCI {
	return b.sci
}
func (b *bridge) Services() *Services {
	return b.services
}
func (b *bridge) Session() *Session {
	return b.session
}
func (b *bridge) Shutdown() *Shutdown {
	return b.shutdown
}
func (b *bridge) SMSGateway() *SMSGateway {
	return b.smsGateway
}
func (b *bridge) System() *System {
	return b.system
}
func (b *bridge) Tool() *Tool {
	return b.tool
}
func (b *bridge) Upgrade() *Upgrade {
	return b.upgrade
}
func (b *bridge) UserGroups() *UserGroups {
	return b.userGroups
}
func (b *bridge) Users() *Users {
	return b.users
}
func (b *bridge) VLANPoolings() *VLANPoolings {
	return b.vlanpoolings
}
