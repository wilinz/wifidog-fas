package model

type WifiDogAuth struct {
	Stage    string `form:"stage"`
	IP       string `form:"ip"`
	MAC      string `form:"mac"`
	Token    string `form:"token"`
	Incoming int64  `form:"incoming"`
	Outgoing int64  `form:"outgoing"`
	GwId     string `form:"gw_id"`
}

//type WifiDogLogin struct {
//	GatewayAddress string `form:"gw_address"`
//	GatewayPort    int    `form:"gw_port"`
//	GatewayID      string `form:"gw_id"`
//	IP             string `form:"ip"`
//	MAC            string `form:"mac"`
//	URL            string `form:"url"`
//}

type WifiDogLogin struct {
	GwAddress   string `form:"gw_address"`
	GwPort      int    `form:"gw_port"`
	GwId        string `form:"gw_id"`
	ChannelPath string `form:"channel_path"`
	Ssid        string `form:"ssid"`
	Ip          string `form:"ip"`
	Mac         string `form:"mac"`
	Url         string `form:"url"`
}

type FasLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
	MAC      string `form:"mac"`
}

type JsonTemplate struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PingInfo struct {
	GwId             string  `form:"gw_id"`
	SysUptime        int     `form:"sys_uptime"`
	SysMemfree       int     `form:"sys_memfree"`
	SysLoad          float64 `form:"sys_load"`
	NfConntrackCount int     `form:"nf_conntrack_count"`
	CpuUsage         string  `form:"cpu_usage"`
	WifidogUptime    int     `form:"wifidog_uptime"`
	OnlineClients    int     `form:"online_clients"`
	OfflineClients   int     `form:"offline_clients"`
	Ssid             string  `form:"ssid"`
	Version          string  `form:"version"`
	Type             string  `form:"type"`
	Name             string  `form:"name"`
	ChannelPath      string  `form:"channel_path"`
	WiredPassed      int     `form:"wired_passed"`
}

type AuthInfo struct {
	GwID    string    `json:"gw_id"`
	Clients []Clients `json:"clients"`
}
type Clients struct {
	ID         int    `json:"id"`
	IP         string `json:"ip"`
	Mac        string `json:"mac"`
	Token      string `json:"token"`
	Name       string `json:"name"`
	Incoming   int64  `json:"incoming"`
	Outgoing   int64  `json:"outgoing"`
	FirstLogin int    `json:"first_login"`
	OnlineTime int    `json:"online_time"`
	IsOnline   bool   `json:"is_online"`
	Wired      bool   `json:"wired"`
}

type AuthInfoQuery struct {
	Stage       string `form:"stage"`
	Ip          string `form:"ip"`
	Mac         string `form:"mac"`
	Token       string `form:"token"`
	Incoming    int    `form:"incoming"`
	Outgoing    int    `form:"outgoing"`
	FirstLogin  int    `form:"first_login"`
	OnlineTime  int    `form:"online_time"`
	GwId        string `form:"gw_id"`
	ChannelPath string `form:"channel_path"`
	Name        string `form:"name"`
	Wired       int    `form:"wired"`
}

type AuthResponse struct {
	GwID   string   `json:"gw_id"`
	AuthOp []AuthOp `json:"auth_op"`
}
type AuthOp struct {
	ID       int `json:"id"`
	AuthCode int `json:"auth_code"`
}
