# HomeKit Alsa

> So, you have a bunch of Raspberry PIs with speakers hooked-up to them


Lets you control the volume of (alsa) linux sound cards with HomeKit.


## TL;DR

```bash
# Look Ma! No CaPs! No r00t! No WrIte! No deps! Multi-arch! ~30MB! Like the gRoWN-ups do!
docker run -d \
    --env HOMEKIT_NAME="My Fancy Speaker" \
    --env HOMEKIT_PIN="87654312" \
    --name speaker \
    --read-only \
    --cap-drop ALL \
    --group-add audio \
    --net host \
    --device /dev/snd \
    --volume /data \
    --rm \
    dubodubonduponey/homekit-alsa:v1
```

### It works!

Cool.

 * open your iPhone
 * hit the (+) button (top right hand)
 * then "Add Accessory"
 * now "I don't have a code or cannot scan"
 * hit "My_Fancy_Speaker"
 * "Add Anyway"
 * type your pin from HOMEKIT_PIN above
 * "Next", "Done"

### It... DOES... NOT... WORK

Then you need to figure out which sound card / device you are targeting.

Start by figuring out which one is your card (really? you have more than one?)

```bash
docker run \
    --read-only \
    --cap-drop ALL \
    --group-add audio \
    --device /dev/snd \
    --entrypoint homekit-alsa \
    dubodubonduponey/homekit-alsa:v1 \
    list
```

Typically `amixer` or `aplay` (from alsa-utils) should be helpful too, and will help you figure out the "device" part.

Then re-run the image from step 1 above, additionally forcing a card and device

```
--env ALSA_CARD 3 \
--env ALSA_DEVICE WhateverYouFound \
```

When ALSA_CARD points to your alsa card number.

And ALSA_DEVICE is typically "PCM" (or "Digital"), or whatever else amixer showed you.

## Roll your own, for the strong and spirited!

You need golang1.13+ to build (probably older versions work as well but can't be bothered to check).

You also need alsa-utils at runtime (I wish the golang alsa lib would also do amixer...).


```
make build
```

```
# Register your device:
./homekit-alsa register --name "Super Speaker" --pin 87654312
```

Not working?

Same as above...

```
# Figure out from this what is your card number:
./homekit-alsa list

# Or dig into what amixer is reporting

# Confirm you got the right card and device by setting the volume manually:
./homekit-alsa set --card 3 --device Something --volume 75

# Now register (with the right card and device)
./homekit-alsa register --card 3 --device Something --name "Super Speaker" --pin 87654312
```

Fire-up that iPhone. Launch the Home app, etc.

## IT'SSSS... A... FAN???

Yeah.

There are no speakers in HomeKit (honest! don't give me this BS about speakers being available - it's only for surveillance cameras).

See https://github.com/nfarina/homebridge/issues/1326#issuecomment-360357404 if you don't believe me.

If you are really not happy with that, you can change the code to use a "light-bulb" instead of a "fan".

I personally prefer a fan, because music is supposed to blow you away ;-).
