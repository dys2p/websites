<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item"><a href="grapheneos-install.html">Installation</a></li>
		<li class="breadcrumb-item active" aria-current="page">Web-Installationsanleitung</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/grapheneos.org/blob/c61bcec301c0a0c180f86f31cb2554ebecb2f8d0/static/install/web.html">c61bcec</a> vom 2025-07-01. Falls Sie Hinweise oder Verbesserungsvorschläge haben, dann <a href="contact.html">schreiben Sie uns gerne</a> oder arbeiten Sie mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-install-web/de.md">GitHub</a> an dieser Übersetzung.
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

<h1 id="web-install">Web-Installationsanleitung</h1>

Dies ist eine Anleitung für die WebUSB-basierte Installation von GrapheneOS und ist die empfohlene Vorgehensweise für die meisten Benutzer. Die [CLI-Installationsanleitung](https://grapheneos.org/install/cli) [[deutsche Übersetzung](grapheneos-install-cli.html)] ist die traditionellere Methode zur Installation von GrapheneOS.

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
      <a href="#obtaining-factory-images">Factory Images beschaffen</a>
    </li>
    <li>
      <a href="#flashing-factory-images">Factory Images flashen</a>
    </li>
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

Für den Web-Installationsprozess benötigen Sie einen Computer mit mindestens 2 GB freiem Arbeitsspeicher und 32 GB freiem Speicherplatz. Die Web-Installation kann im Gegensatz zur CLI-Installation auf einem Android-Smartphone oder -Tablet ausgeführt werden.

Sie benötigen ein USB-Kabel, um das Gerät an einem Computer anzuschließen, der die Installation durchführt. Verwenden Sie nach Möglichkeit das hochwertige, standardkonforme USB-C-Kabel, das dem Gerät beiliegt. Wenn Ihr Computer keine USB-C-Anschlüsse hat, benötigen Sie ein hochwertiges USB-C-zu-USB-A-Kabel. Sie sollten vermeiden, einen USB-Hub zu verwenden, wie z. B. an der Vorderseite eines Desktop-Computergehäuses. Schließen Sie das Kabel direkt an einen rückwärtigen Anschluss eines Desktops oder an die Anschlüsse eines Laptops an. Viele weit verbreitete USB-Kabel und -Hubs sind kaputt und die häufigste Ursache für Probleme bei der Installation von GrapheneOS.

Die Installation von einem Betriebssystem in einer virtuellen Maschine wird nicht empfohlen. USB-Passthrough ist oft nicht zuverlässig. Um diese Probleme auszuschließen, installieren Sie von einem Betriebssystem, das direkt auf einem Gerät läuft. Virtuelle Maschinen sind außerdem oft so konfiguriert, dass sie sehr begrenzten Arbeitsspeicher und Speicherplatz haben.

Offiziell unterstützte Betriebssysteme für die Web-Installationsmethode:

* Windows 10
* Windows 11
* macOS Ventura (13)
* macOS Sonoma (14)
* macOS Sequoia (15)
* Arch Linux
* Debian 11 (Bullseye)
* Debian 12 (Bookworm)
* Ubuntu 20.04 LTS
* Ubuntu 22.04 LTS
* Ubuntu 24.04 LTS
* Ubuntu 24.10
* Linux Mint 20 (nutzen Sie die Anweisungen für Ubuntu 20.04 LTS)
* Linux Mint 21 (nutzen Sie die Anweisungen für Ubuntu 22.04 LTS)
* Linux Mint 22 (nutzen Sie die Anweisungen für Ubuntu 24.04 LTS)
* Linux Mint Debian Edition 6 (nutzen Sie die Anweisungen für Debian 12)
* ChromeOS
* GrapheneOS
* Android 13 mit Play Protect-Zertifizierung
* Android 14 mit Play Protect-Zertifizierung
* Android 15 mit Play Protect-Zertifizierung
* Android 16 mit Play Protect-Zertifizierung

Auch ältere, nicht mehr weiterentwickelte Versionen dieser Betriebssysteme lassen sich einsetzen, allerdings erhalten sie keinen offiziellen Support.

Stellen Sie sicher, dass Ihr Betriebssystem auf dem neuesten Stand ist, bevor Sie fortfahren.

Offiziell unterstützte Browser für die Web-Installationsmethode:

* Chromium (außer unter Ubuntu, weil dort ein kaputtes Snap-Paket ohne funktionierendes WebUSB beiliegt)
* Vanadium (GrapheneOS)
* Google Chrome
* Microsoft Edge
* Brave (mit deaktivierten Brave Shields, da es die Speichernutzung auf einen niedrigen Wert begrenzt, um ein Fingerprinting des verfügbaren Speichers zu vermeiden)

Unter Android müssen Sie den Desktop-Modus für den Browser deaktivieren, da er derzeit verhindert, dass der Web-Installer Android erkennt und nach jedem Neustart die Erlaubnis zur erneuten Verbindung mit dem Gerät einholen muss. Der Desktop-Modus ist auf großen Tablets mit mindestens 8 GB RAM, wie dem Pixel Tablet, standardmäßig aktiviert.

Sie sollten Flatpak- und Snap-Versionen von Browsern vermeiden, da diese dafür bekannt sind, während des Installationsvorgangs Probleme zu verursachen.

Vergewissern Sie sich, dass Ihr Browser auf dem neuesten Stand ist, bevor Sie fortfahren.

Verwenden Sie keinen Inkognito-Modus oder andere private Browser-Modi. Diese Modi verhindern in der Regel, dass das Web-Installationsprogramm über genügend Speicherplatz verfügt, um die heruntergeladene Version zu extrahieren.

Sie benötigen eines der [offiziell unterstützten Geräte](https://grapheneos.org/faq#supported-devices). Um sicherzustellen, dass das Gerät für die Installation von GrapheneOS freigeschaltet werden kann, sollten Sie Netzbetreiber-spezifische Geräte vermeiden. Netzbetreiber-spezifische Varianten der Pixel-Smartphones verwenden dasselbe Stock-Betriebssystem und dieselbe Firmware, aber mit einer Carrier-ID ungleich Null, die bei der Herstellung auf die persistente Partition geflasht wurde. Die Carrier-ID aktiviert eine Netzbetreiber-spezifische Konfiguration im Stock-Betriebssystem, einschließlich einer Deaktivierung der Netzbetreiber- und Bootloader-Entsperrung. Der Netzbetreiber kann dies möglicherweise aus der Ferne deaktivieren, aber seine Support-Mitarbeiter wissen das möglicherweise nicht und werden es wahrscheinlich auch nicht tun. Holen Sie sich ein Netzbetreiber-unabhängiges Gerät, um das Risiko und möglichen Ärger zu vermeiden. Wenn Sie einen Weg finden _können_, ein Netzbetreiber-spezifisches Gerät zu entsperren, ist das kein Problem, da GrapheneOS die Carrier-ID einfach ignorieren kann und die Hardware identisch ist.

Es hat sich bewährt, das Gerät vor der Installation von GrapheneOS zu aktualisieren, um beim Anschluss des Gerätes an den Computer und in der frühen Phase des Installationsprozesses die neueste Firmware zu haben. So oder so, GrapheneOS flasht die neueste Firmware zu Beginn des Installationsprozesses.

<h2 id="enabling-oem-unlocking">OEM-Entsperrung aktivieren</h2>

Die OEM-Entsperrung muss innerhalb des Betriebssystems aktiviert werden.

Aktivieren Sie das Menü „Entwickleroptionen“, indem Sie zu <b>Einstellungen&#160;<span aria-label="and then">></span> Über das Telefon/Tablet</b> gehen und wiederholt auf den Menüeintrag <b>Build-Nummer</b> drücken, bis die Entwickleroptionen aktiviert sind.

Gehen Sie dann zu <b>Einstellungen&#160;<span aria-label="and then">></span> System&#160;<span aria-label="and then">></span> Entwickleroptionen</b> und aktivieren Sie die Einstellung <b>OEM-Entsperrung</b>. Bei Gerätevarianten (SKUs), die von Netzbetreibern als gesperrte Geräte verkauft werden können, erfordert die Aktivierung der <b>OEM-Entsperrung</b> einen Internetzugang, damit das Stock-Betriebssystem prüfen kann, ob das Gerät von einem Netzbetreiber als gesperrt verkauft wurde.

Für das Pixel 6a wird die OEM-Entsperrung nicht mit der ab Werk installierten Version des Stock-Betriebssystems funktionieren. Sie müssen es mit einem Over-the-Air-Update auf die Version von Juni 2022 oder später aktualisieren. Nach der Aktualisierung müssen Sie das Gerät auf die Werkseinstellungen zurücksetzen, um die OEM-Entsperrung zu ermöglichen.

<h2 id="flashing-as-non-root">Flash als Nicht-Root-Benutzer</h2>

Bei traditionellen Linux-Distributionen können USB-Geräte ohne udev-Regeln für den jeweiligen Gerätetyp nicht als Nicht-Root verwendet werden. Für andere Plattformen ist das kein Problem.

Unter Arch Linux installieren Sie das Paket `android-udev`. Unter Debian und Ubuntu installieren Sie das Paket `android-sdk-platform-tools-common`.

<h2 id="working-around-fwupd-bugs-on-linux-distributions">Umgehen der fwupd-Bugs auf Linux-Distributionen</h2>

Die Software fwupd, die häufig auf Linux-Distributionen zur Aktualisierung der Firmware verwendet wird, ist dafür bekannt, dass sie sich fälschlicherweise mit beliebigen Geräten über das Fastboot-Protokoll verbindet, wodurch die Verwendung für den beabsichtigten Zweck blockiert wird. Dies kann dazu führen, dass man eine Fehlermeldung erhält, die besagt, dass das USB-Gerät bereits in Gebrauch ist (Behauptung), wenn man versucht, es für den beabsichtigten Zweck zu verbinden.

Sie können fwupd mit dem folgenden Befehl stoppen:

    sudo systemctl stop fwupd.service

Der Dienst wird dadurch nicht dauerhaft deaktiviert und beim Neustart erneut gestartet.

<h2 id="booting-into-the-bootloader-interface">Booten in den Bootloader-Modus</h2>

Sie müssen Ihr Gerät im Bootloader-Modus booten. Um das zu tun, müssen Sie die Leiser-Taste gedrückt lassen, während das Gerät bootet.

Am einfachsten ist es, das Gerät neu zu starten und die Leiser-Taste gedrückt zu halten, bis es im Bootloader-Modus startet.

Alternativ schalten Sie das Gerät aus und starten Sie es dann, wobei Sie die Leiser-Taste während des Bootvorgangs gedrückt halten. Sie können es entweder mit der Einschalttaste oder durch das Anschließen des Kabels starten, was im nächsten Abschnitt ohnehin nötig ist.

Dieser Schritt ist erst abgeschlossen, wenn Ihr Gerät ein rotes Warndreieck und die Worte „Fastboot-Modus“ anzeigt. Sie dürfen die Einschalttaste des Geräts nicht drücken, um den Menüpunkt „Start“ zu aktivieren, da das Gerät im Fastboot-Modus bleiben muss, damit das Installationsprogramm eine Verbindung herstellen kann.

<h2 id="connecting-phone">Das Gerät verbinden</h2>

Verbinden Sie das Gerät mit dem Computer. Unter Linux müssen Sie das erneut tun, falls Sie die udev-Regeln beim Anschließen des Gerätes noch nicht eingerichtet hatten.

Die aktuellen Windows-Versionen 10 und 11 enthalten einen generischen Treiber, der für fastboot geeignet ist, und erfordert nicht mehr die Installation eines Treibers für die Installation auf dem Pixel 4a (5G) oder neuer. Für die Installation auf älteren Pixel der 4. Generation ist ein Treiber immer noch notwendig, da der generischen Treiber nicht mit fastbootd umgehen kann. Ältere Windows-Versionen benötigen den Treiber auch für nicht-veraltete Pixel-Geräte. Sie können den Treiber über Windows Update beziehen, das ihn als optionales Update erkennt, wenn das Gerät in die Bootloader-Schnittstelle gebootet und mit dem Computer verbunden wird. Öffnen Sie Windows Update, führen Sie eine Prüfung auf Updates durch und öffnen Sie dann „Optionale Updates anzeigen“. Installieren Sie den Treiber für die Android-Bootloader-Schnittstelle als optionales Update, das aufgrund der USB-ID-Überschneidung als „LeMobile Android Device“ angezeigt wird. Ein alternativer Weg, die Windows Fastboot-Treiber zu beziehen, besteht darin, den [neuesten Treiber für Pixels](https://developer.android.com/studio/run/win-usb) von Google herunterzuladen und ihn dann [manuell über den Windows-Geräte-Manager zu installieren](https://developer.android.com/studio/run/oem-usb#InstallingDriver).

Trennen Sie das Pixel-Tablet vom Ladedock mit Lautsprecher, bevor Sie fortfahren. Das Ladedock verwendet USB zum Aufladen und für die Audioausgabe, aber das Tablet bietet keine Unterstützung für die gleichzeitige Verwendung des Ladedock und des USB-Anschlusses.

<h2 id="unlocking-the-bootloader">Den Bootloader entsperren</h2>

Entsperren Sie den Bootloader, damit Sie das Betriebssystem und die Firmware flashen können:

Drücken Sie im Web-Installer unter [Unlocking the bootloader](https://grapheneos.org/install/web#unlocking-the-bootloader) die Schaltfläche <kbd>Unlock bootloader</kbd>.

Der Befehl muss auf dem Gerät bestätigt werden und löscht alle Daten. Verwenden Sie eine der Lautstärketasten, um Ihre Auswahl zu treffen, und die Einschalttaste, um diese zu bestätigen.

<h2 id="obtaining-factory-images">Factory Images beschaffen</h2>

Um mit dem Installationsprozess fortzufahren, müssen Sie die GrapheneOS Factory Image für Ihr Gerät beschaffen.

Drücken Sie im Web-Installer unter [Obtaining factory images](https://grapheneos.org/install/web#obtaining-factory-images) die Schaltfläche <kbd>Download release</kbd>.

<h2 id="flashing-factory-images">Factory Images flashen</h2>

Die Erstinstallation erfolgt durch das Flashen der Factory Images. Das wird das bestehende Betriebssystem ersetzen und alle vorhandenen Daten löschen.

Drücken Sie im Web-Installer unter [Flashing factory images](https://grapheneos.org/install/web#flashing-factory-images) die Schaltfläche <kbd>Flash release</kbd>.

Warten Sie, bis der Flashvorgang abgeschlossen ist. Das Flashen der Firmware, der Neustart in den Bootloader-Modus und das Flashen des Betriebssystems werden automatisch durchgeführt. Vermeiden Sie es, mit dem Gerät zu interagieren, bis das Flash-Skript beendet ist. <a href="#locking-the-bootloader">Sperren Sie im Anschluss den Bootloader</a>, bevor Sie das Gerät verwenden, da beim Sperren die Daten erneut gelöscht werden.

<h2 id="locking-the-bootloader">Den Bootloader sperren</h2>

Das Sperren des Bootloaders ist wichtig, da es einen vollständig verifizierten Bootvorgang ermöglicht. Es verhindert auch die Verwendung von fastboot zum Flashen, Formatieren oder Löschen von Partitionen. Ein verifizierter Bootvorgang erkennt Änderungen an den Betriebssystempartitionen und verhindert das Lesen von veränderten oder beschädigten Daten. Wenn Änderungen erkannt werden, wird mit Hilfe von Fehlerkorrekturdaten versucht, die ursprünglichen Daten zu erhalten, die dann erneut verifiziert werden, was das verifizierte Booten gegen nicht-bösartige Datenveränderungen abhärtet.

Während sich das Gerät im Bootloader-Modus befindet, setzen Sie es auf „locked“:

Drücken Sie im Web-Installer unter [Locking the bootloader](https://grapheneos.org/install/web#locking-the-bootloader) die Schaltfläche <kbd>Lock bootloader</kbd>.

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

Beim Laden eines alternativen Betriebssystems zeigt das Gerät beim Booten einen gelben Hinweis mit der ID des alternativen Betriebssystems an, die auf dem SHA256-Hashwert des öffentlichen Verified-Boot-Schlüssels basiert. Bei Pixels der 4. und 5. Generation werden nur die ersten 32 Bits des Hashwerts angezeigt, sodass Sie diesen Ansatz nicht verwenden können. Ab Pixels der 6. Generation wird der vollständigen Hashwert angezeigt, den Sie mit den offiziellen Verified-Boot-Key-Hashwert von GrapheneOS vergleichen können:

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

Diese Überprüfung ist nach der Installation sinnvoll, aber Sie müssen sie nicht manuell vornehmen, damit verifiziertes Booten funktioniert. Der öffentliche Schlüssel für verifiziertes Booten, der auf dem Sicherheitschip gespeichert wurde, kann nur geändert werden, wenn das Gerät entsperrt ist. Das Entsperren [des Bootloaders, Anm. d. Übers.] des Geräts löscht den Sicherheitschip wie beim Zurücksetzen auf Werkseinstellungen und verhindert ein Wiederherstellen von Daten, selbst wenn die SSD geklont wurde und Ihre Passphrase(n) erlangt wurde(n), da die Verschlüsselungsschlüssel nicht mehr abgeleitet werden können. Der Verified-Boot-Schlüssel ist außerdem eine der Eingaben für die Ableitung der Verschlüsselungsschlüssel, zusätzlich zu den Sperrmethoden des Benutzers und zufälligen Token auf dem Sicherheitschip.

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

Drücken Sie im Web-Installer unter [Replacing GrapheneOS with the stock OS](https://grapheneos.org/install/web#replacing-grapheneos-with-the-stock-os) die Schaltfläche <kbd>Remove non-stock key</kbd>.
