package i2pconfig

import (
	"strings"
)

import (
	. "github.com/eyedeekay/sam3"
)

type I2PConfig struct {
	SamHost string
	SamPort string
	TunName string

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
