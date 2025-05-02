# Integritätsprüfung von USB-Sticks mit Linux

Wir bieten Live-USB-Sticks mit Linux-Betriebssystemen an, die direkt vom USB-Stick starten – zum Ausprobieren, zum Installieren oder zur alltäglichen Nutzung. Das Einrichten des USB-Sticks besteht auf unserer Seite darin, die entsprechende Imagedatei auf den USB-Stick zu schreiben. Vorher prüfen wir, dass die Imagedatei nicht beschädigt oder manipuliert wurde, indem wir eine Prüfsumme der Datei erzeugen und mit der Prüfsumme auf der Webseite der jeweiligen Linux-Distribution vergleichen.

Damit du uns nicht vertrauen musst, zeigen wir dir hier, wie du auch selbst überprüfen kannst, dass das vorinstallierte Betriebssystem nicht beschädigt oder manipuliert wurde. Die angegebenen Befehle sind für die Linux-Kommandozeile gedacht.

## Tails

Tails (_The Amnesic Incognito Live System_) ist ein Linux-Betriebssystem, das dich vor Überwachung und Zensur schützt. Tails wird direkt von einer DVD oder einem USB-Stick gestartet und hinterlässt keine ungewollten Spuren auf dem Computer. Zum Lesen können wir neben der [Tails-Website](https://tails.net) auch die [Tails-Broschüre von Capulcu](https://capulcu.blackblogs.org/neue-texte/bandi/) empfehlen.

Die Integritätsprüfung eines Tails-Sticks ist **nur von einem anderen Betriebssystem aus und nur vor dem ersten Start des Tails-Sticks** möglich, weil Tails beim ersten Start Veränderungen an den Daten auf dem USB-Stick vornimmt.

| Tails-Version | Befehl                                                                                        | Prüfsumme                                                                                                                                                                                                        |
| ------------- | --------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 6.15          | `sudo head --bytes 1589641216 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [f1b9dedda86c26fa054647d9b59cf4443d6f65beaf7755c0e2213b9a4149ce03](https://gitlab.tails.boum.org/tails/tails/-/blob/8bfaa9e4858402012c570e284d16018587e821ed/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.13          | `sudo head --bytes 1589641216 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [91e6a25d9e2b010e5e77c65ecb3adf760785b243f4d64b323012f13460db17e9](https://gitlab.tails.boum.org/tails/tails/-/blob/220dab2ff633206bbb13c249523749922324a993/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.12          | `sudo head --bytes 1589641216 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [9956584cd462380b6cab2da8145279c640a56042a123a7038bbcece9122a7668](https://gitlab.tails.boum.org/tails/tails/-/blob/221e2ca5da049d46ed410ea47fdee9483e7558bc/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.11          | `sudo head --bytes 1589641216 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [230dcd717faaf225d5a738e72fcacdfad6aca05064a2069d85f2ef2e77e5c9f5](https://gitlab.tails.boum.org/tails/tails/-/blob/3db741ac0dcdc5a326a8f2d21bbfd5d6e288e651/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.10          | `sudo head --bytes 1581252608 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [46ff2ce0f3b9d3e64df95c4371601a70c78c1bc4e2977741419593ce14a810a7](https://gitlab.tails.boum.org/tails/tails/-/blob/628804baab28218e09a4b426dfa37f51eaca19db/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.8.1         | `sudo head --bytes 1556086784 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [e9a924b8a13c981d260c08f490a1e101303f17980353c02ad145682138a623d7](https://gitlab.tails.boum.org/tails/tails/-/blob/847a25d0e58636817d5c407b2f6c30096f1860d9/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.7           | `sudo head --bytes 1556086784 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [9417315edc4aa8db7b3f43d8e1e7245ef477353e2b8f942da8ab3e01ac6509cc](https://gitlab.tails.boum.org/tails/tails/-/blob/d03f2c15a177f900727f05db49678335f0edf636/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.6           | `sudo head --bytes 1540358144 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [fe0e7b24e4a1400843eb25e483858378693647e2ab4521c684826e3be1704369](https://gitlab.tails.boum.org/tails/tails/-/blob/209bd74c1f751281390d8e448081ea381cc761c3/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.5           | `sudo head --bytes 1474297856 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [f417b60f5d291ebe7e6c7126e5573d5cd6c8fef78f5fae88a394733d850700aa](https://gitlab.tails.boum.org/tails/tails/-/blob/9a29446af4079e70b9cd81cb2cb358ec21a55fb9/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.4           | `sudo head --bytes 1474297856 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [f8c36fad61a7f8c0fce45202369f85499a1c90f1bc7e5e5b320f2de1c3fa4e8d](https://gitlab.tails.boum.org/tails/tails/-/blob/75cb770ef136de54ee090e01432ce13e7a4cbd58/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.3           | `sudo head --bytes 1474297856 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [81177ab73849b2a8d7a6d9a42867128f36be4fe5abd7920c126515be740eff23](https://gitlab.tails.boum.org/tails/tails/-/blob/36635499fc189d9dedc82d530402377226565d0b/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.2           | `sudo head --bytes 1465909248 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [d93cd5220999d70c32b88f850437774b8cae39a4e9da4845e42a4f53e5e8e6cc](https://gitlab.tails.boum.org/tails/tails/-/blob/9023c841b196be0f92d65832d8e1dec7cc85bd19/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.1           | `sudo head --bytes 1433403392 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [87735e32da9de6592805427546eabb90ae3f52010fb4a4da18791fa630630b8e](https://gitlab.tails.boum.org/tails/tails/-/blob/dc2b1dc65925049878a87de4c8de63e84b3b562a/wiki/src/install/v2/Tails/amd64/stable/latest.json) |

## Was machen die Befehle?

Die Befehle zur Integritätsprüfung bestehen aus diesen Einzelteilen:

* `lsblk -ndo pkname` ermittelt die Gerätedatei des USB-Sticks
* `head --bytes [Größe der Imagedatei]` liest so viele Bytes vom USB-Stick, wie die Imagedatei groß ist
* `sha256sum` berechnet die SHA256-Prüfsumme über den gelesenen Daten
* Wenn die Prüfsumme der Daten (vom USB-Stick) identisch mit der Prüfsumme der Imagedatei (von der Webseite) ist, dann ist davon auszugehen, dass die gelesenen Daten auch identisch mit der Imagedatei sind.
