# 2024-05-17 Die Fehlermeldung "Your device is corrupt" auf Google-Pixel-Smartphones mit GrapheneOS

Auf dem Smartphone eines Kunden, der bei uns ein [Google Pixel 6a mit vorinstalliertem GrapheneOS](https://shop.proxysto.re/category/6) und gesperrtem Bootloader gekauft hatte, erschien vor Kurzem die Fehlermeldung _"Your device is corrupt. It can't be trusted and may not work properly. Press power button to continue."_ Die Fehlermeldung ließ sich zwar durch Drücken des Einschaltknopfs überspringen, aber GrapheneOS startete nicht.

Die Fehlermeldung kann ein Zeichen einer Manipulation oder eines Hardwaredefekts sein. Im unserem Fall hatte sich das Smartphone allerdings über Nacht ausgeschaltet, weil der Akku leer war. Offenbar hatte der plötzliche Stromausfall die Installation eines Systemupdates unterbrochen.

Nachdem wir den Akku aufgeladen hatten, haben wir das Pixel 6a zunächst durch langes Drücken des Power-Buttons ausgeschaltet. Danach haben wir es mit `Power + Volume Down` in den _Fastboot Mode_ versetzt und an einen Linux-Computer angeschlossen. Auf dem Computer haben wir [fastboot installiert](https://grapheneos.org/install/cli#obtaining-fastboot) und uns damit Details anzeigen lassen:

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

Moderne Android-Smartphones [speichern zwei Versionen ihres Betriebssystems](https://wiki.postmarketos.org/wiki/Android_AB_Slots). Deren Speicherorte werden als "Slot A" und "Slot B" bezeichnet. Damit sollen erfolglose Updates rückgängig gemacht werden können. Das hat in diesem Fall anscheinend nicht funktioniert. Beide Slots waren als `ok` und `successful` gekennzeichnet, keiner als `unbootable`. Die `retry-count`s blieben auch nach mehreren Startversuchen unverändert.

Allerdings war bei dem Stromausfall offenbar ein `merging`-Vorgang unterbrochen worden. Der Versuch, diesen Vorgang im _Fastboot Mode_ mit `fastboot snapshot-update merge` fortzusetzen, scheiterte mit der Fehlermeldung `FAILED (remote: 'not allowed in locked state')`.

Im _fastbootd_-Modus funktionierte sie allerdings. Dazu muss im _Fastboot Mode_ zunächst mit den Lautstärketasten und dem Power-Button in den _Recovery Mode_ navigiert werden. Dieser begrüßt uns mit der Meldung "No command". Das Menü erscheint erst, wenn wir bei gedrückter Power-Taste einmal kurz auf die Volume-up-Taste drücken und danach beide Tasten loslassen. Im Recovery-Menü wählen wir _Enter fastboot_. Nun befinden wir uns im _fastbootd_-Modus und können am Computer den unterbrochenen Vorgang abschließen:

```
$ fastboot snapshot-update merge
```

Danach startete GrapheneOS auf dem Google Pixel 6a wieder. Auch die Daten des Benutzers waren noch da.
