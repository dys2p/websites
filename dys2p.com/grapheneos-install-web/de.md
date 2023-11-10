<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item"><a href="grapheneos-install.html">Installation</a></li>
		<li class="breadcrumb-item active" aria-current="page">Web-Installationsanleitung</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/grapheneos.org/blob/fd2170de159bcfe7136fbbb6418a6f09ec26fba3/static/install/cli.html"> 41387d5</a> vom 2023-11-07. Falls du Hinweise oder Verbesserungsvorschläge hast, dann <a href="contact.html">schreib uns gerne</a> oder arbeite mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-install-cli/de.md">GitHub</a> an dieser Übersetzung.
</div>

<!--
Copyright © 2014-2023 GrapheneOS

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
-->

<h1 id="web-install">Web-Installationssanleitung</h1>

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
              <a href="#hardware-based-attestation">Hardware-gestützte Attestierung</a>
            </li>
      	  </ul>
        <li>
          <a href="#replacing-grapheneos-with-the-stock-os">GrapheneOS durch das Standardbetriebssystem ersetzen</a>
        </li>
        <li>
          <a href="#further-information">Weitere Informationen</a>
        </li>
      </ul>
  </ul>
</nav>

<h2 id="prerequisites">Voraussetzungen</h2>

Sie sollten mindestens 2 GB freien Arbeitsspeicher und 32 GB freien Speicherplatz zur Verfügung haben.

Sie benötigen ein USB-Kabel, um das Gerät mit einem Laptop oder Desktop-Computer zu verbinden. Verwenden Sie nach Möglichkeit das hochwertige, standardkonforme USB-C-Kabel, das dem Gerät beiliegt. Wenn Ihr Computer keine USB-C-Anschlüsse hat, benötigen Sie ein hochwertiges USB-C-zu-USB-A-Kabel. Sie sollten vermeiden, einen USB-Hub zu verwenden, wie z. B. an der Vorderseite eines Desktop-Computergehäuses. Schließen Sie das Kabel direkt an einen rückwärtigen Anschluss eines Desktops oder an die Anschlüsse eines Laptops an. Viele weit verbreitete USB-Kabel und -Hubs sind kaputt und die häufigste Ursache für Probleme bei der Installation von GrapheneOS.

Die Installation von einem Betriebssystem in einer virtuellen Maschine wird nicht empfohlen. USB-Passthrough ist oft nicht zuverlässig. Um diese Probleme auszuschließen, installieren Sie von einem Betriebssystem, das direkt auf einem Gerät läuft. Virtuelle Maschinen sind außerdem oft so konfiguriert, dass sie sehr begrenzten Arbeitsspeicher und Speicherplatz haben.

Offiziell unterstützte Betriebssysteme für die Web-Installationsmethode:

* Windows 10
* Windows 11
* macOS Big Sur (11)
* macOS Monterey (12)
* macOS Ventura (13)
* macOS Sonoma (14)
* Arch Linux
* Debian 10 (buster)
* Debian 11 (bullseye)
* Debian 12 (bookworm)
* Ubuntu 20.04 LTS
* Ubuntu 22.04 LTS
* Ubuntu 23.04
* Ubuntu 23.10
* ChromeOS
* GrapheneOS
* Google Android (das Stock-Betriebssystem der Pixel-Smartphones) und andere zertifizierte Android-Varianten

Stellen Sie sicher, dass Ihr Betriebssystem auf dem neuesten Stand ist, bevor Sie fortfahren.

Offiziell unterstützte Browser für die Web-Installationsmethode:

* Chromium (außer unter Ubuntu, weil dort ein kaputtes Snap-Paket ohne funktionierendes WebUSB beiliegt)
* Vanadium (GrapheneOS)
* Google Chrome
* Microsoft Edge
* Brave (mit deaktivierten Brave Shields, da es die Speichernutzung auf einen niedrigen Wert begrenzt, um ein Fingerprinting des verfügbaren Speichers zu vermeiden)

Sie sollten Flatpak- und Snap-Versionen von Browsern vermeiden, da diese dafür bekannt sind, während des Installationsvorgangs Probleme zu verursachen.

Vergewissern Sie sich, dass Ihr Browser auf dem neuesten Stand ist, bevor Sie fortfahren.

