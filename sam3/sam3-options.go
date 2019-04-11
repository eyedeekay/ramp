package sam3

import (
	"fmt"
	"strconv"
    "strings"
)

import (
	. "github.com/eyedeekay/sam3"
)

//Option is a SAM Option
type Option func(*SAM) error

//SetType sets the type of the forwarder server
func SetType(s string) func(*SAM) error {
	return func(c *SAM) error {
		if s == "http" {
			c.Config.I2PConfig.Type = s
			return nil
		} else {
			c.Config.I2PConfig.Type = "server"
			return nil
		}
	}
}

func SetSAMAddress(s string) func(*SAM) error {
	return func(c *SAM) error {
		sp := strings.Split(s, ":")
		if len(sp) > 2 {
			return fmt.Errorf("Invalid address string: %s", sp)
		}
		if len(sp) == 2 {
			c.Config.I2PConfig.SamPort = sp[1]
		}
		c.Config.I2PConfig.SamHost = sp[0]
		return nil
	}
}

//SetSAMHost sets the host of the SAM's SAM bridge
func SetSAMHost(s string) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the SAM's SAM bridge using a string
func SetSAMPort(s string) func(*SAM) error {
	return func(c *SAM) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Config.I2PConfig.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetFromPort sets the FROM_PORT propert to pass to SAM
func SetFromPort(s string) func(*SAM) error {
	return func(c *SAM) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid FROM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Config.I2PConfig.Fromport = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetToPort sets the TO_PORT property to pass to SAM
func SetToPort(s string) func(*SAM) error {
	return func(c *SAM) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Config.I2PConfig.Toport = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetName sets the host of the SAM's SAM bridge
func SetName(s string) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.TunName = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u < 7 && u >= 0 {
			c.Config.I2PConfig.InLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u < 7 && u >= 0 {
			c.Config.I2PConfig.OutLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(*SAM) error {
	return func(c *SAM) error {
		if i < 7 && i > -7 {
			c.Config.I2PConfig.InVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(*SAM) error {
	return func(c *SAM) error {
		if i < 7 && i > -7 {
			c.Config.I2PConfig.OutVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u <= 16 && u > 0 {
			c.Config.I2PConfig.InQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u <= 16 && u > 0 {
			c.Config.I2PConfig.OutQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u < 6 && u >= 0 {
			c.Config.I2PConfig.InBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u < 6 && u >= 0 {
			c.Config.I2PConfig.OutBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.EncryptLeaseSet = "true"
			return nil
		}
		c.Config.I2PConfig.EncryptLeaseSet = "false"
		return nil
	}
}

//SetLeaseSetKey sets the host of the SAM's SAM bridge
func SetLeaseSetKey(s string) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the host of the SAM's SAM bridge
func SetLeaseSetPrivateKey(s string) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the host of the SAM's SAM bridge
func SetLeaseSetPrivateSigningKey(s string) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetMessageReliability sets the host of the SAM's SAM bridge
func SetMessageReliability(s string) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.MessageReliability = s
		return nil
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.InAllowZeroHop = "true"
			return nil
		}
		c.Config.I2PConfig.InAllowZeroHop = "false"
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.OutAllowZeroHop = "true"
			return nil
		}
		c.Config.I2PConfig.OutAllowZeroHop = "false"
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.UseCompression = "true"
			return nil
		}
		c.Config.I2PConfig.UseCompression = "false"
		return nil
	}
}

//SetFastRecieve tells clients to use compression
func SetFastRecieve(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.FastRecieve = "true"
			return nil
		}
		c.Config.I2PConfig.FastRecieve = "false"
		return nil
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.ReduceIdle = "true"
			return nil
		}
		c.Config.I2PConfig.ReduceIdle = "false"
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.ReduceIdleTime = "300000"
		if u >= 6 {
			c.Config.I2PConfig.ReduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.ReduceIdleTime = "300000"
		if u >= 300000 {
			c.Config.I2PConfig.ReduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(*SAM) error {
	return func(c *SAM) error {
		if u < 5 {
			c.Config.I2PConfig.ReduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(*SAM) error {
	return func(c *SAM) error {
		if b {
			c.Config.I2PConfig.CloseIdle = "true"
			return nil
		}
		c.Config.I2PConfig.CloseIdle = "false"
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.CloseIdleTime = "300000"
		if u >= 6 {
			c.Config.I2PConfig.CloseIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(*SAM) error {
	return func(c *SAM) error {
		c.Config.I2PConfig.CloseIdleTime = "300000"
		if u >= 300000 {
			c.Config.I2PConfig.CloseIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetAccessListType tells the system to treat the AccessList as a whitelist
func SetAccessListType(s string) func(*SAM) error {
	return func(c *SAM) error {
		if s == "whitelist" {
			c.Config.I2PConfig.AccessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.Config.I2PConfig.AccessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.Config.I2PConfig.AccessListType = ""
			return nil
		} else if s == "" {
			c.Config.I2PConfig.AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetAccessList tells the system to treat the AccessList as a whitelist
func SetAccessList(s []string) func(*SAM) error {
	return func(c *SAM) error {
		if len(s) > 0 {
			for _, a := range s {
				c.Config.I2PConfig.AccessList = append(c.Config.I2PConfig.AccessList, a)
			}
			return nil
		}
		return nil
	}
}
