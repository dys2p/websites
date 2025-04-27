<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item active" aria-current="page">GrapheneOS</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/grapheneos.org/blob/86173d96cb235a9ed6960a8b63a0326808e13033/static/index.html">86173d9</a> vom 2024-11-13. Falls du Hinweise oder Verbesserungsvorschläge hast, dann <a href="contact.html">schreib uns gerne</a> oder arbeite mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-index/de.md">GitHub</a> an dieser Übersetzung.
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

<h1 id="grapheneos">GrapheneOS</h1>

Das private und sichere mobile Betriebssystem mit Android-App-Kompatibilität. Entwickelt als nicht gewinnorientiertes Open-Source-Projekt.

[GrapheneOS installieren](https://grapheneos.org/install/) [[deutsche Übersetzung](grapheneos-install.html)]

<h2 id="about">Über GrapheneOS</h2>

GrapheneOS ist ein auf Privatsphäre und Sicherheit ausgerichtetes mobiles Betriebssystem mit Android-App-Kompatibilität, das als gemeinnütziges [Open-Source](https://grapheneos.org/source)-Projekt entwickelt wurde. Es konzentriert sich auf die Erforschung und Entwicklung von Privatsphäre- und Sicherheitstechnologien, einschließlich wesentlicher Verbesserungen bei Sandboxing, der Vermeidung von Exploits und dem Berechtigungsmodell. Es wurde im Jahr 2014 gegründet und war [früher als CopperheadOS](https://grapheneos.org/history/copperheados) bekannt.

GrapheneOS verbessert die Privatsphäre und die Sicherheit des Betriebssystems von Grund auf. Es setzt Technologien ein, um ganze Klassen von Schwachstellen zu entschärfen und das Ausnutzen der häufigsten Ursachen von Schwachstellen erheblich zu erschweren. Es verbessert die Sicherheit sowohl des Betriebssystems als auch der darauf ausgeführten Anwendungen. Die App-Sandbox und andere Sicherheitsgrenzen werden verstärkt. GrapheneOS versucht zu vermeiden, dass die Privatsphäre- und Sicherheitsfunktionen das Benutzererlebnis beeinträchtigen. Im Idealfall können die Funktionen so gestaltet werden, dass sie immer aktiviert sind, ohne die Benutzerfreundlichkeit zu beeinträchtigen und ohne zusätzliche Komplexität wie etwa Konfigurationsoptionen. Das ist nicht immer möglich, und GrapheneOS fügt verschiedene Umschalter für Funktionen wie die Netzwerkberechtigung, Sensor-Berechtigungen, Einschränkungen bei gesperrtem Gerät (USB-C / Federkontaktstifte, Kamera, Schnellzugriff) usw. zusammen mit komplexeren benutzerseitigen Privatsphäre- und Sicherheitsfunktionen mit eigener Benutzeroberfläche hinzu.

Die [Unterseite "Features"](https://grapheneos.org/features) bietet einen Überblick über die wesentlichen Verbesserungen der Privatsphäre und der Sicherheit, die GrapheneOS dem Android Open Source Project (AOSP) hinzufügt. Viele unserer früheren Funktionen sind [in AOSP, Linux und andere Projekte eingeflossen, um die Privatsphäre und Sicherheit für Milliarden von Nutzern zu verbessern](https://grapheneos.org/faq#upstream), sodass sie nicht mehr auf unserer Funktionsseite aufgeführt sind.

Die offiziellen Versionen sind auf der [Unterseite "Releases"](https://grapheneos.org/releases) verfügbar, Installationsanleitungen auf der [Unterseite "Install"](https://grapheneos.org/install/) [[deutsche Übersetzung](grapheneos-install.html)].

GrapheneOS entwickelt auch verschiedene Anwendungen und Dienste mit dem Schwerpunkt auf Privatsphäre und Sicherheit. Vanadium ist eine gehärtete Variante des Chromium-Browsers und WebView, die speziell für GrapheneOS erstellt wurde. GrapheneOS umfasst auch unseren minimalistischen, auf Sicherheit ausgerichteten PDF-Viewer, unsere hardwarebasierte Auditor-App und den Attestierungsdienst, der sowohl eine lokale als auch eine entfernte Verifizierung von Geräten ermöglicht, unsere moderne, auf Privatsphäre und Sicherheit ausgerichtete Kamera-App und das extern entwickelte verschlüsselte Backup von Seedvault, das ursprünglich für die Aufnahme in GrapheneOS entwickelt wurde.

<h2 id="never-google-services">Keine Google-Anwendungen oder -Dienste</h2>

GrapheneOS wird weder Google-Play-Dienste noch eine andere Implementierung von Google-Diensten wie microG beinhalten. Es ist möglich, die Play-Dienste in Form von vollständig isolierten Apps ohne besondere Privilegien über unsere [isolierte Google-Play-Kompatibilitätsschicht](https://grapheneos.org/usage#sandboxed-google-play) zu installieren. Weitere Einzelheiten zu unseren Plänen, die Lücken zu schließen, die durch das Fehlen der Play-Dienste und der Google-Apps entstehen, finden Sie im entsprechenden [Abschnitt unserer FAQ](https://grapheneos.org/faq#google-services).

<h2 id="device-support">Unterstützte Geräte</h2>

Siehe [Abschnitt Geräteunterstützung in den FAQ](https://grapheneos.org/faq#device-support).
