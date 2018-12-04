package config

type Config struct {
	MySQL      MySQL      `json:"MySQL,omitempty"`
	Server     Server     `json:"Server,omitempty"`
	MailServer MailServer `json:"Server,omitempty"`
}

type MySQL struct {
	Username  string `json:"Username,omitempty"`
	Password  string `json:"Password,omitempty"`
	Host      string `json:"Host,omitempty"`
	Port      int    `json:"Port,omitempty"`
	Database  string `json:"Database,omitempty"`
	Parameter string `json:"Parameter,omitempty"`
}

type Server struct {
	Port             int    `json:"Port,omitempty"`
	BaseContextPath  string `json:"BaseContextPath,omitempty"`
	PrivatePassPhase string `json:"PrivatePassPhase,omitempty"`
	Salt             string `json:"Salt,omitempty"`
}

type MailServer struct {
	Domain   string `json:"Domain,omitempty"`
	Server   string `json:"Server,omitempty"`
	Port     int    `json:"Port,omitempty"`
	Email    string `json:"Email,omitempty"`
	Password string `json:"Password,omitempty"`
}
