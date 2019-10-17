# HomeKit Alsa

Lets you control the volume of (alsa) sound cards with HomeKit.

## TL;DR

```
make build
```

```
# Register your device:
./homekit-alsa register --name "Super Speaker" --pin 87654312
```

Not working?

Probably your card is not the first one we found and/or your device is not "Digital".

```
# Figure out from this what is your card number, and device name:
./homekit-alsa list

# Confirm you got the right card and device by setting the volume manually:
./homekit-alsa --card 3 --device Something volume 75

# Now register (with the right card and device)
./homekit-alsa register --card 0 --device Digital --name "Super Speaker" --pin 87654312
```

Fire-up that iPhone. Launch the Home app. Add the accessory.

## A... FAN???

Yeah.

There are no speakers in HomeKit.

See https://github.com/nfarina/homebridge/issues/1326#issuecomment-360357404 for some relevant discussion.

If you are really not happy with that, you can change the code to use a light-bulb.
I personally prefer a fan, because music is supposed to blow you away ;-).