Verwenden Sie keinen Inkognito-Modus oder andere private Browsing-Modi. Diese Modi verhindern in der Regel, dass das Web-Installationsprogramm über genügend Speicherplatz verfügt, um die heruntergeladene Version zu extrahieren.

Sie benötigen eines der offiziell unterstützten Geräte. Um sicherzustellen, dass das Gerät für die Installation von GrapheneOS freigeschaltet werden kann, sollten Sie netzbetreiber-spezifische Geräte vermeiden. Netzbetreiber-spezifische Varianten der Pixel-Smartphones verwenden dasselbe Stock-Betriebssystem und dieselbe Firmware, aber mit einer Carrier-ID ungleich Null, die bei der Herstellung auf die persistente Partition geflasht wurde. Die Carrier-ID aktiviert eine netzbetreiber-spezifische Konfiguration im Stock-Betriebssystem, einschließlich einer Deaktivierung der Netzbetreiber- und Bootloader-Entsperrung. Der Netzbetreiber kann dies möglicherweise aus der Ferne deaktivieren, aber seine Support-Mitarbeiter wissen das möglicherweise nicht und werden es wahrscheinlich auch nicht tun. Holen Sie sich ein netzbetreiberunabhängiges Gerät, um das Risiko und möglichen Ärger zu vermeiden. Wenn Sie einen Weg finden _können_, ein netzbetreiber-spezifisches Gerät zu entsperren, ist das kein Problem, da GrapheneOS die Carrier-ID einfach ignorieren kann und die Hardware die gleiche ist.

Es hat sich bewährt, das Gerät vor der Installation von GrapheneOS zu aktualisieren, um beim Anschluss des Gerätes an den Computer und in der frühen Phase des Installationsprozesses die neueste Firmware zu haben. So oder so, GrapheneOS flasht die neueste Firmware zu Beginn des Installationsprozesses.

<h2 id="enabling-oem-unlocking">OEM-Entsperrung aktivieren</h2>

Die OEM-Entsperrung muss innerhalb des Betriebssystems aktiviert werden.

Aktivieren Sie das Menü "Entwickleroptionen", indem Sie zu Einstellungen ➔ Über das Telefon/Tablet gehen und wiederholt auf den Menüeintrag "Build-Nummer" drücken, bis die Entwickleroptionen aktiviert sind.

Gehen Sie dann zu Einstellungen ➔ System ➔ Entwickleroptionen und aktivieren Sie die Einstellung "OEM-Entsperrung". Bei Gerätevarianten (SKUs), die von Netzbetreibern als gesperrte Geräte verkauft werden können, erfordert die Aktivierung der "OEM-Entsperrung" einen Internetzugang, damit das Stock-Betriebssystem prüfen kann, ob das Gerät von einem Netzbetreiber als gesperrt verkauft wurde.

Für das Pixel 6a wird die OEM-Entsperrung nicht mit der ab Werk installierten Version des Stock-Betriebssystems funktionieren. Sie müssen es mit einem Over-the-Air-Update auf die Version von Juni 2022 oder später aktualisieren. Nach der Aktualisierung müssen Sie das Gerät auf die Werkseinstellungen zurücksetzen, um die OEM-Entsperrung zu ermöglichen.

<h2 id="flashing-as-non-root">Flash als Nicht-Root-Benutzer</h2>

Bei traditionellen Linux-Distributionen können USB-Geräte ohne udev-Regeln für den jeweiligen Gerätetyp nicht als Nicht-Root verwendet werden. Für andere Plattformen ist das kein Problem.

Unter Arch Linux installieren Sie das Paket `android-udev`. Unter Debian und Ubuntu installieren Sie das Paket `android-sdk-platform-tools-common`.

<h2 id="booting-into-the-bootloader-interface">Booten in den Bootloader-Modus</h2>

Sie müssen Ihr Gerät in den Bootloader-Modus booten. Um das zu tun, müssen Sie die Leiser-Taste gedrückt lassen, während das Gerät bootet.

Am einfachsten ist es, das Gerät neu zu starten und die Leiser-Taste gedrückt zu halten, bis es im Bootloader-Modus startet.

Alternativ schalten Sie das Gerät aus und starten Sie es dann, wobei Sie die Leiser-Taste während des Bootvorgangs gedrückt halten. Sie können es entweder mit der Einschalttaste oder durch das Anschließen des Kabels starten, was im nächsten Abschnitt ohnehin nötig ist.

