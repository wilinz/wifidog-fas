package dao

type WifiDogAuth struct {
	Stage    string `form:"stage"`
	IP       string `form:"ip"`
	MAC      string `form:"mac"`
	Token    string `form:"token"`
	Incoming int64  `form:"incoming"`
	Outgoing int64  `form:"outgoing"`
	GwId     string `form:"gw_id"`
}

type WifiDogLogin struct {
	GatewayAddress string `form:"gw_address"`
	GatewayPort    int    `form:"gw_port"`
	GatewayID      string `form:"gw_id"`
	IP             string `form:"ip"`
	MAC            string `form:"mac"`
	URL            string `form:"url"`
}

type FasLogin struct {
	Username       string `form:"username"`
	Password       string `form:"password"`
	MAC            string `form:"mac"`
}

type JsonTemplate struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
