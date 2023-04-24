<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item active" aria-current="page">Über die Auditor-App</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/AttestationServer/blob/2482f029190640a3474ca7ee212fe1ac2bfa8d2d/static/about.html">2482f02</a> vom 2023-04-15. Falls du Hinweise oder Verbesserungsvorschläge hast, dann <a href="contact.html">schreib uns gerne</a> oder arbeite mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-attestation-about/de.md">GitHub</a> an dieser Übersetzung.
</div>

<!--
Copyright © 2018-2023 GrapheneOS

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

<h1 id="about">Über die Auditor-App</h1>

Die [Auditor-App](https://play.google.com/store/apps/details?id=app.attestation.auditor) verwendet hardwarebasierte Sicherheitsfunktionen, um die Identität eines Geräts sowie die Authentizität und Integrität des Betriebssystems zu überprüfen. Sie prüft, ob auf dem Gerät das Stock-Betriebssystem mit gesperrtem Bootloader läuft und dass das Betriebssystem nicht manipuliert wurde. Ein Downgrade auf eine frühere Version wird ebenfalls erkannt. Sie baut auf der hardwarebasierten Überprüfung des Betriebssystems auf, welche sie mit der App verknüpft, um softwarebasierte Plausibilitätsprüfungen durchzuführen und zusätzliche Informationen über den Gerätestatus und die Konfiguration zu sammeln, die über das hinausgehen, was die Hardware direkt attestieren kann.

Die Grundlage der Auditor-App ist die Erzeugung eines dauerhaften Schlüssels im [hardwaregestützten Schlüsselspeicher](https://source.android.com/security/keystore/), um die Identität des Geräts zu überprüfen und mittels Verified Boot sicherzustellen, dass das Betriebssystem nicht manipuliert oder downgegraded wurde. Es führt einen Paarungsprozess zwischen dem Gerät, das die Verifizierung durchführt (Auditor), und dem zu verifizierenden Gerät (Auditee) durch, um ein Trust-On-First-Use-Modell (TOFU) zu implementieren. Das Gerät, das die Verifizierung durchführt, kann entweder ein anderes Android-Gerät sein, auf dem die App im Auditor-Modus läuft, oder der Dienst [https://attestation.app/](https://attestation.app) für automatische, planmäßige Überprüfungen mit Unterstützung für E-Mail-Warnungen. Schauen Sie sich dafür die [Auditor-Anleitung](https://attestation.app/tutorial) [[deutsche Übersetzung](grapheneos-attestation-tutorial.html)]  an. Das Protokoll, das sowohl für die lokale als auch für die Remote-Attestierung verwendet wird, ist [im Quellcode dokumentiert](https://github.com/GrapheneOS/Auditor/blob/52/app/src/main/java/app/attestation/auditor/AttestationProtocol.java#L128-L202).

[Verified Boot](https://source.android.com/security/verifiedboot/) validiert die Integrität und Authentizität der Firmware und des gesamten Betriebssystems (Kernel und Userspace) anhand eines unveränderlichen Hardware-Root-of-Trust. Die Ergebnisse werden an den Hardware-gestützten Schlüsselspeicher weitergeleitet und für den Schutz der Schlüssel verwendet.

Die [Key-Attestation](https://source.android.com/security/keystore/attestation)-Funktion des hardwaregestützten Schlüsselspeichers bietet direkte Unterstützung für die Attestierung von Geräteeigenschaften und das Bootstrapping des Trust-On-First-Use-Modells der Auditor-App durch eine grundlegende erste Verifizierung, die mit einem bekannten Root-Zertifikat verkettet wird. Die neueste Version der Key-Attestation liefert ein signierte Ausgabe mit u. a. dem Verified-Boot-Status, dem Verified-Boot-Schlüssel, einem Hash aller durch Verified Boot geschützten Daten und die Versionsangaben der Betriebssystempartitionen. Sie unterstützt auch die Trust-Verkettung mit der Anwendung, die die Attestierung durchführt, was der Auditor-Anwendung als Grundlage für die Prüfungen auf der Softwareebene dient.

Geräte, die mit Android 9 oder höher ausgeliefert werden, _können_ eine [StrongBox-Keymaster](https://developer.android.com/training/articles/keystore#HardwareSecurityModule)-Implementierung enthalten, die es der Auditor-App ermöglicht, die vom Attestierungsprotokoll verwendeten Schlüssel im dedizierten Hardware-Sicherheitsmodul (HSM) zu speichern (wie etwa im [Titan M](https://android-developers.googleblog.com/2018/10/building-titan-better-security-through.html) in Pixel-Telefonen), anstatt das Trusted Execution Environment (TEE) auf dem Hauptprozessor zu verwenden. Dies kann die Angriffsfläche erheblich reduzieren.

Die Sicherheitsverbesserungen, die künftige Hardware-Generationen und künftige Android-Versionen bieten, werden von diesen Projekten genau verfolgt. Der Kern-Workflow und die Funktionen sind bereits implementiert, aber die Grundlage wird regelmäßig verbessert, ebenso wie die Benutzeroberfläche und die Dokumentation. Die App und der Dienst sind so konzipiert, dass sie mittels eines versionierten Protokolls auf- und abwärtskompatibel sind, um spätere wesentliche Änderungen zu ermöglichen.

<h2 id="device-support">Unterstützte Geräte</h2>

Jedes Gerät mit Android 10 oder höher kann die Auditor-App ausführen und sie zur Überprüfung anderer Geräte verwenden. Allerdings haben nur Geräte, die mit Android 8.0 oder höher _auf den Markt gebracht_ wurden, die notwendige Hardware-Unterstützung, um verifiziert zu werden. Außerdem muss jedes Gerätemodell explizit in die App integriert werden. Die folgenden Geräte werden derzeit von der neuesten stabilen Version unterstützt:

* BlackBerry Key2 (BBF100-1 and BBF100-6 models)
* BQ Aquaris X2 Pro
* Google Pixel 2
* Google Pixel 2 XL
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
* Huawei Honor 7A Pro (AUM-L29 model)
* Honor 9 Lite (LLD-L31 model)
* Huawei Honor 10 (COL-L29 model)
* Huawei Honor View 10 (BKL-L04 and BKL-L09 models)
* Huawei Mate 10 (ALP-L29 model)
* Huawei Mate 20 lite (SNE-LX1 model)
* Huawei Mate 20 Pro (LYA-L29 model)
* Huawei P smart 2019 (POT-LX3 model)
* Huawei P20 (EML-L09 model)
* Huawei P20 Pro (CLT-L29 model)
* Huawei Y7 2019 (DUB-LX3 model)
* Huawei Y9 2019 (JKM-LX3 model)
* HTC EXODUS 1
* HTC U12+
* LG Stylo 5 (LM-Q720 model)
* LG Q Stylo 4 (LG-Q710AL model)
* Motorola moto g⁷
* Motorola One Vision
* Nokia 3.1
* Nokia 6.1
* Nokia 6.1 Plus
* Nokia 7.1
* Nokia 7 Plus
* OnePlus 6 (A6003 model)
* OnePlus 6T (A6013 model)
* OnePlus 7 Pro (GM1913 model)
* Oppo R15 Pro (CPH1831 model)
* Oppo A7 (CPH1903 model)
* Oppo A5s (CPH1909 model)
* Realme C2 (RMX1941 model)
* Samsung Galaxy A70 (SM-A705FN model)
* Samsung Galaxy Amp Prime 3 (SM-J337AZ model)
* Samsung Galaxy J2 Core (SM-J260A, SM-J260F and SM-J260T1 models)
* Samsung Galaxy J3 2018 (SM-J337A and SM-J337T models)
* Samsung Galaxy J7 (SM-J737T1 model)
* Samsung Galaxy M20 (SM-M205F model)
* Samsung Galaxy Note 9 (SM-N960F and SM-N960U models)
* Samsung Galaxy Note 10 (SM-N970F and SM-N970U models)
* Samsung Galaxy Note 10+ (SM-N975U model)
* Samsung Galaxy S9 (SM-G960F, SM-G960U, SM-G960U1, SM-G960W and SM-G9600 models)
* Samsung Galaxy S9+ (SM-G965F, SM-G965U, SM-G965U1 and SM-G965W models)
* Samsung Galaxy S10e (SM-G970F model)
* Samsung Galaxy S10+ (SM-G975F model)
* Samsung Galaxy Tab A 10.1 (SM-T510 model)
* Samsung Galaxy Tab S4 (SM-T835 model)
* Sony Xperia XA2 (H3113, H3123 and H4113 models)
* Sony Xperia XZ1 / XZ1 Compact (G8341 and G8342 models)
* Sony Xperia XZ1 Compact (G8441 model)
* Sony Xperia XZ2 (H8216 model)
* Sony Xperia XZ2 Compact (H8314 and H8324 models)
* T-Mobile REVVL 2
* Vivo 1807
* Xiaomi Mi A2
* Xiaomi Mi A2 Lite
* Xiaomi Mi 9
* Xiaomi POCOPHONE F1

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
* Samsung Galaxy Note 10 (SM-N970U model)
* Samsung Galaxy Note 10+ (SM-N975U model)

Die folgenden Geräte unterstützen die Attest-Key-Funktion zur Erzeugung eines paarungsspezifischen Signierschlüssels für die Attestierung:

* Google Pixel 6
* Google Pixel 6 Pro
* Google Pixel 6a
* Google Pixel 7
* Google Pixel 7 Pro

Die Auditor-App unterstützt auch die Überprüfung von alternativen Betriebssystemen auf unterstützten Geräten. Sie kann [GrapheneOS](https://grapheneos.org/) auf den folgenden Geräten überprüfen:

* Google Pixel 2
* Google Pixel 2 XL
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

Bei alternativen Betriebssystemen muss der Verified-Boot-Schlüssel in der Auditor-App bzw. dem Attestierungs-Server enthalten sein. Die App und der Dienst zeigen den Namen des Betriebssystems an, das auf dem Gerät verifiziert wird. Leider bieten die meisten alternativen Betriebssysteme keine Unterstützung für vollständig verifiziertes Booten und die meisten Geräte unterstützen die Verwendung von verifiziertem Booten mit einem benutzerdefinierten Schlüssel nicht. Die App ist außerdem davon abhängig, dass das Betriebssystem das grundlegende Sicherheitsmodell für Erweiterungen über die grundlegende hardwarebasierte Attestierungsunterstützung hinaus beibehält.

GrapheneOS ist ein gehärtetes mobiles Betriebssystem mit Android-App-Kompatibilität, das sich auf die Erforschung und Entwicklung von Privatsphäre- und Sicherheitstechnologien konzentriert, einschließlich wesentlicher Verbesserungen bei Sandboxing, der Vermeidung von Exploits und dem Berechtigungsmodell. GrapheneOS behält auch alle Standard-Sicherheitsfunktionen bei. Versionen sind [auf der GrapheneOS-Unterseite "Releases"](https://grapheneos.org/releases) verfügbar und können mit der Auditor-App und dem Auditor-Server verwendet werden.
