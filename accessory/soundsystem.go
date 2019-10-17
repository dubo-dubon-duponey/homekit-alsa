package accessory

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type SoundSystem struct {
	*accessory.Accessory

	Amplifier *Amplifier
}

func NewSoundSystem(info accessory.Info) *SoundSystem {
	acc := SoundSystem{}
	acc.Accessory = accessory.New(info, accessory.TypeOther)

	acc.Amplifier = NewAmplifier()
	acc.AddService(acc.Amplifier.Service)

	return &acc
}

type Amplifier struct {
	*service.Service

	Volume *characteristic.RotationSpeed

	On *characteristic.On
}

func NewAmplifier() *Amplifier {
	svc := Amplifier{}
	// service.TypeSpeaker is not available for non certified Apple products it seems, so, cheating with a fan
	svc.Service = service.New(service.TypeFan)

	// Adding characteristics on and volume
	svc.Volume = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}