<h2 id="connecting-phone">Das Gerät verbinden</h2>

Verbinden Sie das Gerät mit dem Computer. Unter Linux müssen Sie das erneut tun, falls Sie die udev-Regeln beim Anschließen des Gerätes noch nicht eingerichtet hatten.

Unter Windows müssen Sie einen Treiber für fastboot installieren, sofern Sie ihn nicht bereits haben. Für andere Betriebssysteme ist kein Treiber erforderlich. Sie können den Treiber über Windows Update beziehen, das ihn als optionales Update erkennt, wenn das Gerät im Bootloader-Modus ist und mit dem Computer verbunden ist. Öffnen Sie Windows Update, führen Sie eine Prüfung auf Updates durch und gehen Sie danach auf "Optionale Updates anzeigen". Installieren Sie den Treiber für die Android-Bootloader-Schnittstelle als optionales Update.

Eine alternativer Weg, die Windows Fastboot-Treiber zu beziehen, besteht darin, den [neuesten Treiber für Pixels](https://developer.android.com/studio/run/win-usb) von Google herunterzuladen und ihn dann [manuell über den Windows-Geräte-Manager zu installieren](https://developer.android.com/studio/run/oem-usb#InstallingDriver).

<h2 id="unlocking-the-bootloader">Den Bootloader entsperren</h2>

Entsperren Sie den Bootloader, damit Sie das Betriebssystem und die Firmware flashen können:

Drücken Sie im Web-Installer unter [Unlocking the bootloader](https://grapheneos.org/install/web#unlocking-the-bootloader) die Schaltfläche <kbd>Unlock bootloader</kbd>.

Der Befehl muss auf dem Gerät bestätigt werden und löscht alle Daten. Verwenden Sie eine der Lautstärketasten, um Ihre Auswahl zu treffen, und die Einschalttaste, um diese zu bestätigen.

<h2 id="obtaining-factory-images">Factory Images beschaffen</h2>

Um mit dem Installationsprozess fortzufahren, müssen Sie die GrapheneOS Factory Image für Ihr Gerät beschaffen.

Drücken Sie im Web-Installer unter [Obtaining factory images](https://grapheneos.org/install/web#obtaining-factory-images) die Schaltfläche <kbd>Download release</kbd>.

<h2 id="flashing-factory-images">Factory Images flashen</h2>

Die Erstinstallation erfolgt durch das Flashen der Factory Images. Das wird das bestehende Betriebssystem ersetzen und alle vorhandenen Daten löschen.

Drücken Sie im Web-Installer unter [Locking the bootloader](https://grapheneos.org/install/web#flashing-factory-images) die Schaltfläche <kbd>Flashing factory images</kbd>.

Warten Sie, bis der Flashvorgang abgeschlossen ist. Die Firmware wird automatisch geflasht, das Gerät wird in den Bootloader-Modus gebootet, der Kern des Betriebssystems wird geflasht, das Gerät wird in den Userspace-Fastboot-Modus gebootet, der Rest des Betriebssystems wird geflasht und schließlich wird es wieder in den Bootloader-Modus gebootet. Vermeiden Sie es, mit dem Gerät zu interagieren, bis das Flash-Skript beendet ist und das Gerät wieder im Bootloader-Modus ist. <a href="#locking-the-bootloader">Sperren Sie im Anschluss den Bootloader</a>, bevor Sie das Gerät verwenden, da beim Sperren die Daten erneut gelöscht werden.

<h2 id="locking-the-bootloader">Den Bootloader sperren</h2>

Das Sperren des Bootloaders ist wichtig, da es einen vollständig verifizierten Bootvorgang ermöglicht. Es verhindert auch die Verwendung von fastboot zum Flashen, Formatieren oder Löschen von Partitionen. Ein verifizierter Bootvorgang erkennt Änderungen an den Betriebssystempartitionen und verhindert das Lesen von veränderten oder beschädigten Daten. Wenn Änderungen erkannt werden, wird mit Hilfe von Fehlerkorrekturdaten versucht, die ursprünglichen Daten zu erhalten, die dann erneut verifiziert werden, was das verifizierte Booten gegen nicht-bösartige Datenveränderungen abhärtet.

Während sich das Gerät im Bootloader-Modus befindet, setzen Sie es auf "locked":

Drücken Sie im Web-Installer unter [Locking the bootloader](https://grapheneos.org/install/web#locking-the-bootloader) die Schaltfläche <kbd>Lock bootloader</kbd>.

Der Befehl muss auf dem Gerät bestätigt werden und löscht alle Daten. Verwenden Sie eine der Lautstärketasten, um Ihre Auswahl zu treffen, und die Einschalttaste, um diese zu bestätigen.

<h2 id="post-installation">Nach der Installation</h2>

<h3 id="booting">Booten</h3>

Sie haben GrapheneOS nun erfolgreich installiert und können es booten. Wenn Sie die Einschalttaste drücken und die Standardoption "Start" im Bootloader-Modus ausgewählt lassen, wird das Betriebssystem gestartet.

<h3 id="disabling-oem-unlocking">OEM-Entsperrung deaktivieren</h3>

Die OEM-Entsperrung kann nach dem Neustart des Betriebssystems im Menü "Entwickleroptionen" wieder deaktiviert werden.

Nach der Deaktivierung der OEM-Entsperrung empfehlen wir, die Entwickleroptionen insgesamt zu deaktivieren, falls das Gerät nicht zur App- oder Betriebssystementwicklung verwendet wird.

<h3 id="verifying-installation">Überprüfung der Installation</h3>

Die von den unterstützten Geräten bereitgestellten Funktionen "Verified Boot" und "Attestation" können verwendet werden, um zu überprüfen, ob die Hardware, die Firmware und die GrapheneOS-Installation unverfälscht sind. Selbst wenn der Computer, den Sie zum Flashen von GrapheneOS verwendet haben, kompromittiert wurde und ein Angreifer GrapheneOS durch sein eigenes bösartiges Betriebssystem ersetzt hat, kann dies mit diesen Funktionen erkannt werden.

Verified Boot überprüft bei jedem Startvorgang sämtliche Firmware- und Betriebssystem-Images. Der öffentliche Schlüssel für die Firmware-Images ist werkseitig in elektronische Sicherungen im SoC eingebrannt. Firmware-Sicherheitsupdates können auch den in die Sicherungen eingebrannten Rollback-Index aktualisieren, um einen Rollback-Schutz zu gewährleisten.

Die letzte Bootphase der Firmware vor dem Betriebssystem ist für die Verifizierung des Betriebssystems zuständig. Für das Stock-Betriebssystem wird ein fest eingebauter öffentlicher Schlüssel verwendet. Bei der Installation von GrapheneOS wird der öffentliche Verified-Boot-Schlüssel von GrapheneOS in den Sicherheitschip geflasht. Bei jedem Start wird dieser Schlüssel geladen und zur Verifizierung des Betriebssystems verwendet. Sowohl für das Stock-Betriebssystem als auch für GrapheneOS wird ein auf dem Sicherheits-Patch-Level basierender Rollback-Index aus dem Sicherheitschip geladen, um einen Rollback-Schutz zu gewährleisten.

<h4 id="verified-boot-key-hash">Hash des öffentlichen Verified-Boot-Schlüssels</h4>

Beim Laden eines alternativen Betriebssystems zeigt das Gerät beim Booten einen gelben Hinweis mit der ID des alternativen Betriebssystems an, die auf dem SHA256-Hashwert des öffentlichen Verified-Boot-Schlüssels basiert. Bei Pixels der 4. und 5. Generation werden nur die ersten 32 Bits des Hashwerts angezeigt, sodass Sie diesen Ansatz nicht verwenden können. Ab Pixels der 6. Generation wird der vollständigen Hashwert angezeigt, den Sie mit den offiziellen Verified-Boot-Hashwerten von GrapheneOS vergleichen können:

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

Diese Überprüfung ist nach der Installation sinnvoll, aber Sie müssen sie nicht manuell vornehmen, damit verifiziertes Booten funktioniert. Der öffentliche Schlüssel für verifiziertes Booten, der auf dem Sicherheitschip gespeichert wurde, kann nur geändert werden, wenn das Gerät entsperrt ist. Das Entsperren [des Bootloaders, Anm. d. Übers.] des Geräts löscht den Sicherheitschip auf die selbe Weise wie ein Zurücksetzen auf Werkseinstellungen und verhindert ein Wiederherstellen von Daten, selbst wenn die SSD geklont wurde und Ihre Passphrase(n) erlangt wurde(n), da die Verschlüsselungsschlüssel nicht mehr abgeleitet werden können. Der Verified-Boot-Schlüssel ist außerdem einer der Inputs für die Ableitung der Verschlüsselungsschlüssel, zusätzlich zu den Sperrmethode(n) des Benutzers und zufälligen Token(s) auf dem Sicherheitschip.

<h4 id="hardware-based-attestation">Hardware-gestützte Attestierung</h4>

GrapheneOS enthält unsere Auditor-App, die eine Kombination aus Verified-Boot- und Attestierungsfunktionen verwendet, um die Unverfälschtheit der Hardware, der Firmware und des Betriebssystems zu überprüfen und weitere nützliche Daten der Hardware und des Betriebssystems bereitzustellen.

Da der Zweck von Auditor darin besteht, Informationen über das Gerät zu erhalten, ohne dem Gerät vertrauen zu müssen, werden die Ergebnisse nicht auf dem zu überprüfenden Gerät angezeigt. Sie benötigen ein zweites Android-Gerät, auf dem der Auditor zur lokalen QR-Code-basierte Überprüfung läuft. Sie können auch unseren optionalen Dienst zur Überwachung der Geräteintegrität für automatische, planmäßige Überprüfungen mit Unterstützung für E-Mail-Warnungen verwenden.

Eine Anleitung dazu finden Sie im [Auditor-Tutorial](https://attestation.app/tutorial) [[deutsche Übersetzung](grapheneos-attestation-tutorial.html)].

Auditor basiert in erster Linie auf einem Pairing-Modell, bei dem ein hardwaregestützter Signierschlüssel und ein hardwaregestützter Attestierungs-Signierschlüssel generiert und als Teil der ersten Verifizierung festgehalten werden. Die erste Verifizierung entsteht auf Basis einer Chain of Trust zu einer der Android-Attestierungs-Root-Keys. Nach der ersten Verifizierung bietet er ein hochgradig sicheres System für den Erhalt weiterer Informationen über das Gerät. Ein Angreifer könnte die erste Verifizierung mit einem geleakten Attestation-Key umgehen oder durch die Weiterleitung an ein anderes Gerät mit dem Gerätemodell, dem Betriebssystem und dem Patch-Level, das der Benutzer erwartet. Das Weiterleiten auf ein anderes Gerät wird in Zukunft mit einer optionalen Unterstützung für die Attestierung der Hardware-Seriennummer angegangen werden.

<h3 id="replacing-grapheneos-with-the-stock-os">GrapheneOS durch das Standardbetriebssystem ersetzen</h3>

Die Installation des Stock-Betriebssystems mittels Factory Images ist ähnlich wie der oben beschriebene Prozess, aber mit [Googles Web-Flashing-Tool](https://flash.android.com/back-to-public). Vor dem Flashen und Sperren gibt es jedoch einen zusätzlichen Schritt, um das Gerät vollständig in einen sauberen Werkszustand zurückzusetzen.

Die GrapheneOS Factory Images flashen einen nicht-standardmäßigen Android-Verified-Boot-Schlüssel, der gelöscht werden muss, um das Gerät vollständig in den Werkszustand zurückzuversetzen. Bevor Sie die Factory Images flashen und den Bootloader sperren, sollten Sie den benutzerdefinierten Android-Verified-Boot-Schlüssel löschen, um ihm nicht länger zu vertrauen:

Drücken Sie im Web-Installer unter [Replacing GrapheneOS with the stock OS](https://grapheneos.org/install/web#replacing-grapheneos-with-the-stock-os) die Schaltfläche <kbd>Remove non-stock key</kbd>.

<h3 id="further-information">Weitere Informationen</h3>

Bitte schauen Sie sich das [Benutzerhandbuch](https://grapheneos.org/usage) und die [FAQ](https://grapheneos.org/faq) für weitere Informationen an. Wenn Sie weitere Fragen haben, die nicht von der Website abgedeckt werden, treten Sie den [offiziellen GrapheneOS-Chat-Kanälen](https://grapheneos.org/contact#community) bei und stellen Sie Ihre Fragen in dem entsprechenden Kanal.
