package homekitamp

import (
	"github.com/brutella/hc"
	hcacc "github.com/brutella/hc/accessory"
	"github.com/dubo-dubon-duponey/homekit-alsa/accessory"
	"github.com/dubo-dubon-duponey/homekit-alsa/ampcontrol"
)

type HomeKitAmp struct {
	HomeKit    *accessory.SoundSystem
	Controller *ampcontrol.AmpControl
}

func NewHomeKitAmp(card uint64, device string, info hcacc.Info) *HomeKitAmp {
	sosy := HomeKitAmp{
		accessory.NewSoundSystem(info),
		ampcontrol.NewAmpControl(card, device),
	}

	sosy.HomeKit.Amplifier.On.OnValueRemoteUpdate(sosy.Controller.SetOn)
	sosy.HomeKit.Amplifier.On.OnValueRemoteGet(sosy.Controller.GetOn)

	sosy.HomeKit.Amplifier.Volume.OnValueRemoteUpdate(sosy.Controller.SetVolume)
	sosy.HomeKit.Amplifier.Volume.OnValueRemoteGet(sosy.Controller.GetVolume)

	return &sosy
}

func (a *HomeKitAmp) Start(pin string) error {

	t, err := hc.NewIPTransport(hc.Config{Pin: pin}, a.HomeKit.Accessory)
	if err != nil {
		return err
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()

	return nil
}
