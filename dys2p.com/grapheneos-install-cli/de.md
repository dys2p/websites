<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item"><a href="grapheneos-install.html">Installation</a></li>
		<li class="breadcrumb-item active" aria-current="page">CLI-Installationsanleitung</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/grapheneos.org/blob/73b2e2cda384310f4c26ebff049afebb31906a37/static/install/cli.html">73b2e2c</a> vom 2025-11-28. Falls Sie Hinweise oder Verbesserungsvorschläge haben, dann <a href="contact.html">schreiben Sie uns gerne</a> oder arbeiten Sie mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-install-cli/de.md">GitHub</a> an dieser Übersetzung.
</div>

<!--
Copyright © 2014-2025 GrapheneOS

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

https://grapheneos.org/LICENSE.txt
-->

<h1 id="cli-install">CLI-Installationsanleitung</h1>

Dies ist eine Anleitung für die Installation von GrapheneOS auf den [offiziell unterstützten Geräten](https://grapheneos.org/faq#supported-devices). Sie kann sowohl für die [offiziellen Versionen](https://grapheneos.org/releases) als auch für [eigene Builds](https://grapheneos.org/build) verwendet werden. Der [Web-Installer](https://grapheneos.org/install/web) [[deutsche Übersetzung](grapheneos-install-web.html)] ist eine einfachere Methode zur Installation der offiziellen Versionen über einen Browser mit WebUSB-Unterstützung.

Wir empfehlen dringend, diesen offiziellen Anweisungen zu folgen. Für die offizielle Anleitung wurde viel gemeinsame Arbeit in die Abdeckung aller Grenzfälle gesteckt. Sie wird regelmäßig von vielen Personen auf allen unterstützten Betriebssystemen getestet. Das Befolgen dieser Anweisungen, ohne irgendwelche Schritte zu überspringen, umzuordnen oder hinzuzufügen, gibt Ihnen eine ordentliche GrapheneOS-Installation, außer, es gibt ein Hardwareproblem. Wir raten dringend davon ab, inoffiziellen Anleitungen zu folgen, die in irgendeiner Weise von den offiziellen Anleitungen abweichen.

Wenn Sie Probleme mit dem Installationsprozess haben, fragen Sie im [offiziellen GrapheneOS-Chat](https://grapheneos.org/contact#community) um Hilfe. Es gibt fast immer Leute, die bereit sind zu helfen. Bevor Sie um Hilfe bitten, versuchen Sie, der Anleitung selbst zu folgen und bitten Sie erst dann um Hilfe, wenn Sie selbst nicht weiterkommen.

<h2 id="table-of-contents">Inhaltsverzeichnis</h2>

<nav>
  <ul>
    <li>
      <a href="#prerequisites">Voraussetzungen</a>
    </li>
    <li>
      <a href="#enabling-oem-unlocking">OEM-Entsperrung aktivieren</a>
    </li>
    <li>
      <a href="#opening-terminal">Ein Terminal öffnen</a>
    </li>
    <li>
      <a href="#obtaining-fastboot">fastboot beschaffen</a>
    </li>
      <ul>
        <li>
          <a href="#standalone-platform-tools">Standalone Platform Tools</a>
        </li>
      </ul>
    <li>
      <a href="#checking-fastboot-version">Überprüfen der Fastboot-Version</a>
    </li>
    <li>
      <a href="#flashing-as-non-root">Flash als Nicht-Root-Benutzer</a>
    </li>
    <li>
      <a href="#working-around-fwupd-bugs-on-linux-distributions">Umgehen der fwupd-Bugs auf Linux-Distributionen</a>
    </li>
    <li>
      <a href="#booting-into-the-bootloader-interface">Booten in den Bootloader-Modus</a>
    </li>
    <li>
      <a href="#connecting-phone">Das Gerät verbinden</a>
    </li>
    <li>
      <a href="#unlocking-the-bootloader">Den Bootloader entsperren</a>
    </li>
    <li>
      <a href="#obtaining-openssh">OpenSSH beschaffen</a>
    </li>
    <li>
      <a href="#obtaining-factory-images">Factory Images beschaffen</a>
    </li>
    <li>
      <a href="#flashing-factory-images">Factory Images flashen</a>
    </li>
      <ul>
        <li>
          <a href="#troubleshooting">Fehlersuche</a>
        </li>
      </ul>
    <li>
      <a href="#locking-the-bootloader">Den Bootloader sperren</a>
    </li>
    <li>
      <a href="#post-installation">Nach der Installation</a>
    </li>
      <ul>
        <li>
          <a href="#booting">Booten</a>
        </li>
        <li>
          <a href="#disabling-oem-unlocking">OEM-Entsperrung deaktivieren</a>
        </li>
        <li>
          <a href="#verifying-installation">Überprüfung der Installation</a>
        </li>
      	  <ul>
            <li>
              <a href="#verified-boot-key-hash">Hash des öffentlichen Verified-Boot-Schlüssels</a>
            </li>
            <li>
              <a href="#hardware-based-attestation">Hardwaregestützte Attestierung</a>
            </li>
      	  </ul>
        <li>
          <a href="#further-information">Weitere Informationen</a>
        </li>
        <li>
          <a href="#replacing-grapheneos-with-the-stock-os">GrapheneOS durch das Standardbetriebssystem ersetzen</a>
        </li>
      </ul>
  </ul>
</nav>

<h2 id="prerequisites">Voraussetzungen</h2>

Für den CLI-Installationsprozess benötigen Sie einen Computer mit mindestens 2 GB freiem Arbeitsspeicher und 32 GB freiem Speicherplatz. Die Web-Installation kann im Gegensatz zur CLI-Installation auf einem Android-Smartphone oder -Tablet ausgeführt werden.

Sie benötigen ein USB-Kabel, um das Gerät an einem Computer anzuschließen, der die Installation durchführt. Verwenden Sie nach Möglichkeit das hochwertige, standardkonforme USB-C-Kabel, das dem Gerät beiliegt. Wenn Ihr Computer keine USB-C-Anschlüsse hat, benötigen Sie ein hochwertiges USB-C-zu-USB-A-Kabel. Sie sollten vermeiden, einen USB-Hub zu verwenden, wie z. B. an der Vorderseite eines Desktop-Computergehäuses. Schließen Sie das Kabel direkt an einen rückwärtigen Anschluss eines Desktops oder an die Anschlüsse eines Laptops an. Viele weit verbreitete USB-Kabel und -Hubs sind kaputt und die häufigste Ursache für Probleme bei der Installation von GrapheneOS.

Die Installation von einem Betriebssystem in einer virtuellen Maschine wird nicht empfohlen. USB-Passthrough ist oft nicht zuverlässig. Um diese Probleme auszuschließen, installieren Sie von einem Betriebssystem, das direkt auf einem Gerät läuft. Virtuelle Maschinen sind außerdem oft so konfiguriert, dass sie sehr begrenzten Arbeitsspeicher und Speicherplatz haben.

Offiziell unterstützte Betriebssysteme für die CLI-Installationsmethode:

* Windows 10
* Windows 11
* macOS Sonoma (14)
* macOS Sequoia (15)
* macOS Tahoe (26)
* Arch Linux
* Debian 12 (Bookworm)
* Debian 13 (Trixie)
* Ubuntu 22.04 LTS
* Ubuntu 24.04 LTS
* Ubuntu 25.04
* Linux Mint 21 (nutzen Sie die Anweisungen für Ubuntu 22.04 LTS)
* Linux Mint 22 (nutzen Sie die Anweisungen für Ubuntu 24.04 LTS)
* Linux Mint Debian Edition 6 (nutzen Sie die Anweisungen für Debian 12)

Auch ältere, nicht mehr weiterentwickelte Versionen dieser Betriebssysteme lassen sich einsetzen, allerdings erhalten sie keinen offiziellen Support.

Stellen Sie sicher, dass Ihr Betriebssystem auf dem neuesten Stand ist, bevor Sie fortfahren.

Der [Web-Installer](https://grapheneos.org/install/web) [[deutsche Übersetzung](grapheneos-install-web.html)] ist portabler und kann von Android, ChromeOS und GrapheneOS selbst verwendet werden, da er überall ausgeführt werden kann, wo ein Browser mit funktionierender WebUSB-Unterstützung zur Verfügung steht.

Sie benötigen eines der [offiziell unterstützten Geräte](https://grapheneos.org/faq#supported-devices). Um sicherzustellen, dass das Gerät für die Installation von GrapheneOS freigeschaltet werden kann, sollten Sie Netzbetreiber-spezifische Geräte vermeiden. Netzbetreiber-spezifische Varianten der Pixel-Smartphones verwenden dasselbe Stock-Betriebssystem und dieselbe Firmware, aber mit einer Carrier-ID ungleich Null, die bei der Herstellung auf die persistente Partition geflasht wurde. Die Carrier-ID aktiviert eine Netzbetreiber-spezifische Konfiguration im Stock-Betriebssystem, einschließlich einer Deaktivierung der Netzbetreiber- und Bootloader-Entsperrung. Der Netzbetreiber kann dies möglicherweise aus der Ferne deaktivieren, aber seine Support-Mitarbeiter wissen das möglicherweise nicht und werden es wahrscheinlich auch nicht tun. Holen Sie sich ein Netzbetreiber-unabhängiges Gerät, um das Risiko und möglichen Ärger zu vermeiden. Wenn Sie einen Weg finden _können_, ein Netzbetreiber-spezifisches Gerät zu entsperren, ist das kein Problem, da GrapheneOS die Carrier-ID einfach ignorieren kann und die Hardware die gleiche ist.

Es hat sich bewährt, das Gerät vor der Installation von GrapheneOS zu aktualisieren, um beim Anschluss des Gerätes an den Computer und in der frühen Phase des Installationsprozesses die neueste Firmware zu haben. So oder so, GrapheneOS flasht die neueste Firmware zu Beginn des Installationsprozesses.

<h2 id="enabling-oem-unlocking">OEM-Entsperrung aktivieren</h2>

Die OEM-Entsperrung muss innerhalb des Betriebssystems aktiviert werden.

Aktivieren Sie das Menü „Entwickleroptionen“, indem Sie zu <b>Einstellungen&#160;<span aria-label="and then">></span> Über das Telefon/Tablet</b> gehen und wiederholt auf den Menüeintrag <b>Build-Nummer</b> drücken, bis die Entwickleroptionen aktiviert sind.

Gehen Sie dann zu <b>Einstellungen&#160;<span aria-label="and then">></span> System&#160;<span aria-label="and then">></span> Entwickleroptionen</b> und aktivieren Sie die Einstellung <b>OEM-Entsperrung</b>. Bei Gerätevarianten (SKUs), die von Netzbetreibern als gesperrte Geräte verkauft werden können, erfordert die Aktivierung der <b>OEM-Entsperrung</b> einen Internetzugang, damit das Stock-Betriebssystem prüfen kann, ob das Gerät von einem Netzbetreiber als gesperrt verkauft wurde.

Für das Pixel 6a wird die OEM-Entsperrung nicht mit der ab Werk installierten Version des Stock-Betriebssystems funktionieren. Sie müssen es mit einem Over-the-Air-Update auf die Version von Juni 2022 oder später aktualisieren. Nach der Aktualisierung müssen Sie das Gerät auf die Werkseinstellungen zurücksetzen, um die OEM-Entsperrung zu ermöglichen.

<h2 id="opening-terminal">Ein Terminal öffnen</h2>

Diese Anleitung verwendet Kommandozeilen-Werkzeuge. Starten Sie das Terminal wie jede andere Anwendung. Falls Sie Windows verwenden, starten Sie eine normale Nicht-Administrator-Instanz des PowerShell-Terminals. Verwenden Sie nicht die alte Eingabeaufforderung und auch nicht die PowerShell als Administrator.

Verwenden Sie für den gesamten Installationsvorgang das gleiche Terminal. Wenn Sie es schließen, geht die Einrichtung der Umgebungsvariablen für die Installation verloren.

Führen Sie unter Windows den folgenden Befehl aus, um das veraltete curl-Alias der PowerShell für die aktuelle Shell zu entfernen, damit Sie `curl` statt `curl.exe` verwenden können:

    Remove-Item Alias:Curl

<h2 id="obtaining-fastboot">fastboot beschaffen</h2>

Sie benötigen eine aktuelle Kopie des `fastboot`-Tools und das Verzeichnis, das es enthält, muss in die Umgebungsvariable `PATH` aufgenommen werden. Mit `fastboot --version` können Sie die aktuell installierte Version ermitteln. Sie muss mindestens `35.0.1` sein. Sie können dafür die Paketverwaltung Ihrer genutzten Distribution verwenden, aber die meisten Distributionen paketieren fälschlicherweise Entwicklungs-Snapshots von fastboot, verstecken das Standard-Versionsschema der Platform Tools (adb, fastboot usw.) hinter ihrem eigenen Schema und halten es nicht auf dem neuesten Stand, obwohl das entscheidend ist.

Unter Arch Linux installieren Sie `android-tools` und überspringen Sie den folgenden Abschnitt über das Verwenden der eigenständigen Versionen der Android Platform Tools:

    sudo pacman -S android-tools

Debian und Ubuntu haben kein brauchbares Paket für fastboot. Ihre Pakete für diese Werkzeuge sind sowohl kaputt als auch viele Jahre veraltet. Folgen Sie den nachfolgenden Anweisungen für Plattformen ohne ein geeignetes Paket.

<h3 id="standalone-platform-tools">Standalone Platform Tools</h3>

Wenn Ihr Betriebssystem keine brauchbare Version von fastboot enthält, können Sie die offiziellen Standalone-Versionen der Platform Tools verwenden. Das ist unsere Empfehlung für die meisten Benutzer. Das Flashen funktioniert nur, wenn Sie diese Anweisungen befolgen, einschließlich der Einrichtung von PATH.

Zum Herunterladen, Verifizieren und Entpacken der Standalone Platform Tools unter Debian und Ubuntu:

	sudo apt install libarchive-tools
	curl -O https://dl.google.com/android/repository/platform-tools_r35.0.2-linux.zip
	echo 'acfdcccb123a8718c46c46c059b2f621140194e5ec1ac9d81715be3d6ab6cd0a platform-tools_r35.0.2-linux.zip' | sha256sum -c
	bsdtar xvf platform-tools_r35.0.2-linux.zip

Zum Herunterladen, Verifizieren und Entpacken der Standalone Platform Tools unter macOS:

	curl -O https://dl.google.com/android/repository/platform-tools_r35.0.2-darwin.zip
	echo 'SHA256 (platform-tools_r35.0.2-darwin.zip) = 1820078db90bf21628d257ff052528af1c61bb48f754b3555648f5652fa35d78' | shasum -c
	tar xvf platform-tools_r35.0.2-darwin.zip

Zum Herunterladen, Verifizieren und Entpacken der Standalone Platform Tools unter Windows:

	curl -O https://dl.google.com/android/repository/platform-tools_r35.0.2-win.zip
	(Get-FileHash platform-tools_r35.0.2-win.zip).hash -eq "2975a3eac0b19182748d64195375ad056986561d994fffbdc64332a516300bb9"
	tar xvf platform-tools_r35.0.2-win.zip
	
Als Nächstes fügen Sie die Tools zu Ihrem `PATH` in der aktuellen Shell hinzu, sodass sie ohne Angabe des Dateipfads verwendet werden können, was die Verwendung durch das Flash-Skript ermöglicht.

Unter Debian, Ubuntu und macOS:

    export PATH="$PWD/platform-tools:$PATH"

Unter Windows:

    $env:Path = "$pwd\platform-tools;$env:Path"

Dies ändert nur den `PATH` für die aktuelle Shell und muss erneut durchgeführt werden, wenn Sie ein neues Terminal öffnen.

<h2 id="checking-fastboot-version">Überprüfen der Fastboot-Version</h2>

Überprüfen Sie die Ausgabe von `fastboot --version`, bevor Sie fortfahren.

Beispiel für die Ausgabe nach dem Ausführen der obigen Anweisungen für die Standalone Platform Tools:

    fastboot version 35.0.2-12147458
    Installed as /home/username/platform-tools/fastboot

<h2 id="flashing-as-non-root">Flash als Nicht-Root-Benutzer</h2>

Bei traditionellen Linux-Distributionen können USB-Geräte ohne udev-Regeln für den jeweiligen Gerätetyp nicht als Nicht-Root verwendet werden. Für andere Plattformen ist das kein Problem.

Unter Arch Linux:

    sudo pacman -S android-udev

Unter Debian und Ubuntu:

    sudo apt install android-sdk-platform-tools-common

Die udev-Regeln unter Debian und Ubuntu sind sehr veraltet, aber das Paket enthält dennoch die Regeln, die für Pixel-Geräte benötigt werden, da die gleichen USB-IDs seit vielen Jahren verwendet werden.

<h2 id="working-around-fwupd-bugs-on-linux-distributions">Umgehen der fwupd-Bugs auf Linux-Distributionen</h2>

Die Software fwupd, die häufig auf Linux-Distributionen zur Aktualisierung der Firmware verwendet wird, ist dafür bekannt, dass sie sich fälschlicherweise mit beliebigen Geräten über das fastboot-Protokoll verbindet, wodurch die Verwendung für den beabsichtigten Zweck blockiert wird. Dies kann dazu führen, dass man eine Fehlermeldung erhält, die besagt, dass das USB-Gerät bereits in Gebrauch ist (Behauptung), wenn man versucht, es für den beabsichtigten Zweck zu verbinden.

Sie können fwupd mit dem folgenden Befehl stoppen:

    sudo systemctl stop fwupd.service

Der Dienst wird dadurch nicht dauerhaft deaktiviert und beim Neustart erneut gestartet.

<h2 id="booting-into-the-bootloader-interface">Booten in den Bootloader-Modus</h2>

Sie müssen Ihr Gerät in den Bootloader-Modus booten. Um das zu tun, müssen Sie die Leiser-Taste gedrückt lassen, während das Gerät bootet.

Am einfachsten ist es, das Gerät neu zu starten und die Leiser-Taste gedrückt zu halten, bis es im Bootloader-Modus startet.

Alternativ schalten Sie das Gerät aus und starten Sie es dann, wobei Sie die Leiser-Taste während des Bootvorgangs gedrückt halten. Sie können es entweder mit der Einschalttaste oder durch das Anschließen des Kabels starten, was im nächsten Abschnitt ohnehin nötig ist.

Dieser Schritt ist erst abgeschlossen, wenn Ihr Gerät ein rotes Warndreieck und die Worte „Fastboot Mode“ anzeigt. Sie dürfen die Einschalttaste des Geräts nicht drücken, um den Menüpunkt „Start“ zu aktivieren, da das Gerät im Fastboot-Modus bleiben muss, damit das Installationsprogramm eine Verbindung herstellen kann.

<h2 id="connecting-phone">Das Gerät verbinden</h2>

Verbinden Sie das Gerät mit dem Computer. Unter Linux müssen Sie das erneut tun, falls Sie die udev-Regeln beim Anschließen des Gerätes noch nicht eingerichtet hatten.

Die aktuellen Windows-Versionen 10 und 11 enthalten einen generischen Treiber, der für fastboot geeignet ist, und erfordern nicht mehr die Installation eines Treibers für die Installation auf dem Pixel 4a (5G) oder neuer. Für die Installation auf älteren Pixel der 4. Generation ist ein Treiber immer noch notwendig, da der generischen Treiber nicht mit fastbootd umgehen kann. Ältere Windows-Versionen benötigen den Treiber auch für nicht-veraltete Pixel-Geräte. Sie können den Treiber über Windows Update beziehen, das ihn als optionales Update erkennt, wenn das Gerät in die Bootloader-Schnittstelle gebootet und mit dem Computer verbunden wird. Öffnen Sie Windows Update, führen Sie eine Prüfung auf Updates durch und öffnen Sie dann „Optionale Updates anzeigen“. Installieren Sie den Treiber für die Android-Bootloader-Schnittstelle als optionales Update, das aufgrund der USB-ID-Überschneidung als „LeMobile Android Device“ angezeigt wird. Ein alternativer Weg, die Windows Fastboot-Treiber zu beziehen, besteht darin, den [neuesten Treiber für Pixels](https://developer.android.com/studio/run/win-usb) von Google herunterzuladen und ihn dann [manuell über den Windows-Geräte-Manager zu installieren](https://developer.android.com/studio/run/oem-usb#InstallingDriver).

Trennen Sie das Pixel-Tablet vom Ladedock mit Lautsprecher, bevor Sie fortfahren. Das Ladedock verwendet USB zum Aufladen und für die Audioausgabe, aber das Tablet bietet keine Unterstützung für die gleichzeitige Verwendung des Ladedock und des USB-Anschlusses.

<h2 id="unlocking-the-bootloader">Den Bootloader entsperren</h2>

Entsperren Sie den Bootloader, damit Sie das Betriebssystem und die Firmware flashen können:

    fastboot flashing unlock

Der Befehl muss auf dem Gerät bestätigt werden und löscht alle Daten. Verwenden Sie eine der Lautstärketasten, um Ihre Auswahl zu treffen, und die Einschalttaste, um diese zu bestätigen.

<h2 id="obtaining-openssh">OpenSSH beschaffen</h2>

OpenSSH wird verwendet, um den Download des Betriebssystems über die durch HTTPS gebotene Sicherheit hinaus zu überprüfen.

macOS und Windows enthalten OpenSSH bereits in ihrer Grundinstallation.

Unter Arch Linux:

    sudo pacman -S openssh

Unter Debian und Ubuntu:

    sudo apt install openssh-client

<h2 id="obtaining-factory-images">Factory Images beschaffen</h2>

Um mit dem Installationsprozess fortzufahren, müssen Sie die GrapheneOS Factory Images für Ihr Gerät beschaffen.

Sie können die Dateien entweder mit Ihrem Browser oder mit einem Befehl wie `curl` herunterladen. In der Regel ist es einfacher, die Kommandozeile zu verwenden, da Sie sie bereits für den Rest des Installationsprozesses nutzen. Daher verwendet diese Anleitung `curl`.

Laden Sie [den öffentlichen Schlüssel der Factory Images (allowed_signers)](https://releases.grapheneos.org/allowed_signers) herunter, um die Factory Images zu verifizieren:

    curl -O https://releases.grapheneos.org/allowed_signers

Das ist der Inhalt von `allowed_signers`:

    contact@grapheneos.org ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIIUg/m5CoP83b0rfSCzYSVA4cw4ir49io5GPoxbgxdJE

Andere Stellen, an denen Sie den Signierschlüssel erhalten können:
<ul>
<li><a href="https://bsky.app/profile/grapheneos.org/post/3kleyygkptm2x">Bluesky</a></li>
<li><a href="https://x.com/GrapheneOS/status/1757758688952009209">X</a></li>
<li><a href="https://github.com/GrapheneOS/releases.grapheneos.org/blob/main/static/allowed_signers">GitHub</a></li>
</ul>

Wenn der aktuelle Signierschlüssel ersetzt wird, wird der neue Schlüssel mit diesem signiert.

Laden Sie die Factory Images für das Gerät von der [Unterseite „Releases“](https://grapheneos.org/releases) herunter. Um zum Beispiel die `VERSION`-Version für ein Gerät mit dem Codenamen `DEVICE_NAME` herunterzuladen:

    curl -O https://releases.grapheneos.org/DEVICE_NAME-install-VERSION.zip
    curl -O https://releases.grapheneos.org/DEVICE_NAME-install-VERSION.zip.sig

Überprüfen Sie anschließend die Images anhand der Signatur.

Unter Linux und macOS:

    ssh-keygen -Y verify -f allowed_signers -I contact@grapheneos.org -n "factory images" -s DEVICE_NAME-install-VERSION.zip.sig < DEVICE_NAME-install-VERSION.zip

Unter Windows:

    cmd /c 'ssh-keygen -Y verify -f allowed_signers -I contact@grapheneos.org -n "factory images" -s DEVICE_NAME-install-VERSION.zip.sig < DEVICE_NAME-install-VERSION.zip'

Wenn die Überprüfung erfolgreich war, wird die folgende Ausgabe angezeigt:

    Good "factory images" signature for contact@grapheneos.org with ED25519 key SHA256:AhgHif0mei+9aNyKLfMZBh2yptHdw/aN7Tlh/j2eFwM

<h2 id="flashing-factory-images">Factory Images flashen</h2>

Die Erstinstallation erfolgt durch das Flashen der Factory Images. Dadurch wird das bestehende Betriebssystem ersetzt und alle vorhandenen Daten werden gelöscht.

Extrahieren Sie als Nächstes die Factory Images.

Unter Linux:

    bsdtar xvf DEVICE_NAME-install-VERSION.zip

Unter macOS und Windows:

    tar xvf DEVICE_NAME-install-VERSION.zip

Wechseln Sie in das Verzeichnis:

    cd DEVICE_NAME-install-VERSION

Flashen Sie die Images mit dem flash-all-Script in dem Verzeichnis.

Unter Linux und macOS:

    bash flash-all.sh

Unter Windows:

    ./flash-all.bat

Warten Sie, bis der Flashvorgang abgeschlossen ist. Das Flashen der Firmware, der Neustart in den Bootloader-Modus und das Flashen des Betriebssystems werden automatisch durchgeführt. Vermeiden Sie es, mit dem Gerät zu interagieren, bis das Flash-Skript beendet ist. <a href="#locking-the-bootloader">Sperren Sie anschließend den Bootloader</a>, bevor Sie das Gerät verwenden, da durch das Sperren die Daten erneut gelöscht werden.

<h3 id="troubleshooting">Fehlersuche</h3>

Die Textausgabe eines fehlgeschlagenen Flashversuchs enthält wertvolle Diagnoseinformationen, die wichtig sind, um zu wissen, wo und wie der Prozess schief gelaufen ist. Bitte geben Sie diese Informationen an, wenn Sie im [GrapheneOS-Chatraum](https://grapheneos.org/contact#community) um Hilfe bitten.

Ein häufiges Problem bei Linux-Distributionen ist, dass sie das Standardverzeichnis für temporäre Dateien `/tmp` als tmpfs mounten, was dazu führt, dass der Arbeitsspeicher und Swap anstatt des dauerhaften Speicher verwendet wird. Standardmäßig beträgt die Größe 50% des verfügbaren virtuellen Speichers. Dies ist für den Flash-Prozess oft nicht ausreichend, zumal `/tmp` von Anwendungen und Benutzern gemeinsam genutzt wird. Verwenden Sie ein anderes temporäres Verzeichnis, falls Ihr `/tmp` nicht genügend Platz bietet:

    mkdir tmp && TMPDIR="$PWD/tmp" ./flash-all.sh

<h2 id="locking-the-bootloader">Den Bootloader sperren</h2>

Das Sperren des Bootloaders ist wichtig, da es einen vollständig verifizierten Bootvorgang ermöglicht. Es verhindert auch die Verwendung von fastboot zum Flashen, Formatieren oder Löschen von Partitionen. Ein verifizierter Bootvorgang erkennt Änderungen an den Betriebssystempartitionen und verhindert das Lesen von veränderten oder beschädigten Daten. Wenn Änderungen erkannt werden, wird mit Hilfe von Fehlerkorrekturdaten versucht, die ursprünglichen Daten zu erhalten, die dann erneut verifiziert werden, was das verifizierte Booten gegen nicht-bösartige Datenveränderungen abhärtet.

Während sich das Gerät im Bootloader-Modus befindet, setzen Sie es auf „locked“:

    fastboot flashing lock

Der Befehl muss auf dem Gerät bestätigt werden und löscht alle Daten. Verwenden Sie eine der Lautstärketasten, um Ihre Auswahl zu treffen, und die Einschalttaste, um diese zu bestätigen.

<h2 id="post-installation">Nach der Installation</h2>

<h3 id="booting">Booten</h3>

Sie haben GrapheneOS nun erfolgreich installiert und können es booten. Wenn Sie die Einschalttaste drücken und die Standardoption „Start“ im Bootloader-Modus ausgewählt lassen, wird das Betriebssystem gestartet.

<h3 id="disabling-oem-unlocking">OEM-Entsperrung deaktivieren</h3>

Bei der Ersteinrichtung enthält der letzte Schritt einen Schalter für die empfohlene Deaktivierung der OEM-Entsperrung, der standardmäßig aktiviert ist.

Wenn Sie die OEM-Entsperrung in Zukunft aktivieren oder deaktivieren müssen, können Sie dies im Menü „Entwickleroptionen“ des Betriebssystems tun.

<h3 id="verifying-installation">Überprüfung der Installation</h3>

Die von den unterstützten Geräten bereitgestellten Funktionen „Verified Boot“ und „Attestation“ können verwendet werden, um zu überprüfen, ob die Hardware, die Firmware und die GrapheneOS-Installation unverfälscht sind. Selbst wenn der Computer, den Sie zum Flashen von GrapheneOS verwendet haben, kompromittiert wurde und ein Angreifer GrapheneOS durch sein eigenes bösartiges Betriebssystem ersetzt hat, kann dies mit diesen Funktionen erkannt werden.

Verified Boot überprüft bei jedem Startvorgang sämtliche Firmware- und Betriebssystem-Images. Der öffentliche Schlüssel für die Firmware-Images ist werkseitig in elektronische Sicherungen im SoC eingebrannt. Firmware-Sicherheitsupdates aktualisieren auch den in die Sicherungen eingebrannten Rollback-Index, um einen Rollback-Schutz zu gewährleisten.

Die letzte Bootphase der Firmware vor dem Betriebssystem ist für die Verifizierung des Betriebssystems zuständig. Für das Stock-Betriebssystem wird ein fest eingebauter öffentlicher Schlüssel verwendet. Bei der Installation von GrapheneOS wird der öffentliche Verified-Boot-Schlüssel von GrapheneOS in den Sicherheitschip geflasht. Bei jedem Start wird dieser Schlüssel geladen und zur Verifizierung des Betriebssystems verwendet. Sowohl für das Stock-Betriebssystem als auch für GrapheneOS wird ein auf dem Sicherheits-Patchlevel basierender Rollback-Index aus dem Sicherheitschip geladen, um einen Rollback-Schutz zu gewährleisten.

<h4 id="verified-boot-key-hash">Hash des öffentlichen Verified-Boot-Schlüssels</h4>

Beim Laden eines alternativen Betriebssystems zeigt das Gerät beim Booten einen gelben Hinweis mit der ID des alternativen Betriebssystems an, die auf dem SHA256-Hashwert des öffentlichen Verified-Boot-Schlüssels basiert. Bei Pixels der 4. und 5. Generation werden nur die ersten 32 Bit des Hashwerts angezeigt, sodass Sie diesen Ansatz nicht verwenden können. Bei Pixels ab der 6. Generation wird der vollständige Hashwert angezeigt, den Sie mit dem offiziellen, von GrapheneOS verifizierten Boot-Key-Hashwert vergleichen können:

* Pixel 10 Pro Fold: `55a2d44103e56d5ec65496399c417987ba77730e6488fc60ba058d09fc3caee3`
* Pixel 10 Pro XL: `141d7fc32af7958a416f2661b37cf6f27bfb376fb5ce616aeaa27a82c7a04f74`
* Pixel 10 Pro: `4e8ee8f717754052198ca6d2d3aaa232e2461b4293c0d6f297e519cc778de093`
* Pixel 10: `3f7415ea26f5df5b14ea6d153256071a7a1af9ce7b0970b7311cc463c7ea02c7`
* Pixel 9a: `0508de44ee00bfb49ece32c418af1896391abde0f05b64f41bc9a2dfb589445b`
* Pixel 9 Pro Fold: `af4d2c6e62be0fec54f0271b9776ff061dd8392d9f51cf6ab1551d346679e24c`
* Pixel 9 Pro XL: `55d3c2323db91bb91f20d38d015e85112d038f6b6b5738fe352c1a80dba57023`
* Pixel 9 Pro: `f729cab861da1b83fdfab402fc9480758f2ae78ee0b61c1f2137dd1ab7076e86`
* Pixel 9: `9e6a8f3e0d761a780179f93acd5721ba1ab7c8c537c7761073c0a754b0e932de`
* Pixel 8a: `096b8bd6d44527a24ac1564b308839f67e78202185cbff9cfdcb10e63250bc5e`
* Pixel 8 Pro: `896db2d09d84e1d6bb747002b8a114950b946e5825772a9d48ba7eb01d118c1c`
* Pixel 8: `cd7479653aa88208f9f03034810ef9b7b0af8a9d41e2000e458ac403a2acb233`
* Pixel Fold: `ee0c9dfef6f55a878538b0dbf7e78e3bc3f1a13c8c44839b095fe26dd5fe2842`
* Pixel Tablet: `94df136e6c6aa08dc26580af46f36419b5f9baf46039db076f5295b91aaff230`
* Pixel 7a: `508d75dea10c5cbc3e7632260fc0b59f6055a8a49dd84e693b6d8899edbb01e4`
* Pixel 7 Pro: `bc1c0dd95664604382bb888412026422742eb333071ea0b2d19036217d49182f`
* Pixel 7: `3efe5392be3ac38afb894d13de639e521675e62571a8a9b3ef9fc8c44fd17fa1`
* Pixel 6a: `08c860350a9600692d10c8512f7b8e80707757468e8fbfeea2a870c0a83d6031`
* Pixel 6 Pro: `439b76524d94c40652ce1bf0d8243773c634d2f99ba3160d8d02aa5e29ff925c`
* Pixel 6: `f0a890375d1405e62ebfd87e8d3f475f948ef031bbf9ddd516d5f600a23677e8`

Diese Überprüfung ist nach der Installation sinnvoll, aber Sie müssen sie nicht manuell vornehmen, damit verifiziertes Booten funktioniert. Der öffentliche Schlüssel für verifiziertes Booten, der auf dem Sicherheitschip gespeichert wurde, kann nur geändert werden, wenn das Gerät entsperrt ist. Das Entsperren [des Bootloaders, Anm. d. Übers.] des Geräts löscht den Sicherheitschip wie beim Zurücksetzen auf Werkseinstellungen und verhindert ein Wiederherstellen von Daten, selbst wenn die SSD geklont wurde und Ihre Passphrase(n) erlangt wurde(n), da die Verschlüsselungsschlüssel nicht mehr abgeleitet werden können. Der Verified-Boot-Schlüssel ist außerdem einer der Eingaben für die Ableitung der Verschlüsselungsschlüssel, zusätzlich zu den Sperrmethoden des Benutzers und zufälligen Token auf dem Sicherheitschip.

<h4 id="hardware-based-attestation">Hardwaregestützte Attestierung</h4>

GrapheneOS enthält unsere Auditor-App, die eine Kombination aus Verified-Boot- und Attestierungsfunktionen verwendet, um die Unverfälschtheit der Hardware, der Firmware und des Betriebssystems zu überprüfen und weitere nützliche Daten der Hardware und des Betriebssystems bereitzustellen.

Da der Zweck von Auditor darin besteht, Informationen über das Gerät zu erhalten, ohne dem Gerät vertrauen zu müssen, werden die Ergebnisse nicht auf dem zu überprüfenden Gerät angezeigt. Sie benötigen ein zweites Android-Gerät, auf dem der Auditor zur lokalen QR-Code-basierten Überprüfung läuft. Sie können auch unseren optionalen Dienst zur Überwachung der Geräteintegrität für automatische, planmäßige Überprüfungen mit Unterstützung für E-Mail-Warnungen verwenden.

Eine Anleitung dazu finden Sie im [Auditor-Tutorial](https://attestation.app/tutorial) [[deutsche Übersetzung](grapheneos-attestation-tutorial.html)].

Auditor basiert in erster Linie auf einem Pairing-Modell, bei dem ein hardwaregestützter Signierschlüssel und ein hardwaregestützter Attestierungs-Signierschlüssel generiert und als Teil der ersten Verifizierung festgehalten werden. Die erste Verifizierung entsteht auf Basis einer Chain of Trust zu einem der Android-Attestierungs-Root-Keys. Nach der ersten Verifizierung bietet er ein hochgradig sicheres System für den Erhalt weiterer Informationen über das Gerät. Ein Angreifer könnte die erste Verifizierung mit einem geleakten Attestation-Key umgehen oder durch die Weiterleitung an ein anderes Gerät mit dem Gerätemodell, dem Betriebssystem und dem Patchlevel, das der Benutzer erwartet. Das Weiterleiten auf ein anderes Gerät wird in Zukunft mit einer optionalen Unterstützung für die Attestierung der Hardware-Seriennummer angegangen werden.

<h3 id="further-information">Weitere Informationen</h3>

Bitte schauen Sie sich das [Benutzerhandbuch](https://grapheneos.org/usage) und die [FAQ](https://grapheneos.org/faq) für weitere Informationen an. Wenn Sie weitere Fragen haben, die nicht von der Website abgedeckt werden, treten Sie den [offiziellen GrapheneOS-Chat-Kanälen](https://grapheneos.org/contact#community) bei und stellen Sie Ihre Fragen in dem entsprechenden Kanal.

<h3 id="replacing-grapheneos-with-the-stock-os">GrapheneOS durch das Standardbetriebssystem ersetzen</h3>

Die Installation des Stock-Betriebssystems mittels Factory Images ist derselbe Prozess wie oben beschrieben. Vor dem Flashen und Sperren gibt es jedoch einen zusätzlichen Schritt, um das Gerät vollständig in einen sauberen Werkszustand zurückzusetzen.

Die GrapheneOS Factory Images flashen einen nicht-standardmäßigen Android-Verified-Boot-Schlüssel, der gelöscht werden muss, um das Gerät vollständig in den Werkszustand zurückzuversetzen. Bevor Sie die Factory Images flashen, sollten Sie das Gerät in den Fastboot-Modus booten und sicherstellen, dass der Bootloader entsperrt ist. Löschen Sie dann den benutzerdefinierten Android-Verified-Boot-Schlüssel, um ihm nicht länger zu vertrauen:

    fastboot erase avb_custom_key
