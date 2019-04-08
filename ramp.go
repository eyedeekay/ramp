package ramp

import (
	"log"
	"strings"
)

import (
	. "github.com/eyedeekay/sam3"
)

type Ramp struct {
	*SAM
	SAMConn
	StreamSession
	StreamListener
	SamHost string
	SamPort string
	TunName string

	Type string

	encryptLeaseSet           string
	leaseSetKey               string
	leaseSetPrivateKey        string
	leaseSetPrivateSigningKey string
	LeaseSetKeys              I2PKeys
	inAllowZeroHop            string
	outAllowZeroHop           string
	inLength                  string
	outLength                 string
	inQuantity                string
	outQuantity               string
	inVariance                string
	outVariance               string
	inBackupQuantity          string
	outBackupQuantity         string
	fastRecieve               string
	useCompression            string
	messageReliability        string
	closeIdle                 string
	closeIdleTime             string
	reduceIdle                string
	reduceIdleTime            string
	reduceIdleQuantity        string
	//Streaming Library options
	accessListType string
	accessList     []string
}

func (f Ramp) sam() string {
	return f.SamHost + ":" + f.SamPort
}

func (f Ramp) leasesetsettings() (string, string, string) {
	var r, s, t string
	if f.leaseSetKey != "" {
		r = "i2cp.leaseSetKey=" + f.leaseSetKey
	}
	if f.leaseSetPrivateKey != "" {
		s = "i2cp.leaseSetPrivateKey=" + f.leaseSetPrivateKey
	}
	if f.leaseSetPrivateSigningKey != "" {
		t = "i2cp.leaseSetPrivateSigningKey=" + f.leaseSetPrivateSigningKey
	}
	return r, s, t
}

func (f Ramp) print() []string {
	lsk, lspk, lspsk := f.leasesetsettings()
	return []string{
		//f.targetForPort443(),
		"inbound.length=" + f.inLength,
		"outbound.length=" + f.outLength,
		"inbound.lengthVariance=" + f.inVariance,
		"outbound.lengthVariance=" + f.outVariance,
		"inbound.backupQuantity=" + f.inBackupQuantity,
		"outbound.backupQuantity=" + f.outBackupQuantity,
		"inbound.quantity=" + f.inQuantity,
		"outbound.quantity=" + f.outQuantity,
		"inbound.allowZeroHop=" + f.inAllowZeroHop,
		"outbound.allowZeroHop=" + f.outAllowZeroHop,
		"i2cp.fastRecieve=" + f.fastRecieve,
		"i2cp.gzip=" + f.useCompression,
		"i2cp.reduceOnIdle=" + f.reduceIdle,
		"i2cp.reduceIdleTime=" + f.reduceIdleTime,
		"i2cp.reduceQuantity=" + f.reduceIdleQuantity,
		"i2cp.closeOnIdle=" + f.closeIdle,
		"i2cp.closeIdleTime=" + f.closeIdleTime,
		"i2cp.messageReliability" + f.messageReliability,
		"i2cp.encryptLeaseSet=" + f.encryptLeaseSet,
		lsk, lspk, lspsk,
		f.accesslisttype(),
		f.accesslist(),
	}
}

func (f Ramp) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f Ramp) accesslist() string {
	if f.accessListType != "" && len(f.accessList) > 0 {
		r := ""
		for _, s := range f.accessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func NewRamp(opts ...func(*Ramp) error) (*Ramp, error) {
	var ramp Ramp
	var err error
	ramp.SamHost = "127.0.0.1"
	ramp.SamPort = "7656"
	ramp.TunName = "Ramp"
	ramp.Type = "server"
	ramp.inLength = "3"
	ramp.outLength = "3"
	ramp.inQuantity = "2"
	ramp.outQuantity = "2"
	ramp.inVariance = "1"
	ramp.outVariance = "1"
	ramp.inBackupQuantity = "3"
	ramp.outBackupQuantity = "3"
	ramp.inAllowZeroHop = "false"
	ramp.outAllowZeroHop = "false"
	ramp.encryptLeaseSet = "false"
	ramp.leaseSetKey = ""
	ramp.leaseSetPrivateKey = ""
	ramp.leaseSetPrivateSigningKey = ""
	ramp.fastRecieve = "false"
	ramp.useCompression = "true"
	ramp.reduceIdle = "false"
	ramp.reduceIdleTime = "15"
	ramp.reduceIdleQuantity = "4"
	ramp.closeIdle = "false"
	ramp.closeIdleTime = "300000"
	ramp.messageReliability = "none"
	for _, o := range opts {
		if err := o(&ramp); err != nil {
			return nil, err
		}
	}
	if ramp.SAM, err = NewSAM(ramp.sam()); err != nil {
		return nil, err
	}
	/*log.Println("SAM Bridge connection established.")
	if ramp.save {
		log.Println("Saving i2p keys")
	}
	if ramp.SamKeys, err = i2pkeyramp.Load(ramp.FilePath, ramp.TunName, ramp.passfile, ramp.samConn, ramp.save); err != nil {
		return nil, err
	}*/
	log.Println("Destination keys generated, tunnel name:", ramp.TunName)
	/*if ramp.save {
		if err := i2pkeyramp.Save(ramp.FilePath, ramp.TunName, ramp.passfile, ramp.SamKeys); err != nil {
			return nil, err
		}
		log.Println("Saved tunnel keys for", ramp.TunName)
	}*/
	return &ramp, nil
}
