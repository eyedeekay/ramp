package i2pconfig

import (
	"strings"
)

import (
	. "github.com/eyedeekay/sam3"
)

// I2PConfig is a struct which manages I2P configuration options
type I2PConfig struct {
	SamHost string
	SamPort string
	TunName string

    SamMin string
    samMax string

	Type string

	EncryptLeaseSet           string
	LeaseSetKey               string
	LeaseSetPrivateKey        string
	LeaseSetPrivateSigningKey string
	LeaseSetKeys              I2PKeys
	InAllowZeroHop            string
	OutAllowZeroHop           string
	InLength                  string
	OutLength                 string
	InQuantity                string
	OutQuantity               string
	InVariance                string
	OutVariance               string
	InBackupQuantity          string
	OutBackupQuantity         string
	FastRecieve               string
	UseCompression            string
	MessageReliability        string
	CloseIdle                 string
	CloseIdleTime             string
	ReduceIdle                string
	ReduceIdleTime            string
	ReduceIdleQuantity        string
	//Streaming Library options
	AccessListType string
	AccessList     []string
}

func (f I2PConfig) Leasesetsettings() (string, string, string) {
	var r, s, t string
	if f.LeaseSetKey != "" {
		r = "i2cp.leaseSetKey=" + f.LeaseSetKey
	}
	if f.LeaseSetPrivateKey != "" {
		s = "i2cp.leaseSetPrivateKey=" + f.LeaseSetPrivateKey
	}
	if f.LeaseSetPrivateSigningKey != "" {
		t = "i2cp.leaseSetPrivateSigningKey=" + f.LeaseSetPrivateSigningKey
	}
	return r, s, t
}

func (f I2PConfig) Print() []string {
	lsk, lspk, lspsk := f.Leasesetsettings()
	return []string{
		//f.targetForPort443(),
		"inbound.length=" + f.InLength,
		"outbound.length=" + f.OutLength,
		"inbound.lengthVariance=" + f.InVariance,
		"outbound.lengthVariance=" + f.OutVariance,
		"inbound.backupQuantity=" + f.InBackupQuantity,
		"outbound.backupQuantity=" + f.OutBackupQuantity,
		"inbound.quantity=" + f.InQuantity,
		"outbound.quantity=" + f.OutQuantity,
		"inbound.allowZeroHop=" + f.InAllowZeroHop,
		"outbound.allowZeroHop=" + f.OutAllowZeroHop,
		"i2cp.fastRecieve=" + f.FastRecieve,
		"i2cp.gzip=" + f.UseCompression,
		"i2cp.reduceOnIdle=" + f.ReduceIdle,
		"i2cp.reduceIdleTime=" + f.ReduceIdleTime,
		"i2cp.reduceQuantity=" + f.ReduceIdleQuantity,
		"i2cp.closeOnIdle=" + f.CloseIdle,
		"i2cp.closeIdleTime=" + f.CloseIdleTime,
		"i2cp.messageReliability" + f.MessageReliability,
		"i2cp.encryptLeaseSet=" + f.EncryptLeaseSet,
		lsk, lspk, lspsk,
		f.Accesslisttype(),
		f.Accesslist(),
	}
}

func (f I2PConfig) Accesslisttype() string {
	if f.AccessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.AccessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.AccessListType == "none" {
		return ""
	}
	return ""
}

func (f I2PConfig) Accesslist() string {
	if f.AccessListType != "" && len(f.AccessList) > 0 {
		r := ""
		for _, s := range f.AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func NewConfig(opts ...func(*I2PConfig) error) (*I2PConfig, error) {
	var config I2PConfig
	config.SamHost = "127.0.0.1"
	config.SamPort = "7656"
	config.TunName = "Ramp"
	config.Type = "server"
	config.InLength = "3"
	config.OutLength = "3"
	config.InQuantity = "2"
	config.OutQuantity = "2"
	config.InVariance = "1"
	config.OutVariance = "1"
	config.InBackupQuantity = "3"
	config.OutBackupQuantity = "3"
	config.InAllowZeroHop = "false"
	config.OutAllowZeroHop = "false"
	config.EncryptLeaseSet = "false"
	config.LeaseSetKey = ""
	config.LeaseSetPrivateKey = ""
	config.LeaseSetPrivateSigningKey = ""
	config.FastRecieve = "false"
	config.UseCompression = "true"
	config.ReduceIdle = "false"
	config.ReduceIdleTime = "15"
	config.ReduceIdleQuantity = "4"
	config.CloseIdle = "false"
	config.CloseIdleTime = "300000"
	config.MessageReliability = "none"
	for _, o := range opts {
		if err := o(&config); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
