package sam3

import (
	"fmt"
	"net"
	"strings"
)

import (
	. "github.com/eyedeekay/sam3"
)

func NewSAMOpts(opts ...func(*SAM) error) (*SAM, error) {
	var sam SAM
	var err error
	for _, o := range opts {
		if err := o(&sam); err != nil {
			return nil, err
		}
	}
	if conn, err := net.Dial("tcp", sam.Config.I2PConfig.Sam()); err == nil {
		if _, err = conn.Write(sam.Config.HelloBytes()); err == nil {
			buf := make([]byte, 256)
			n, err := conn.Read(buf)
			if err == nil {
				if strings.Contains(string(buf[:n]), "HELLO REPLY RESULT=OK") {
					sam.Conn = conn
					sam.Config.I2PConfig.DestinationKeys = nil
					sam.Resolver, err = NewSAMResolver(&sam)
					if err == nil {
    					return &sam, nil
					}
				} else if string(buf[:n]) == "HELLO REPLY RESULT=NOVERSION\n" {
					err = fmt.Errorf("That SAM bridge does not support SAMv3.")
				} else {

					err = fmt.Errorf("%s", string(buf[:n]))
				}
			}
		}
		conn.Close()
	}
	return nil, err

}

/*func NewKeys(opts ...func(*I2PKeys) error) (*I2PKeys, error){

}*/

/*func newGenericSession(opts ...func(*GenericSession) error) (*GenericSession, error){

}*/
