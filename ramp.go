package ramp

import (
	"log"
)

import (
	. "github.com/eyedeekay/ramp/config"
	. "github.com/eyedeekay/sam3"
)

type Ramp struct {
	*SAM
	SAMConn
	StreamSession
	StreamListener
	I2PConfig
}

func NewRamp(opts ...func(*Ramp) error) (*Ramp, error) {
	var ramp Ramp
	var err error
	ramp.I2PConfig.SamHost = "127.0.0.1"
	ramp.I2PConfig.SamPort = "7656"
	ramp.I2PConfig.TunName = "Ramp"
	ramp.I2PConfig.Type = "server"
	ramp.I2PConfig.InLength = "3"
	ramp.I2PConfig.OutLength = "3"
	ramp.I2PConfig.InQuantity = "2"
	ramp.I2PConfig.OutQuantity = "2"
	ramp.I2PConfig.InVariance = "1"
	ramp.I2PConfig.OutVariance = "1"
	ramp.I2PConfig.InBackupQuantity = "3"
	ramp.I2PConfig.OutBackupQuantity = "3"
	ramp.I2PConfig.InAllowZeroHop = "false"
	ramp.I2PConfig.OutAllowZeroHop = "false"
	ramp.I2PConfig.EncryptLeaseSet = "false"
	ramp.I2PConfig.LeaseSetKey = ""
	ramp.I2PConfig.LeaseSetPrivateKey = ""
	ramp.I2PConfig.LeaseSetPrivateSigningKey = ""
	ramp.I2PConfig.FastRecieve = "false"
	ramp.I2PConfig.UseCompression = "true"
	ramp.I2PConfig.ReduceIdle = "false"
	ramp.I2PConfig.ReduceIdleTime = "15"
	ramp.I2PConfig.ReduceIdleQuantity = "4"
	ramp.I2PConfig.CloseIdle = "false"
	ramp.I2PConfig.CloseIdleTime = "300000"
	ramp.I2PConfig.MessageReliability = "none"
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
