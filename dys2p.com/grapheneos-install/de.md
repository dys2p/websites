<nav aria-label="breadcrumb">
	<ol class="breadcrumb">
		<li class="breadcrumb-item"><a href="grapheneos-preface.html">Übersetzungen zu GrapheneOS</a></li>
		<li class="breadcrumb-item active" aria-current="page">Installation</li>
	</ol>
</nav>

<div class="alert alert-primary">
	Diese Übersetzung basiert auf dem Commit <a href="https://github.com/GrapheneOS/grapheneos.org/blob/65117f9794eba2bdf20b80a2fe5884dff8938f1e/static/install/index.html">65117f9</a> vom 2024-09-23. Falls du Hinweise oder Verbesserungsvorschläge hast, dann <a href="contact.html">schreib uns gerne</a> oder arbeite mit uns auf <a href="https://github.com/dys2p/websites/blob/main/dys2p.com/grapheneos-install/de.md">GitHub</a> an dieser Übersetzung.
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

<h1 id="install">Installation</h1>

GrapheneOS hat zwei offiziell unterstützte Installationsmethoden. Sie können entweder die [WebUSB-basierte Installation](https://grapheneos.org/install/web) [[deutsche Übersetzung](grapheneos-install-web.html)] verwenden, die für die meisten Benutzer empfohlen wird, oder die [Installation in der Kommandozeile](https://grapheneos.org/install/cli) [[deutsche Übersetzung](grapheneos-install-cli.html)], die für technisch versiertere Benutzer gedacht ist.

Wir empfehlen dringend, eine der offiziellen Installationsmethoden zu verwenden. Die Installationsanleitungen von Dritten sind in der Regel veraltet und enthalten oft irreführende Ratschläge und Fehler.

Wenn Sie Probleme mit dem Installationsprozess haben, bitten Sie im [offiziellen GrapheneOS-Chat](https://grapheneos.org/contact#community) um Hilfe. Dort gibt es fast immer Leute, die bereit sind zu helfen. Bevor Sie um Hilfe bitten, versuchen Sie, der Anleitung selbst zu folgen und bitten Sie erst dann um Hilfe, wenn Sie nicht weiterkommen.

Der Weg über die Kommandozeile erfordert ein Betriebssystem mit den richtigen fastboot- und OpenSSH-Paketen sowie ein ausreichendes Verständnis der Abläufe, um nicht blindlings den Anweisungen unserer Website zu vertrauen. Bei der webbasierten Installation benötigen Sie außer einem Browser mit WebUSB-Unterstützung keine weitere Software. Sie können trotzdem vermeiden, unserer Serverinfrastruktur vertrauen zu müssen, indem Sie den Hash des Verified-Boot-Schlüssels überprüfen.
