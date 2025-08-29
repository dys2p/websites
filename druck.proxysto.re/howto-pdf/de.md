# How to PDF? Druckfertige PDF-Dateien für alle!

<div class="alert alert-success my-3 text-center"><a rel="noreferrer" target="_blank" href="https://www.flyeralarm.com/de/content/index/open/id/911/druckdaten.html">Eine ähnliche Anleitung findet ihr auch bei Flyeralarm</a>.</div>

Wir lieben DIY -- aber wenn unsere Plakate, Flyer oder Aufkleber besonders schick werden sollen, reicht der Kopierer im Copyshop nicht aus. Doch (Online-)Druckereien haben meist gewisse Anforderungen an eure Dateien. Warum sie diese aus gutem Grund haben und wie ihr druckbare Dateien erstellt, lernt ihr jetzt.

Um druckbare PDF-Dateien zu erstellen, verwenden wir die freie Software [Scribus](https://www.scribus.net). Ihr werdet euch aber auch mit anderen Programmen zurechtfinden. Damit puzzeln wir Bilder, Vektorgrafiken und Textfelder zusammen. Eine Vektorgrafik ist alles, was keine Pixel hat.

## Anschnitt und Rand

Weiße Ränder waren gestern -- im 21. Jahrhundert erwarten wir randlos bedrucktes Papier. Doch die Maschinen in der Druckerei schneiden nicht ganz exakt. Würden wir dein Motiv auf ein weißes Blatt Papier drucken und dann ausschneiden, dann würden wir an der einen Seite einen Millimeter deines Motivs abschneiden, während an der anderen Seite ein Millimeter weißes Papier übrig bliebe.

Die Lösung: Du gestaltest die Datei so, dass der Hintergrund etwas größer ist. Diese zusätzliche Fläche wird **Anschnitt** genannt. Sie wird mitgedruckt und dann (mehr oder weniger genau) abgeschnitten. Weiße Ränder gehören damit der Vergangenheit an. Deine druckbare Datei ist also einige Millimeter breiter und höher als das Endformat. Die genauen Maße erfährst du bei der Druckerei. Sie stehen meist im **Datenblatt** des Artikels.

Wichtige Informationen oder Gestaltungselemente sollten außerdem einen **Sicherheitsabstand** zum Rand einhalten. Ansonsten werden sie beim (ungenauen) Schneiden möglicherweise abgeschnitten.

## Auflösung

Die Realität ist schärfer als dein Bildschirm. Damit menschliche Augen auf dem Papier keine Pixel (Bildpunkte) mehr erkennen, erwarten die meisten Druckereien bei Bildern eine Mindestauflösung von 300 dpi. Die Einheit **dpi** bedeutet "dots per inch", also Pixel pro Zoll. Wichtig ist also nicht nur die Größe des Bildes (wie etwa 2000 x 3000 Pixel), sondern auch, in welcher Größe es gedruckt wird (beispielsweise 10 x 15 cm).

Natürlich kannst du auch zu niedrig aufgelöste Bilder in dein Motiv einbetten, die Pixel auf dem Druckprodukt gehen dann aber auf deine Kappe. Versuche daher, Fotos immer in bestmöglicher Auflösung zu erhalten, und bedenke, dass manche Messenger-Apps Bilder beim Versenden herunterskalieren.

Wenn du ein niedrig aufgelöstes Bild nur für den Hintergrund brauchst, kannst du es hochskalieren und danach mit einem Weichzeichnungs-Filter bearbeiten. Unscharfe Bilder sehen immer noch besser aus als pixelige Bilder.

## Farbmanagement

Dein Bildschirm mischt Farben **additiv** aus rotem, grünem und blauem Licht (**RGB**) zusammen. Die Farben auf Papier entstehen hingegen durch **subtraktive Farbmischung**: Das Papier ist mit Punkten in cyanblau, magentarot, gelb und schwarz (cyan, magenta, yellow and key: **CMYK**) bedruckt. Wenn weißes Licht auf das Papier fällt, verschluckt die Farbe einen Teil davon und reflektieren den Rest.

Es gibt Farben, die dein Bildschirm anzeigen kann, die deine Druckerei aber nicht zu Papier bringen kann -- und umgekehrt. Die Menge der Farben, den ein Medium darstellen kann, wird als **Farbraum** bezeichnet. Die Unterschiede liegen nicht nur in der Farbmischung. Eine RGB-Farbe wird auf zwei Bildschirmmodellen wahrscheinlich ein bisschen unterschiedlich aussehen, eine CMYK-Farbe auf zwei verschiedenen Papiersorten auch.

Um dieses Chaos in den Griff zu bekommen, existieren **Farbprofile**. Das sind Dateien mit der Dateiendung **.icc**, die den Farbraum eines Mediums beschreiben. Gute Druckereien stellen dir die Farbprofile ihrer Druckmaschinen als Download oder Link zur Verfügung. ([Die von Flyeralarm genutzten Farbprofile findet ihr hier.](https://www.flyeralarm.com/de/content/index/open/id/27240/farbe.html)) Zum Umrechnen der Farben ist noch ein zweites Farbprofil nötig: das deines Bildschirms. Wenn\'s schnell gehen muss, nimm das beiliegende "SRGB"-Profil. Profis haben spezielle Bildschirme oder Geräte zur Farbkalibrierung.

Viele Grafikprogramme bieten, nachdem du beide Farbprofile -- für Druck und Bildschirm -- eingestellt hast, die Möglichkeiten, die Druckerfarben auf dem Bildschirm zu simulieren und Farben außerhalb des druckbaren Farbraums zu markieren. Außerdem kannst du einstellen, was mit nicht-druckbaren Farben passieren soll: "Absolut farbmetrisch" ändert nur nicht-druckbare Farben in die nächstgelegene druckbare Farbe. "Relativ farbmetrisch" passt auch die druckbaren Farben an, damit die Farben auf dem Papier im richtigen Verhältnis zueinander stehen. "Wahrnehmung" verwendet einen Algorithmus, damit es für dein Auge irgendwie stimmig aussieht. Oft werden Bilder nach "Wahrnehmung" umgerechnet und Füllfarben in Vektorgrafiken "absolut farbmetrisch" -- in dieser Kombination ist es allerdings möglich, dass zwei Farben auf dem Bildschirm gleich aussehen, sich auf dem Papier aber unterscheiden. Choose your weapon, wir drücken die Daumen!

## PDF/X-3

Willkommen in der Praxis! Damit der Kram mit den Farbprofilen funktioniert, erwarten die meisten Druckereien PDF-Dateien nach dem **PDF/X-3**-Standard. Dieser Standard verbietet eine ganze Menge Dinge (wie etwa Transparenzeffekte oder fehlende Schriftarten) und trägt damit dazu bei, dass das Motiv bei der Druckerei genauso aussieht wie bei dir. Viele Grafikprogramme ermöglichen einen Export im PDF/X-3-Format.

**Manchmal wird auch PDF/X-4 bevorzugt, etwa beim Bedrucken von Textilien.** PDF/X-4 unterstützt Ebenen und Transparenzen. Ob PDF/X-3 oder PDF/X-4 erwartet wird, steht im Datenblatt des ausgewählten Produkts.

Falls du dir Probleme mit Transparenz und Schriftarten ersparen möchtest, kannst du dein Motiv zunächst als RGB-Pixelgrafik exportieren, etwa im PNG-Format. Importiere danach die PNG-Datei in ein neues Dokument, stelle das Farbmanagement ein und exportiere es als PDF/X-3.

## Die wichtigsten Einstellungen in Scribus

[Scribus](https://www.scribus.net) ist eine freie *Desktop-Publishing*-Software. Sie eignet sich gut zum Erzeugen von druckbaren PDF-Dateien.

![Scribus: Anschnitt einstellen](/static/images/scribus-anschnitt-1-mark.png)

Beim Erstellen oder Bearbeiten einer Datei in Scribus kannst du den Anschnitt angeben. Scribus rechnet ihn dann hinzu. Er erscheint im Dokument außerhalb der Seitenränder.

Du musst die Anschnitt-Einstellung nicht verwenden, sondern kannst auch selbst rechnen und Endformat plus Anschnitt als benutzerdefinierte Dokumentengröße einstellen.

![Scribus: Anschnitt](/static/images/scribus-anschnitt-2-mark.png)

Achte darauf, dass der Hintergrund deiner Datei sich auch auf den Anschnitt erstreckt.

![Scribus: Export, als PDF speichern](/static/images/scribus-export-1.png)

Klicke auf <kbd>Datei</kbd> › <kbd>Exportieren</kbd> › <kbd>Als PDF speichern</kbd>.

![Scribus: Export, PDF/X-3 und Kompressionsqualität](/static/images/scribus-export-2-mark.png)

Wähle <kbd>PDF/X-3</kbd> und <kbd>Kompressionsqualität: Maximal</kbd> aus.

![Scribus: Export, Pre-Press](/static/images/scribus-export-3-mark.png)

Im Reiter *Pre-Press* muss <kbd>Dokumenteneinstellungen für den Anschnitt verwenden</kbd> aktiviert sein. Nur dann landet der Anschnitt auch mit in der PDF-Datei.

Wähle aus Ausgabeprofil <kbd>ISO Coated v2 300%</kbd>. Damit du die PDF-Datei speichern kannst, musst du außerdem irgendeinen <kbd>Infotext</kbd> eingeben.
