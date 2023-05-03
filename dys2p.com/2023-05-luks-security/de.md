# Zur Sicherheit der Linux-Datenträgerverschlüsselung LUKS

_2023-05-03_

In den vergangenen Tagen gab es Unklarheiten und Sorgen zur Datenträgerverschlüsselung LUKS ("Linux Unified Key Setup"), die unter Linux weit verbreitet ist. Wir veröffentlichen hier unsere Einschätzung dazu.

## Inhaltsverzeichnis

* [Vorgeschichte](#background)
* [Diskussion auf Reddit](#discussion-reddit)
* [Schlüsselableitungsfunktionen](#kdf)
* [Angriffe auf verschlüsselte Daten(träger)](#attacks)
* [Empfehlungen](#recommendations)
  * [Upgraden verschlüsselter Laufwerke](#upgrade)
  * [Schlüsselableitungsfunktionen von Linux-Distributionen](#history)
  * [Neuinstallation statt Upgrade](#reinstall)
  * [Schutzmaßnahmen gegen Zugriffe auf entsperrte Geräte](#unlocked-devices)
  * [Schutzmaßnahmen gegen Angriffe auf Passwörter und Passphrasen](#passphrases)
* [Fazit](#conclusion)

<h2 id="background">Vorgeschichte</h2>

Am 17. April 2023 veröffentlichte Matthew "mjg59" Garrett den Aufruf, die Schlüsselableitungsfunktion (Key Derivation Function, KDF) von LUKS-verschlüsselten Datenträger zu ändern: [PSA: upgrade your LUKS key derivation function](https://mjg59.dreamwidth.org/66429.html). Darin bezieht er sich auf den Artikel "[Une lettre d’Ivan, enfermé à la prison de Villepinte : perquisitions et disques durs déchiffrés](https://nantes.indymedia.org/posts/87395/une-lettre-divan-enferme-a-la-prison-de-villepinte-perquisitions-et-disques-durs-dechiffres/)" (Crosspost von [paris-luttes.info](https://paris-luttes.info/une-lettre-d-ivan-enferme-a-la-16935)).

Demzufolge ist es gelungen, zwei verschlüsselte Computer eines Beschuldigten zu entschlüsseln – ein mit BitLocker verschlüsselter Windows-Computer und ein LUKS-verschlüsselter Computer mit Ubuntu 18. (Ob es sich um Ubuntu 18.04 oder 18.10 handelt, geht aus dem Text nicht hervor, ist aber auch nicht relevant.)

Der Windows-Computer sei von einem bootfähigen USB-Stick gestartet worden, dann sei mit der Software "AccessData FTK imager 3.3.05" (inzwischen "FTK® Imager von Exterro", welche AccessData im Jahr 2020 aufgekauft haben) eine Kopie der (verschlüsselten) Festplatte erstellt worden. Das Erstellen einer Kopie, um mit dieser weiterzuarbeiten, ist in der IT-Forensik normal und sagt nichts über die Sicherheit von Verschlüsselungsverfahren aus. Mit BitLocker möchten wir uns an dieser Stelle nicht weiter beschäftigen. Verwiesen sei aber auf Stolperfallen wie etwa Schwachstellen in Implementationen des Trusted Platform Module (TPM), die beispielsweise [TPM-Sniffing-Angriffe](https://pulsesecurity.co.nz/articles/TPM-sniffing) ermöglichen, oder Schlagzeilen wie [Windows 10: Microsoft kriegt Bitlocker-Nachschlüssel frei Cloud (2015)](https://www.heise.de/news/Windows-10-Microsoft-kriegt-Bitlocker-Nachschluessel-frei-Cloud-3056977.html).

Der zweite Computer, mit Ubuntu 18, sei mit LUKS1 verschlüsselt gewesen unter Verwendung eines Passworts aus mehr als zwanzig Zeichen, bestehend aus Buchstaben, Zahlen und Sonderzeichen. Garrett hält es für möglich, dass kein "opsec failure" vorliegt und der Grund für den erfolgreichen Zugriff auf den Ubuntu-Datenträger demzufolge die Datenträgerverschlüsselung LUKS1 und dessen Schlüsselableitungsfunktion PBKDF2 ist. Wir gehen weiter unten darauf ein.

Es könnte auch andere Gründe für die Entschlüsselung geben. Aus dem Text geht z. B. nicht hervor, in welchem Zustand der Ubuntu-Computer beschlagnahmt wurde. Eine Beschlagnahmung im eingeschalteten Zustand mit gesperrtem Bildschirm oder im Suspend-to-RAM ("Standby") ermöglicht das Auslesen der Schlüssel aus dem Arbeitsspeicher mittels einer [Cold Boot Attack](https://sarwiki.informatik.hu-berlin.de/Cold_Boot_Attack). Da die Entschlüsselung für die betroffene Person wohl überraschend kam, würden wir zumindest ausschließen, dass das Passwort irgendwo handschriftlich notiert war – eine von vielen Möglichkeiten, die wir später noch benennen.

<h2 id="discussion-reddit">Diskussion auf Reddit</h2>

Auf Reddit entwickelte sich [eine kurze Diskussion](https://www.reddit.com/r/linux/comments/12q51ce/comment/jgpvsqc/) zwischen dem LUKS-Entwickler Clemens Fruhwirth und Matthew "mjg59" Garrett (im Web Archive: [Fruhwirth](https://web.archive.org/web/20230420155735/https://www.reddit.com/r/linux/comments/12q51ce/comment/jgpvsqc/), [Garrett](https://web.archive.org/web/20230418093714/https://www.reddit.com/r/linux/comments/12q51ce/comment/jgq41pu/), [Fruhwirth](https://web.archive.org/web/20230418204541/https://www.reddit.com/r/linux/comments/12q51ce/comment/jgskur0/)).

Clemens Fruhwirth vermutet nicht LUKS1, sondern andere Gründe als Ursache. Selbst bei LUKS1 mit PBKDF2 und einem Passwort aus 13 zufälligen Groß- und Kleinbuchstaben und Ziffern würde ein Angriff mit der aktuellen Rechenkraft des gesamten Bitcoin-Netzwerks (unter der Annahme, dass es auf diesem Stand bleibt und es keinen weiteren technischen Fortschritt gibt) 77 Jahre dauern. Unter der Annahme, dass die zur Verfügung stehende Rechenleistung sich alle 24 Monate verdoppelt, wären es 10-12 Jahre. Garrett ist der Auffassung, das Passwörter selten wirklich zufällig sind und damit schwächer sind, als es theoretisch möglich wäre. Er merkt auch an, dass die Rechenleistung von Grafikkarten und von möglicher Spezialhardware (ASICs) schneller als prognostiziert zunimmt.

Anzumerken ist, dass Fruhwirth in seinem Rechenbeispiel von einem PBKDF2-Schwierigkeitsgrad von 3.000.000 Wiederholungen ausgeht. Bei Laufwerken, die vor vielen Jahren verschlüsselt wurden, kann der Schwierigkeitsgrad durchaus um den Faktor 20 bis 30 geringer sein. Mehr dazu im nächsten Abschnitt.

<h2 id="kdf">Schlüsselableitungsfunktionen</h2>

Zur Verschlüsselung von Datenträgern werden symmetrische Verschlüsselungsverfahren verwendet, beispielsweise AES ("Advanced Encryption Standard"). Symmetrische Verschlüsselungsverfahren nutzen zum Verschlüsseln und zum Entschlüsseln den selben kryptografischen Schlüssel. Dieser muss eine bestimmte Länge haben, beispielsweise 256 Bit.

In der Praxis wird eine Passphrase nicht direkt als kryptografischer Schlüssel verwendet. Stattdessen wird eine Schlüsselableitungsfunktion ausgewählt, die aus der Passphrase einen kryptografischen Schlüssel erzeugt. Das hat mehrere Vorteile:

* Die Passphrase muss nicht 256 Bit lang sein, sondern kann eine andere Länge haben.
* Die Passphrase kann geändert werden, ohne den ganzen Datenträger neu verschlüsseln zu müssen.
* Wir können den Zeitaufwand festlegen, der zum Erzeugen des kryptografischen Schlüssels nötig ist.

Der letzte Punkt ist hier relevant. Um eine Passphrase auszuprobieren, muss die Schlüsselableitungsfunktion ausgeführt werden. Je aufwändiger die verwendete Funktion ist, umso aufwändiger ist es dementsprechend auch, eine große Menge möglicher Passphrasen durchzuprobieren. Moderne Grafikkarten können tausende Rechenoperationen parallel ausführen. Deshalb werden mittlerweile Schlüsselableitungsfunktionen wie etwa **Argon2id** bevorzugt, die resistent gegen Angriffe mit spezieller Hardware (Grafikkarten oder ASICs) sind, indem sie nicht nur viel Rechenzeit, sondern auch viel Arbeitsspeicher benötigen und sich nicht gut parallel ausführen lassen.

Beim Anlegen eines verschlüsselten Datenträgers wird der Schwierigkeitsgrad meist so gewählt, dass wir etwa eine Sekunde lang warten müssen. Ein Datenträger, der vor vielen Jahren auf langsamer Hardware verschlüsselt wurde, wird daher eine ältere Schlüsselableitungsfunktion und einen geringeren Schwierigkeitsgrad verwenden als ein Datenträger, den wir heute auf einem modernen Computer verschlüsseln. Deshalb kann es sinnvoll sein, den Schwierigkeitsgrad der Funktion gelegentlich zu erhöhen oder auf eine modernere Funktion umzusteigen. Genau das empfiehlt Matthew Garrett in seinem Aufruf.

Das ältere LUKS1 unterstützt nur die Schlüsselableitungsfunktion [PBKDF2](https://web.archive.org/web/20170411220929/https://www.emc.com/collateral/white-papers/h11302-pkcs5v2-1-password-based-cryptography-standard-wp.pdf), die zwar eine einstellbare Rechenzeit, aber kaum Arbeitsspeicher benötigt. PBKDF2 ist lediglich eine einstellbare Anzahl Wiederholungen ("Iterationen") des SHA256- oder SHA512-Hashverfahrens. Das neuere LUKS2 verwendet standardmäßig die Funktion [Argon2](https://de.wikipedia.org/wiki/Argon2#Struktur) mit einstellbarer Rechenzeit und einstellbarem Arbeitsspeicherbedarf. Sowohl Garrett als auch beispielsweise das [OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) und [KeePassXC](https://keepassxc.org/docs/KeePassXC_UserGuide) empfehlen die Variante _Argon2id_. Auch das Forensik-Unternehmen elcomsoft [stellt fest](https://blog.elcomsoft.com/2022/08/probing-linux-disk-encryption-luks2-argon-2-and-gpu-acceleration/), dass LUKS2 Brute-Force-Angriffe mit Grafikkarten erheblich schwerer macht.

Auch mit LUKS1 und dessen Schlüsselableitungsfunktion PBKDF2 besteht aktuell kein Grund zur Sorge, wenn die verwendete Passphrase ausreichend komplex ist.

<h2 id="attacks">Angriffe auf verschlüsselte Daten(träger)</h2>

Einen Eindruck über die Methoden und Wege, Passwörter und Passphrasen zu erlangen bzw. zu umgehen, liefert beispielsweise das Kapitel "Analyzing Filesystem Encryption" des Buches "Practical Linux Forensics: A Guide for Digital Investigators" (übersetzt):

> Die Entschlüsselung geschützter Daten erfordert ein/e Passwort/Passphrase oder eine Kopie des kryptografischen Schlüssels (eine Zeichenkette oder Schlüsseldatei). Die forensische Herausforderung besteht darin, den Entschlüsselungsschlüssel zu finden. Einige bekannte Methoden zur Wiederherstellung von Passwörtern/Schlüsseln sind:
>
> - Bruteforce mit wörterbuchbasierten Angriffen zum Auffinden einfacher Passwörter
> - Bruteforce mit GPU-Clustern für die schnelle, umfassende Passwortsuche
> - Kryptoanalyse (mathematische Schwäche, Reduzierung des Schlüsselraums)
> - Auffinden von Passwörtern, die zuvor gespeichert, geschrieben oder übertragen wurden
> - Wiederverwendung von Passwörtern über mehrere Konten oder Geräte hinweg
> - Rechtliche Verpflichtung zur Vorlage von Passwörtern vor Gericht (Anmkerung: Gibt es in einigen Ländern.)
> - Kooperativer Systembesitzer oder Komplize mit dem Passwort
> - Schlüsselsicherung/-entnahme in Unternehmensumgebungen
> - Ausnutzung von Geräten, Schwachstellen oder Hintertüren
> - Keylogger oder Tastatursichtbarkeit (Videokameras oder Teleskop)
> - Regenbogen-Tabellen: Vorberechnete Tabelle mit kryptografischen Hashes
> - Schlüssel aus dem Speicher extrahieren: PCIbus-DMA-Angriffe, Hibernation
> - Man-in-the-Middle-Angriff (MITM-Angriff, auch Machine-in-the-Middle) auf den Netzwerkverkehr
> - Social Engineering
> - Erzwungene Nutzung biometrischer Daten
> - Folter, Erpressung oder Nötigung

Wie ist die Einschätzung anderer zu diesem Thema? Mikko Hyppönen schrieb in "If It's Smart, It's Vulnerable" im Kapitel "Cracking Passwords" (übersetzt):

> Die Behörden verfügen zu diesem Zweck über ziemlich beeindruckende Entschlüsselungssysteme. Ein Bürogebäude in Den Haag beispielsweise verfügt über Entschlüsselungshardware von der Größe eines Supercomputers, die ihr eigenes Kraftwerk benötigt. Mit einer solchen Hardware können Millionen von Passwortoptionen pro Sekunde getestet werden. Dennoch kann es Monate dauern, eine einzige verschlüsselte Datei [oder ein verschlüsseltes Laufwerk] zu öffnen.

> Automatisierte Entschlüsselungssysteme verwenden eine clevere Taktik, um diesen Vorgang zu beschleunigen. Wird eine kennwortgeschützte Datei auf der Festplatte eines Verdächtigen gefunden, werden alle Dateien auf dem Laufwerk indiziert und alle einzelnen Wörter aus jeder Datei gesammelt, um sie als Kennwörter zu testen. Wenn dies nicht funktioniert, werden alle gefundenen Wörter in umgekehrter Reihenfolge getestet und das Laufwerk nach ungenutzten Bereichen und gelöschten Dateien durchsucht, und die darin enthaltenen Wörter werden ausprobiert. In erstaunlich vielen Fällen lassen sich die Dateien auf diese Weise entschlüsseln.

Noch bessere Erfolgsaussichten biete jedoch eine andere Methode:

> Die beste Taktik besteht darin, einen Kriminellen abzulenken und ihn daran zu hindern, seine Geräte bei der Verhaftung zu zerstören oder zu sperren. Bei einer von EUROPOL koordinierten Verhaftung saß der Verdächtige in einem Café mit einem Laptop, als sich eine verdeckte Ermittlerin an denselben Tisch setzte und einige Augenblicke später ihren Kaffee verschüttete. Die Verdächtige stand auf, woraufhin ein männlicher Beamter, der hinter ihr wartete, vortrat und den unverschlossenen Computer zur forensischen Analyse mitnahm.

Auch der Laptop von Ross Ulbricht, dem Betreiber der Darknet-Plattform "Silk Road", [wurde im entsperrten Zustand beschlagnahmt](https://en.wikipedia.org/wiki/Ross_Ulbricht#Arrest). In einem anderen Fall wurde der Betreiber des Neonazi-Forums "Thiazi" mit einem Telefonanruf dazu gebracht, kurz vor dem Zugriff seinen Computer hochzufahren und zu entsperren.

Trotz allem gilt das Fazit des [OWASP Password Storage Cheat Sheets](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) (das sich zwar an Onlinedienste richtet, sich aber auf die Verschlüsselung gespeicherter Daten übertragen lässt):

> Starke Kennwörter, die mit modernen Hashing-Algorithmen und unter Verwendung bewährter Hashing-Verfahren gespeichert werden, sollten für einen Angreifer praktisch unmöglich zu knacken sein.

<h2 id="recommendations">Empfehlungen</h2>

Für verschlüsselte Datenträger, die sich in deinem Besitz befinden, teilen wir die Empfehlung von mjg59, auf LUKS2 zu upgraden und die Schlüsselableitungsfunktion zu _Argon2id_ zu ändern (übersetzt):

> Was die Sache noch schlimmer macht, ist die Tatsache, dass die Distributionen [die Schlüsselableitungsfunktion] im Allgemeinen in keiner Weise aktualisieren. Wenn Sie Ihr System installiert haben und pbkdf2 als KDF erhalten haben, verwenden Sie wahrscheinlich immer noch pbkdf2, selbst wenn Sie auf ein System aktualisiert haben, das bei einer Neuinstallation argon2id verwenden würde. Glücklicherweise kann das alles an Ort und Stelle behoben werden. Beachten Sie aber, dass Sie, wenn hier etwas schief geht, den Zugriff auf alle Ihre verschlüsselten Daten verlieren könnten. **Bevor Sie also irgendetwas tun, stellen Sie sicher, dass alle Daten gesichert sind (und finden Sie heraus, wie Sie das Backup sicher aufbewahren können, damit Ihre Daten nicht auf diese Weise beschlagnahmt werden).**

Für Datenträger, die sich außerhalb deiner Reichweite befinden, haben wir eine Übersicht über gängige Linux-Distributionen und die von ihnen verwendeten Schlüsselableitungsfunktionen aufgestellt.

<h3 id="upgrade">Upgraden verschlüsselter Laufwerke</h3>

mjg59 schreibt (übersetzt):

> Stellen Sie zunächst sicher, dass Sie eine möglichst aktuelle Version Ihrer Distribution verwenden. Wenn Sie Werkzeuge haben, die das LUKS2-Format unterstützen, bedeutet das nicht, dass Ihre Distribution das alles integriert hat, und alte Versionen Ihrer Distribution erlauben es Ihnen vielleicht, Ihr LUKS-Setup zu aktualisieren, ohne das ein Booten von diesem Format zu unterstützen. Wenn Sie außerdem ein verschlüsseltes `/boot` verwenden, sollten Sie jetzt damit aufhören - sehr aktuelle Versionen von grub2 unterstützen LUKS2, aber nicht argon2id, und das macht Ihr System unbootbar.
>
> Als nächstes müssen Sie herausfinden, welches Gerät unter /dev zu Ihrer verschlüsselten Partition gehört. Führen Sie folgene Eingabe aus
>
> `lsblk`
>
> und suchen Sie nach Einträgen, die den Typ "crypt" aufweisen. Das Gerät darüber in der Baumstruktur ist das eigentliche verschlüsselte Gerät. Notieren Sie diesen Namen und führen Sie
>
> `sudo cryptsetup luksHeaderBackup /dev/whatever --header-backup-file /tmp/luksheader`
>
> aus und kopieren Sie es auf einen USB-Stick oder ähnliches. Wenn hier etwas schief geht, können Sie ein Live-Image booten und Folgendes ausführen
>
> `sudo cryptsetup luksHeaderRestore /dev/whatever --header-backup-file luksheader`
>
> um es wiederherzustellen.
>
> (Ergänzung: Sobald alles funktioniert, löschen Sie diese Sicherung! Es enthält den alten schwachen Schlüssel, und jemand kann ihn möglicherweise dazu verwenden, Ihren Festplattenverschlüsselungsschlüssel mit dem alten KDF zu erzwingen, selbst wenn Sie das KDF auf der Festplatte aktualisiert haben).
>
> Als nächstes führen Sie folgendes aus
>
> `sudo cryptsetup luksDump /dev/whatever`
>
> und suchen Sie die Zeile `Version:`. Wenn es Version 1 ist, müssen Sie den Header auf LUKS2 aktualisieren. Führen Sie dazu aus
>
> `sudo cryptsetup convert /dev/whatever --type luks2`
>
> und folgen Sie den Eingabeaufforderungen. Vergewissern Sie sich, dass Ihr System noch bootet, und wenn nicht, gehen Sie zurück und stellen Sie die Sicherung Ihres Headers wieder her. Wenn zu diesem Zeitpunkt alles in Ordnung ist, führen Sie
>
> `sudo cryptsetup luksDump /dev/whatever`
>
> und suchen Sie nach der PBKDF: Zeile in jedem Keyslot (achten Sie nur auf die Keyslots, ignorieren Sie alle Verweise auf pbkdf2, die nach der `Digests:`-Zeile kommen). Wenn die PBKDF entweder "pbkdf2" oder "argon2i" lautet, sollten Sie sie in argon2id umwandeln. Führen Sie dazu Folgendes aus
>
> `sudo cryptsetup luksConvertKey /dev/whatever --pbkdf argon2id`
>
> und folgen Sie den Aufforderungen. Wenn Ihr Laufwerk mit mehreren Passwörtern verknüpft ist, haben Sie mehrere Keyslots, und Sie müssen diesen Vorgang für jedes Passwort wiederholen.

Das Upgrade des [Persistent Storage in Tails](https://tails.boum.org/doc/persistent_storage/) erfordert Administratorrechte, die du erhälst, indem du beim
Anmeldebildschirm ein [Administrationspasswort](https://tails.boum.org/doc/first_steps/welcome_screen/administration_password/index.de.html) vergibst. Außerdem darf beim Upgrade von LUKS1 auf LUKS2 der persistente Speicher nicht eingebunden sein.

<h3 id="history">Schlüsselableitungsfunktionen von Linux-Distributionen</h3>

Das Systemwerkzeug `cryptsetup` führte in [Version 2.0.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.0.0-ReleaseNotes) (veröffentlicht 2017-12-10) die Unterstützung für das LUKS2-Format ein und nutzt dafür standardmäßig _Argon2i_ mit 128 MB Arbeitsspeicherbedarf und 800 ms Rechenzeit (auf dem verwendeten Computer). In [Version 2.0.1](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.0.1-ReleaseNotes) (veröffentlicht 2018-01-21) wurde der Arbeitsspeicherbedarf auf 1 GiB erhöht. Das LUKS2-Format selbst ist jedoch erst seit [Version 2.1.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.1.0-ReleaseNotes) (2019-02-08) Standard. Seit [Version 2.4.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.4.0-ReleaseNotes) (veröffentlicht 2021-08-18) verwendet cryptsetup standardmäßig Argon2id. Linux-Distributionen enthalten allerdings oftmals nicht die neuesten Softwareversionen und verwenden auch nicht zwangsläufig die Standardoptionen.

Entscheidend ist der Zeitpunkt der Verschlüsselung bzw. die zur Verschlüsselung verwendete Version des Betriebssystems. Falls du beispielsweise einen Datenträger im Rahmen der Installation von Ubuntu 18.04 verschlüsselt hast und dein Betriebssystem später auf Ubuntu 22.04 aktualisiert hast, wird er wahrscheinlich mit dem älteren LUKS1 verschlüsselt sein. Die folgende Tabelle zeigt, welche Verschlüsselungsverfahren die **Installationsprogramme** bzw. Installationsanleitungen der untersuchten Linux-Distributionen verwenden.

<table class="table">
  <thead>
    <tr>
      <th>Linux-Distribution</th>
      <th>Version der Distribution</th>
      <th>Release Date</th>
      <th>Verfahren</th>
      <th>Diskussion</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Arch Linux</td>
      <td>
        Wiki ab <a href="https://wiki.archlinux.org/index.php?title=Dm-crypt/Device_encryption&amp;oldid=506702">2018-01-08</a>, standardmäßig ab 2019.03.01
      </td>
      <td>-</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>
        Wiki ab <a href="https://wiki.archlinux.org/index.php?title=Dm-crypt/Device_encryption&amp;oldid=692464">2021-08-22</a>, standardmäßig ab 2021.09.01
      </td>
      <td>-</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Debian</td>
      <td>
        <a href="https://www.debian.org/releases/buster/amd64/release-notes/ch-whats-new.en.html#cryptsetup-luks2">ab 10</a>
      </td>
      <td>2019-07-06</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>11.6.0</td>
      <td>2022-12-17</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td>Fedora Linux</td>
      <td>
        <a href="https://fedoraproject.org/wiki/Releases/30/ChangeSet#Switch_cryptsetup_default_metadata_format_to_LUKS2">ab 30</a>
      </td>
      <td>2019-04-30</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>38</td>
      <td>2023-04-18</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Kali Linux</td>
      <td>2023.1</td>
      <td>2023-03-13</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Linux Mint</td>
      <td>20</td>
      <td>2020-06-27</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>21.1</td>
      <td>2022-12-20</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Manjaro</td>
      <td>22.1.0 (getestet mit manjaro-xfce-22.1.0-230421)</td>
      <td>2023-04-21</td>
      <td>LUKS1 mit PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td>openSUSE</td>
      <td>Leap 15.4</td>
      <td>2022-06-08</td>
      <td>LUKS1 mit PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>Tumbleweed Snapshot 20230501</td>
      <td>2023-05-01</td>
      <td>LUKS1 mit PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td>Pop!_OS</td>
      <td>22.04</td>
      <td>2022-04-25</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Qubes OS</td>
      <td>
        <a href="https://www.qubes-os.org/doc/upgrade/4.1/">ab 4.1</a>
      </td>
      <td>2022-02-04</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>4.1.2</td>
      <td>2023-03-14</td>
      <td>LUKS2 mit Argon2i</td>
      <td>
        <a href="https://forum.qubes-os.org/t/18071">Qubes-Forum</a>, <a href="https://www.reddit.com/r/Qubes/comments/12r4w1v/">Reddit</a>
      </td>
    </tr>
    <tr>
      <td>Tails</td>
      <td>5.12</td>
      <td>2023-04-19</td>
      <td>LUKS1 mit PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>5.13 (geplant)</td>
      <td>2023-05</td>
      <td>LUKS2</td>
      <td>
        <a href="https://gitlab.tails.boum.org/tails/tails/-/issues/19615">Issue 19615</a>, <a href="https://gitlab.tails.boum.org/tails/tails/-/issues/19633">Issue 19633</a>
      </td>
    </tr>
    <tr>
      <td>Ubuntu Linux</td>
      <td>18.04 LTS</td>
      <td>2018-04-23</td>
      <td>LUKS1 mit PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>19.04</td>
      <td>2019-04-18</td>
      <td>LUKS2 mit Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>22.04.2 LTS</td>
      <td>2023-02-23</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>23.04</td>
      <td>2023-04-20</td>
      <td>LUKS2 mit Argon2id</td>
      <td></td>
    </tr>
  </tbody>
</table>

<h3 id="reinstall">Neuinstallation statt Upgrade</h3>

Wie auch Garrett andeutet, ist es möglich, dass Linux-Installationen ein Upgrade auf LUKS2 oder den Umstieg auf Argon2id nicht unterstützen bzw. danach nicht mehr starten können. Um dieses Risiko auszuschließen, kannst du auch deine Daten verschlüsselt sichern und dein Betriebssystem neu aufsetzen. (Ob sich das lohnt, siehst du in der obigen Tabelle. Falls du Tails verwendest, warte damit noch bis zum kommenden Release.) Auch von verschlüsselten Datenträgern, von denen kein Betriebssystem startet, solltest du vor dem Upgrade ein Backup machen. Bei der Gelegenheit kannst du auch nicht mehr benötigte Daten löschen – denn nicht vorhandene Daten können auch nicht entschlüsselt werden.

Für jene, die sich nicht allein auf ein Verfahren verlassen möchten, empfehlen wir: Verschlüsselt eure Laufwerk zuerst mit LUKS2 und speichert besonders schützenswerte Inhalte in zusätzlichen [VeraCrypt-Containern](https://veracrypt.fr/en/VeraCrypt%20Volume.html) mit einem anderen sicheren Passwort und einem anderen Verschlüsselungsverfahren. Dadurch sind eure Daten auch dann noch geschützt, wenn Schwachstellen in einem Verfahren entdeckt werden. Allerdings nutzt VeraCrypt derzeit auch noch [PBKDF2 als Schlüsselableitungsfunktion](https://github.com/veracrypt/VeraCrypt/issues/127).

<h3 id="unlocked-devices">Schutzmaßnahmen gegen Zugriffe auf entsperrte Geräte</h3>

Eine solche zweistufige Verschlüsselung kann auch gegen Zugriffe auf entsperrte Geräte helfen, indem besonders schützenswerte Inhalte seperat verschlüsselt sind und nur entschlüsselt werden, wenn es erforderlich ist.

Ebenfalls nützlich kann es sein, einen Computer z. B. durch den Kippschalter einer Steckdosenleiste oder das Ziehen des Netzstecker und der damit verbunden Unterbrechung der Stromversorgung auszuschalten. Bei Laptops sollte natürlich zuvor der Akku entfernt werden. Eine andere Möglichkeit, einen Computer schnell auszuschalten oder sogar zu löschen, bieten Geräte wie [BusKill](https://www.buskill.in/luks-self-destruct/).

<h3 id="passphrases">Schutzmaßnahmen gegen Angriffe auf Passwörter und Passphrasen</h3>

Den oben genannten Angriffen auf verschlüsselte Daten(träger) stehen einige Gegenmaßnahmen gegenüber. Verwendete Verfahren und Software sollte möglichst auf dem aktuellen Stand sein (z. B. LUKS2 mit Argon2id). Daneben können insbesondere bei der Wahl von und dem Umgang mit Passwörtern bzw. Passphrasen Fehler vermieden werden.

Wähle ein Passwort bzw. eine Passphrase mit hoher [Entropie](https://de.wikipedia.org/wiki/Entropie_(Informationstheorie)) – oder mit anderen Worten: so dass genügend mögliche Passwörter bzw. Passphrasen in Frage kommen.

Dafür bieten sich Verfahren wie [Diceware](https://de.wikipedia.org/wiki/Diceware) an. Die Würfel sorgen für den notwendigen Zufall, während sich die Wörter mittels Eselsbrücken gut merken lassen. Die Electronic Frontier Foundation (EFF) empfiehlt in [Creating Strong Passwords Using Dice](https://ssd.eff.org/module/creating-strong-passwords#creating-strong-passwords-using-dice) die Verwendung von mindestens sechs zufällig ausgewählten Wörtern aus einer 7776 Wörter umfassenden Liste. Die EFF stellt dazu eine englischsprachige Wortliste bereit. Wir haben eine [deutschsprachige Diceware-Wortliste "de-7776"](https://github.com/dys2p/wordlists-de) zusammengestellt, die demnächst noch einmal überarbeitet wird.

Führe das Verfahren am besten offline mit Würfeln durch. Ein Beispiel für sechs Wörter mit jeweils fünf Würfeln:

1. 36433 jugend
2. 64355 weinkarte
3. 36541 kampagne
4. 63656 wahnsinn
5. 23654 entkernen
6. 34466 hemmschwelle

Die im Beispiel erzeugte Passphrase lautet demnach: "jugend weinkarte kampagne wahnsinn entkernen hemmschwelle".

Die Wortlisten sind 7776 Wörter lang, weil beim Werfen von fünf sechsseitige Würfeln `6^5 = 7776` verschiedene Ergebnisse möglich sind. Jedes Wort hat eine Entropie von `log2(7776) ≈ 13 Bit`. Eine Passphrase aus sechs ausgewürfelten Wörtern ist folglich `5 × 13 ≈ 78 Bit` stark. Es gibt `2^78 ≈ 3 × 10^23` verschiedene mögliche Passphrasen aus sechs Wörtern. Ein Passwort aus dreizehn _zufällig_ ausgewählten Groß- und Kleinbuchstaben sowie Ziffern hat ebenfalls eine Entropie von 78 Bit: eines von 62 möglichen Zeichen bringt `log2(62) ≈ 6 Bit` Entropie, dreizehn zufällige Zeichen demzufolge 78 Bit.

Da die Schlüsselableitungsfunktion PBKDF2 und das Bitcoin-Netzwerk den gleichen bzw. einen ähnlichen Hash-Algorithmus verwenden, lassen sich ihre Leistungen gut vergleichen. Das Bitcoin-Netzwerk, das in der Diskussion auf Reddit als Referenz herangezogen wurde, berechnet derzeit `3 × 10^20` Hashes pro Sekunde. Für eine PBKDF2-Schlüsselableitung mit 100.000 Iterationen könnte das Netzwerk damit `3 × 10^15` Passwörter bzw. Passphrasen pro Sekunde ausprobieren. Um alle möglichen Passphrasen aus sechs Wörtern auszuprobieren, bräuchte es bei konstanter Rechenleistung folglich `(3 × 10^23) / (3 × 10^15) = 10^8` Sekunden, oder rund drei Jahre. Im Mittel muss die Hälfte aller Möglichkeiten ausprobiert werden, was rund eineinhalb Jahre dauern würde. Bei einer Passphrase aus fünf statt sechs Wörtern wären es im Durchschnitt nur knapp zwei Stunden. Allerdings gehen wir davon aus, dass die meisten Angreifer nicht über eine derart hohe Rechenleistung verfügen – erst recht nicht für einen langen Zeitraum. Mit Argon2 statt PBKDF dürften beide Angriffe um mehrere Größenordnungen schwerer sein.

Bedenke, dass Passwörter wie `IgjAu23UiB!`, die sich beispielsweise aus den Anfangsbuchstaben eines Satzes ("Ich gehe jeden Abend um 23 Uhr ins Bett!") zusammensetzen, nicht zufällig sind.

Der Onlinedienst 1Password hat im Jahr 2021 Beispiele für Passwörter und den Aufwand, um diese zu knacken, aufgestellt: [How strong should your account password be? Here's what we learned](https://blog.1password.com/cracking-challenge-update/). Die genannten Kosten könnten jedoch schon veraltet sein. Weitere Texte zur Wahl von Passwörtern finden sich u. a. bei [digitalcourage.de](https://digitalcourage.de/digitale-selbstverteidigung/sicherheit-beginnt-mit-starken-passwoertern) und [kuketz-blog.de](https://www.kuketz-blog.de/sicheres-passwort-waehlen-der-zufall-entscheidet/). Die FAQ von [cryptsetup](https://gitlab.com/cryptsetup/cryptsetup/-/wikis/requentlyAskedQuestions#5-security-aspects) bieten ebenfalls Informationen und Empfehlungen.

Für den Umgang mit Passwörtern bzw. Passphrasen empfehlen wir:

* Verwende ein Passwort bzw. eine Passphrase und Abwandlungen davon für nichts anderes, sondern nur für ein Gerät und Anwendungsfall.
* Schreibe ein Passwort bzw. eine Passphrase und Abwandlungen davon nirgends auf.
* Teile ein Passwort bzw. eine Passphrase und Abwandlungen davon niemanden mit.
* Achte darauf, dass dich bei der Eingabe niemand heimlich beobachten oder filmen kann.
* Verwende keine biometrischen Informationen.

Möglicherweise begegnen dir in diesem Zusammenhang auch Verfahren der "glaubhaften Abstreitbarkeit", um das Vorhandensein verschlüsselter Daten zu verschleiern. VeraCrypt bietet dafür die Funktionen [Hidden Volume](https://www.veracrypt.fr/en/Hidden%20Volume.html) und [Hidden Operating System](https://www.veracrypt.fr/en/Hidden%20Operating%20System.html). Bedenken und Nacheile dazu werden u. a. hier benannt:

* [Bruce Schneier: Defeating Encrypted and Deniable File Systems: TrueCrypt v5.1a and the Case of the Tattling OS and Applications](https://www.schneier.com/wp-content/uploads/2016/02/paper-truecrypt-dfs.pdf)
* [Cryptsetup: 5.18 What about Plausible Deniability?](https://gitlab.com/cryptsetup/cryptsetup/-/wikis/FrequentlyAskedQuestions#5-security-aspects)
* [Kedziora, Chow, Susilo: Defeating Plausible Deniability of VeraCrypt Hidden Operating Systems (2017)](https://ro.uow.edu.au/cgi/viewcontent.cgi?article=1542&context=eispapers1)

<h2 id="conclusion">Fazit</h2>

Unserer Einschätzung nach kann, basierend auf den vorhandenen Informationen, nicht ausschließlich PBKDF2 und LUKS1 für die Entschlüsselung verantwortlich gemacht werden. Es ist wichtig, ein starkes Passwort bzw. Passphrase zu verwenden und die beschriebenen Empfehlungen im Umgang mit ihnen zu beachten. Die Verwendung aktueller Software und Verfahren erhöht die Sicherheit und reduziert eventuelle Angriffsflächen.

Das Upgrade der Schlüsselableitungsfunktion ist keine einmalige Aufgabe. Du solltest alle paar Jahre prüfen, ob die verwendete Funktion noch als sicher eingeschätzt wird, und den Schwierigkeitsgrad anpassen. Dies betrifft nicht nur LUKS, sondern alle passwortbasierten Verschlüsselungswerkzeuge wie etwa VeraCrypt oder KeePassXC.

## Weitere Diskussionen und Texte

* [Heise: LUKS: Alte verschlüsselte Container unsicher? Ein Ratgeber für Updates](https://www.heise.de/news/Alte-LUKS-Container-unsicher-Ein-kleiner-Update-Ratgeber-8981054.html)
* [Systemli: Ist die Linux-Festplattenverschlüsselung geknackt?](https://www.systemli.org/2023/04/30/ist-die-linux-festplattenshyverschl%C3%BCsselung-geknackt/)
* [Diskussion auf Hacker News](https://news.ycombinator.com/item?id=35611425)
* [Diskussion auf lobste.rs](https://lobste.rs/s/ik7j1s/psa_upgrade_your_luks_key_derivation)
* [Diskussion auf lwn.net](https://lwn.net/Articles/929343/)
