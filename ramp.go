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
	if ramp.SamKeys, err = i2pkeyramp.Load(ramp.FilePath, ramp.ID(), ramp.passfile, ramp.samConn, ramp.save); err != nil {
		return nil, err
	}*/
	log.Println("Destination keys generated, tunnel name:", ramp.ID())
	/*if ramp.save {
		if err := i2pkeyramp.Save(ramp.FilePath, ramp.ID(), ramp.passfile, ramp.SamKeys); err != nil {
			return nil, err
		}
		log.Println("Saved tunnel keys for", ramp.ID())
	}*/
	return &ramp, nil
}
