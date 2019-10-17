package ampcontrol

import (
	"errors"
	"github.com/dubo-dubon-duponey/homekit-alsa/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type AmpControl struct {
	Card   uint64
	Device string
}

func (a *AmpControl) GetMute() bool {
	out, err := utils.ExecCmd([]string{"amixer", "-c", strconv.FormatUint(a.Card, 10), "get", a.Device})
	if err != nil {
		log.Fatalf("Failed getting volume. Amixer output: %s - error was: %s", out, err)
		return false
	}
	muted, err := parseMuted(string(out))
	if err != nil {
		log.Fatalf("Failed parsing mute. Amixer output: %s - error was: %s", out, err)
		return false
	}

	return muted
}

func (a *AmpControl) SetMute(value bool) {
	var com string
	if value {
		com = "mute"
	} else {
		com = "unmute"
	}
	out, err := utils.ExecCmd([]string{"amixer", "-c", strconv.FormatUint(a.Card, 10), "set", a.Device, com})
	if err != nil {
		log.Fatalf("Failed setting mute. Amixer output: %s - error was: %s", out, err)
	}
}

func (a *AmpControl) GetVolume() float64 {
	out, err := utils.ExecCmd([]string{"amixer", "-c", strconv.FormatUint(a.Card, 10), "get", a.Device})
	if err != nil {
		log.Fatalf("Failed getting volume. Amixer output: %s - error was: %s", out, err)
		return 0
	}
	vol, err := parseVolume(string(out))
	if err != nil {
		log.Fatalf("Failed parsing volume. Amixer output: %s - error was: %s", out, err)
		return 0
	}
	return float64(vol)
}

func (a *AmpControl) SetVolume(value float64) {
	out, err := utils.ExecCmd([]string{"amixer", "-c", strconv.FormatUint(a.Card, 10), "set", a.Device, strconv.FormatFloat(value, 'f', 0, 64) + "%"})
	if err != nil {
		log.Fatalf("Failed setting volume. Amixer output: %s - error was: %s", out, err)
	}
}

func NewAmpControl(card uint64, device string) *AmpControl {
	return &AmpControl{
		card,
		device,
	}
}

var volumePattern = regexp.MustCompile(`\d+%`)

func parseVolume(out string) (int, error) {
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")
		if strings.Contains(s, "Playback") && strings.Contains(s, "%") {
			volumeStr := volumePattern.FindString(s)
			return strconv.Atoi(volumeStr[:len(volumeStr)-1])
		}
	}
	return 0, errors.New("no volume found")
}

func parseMuted(out string) (bool, error) {
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")
		if strings.Contains(s, "Playback") && strings.Contains(s, "%") {
			if strings.Contains(s, "[off]") || strings.Contains(s, "yes") {
				return true, nil
			} else if strings.Contains(s, "[on]") || strings.Contains(s, "no") {
				return false, nil
			}
		}
	}
	return false, errors.New("no muted information found")
}
