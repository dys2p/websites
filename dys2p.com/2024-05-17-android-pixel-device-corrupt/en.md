# 2024-05-17 Fixing the "Your device is corrupt" error on Google Pixel smartphones with GrapheneOS

A short while ago, the error message _"Your device is corrupt. It can't be trusted and may not work properly. Press power button to continue."_ appeared on the smartphone of a customer who bought a [Google Pixel 6a with pre-installed GrapheneOS](https://shop.proxysto.re/category/6) and locked bootloader from us. It was possible to skip the error by pressing the power button, but GrapheneOS did not start.

This error message can be a sign of tampering or a hardware defect. In our case, however, the smartphone had switched off overnight because the battery was empty. Apparently, the sudden power failure had interrupted the installation of a system update.

After charging the battery, we first switched off the Pixel 6a by pressing and holding the power button. Then we put it into _fastboot mode_ with `Power + Volume Down` and connected it to a Linux computer. We installed [fastboot](https://grapheneos.org/install/cli#obtaining-fastboot) on the computer and used it to display details:

```
$ fastboot getvar all
[...]
(bootloader) slot-count:2
(bootloader) slot-fastboot-ok:a:yes
(bootloader) slot-fastboot-ok:b:yes
(bootloader) slot-retry-count:a:1
(bootloader) slot-retry-count:b:1
(bootloader) slot-successful:a:yes
(bootloader) slot-successful:b:yes
(bootloader) slot-suffixes:_a,_b
(bootloader) slot-unbootable:a:no
(bootloader) slot-unbootable:b:no
(bootloader) snapshot-update-status:merging
[...]
```

Modern Android smartphones [store two versions of their operating system](https://wiki.postmarketos.org/wiki/Android_AB_Slots). Their storage locations are called "Slot A" and "Slot B". This is to enable unsuccessful updates to be undone. But it apparently did not work in this case. Both slots were marked as `ok` and `successful`. None was marked as `unbootable`. The `retry-count`s remained unchanged even after several boot attempts.

However, a `merging` process had seemingly been interrupted during the power failure. The attempt to continue this process in _fastboot mode_ with `fastboot snapshot-update merge` failed with the error message `FAILED (remote: 'not allowed in locked state')`.

However, it worked in _fastbootd_ mode. To get there, go to _fastboot mode_, then navigate to _recovery mode_ using the volume keys and the power button. It welcomes us with the message "No command". The menu only appears when we briefly press the volume up button once while holding down the power button and then release both buttons. In the recovery menu, we select _Enter fastboot_. We are now in _fastbootd_ mode and can complete the interrupted process on the computer:

```
$ fastboot snapshot-update merge
```

After that, GrapheneOS succeeded to boot again. The user's data was also still there.
