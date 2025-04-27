<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item active" aria-current="page">Über die Auditor-App</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/AttestationServer/blob/f366591b639c2ae339913c630fff5590f1dd900e/static/about.html">f366591</a> vom 2025-04-21. Falls du Hinweise oder Verbesserungsvorschläge hast, dann <a href="contact.html">schreib uns gerne</a> oder arbeite mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-attestation-about/de.md">GitHub</a> an dieser Übersetzung.
</div>

<!--
Copyright © 2018-2025 GrapheneOS

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

https://attestation.app/LICENSE.txt
-->

<h1 id="about">Über die Auditor-App</h1>

Die [Auditor-App](https://play.google.com/store/apps/details?id=app.attestation.auditor.play) verwendet hardwarebasierte Sicherheitsfunktionen, um die Identität eines Geräts sowie die Authentizität und Integrität des Betriebssystems zu überprüfen. Sie stellt sicher, dass auf dem Gerät ein verifiziertes Betriebssystem mit gesperrtem Bootloader läuft und dass keine Manipulationen vorgenommen wurden. Auditor erkennt auch ein Downgrade auf eine frühere Betriebssystem-, Patch- oder App-Version. Auditor erweitert die hardwarebasierte Betriebssystemüberprüfung, indem es die Überprüfung an die App bindet, softwarebasierte Plausibilitätsprüfungen durchführt und zusätzliche Informationen über den Gerätestatus und die Konfiguration sammelt, die über das hinausgehen, was die Hardware direkt attestieren kann.

Die Grundlage der Auditor-App ist die Erzeugung eines dauerhaften Schlüssels im [hardwaregestützten Schlüsselspeicher](https://source.android.com/security/keystore/), um die Identität des Geräts zu überprüfen und mittels Verified Boot sicherzustellen, dass das Betriebssystem nicht manipuliert oder downgegraded wurde. Es führt einen Paarungsprozess zwischen dem Gerät, das die Verifizierung durchführt (Auditor), und dem zu verifizierenden Gerät (Auditee) durch, um ein Trust-On-First-Use-Modell (TOFU) zu implementieren. Das Gerät, das die Verifizierung durchführt, kann entweder ein anderes Android-Gerät sein, auf dem die App im Auditor-Modus läuft, oder der Dienst [https://attestation.app/](https://attestation.app) für automatische, planmäßige Überprüfungen mit Unterstützung für E-Mail-Warnungen. Schauen Sie sich dafür die [Auditor-Anleitung](https://attestation.app/tutorial) [[deutsche Übersetzung](grapheneos-attestation-tutorial.html)]  an. Das Protokoll, das sowohl für die lokale als auch für die Remote-Attestierung verwendet wird, ist [im Quellcode dokumentiert](https://github.com/GrapheneOS/Auditor/blob/80/app/src/main/java/app/attestation/auditor/AttestationProtocol.java#L120-L193).

[Verified Boot](https://source.android.com/security/verifiedboot/) validiert die Integrität und Authentizität der Firmware und des gesamten Betriebssystems (Kernel und Userspace) anhand eines unveränderlichen Hardware-Root-of-Trust. Die Ergebnisse werden an den Hardware-gestützten Schlüsselspeicher weitergeleitet und für den Schutz der Schlüssel verwendet.

Die [Key-Attestation](https://source.android.com/security/keystore/attestation)-Funktion des hardwaregestützten Schlüsselspeichers bietet direkte Unterstützung für die Attestierung von Geräteeigenschaften und das Bootstrapping des Trust-On-First-Use-Modells der Auditor-App durch eine grundlegende erste Verifizierung, die mit einem bekannten Root-Zertifikat verkettet wird. Die neueste Version der Key-Attestation liefert ein signierte Ausgabe mit u. a. dem Verified-Boot-Status, dem Verified-Boot-Schlüssel, einem Hash aller durch Verified Boot geschützten Daten und die Versionsangaben der Betriebssystempartitionen. Sie unterstützt auch die Trust-Verkettung mit der Anwendung, die die Attestierung durchführt, was der Auditor-Anwendung als Grundlage für die Prüfungen auf der Softwareebene dient.

Geräte, die mit Android 9 oder höher ausgeliefert werden, _können_ eine [StrongBox-Keymaster](https://developer.android.com/training/articles/keystore#HardwareSecurityModule)-Implementierung enthalten, die es der Auditor-App ermöglicht, die vom Attestierungsprotokoll verwendeten Schlüssel im dedizierten Hardware-Sicherheitsmodul (HSM) zu speichern (wie etwa im [Titan M](https://android-developers.googleblog.com/2018/10/building-titan-better-security-through.html) in Pixel-Geräten), anstatt das Trusted Execution Environment (TEE) auf dem Hauptprozessor zu verwenden. Dies kann die Angriffsfläche erheblich reduzieren.

Die Sicherheitsverbesserungen, die künftige Hardware-Generationen und künftige Android-Versionen bieten, werden von diesen Projekten genau verfolgt. Der Kern-Workflow und die Funktionen sind bereits implementiert, aber die Grundlage wird regelmäßig verbessert, ebenso wie die Benutzeroberfläche und die Dokumentation. Die App und der Dienst sind so konzipiert, dass sie mittels eines versionierten Protokolls auf- und abwärtskompatibel sind, um spätere wesentliche Änderungen zu ermöglichen.

<h2 id="device-support">Unterstützte Geräte</h2>

Jedes Gerät mit Android 10 oder höher kann die Auditor-App ausführen und sie zur Überprüfung anderer Geräte verwenden. Allerdings haben nur Geräte, die mit Android 8.0 oder höher _auf den Markt gebracht_ wurden, die notwendige Hardware-Unterstützung, um verifiziert zu werden. Außerdem muss jedes Gerätemodell explizit in die App integriert werden. Die folgenden Geräte werden derzeit von der neuesten stabilen Version unterstützt:

* Google Pixel 3
* Google Pixel 3 XL
* Google Pixel 3a
* Google Pixel 3a XL
* Google Pixel 4
* Google Pixel 4 XL
* Google Pixel 4a
* Google Pixel 4a (5G)
* Google Pixel 5
* Google Pixel 5a
* Google Pixel 6
* Google Pixel 6 Pro
* Google Pixel 6a
* Google Pixel 7
* Google Pixel 7 Pro
* Google Pixel 7a
* Google Pixel Tablet
* Google Pixel Fold
* Google Pixel 8
* Google Pixel 8 Pro
* Google Pixel 8a
* Google Pixel 9
* Google Pixel 9 Pro
* Google Pixel 9 Pro XL
* Google Pixel 9 Pro Fold
* Google Pixel 9a
* OnePlus 7 Pro (GM1913 model)
* Samsung Galaxy Note 10 (SM-N970F and SM-N970U models)
* Samsung Galaxy Note 10+ (SM-N975U model)
* Samsung Galaxy S10e (SM-G970F model)
* Samsung Galaxy S10+ (SM-G975F model)

Die folgenden Geräte bieten ein HSM mit StrongBox-Unterstützung, das von Auditor unterstützt wird:

* Google Pixel 3
* Google Pixel 3 XL
* Google Pixel 3a
* Google Pixel 3a XL
* Google Pixel 4
* Google Pixel 4 XL
* Google Pixel 4a
* Google Pixel 4a (5G)
* Google Pixel 5
* Google Pixel 5a
* Google Pixel 6
* Google Pixel 6 Pro
* Google Pixel 6a
* Google Pixel 7
* Google Pixel 7 Pro
* Google Pixel 7a
* Google Pixel Tablet
* Google Pixel Fold
* Google Pixel 8
* Google Pixel 8 Pro
* Google Pixel 8a
* Google Pixel 9
* Google Pixel 9 Pro
* Google Pixel 9 Pro XL
* Google Pixel 9 Pro Fold
* Google Pixel 9a
* Samsung Galaxy Note 10 (SM-N970U model)
* Samsung Galaxy Note 10+ (SM-N975U model)

Die folgenden Geräte unterstützen die Attest-Key-Funktion zur Erzeugung eines paarungsspezifischen Signierschlüssels für die Attestierung:

* Google Pixel 6
* Google Pixel 6 Pro
* Google Pixel 6a
* Google Pixel 7
* Google Pixel 7 Pro
* Google Pixel 7a
* Google Pixel Tablet
* Google Pixel Fold
* Google Pixel 8
* Google Pixel 8 Pro
* Google Pixel 8a
* Google Pixel 9
* Google Pixel 9 Pro
* Google Pixel 9 Pro XL
* Google Pixel 9 Pro Fold
* Google Pixel 9a

Die Auditor-App unterstützt auch die Überprüfung von alternativen Betriebssystemen auf unterstützten Geräten. Sie kann [GrapheneOS](https://grapheneos.org/) auf den folgenden Geräten überprüfen:

* Google Pixel 3
* Google Pixel 3 XL
* Google Pixel 3a
* Google Pixel 3a XL
* Google Pixel 4
* Google Pixel 4 XL
* Google Pixel 4a
* Google Pixel 4a (5G)
* Google Pixel 5
* Google Pixel 5a
* Google Pixel 6
* Google Pixel 6 Pro
* Google Pixel 6a
* Google Pixel 7
* Google Pixel 7 Pro
* Google Pixel 7a
* Google Pixel Tablet
* Google Pixel Fold
* Google Pixel 8
* Google Pixel 8 Pro
* Google Pixel 8a
* Google Pixel 9
* Google Pixel 9 Pro
* Google Pixel 9 Pro XL
* Google Pixel 9 Pro Fold
* Google Pixel 9a

Bei alternativen Betriebssystemen muss der Verified-Boot-Schlüssel in der Auditor-App bzw. dem Attestierungs-Server enthalten sein. Die App und der Dienst zeigen den Namen des Betriebssystems an, das auf dem Gerät verifiziert wird. Leider bieten die meisten alternativen Betriebssysteme keine Unterstützung für vollständig verifiziertes Booten und die meisten Geräte unterstützen die Verwendung von verifiziertem Booten mit einem benutzerdefinierten Schlüssel nicht. Die App ist außerdem davon abhängig, dass das Betriebssystem das grundlegende Sicherheitsmodell für Erweiterungen über die grundlegende hardwarebasierte Attestierungsunterstützung hinaus beibehält.

GrapheneOS ist ein gehärtetes mobiles Betriebssystem mit Android-App-Kompatibilität, das sich auf die Erforschung und Entwicklung von Privatsphäre- und Sicherheitstechnologien konzentriert, einschließlich wesentlicher Verbesserungen bei Sandboxing, der Vermeidung von Exploits und dem Berechtigungsmodell. GrapheneOS behält auch alle Standard-Sicherheitsfunktionen bei. Versionen sind [auf der GrapheneOS-Unterseite "Releases"](https://grapheneos.org/releases) verfügbar und können mit der Auditor-App und dem Auditor-Server verwendet werden.
