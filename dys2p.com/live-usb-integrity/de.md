# Integritätsprüfung von USB-Sticks mit Linux

Wir bieten Live-USB-Sticks mit Linux-Betriebssystemen an, die direkt vom USB-Stick starten – zum Ausprobieren, zum Installieren oder zur alltäglichen Nutzung. Das Einrichten des USB-Sticks besteht auf unserer Seite darin, die entsprechende Imagedatei auf den USB-Stick zu schreiben. Vorher prüfen wir, dass die Imagedatei nicht beschädigt oder manipuliert wurde, indem wir eine Prüfsumme der Datei erzeugen und mit der Prüfsumme auf der Webseite der jeweiligen Linux-Distribution vergleichen.

Damit du uns nicht vertrauen musst, zeigen wir dir hier, wie du auch selbst überprüfen kannst, dass das vorinstallierte Betriebssystem nicht beschädigt oder manipuliert wurde. Die angegebenen Befehle sind für die Linux-Kommandozeile gedacht.

## Tails

Tails (_The Amnesic Incognito Live System_) ist ein Linux-Betriebssystem, das dich vor Überwachung und Zensur schützt. Tails wird direkt von einer DVD oder einem USB-Stick gestartet und hinterlässt keine ungewollten Spuren auf dem Computer. Zum Lesen können wir neben der [Tails-Website](https://tails.net) auch die [Tails-Broschüre von Capulcu](https://capulcu.blackblogs.org/neue-texte/bandi/) empfehlen.

Die Integritätsprüfung eines Tails-Sticks ist **nur von einem anderen Betriebssystem aus und nur vor dem ersten Start des Tails-Sticks** möglich, weil Tails beim ersten Start Veränderungen an den Daten auf dem USB-Stick vornimmt.

| Tails-Version | Befehl                                                                                        | Prüfsumme                                                                                                                                                                                                        |
| ------------- | --------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 6.1           | `sudo head --bytes 1433403392 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [87735e32da9de6592805427546eabb90ae3f52010fb4a4da18791fa630630b8e](https://gitlab.tails.boum.org/tails/tails/-/blob/dc2b1dc65925049878a87de4c8de63e84b3b562a/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.2           | `sudo head --bytes 1465909248 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [d93cd5220999d70c32b88f850437774b8cae39a4e9da4845e42a4f53e5e8e6cc](https://gitlab.tails.boum.org/tails/tails/-/blob/9023c841b196be0f92d65832d8e1dec7cc85bd19/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.3           | `sudo head --bytes 1474297856 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [81177ab73849b2a8d7a6d9a42867128f36be4fe5abd7920c126515be740eff23](https://gitlab.tails.boum.org/tails/tails/-/blob/36635499fc189d9dedc82d530402377226565d0b/wiki/src/install/v2/Tails/amd64/stable/latest.json) |
| 6.4           | `sudo head --bytes 1474297856 /dev/$(lsblk -ndo pkname /dev/disk/by-label/TAILS) | sha256sum` | [f8c36fad61a7f8c0fce45202369f85499a1c90f1bc7e5e5b320f2de1c3fa4e8d](https://gitlab.tails.boum.org/tails/tails/-/blob/75cb770ef136de54ee090e01432ce13e7a4cbd58/wiki/src/install/v2/Tails/amd64/stable/latest.json) |

## Ventoy

[Ventoy](https://www.ventoy.net/) ist _die_ Lösung, um mehrere Betriebssystem-Imagedateien auf einem USB-Stick zu speichern. Du kopierst die Imagedateien einfach als Dateien auf den Stick. Wenn du einen Computer von dem Ventoy-USB-Stick startest, wählst du eine Imagedatei aus und Ventoy startet sie. Auf der [Liste der unterstützten Imagedateien](https://www.ventoy.net/en/isolist.html) stehen viele Linux-Distributionen und mehrere Windows-Versionen. Bitte beachte, dass zum Starten von Ventoy _Secure Boot_ deaktiviert sein muss.

| Ventoy-Version | Befehl                                 | Prüfsumme                                                        |
| -------------- | -------------------------------------- | ---------------------------------------------------------------- |
| 1.0.98         | `sha256sum /dev/disk/by-label/VTOYEFI` | d2aad0774035fbd8d6de162424a971fb34336a185c89e2ce290861fc1cea2750 |

## Was machen die Befehle?

Die Befehle zur Integritätsprüfung bestehen aus diesen Einzelteilen:

* `lsblk -ndo pkname` ermittelt die Gerätedatei des USB-Sticks
* `head --bytes [Größe der Imagedatei]` liest so viele Bytes vom USB-Stick, wie die Imagedatei groß ist
* `sha256sum` berechnet die SHA256-Prüfsumme über den gelesenen Daten
* Wenn die Prüfsumme der Daten (vom USB-Stick) identisch mit der Prüfsumme der Imagedatei (von der Webseite) ist, dann ist davon auszugehen, dass die gelesenen Daten auch identisch mit der Imagedatei sind.
