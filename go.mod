module github.com/dubo-dubon-duponey/homekit-alsa

require github.com/brutella/hc v1.1.0

replace github.com/yobert/alsa => github.com/dubo-dubon-duponey/alsa v0.0.0-20191017073806-461af7b4c18f

replace github.com/go-xorm/core => xorm.io/core v0.6.3

require (
	github.com/urfave/cli v1.22.1
	github.com/yobert/alsa v0.0.0-00010101000000-000000000000
)

go 1.13
