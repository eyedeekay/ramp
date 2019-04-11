package ramp

func (f Ramp) sam() string {
	return f.SAM.Config.I2PConfig.Sam()
}

func (f Ramp) leasesetsettings() (string, string, string) {
	return f.SAM.Config.I2PConfig.Leasesetsettings()
}

func (f Ramp) print() []string {
	return f.SAM.Config.I2PConfig.Print()
}

func (f Ramp) accesslisttype() string {
	return f.SAM.Config.I2PConfig.Accesslisttype()
}

func (f Ramp) accesslist() string {
	return f.SAM.Config.I2PConfig.Accesslist()
}
