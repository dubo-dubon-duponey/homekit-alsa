package main

import (
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/dubo-dubon-duponey/homekit-alsa/ampcontrol"
	"github.com/dubo-dubon-duponey/homekit-alsa/homekitamp"
	"github.com/dubo-dubon-duponey/homekit-alsa/utils"
	"github.com/urfave/cli"
	"github.com/yobert/alsa"
	"log"
	"os"
)

var amp *homekitamp.HomeKitAmp

func getFirstCard() (uint, error) {
	cards, err := alsa.OpenCards()
	if err != nil {
		return 0, err
	}
	defer alsa.CloseCards(cards)
	return uint(cards[0].Number), nil
}

func listCards(c *cli.Context) error {
	cards, err := alsa.OpenCards()
	if err != nil {
		return err
	}
	defer alsa.CloseCards(cards)

	for _, card := range cards {
		devices, err := card.Devices()
		if err != nil {
			return err
		}
		fmt.Printf("Card number: %d - title: %s - path: %s\n", card.Number, card.Title, card.Path)
		for _, device := range devices {
			fmt.Printf(" > Device number: %d - title: %s - path: %s - type: %s\n", device.Number, device.Title, device.Path, device.Type)
		}
	}
	return nil
}

func getVolume(c *cli.Context) error {
	card := c.Uint64("card")
	device := c.String("device")
	ac := ampcontrol.NewAmpControl(card, device)

	v := ac.GetVolume()
	fmt.Printf("Volume: %f\n", v)
	return nil
}

func setVolume(c *cli.Context) error {
	card := c.Uint64("card")
	device := c.String("device")
	ac := ampcontrol.NewAmpControl(card, device)

	volume := c.Float64("volume")
	ac.SetVolume(volume)
	return nil
}

func getMute(c *cli.Context) error {
	card := c.Uint64("card")
	device := c.String("device")
	ac := ampcontrol.NewAmpControl(card, device)

	v := ac.GetMute()
	fmt.Printf("Mute: %t\n", v)
	return nil
}

func setMute(c *cli.Context) error {
	card := c.Uint64("card")
	device := c.String("device")
	ac := ampcontrol.NewAmpControl(card, device)

	value := c.Bool("true")
	ac.SetMute(value)
	return nil
}

func register(c *cli.Context) error {
	card := c.Uint64("card")
	device := c.String("device")
	pin := c.String("pin")
	storage := c.String("data-path")

	info := accessory.Info{
		Name:             c.String("name"),
		Manufacturer:     c.String("manufacturer"),
		SerialNumber:     c.String("serial"),
		Model:            c.String("model"),
		FirmwareRevision: c.String("version"),
	}

	amp = homekitamp.NewHomeKitAmp(card, device, info)

	t, err := hc.NewIPTransport(hc.Config{
		Pin:         pin,
		StoragePath: storage,
	}, amp.HomeKit.Accessory)
	if err != nil {
		return err
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "HomeKit Speaker"
	app.Usage = "Control your alsa devices over HomeKit"

	uuid, _ := utils.GenerateUUID()
	cardid, _ := getFirstCard()

	app.Commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "list available cards",
			Action: listCards,
		},
		{
			Name:   "get",
			Usage:  "get the current volume for that card / device",
			Action: getVolume,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "card, c",
					Value: cardid,
					Usage: "card",
				},
				cli.StringFlag{
					Name:  "device, d",
					Value: "Digital",
					Usage: "device",
				},
			},
		},
		{
			Name:   "set",
			Usage:  "set the current volume for that card / device",
			Action: setVolume,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "card, c",
					Value: cardid,
					Usage: "card",
				},
				cli.StringFlag{
					Name:  "device, d",
					Value: "Digital",
					Usage: "device",
				},
				cli.Float64Flag{
					Name:  "volume, v",
					Value: 0,
					Usage: "volume",
				},
			},
		},
		{
			Name:   "getmute",
			Usage:  "get mute",
			Action: getMute,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "card, c",
					Value: cardid,
					Usage: "card",
				},
				cli.StringFlag{
					Name:  "device, d",
					Value: "Digital",
					Usage: "device",
				},
			},
		},
		{
			Name:   "setmute",
			Usage:  "set mute",
			Action: setMute,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "card, c",
					Value: cardid,
					Usage: "card",
				},
				cli.StringFlag{
					Name:  "device, d",
					Value: "Digital",
					Usage: "device",
				},
				cli.BoolFlag{
					Name:  "true",
					Usage: "mute",
				},
			},
		},
		{
			Name:   "register",
			Usage:  "register a HomeKit device",
			Action: register,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "card, c",
					Value: cardid,
					Usage: "Alsa card",
				},
				cli.StringFlag{
					Name:  "device, d",
					Value: "Digital",
					Usage: "Alsa device",
				},
				cli.StringFlag{
					Name:  "pin, p",
					Value: "14041976",
					Usage: "Pin code for your device (8 characters)",
				},
				cli.StringFlag{
					Name:  "name",
					Value: "Dubo Dubon Duponey Amp",
					Usage: "Name of your amplifier",
				},
				cli.StringFlag{
					Name:  "data-path",
					Value: "/tmp/dubo-amp",
					Usage: "Where to store the data files for that device",
				},
				cli.StringFlag{
					Name:  "manufacturer",
					Value: "Alsa",
					Usage: "Manufacturer of your amplifier",
				},
				cli.StringFlag{
					Name:  "serial",
					Value: uuid,
					Usage: "Serial number of your amplifier",
				},
				cli.StringFlag{
					Name:  "model",
					Value: "-",
					Usage: "Model of your amplifier",
				},
				cli.StringFlag{
					Name:  "version",
					Value: "1",
					Usage: "Firmware version of your amplifier",
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
