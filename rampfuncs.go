package ramp

func (f Ramp) sam() string {
	return f.SamHost + ":" + f.SamPort
}

func (f Ramp) leasesetsettings() (string, string, string) {
	return f.I2PConfig.Leasesetsettings()
}

func (f Ramp) print() []string {
	return f.I2PConfig.Print()
}

func (f Ramp) accesslisttype() string {
	return f.I2PConfig.Accesslisttype()
}

func (f Ramp) accesslist() string {
	return f.I2PConfig.Accesslist()
}
