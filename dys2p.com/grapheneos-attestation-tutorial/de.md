<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item active" aria-current="page">Auditor-App-Anleitung</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/AttestationServer/blob/bcad62e7aa75e593e63e01348e40ed11d7930cb0/static/tutorial.html">bcad62e</a> vom 2021-11-27. Falls du Hinweise oder Verbesserungsvorschläge hast, dann <a href="contact.html">schreib uns gerne</a> oder arbeite mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-attestation-tutorial/de.md">GitHub</a> an dieser Übersetzung.
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

<h1 id="tutorial">Anleitung</h1>

<h2 id="installation">Installation</h2>

Auditor kann entweder über den [Google Play Store](https://play.google.com/store/apps/details?id=app.attestation.auditor) oder über die [Releases auf GitHub](https://github.com/GrapheneOS/Auditor/releases) installiert werden. Die Versionen aus beiden Quellen sind identisch und verwenden denselben Signierschlüssel, sodass Sie Updates aus beiden Quellen beziehen können, unabhängig davon, wie Sie die App installiert haben.

[GrapheneOS](https://grapheneos.org/) hat Auditor bereits vorinstalliert.

<h2 id="local-verification">Lokale Überprüfung</h2>

Das zu überprüfende Gerät (Auditee) muss eines der unterstützten Geräte sein. Android Developer Previews werden nicht unterstützt, da die verifizierte Hardwareversion bei ihnen auf einen Platzhalterwert gesetzt ist. Das Gerät, das die Überprüfung durchführt (Auditor), muss ein beliebiges Android 8.0+-kompatibles Gerät mit einer Kamera sein.

1. Drücken Sie "Auditor" auf dem Gerät, das die Überprüfung des Auditee durchführen soll.
2. Drücken Sie "Auditee" auf dem Gerät, das überprüft werden soll.
3. Richten Sie die Kamera des Auditees auf den QR-Code des Auditor-Geräts, um die [kryptographische, Anm. d. Übers.] Challenge zu lesen.
4. Tippen Sie auf den QR-Code auf dem Auditor-Gerät, um fortzufahren (wenn Sie dies zu früh tun, können Sie zurück gehen).
5. Richten Sie die Kamera des Auditor-Geräts auf den QR-Code des Auditees, um die Attestierung zu lesen.
6. Sie können das Ergebnis der Attestierung einsehen.

Ein Auditor-Gerät kann beliebig viele Auditees überprüfen. Es zeigt einen Fingerabdruck und die Zeitpunkte der ersten und letzten erfolgreichen Attestierung des gepaarten Geräts in den Ergebnissen. Ein Auditee kann von einer beliebigen Anzahl von Auditoren verifiziert werden, aber für jede eindeutige Paarung wird ein anderer Fingerabdruck angezeigt und nicht derselbe Fingerabdruck bei jedem Auditor für denselben Auditee.

<h2 id="scheduled-remote-verification">Planmäßige Remote Verification</h2>

So richten Sie eine regelmäßige Remote Verification über den Remote Attestation Service ein:

1. Erstellen Sie von einem anderen Gerät einen Account auf [https://attestation.app](https://attestation.app).
2. Drücken Sie den Menü-Button in der App.
3. Drücken Sie im Menü die Funktion "Enable remote verification".
4. Scannen Sie den QR-Code Ihres Kontos, der auf https://attestation.app angezeigt wird.
5. Geben Sie eine E-Mail-Adresse für den Empfang von Warnmeldungen an, falls das Gerät nicht rechtzeitig ein erfolgreiches Attestierungsergebnis liefert.
6. Aktualisieren Sie https://attestation.app, um das Ergebnis der ersten Attestierung anzuzeigen.

<h2 id="expanding-device-support">Ausbau der Geräteunterstützung</h2>

Die Unterstützung für die Verifizierung eines Gerätemodells muss der App auf Grundlage von mindestens einem gültigen Key Attestation Sample des Standardbetriebssystems mit gesperrtem Bootloader hinzugefügt werden. Die Auditor-App kann theoretisch die Verifizierung aller Android-Geräte unterstützen, die mit Android 8 oder höher auf den Markt gekommen sind. Ein Upgrade auf Android 8 ist nicht ausreichend, da die Hardwareunterstützung der Attestierung erforderlich, deren Mindestanforderungen erst mit Android 8 verbindlich wurden.

Um ein Sample einzureichen, öffnen Sie das Menü in der Aktionsleiste und wählen Sie "Submit sample data". Dadurch werden eine Attestierungs-Stichprobe und Geräteinformationen übermittelt, was die Entwicklung der Unterstützung für das Gerätemodell ermöglicht. Es kann ein paar Wochen dauern, bis die Unterstützung in einer neuen Version der App bereitgestellt wird.

Eine gültiges Key Attestation Sample für jedes Gerätemodell zusammen mit einer auf das Wesentliche reduzierten Teilmenge der Systemeigenschaften ist [auf GitHub](https://github.com/GrapheneOS/AttestationSamples) veröffentlicht, um anderen Projekten zu helfen, die an der Verwendung von schlüsselbasierter Attestierung interessiert sind.
