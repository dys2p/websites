// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de_DE": &dictionary{index: de_DEIndex, data: de_DEData},
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"All Services and Projects":       10,
	"Battery Disposal":                35,
	"Cancellation Policy":             4,
	"Cash":                            6,
	"Cash by mail":                    37,
	"Cash by mail in 20 currencies":   7,
	"Cash payment in local store":     36,
	"Cash payment in our store":       31,
	"Closed: June 8 – June 16, 2024":  20,
	"Concept":                         33,
	"Contact & News":                  17,
	"Contact us":                      18,
	"DHL parcel, franked digitally":   29,
	"DHL parcel, franked handwritten": 28,
	"Delivery":                        25,
	"Digital Goods":                   13,
	"Germany":                         19,
	"Got an idea or found an error? Drop us a note!": 24,
	"Imprint":                               3,
	"Legal":                                 0,
	"Local Store":                           12,
	"Mon+Thu 2pm-6pm":                       22,
	"Monero and Bitcoin":                    8,
	"New opening hours!":                    21,
	"Online printing":                       16,
	"Online shop":                           14,
	"Order Service":                         15,
	"PayPal":                                32,
	"Payment":                               5,
	"Pickup from locker in our local store": 27,
	"Pickup in our local store":             26,
	"Privacy policy":                        2,
	"SEPA bank transfer":                    9,
	"Support us":                            34,
	"Terms and Conditions":                  1,
	"Tue+Wed+Fri+Sat 10am-2pm":              23,
	"UPS parcel directly from our contractor": 30,
	"Why?": 11,
}

var de_DEIndex = []uint32{ // 39 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000c, 0x00000010, 0x0000001c,
	0x00000026, 0x00000039, 0x00000043, 0x0000004b,
	0x0000006d, 0x00000080, 0x00000092, 0x000000ad,
	0x000000b4, 0x000000c3, 0x000000d3, 0x000000de,
	0x000000ed, 0x000000fd, 0x0000010c, 0x00000114,
	0x00000120, 0x0000013e, 0x00000154, 0x00000164,
	0x0000017a, 0x0000019d, 0x000001a5, 0x000001c0,
	0x000001ec, 0x00000210, 0x0000022c, 0x0000024f,
	// Entry 20 - 3F
	0x0000026c, 0x00000273, 0x0000027b, 0x0000028c,
	0x000002ac, 0x000002c9, 0x000002da,
} // Size: 180 bytes

const de_DEData string = "" + // Size: 730 bytes
	"\x02Rechtliches\x02AGB\x02Datenschutz\x02Impressum\x02Widerrufsbelehrung" +
	"\x02Bezahlung\x02Bargeld\x02Bargeld per Post in 20 Währungen\x02Monero u" +
	"nd Bitcoin\x02SEPA-Überweisung\x02Alle Angebote und Projekte\x02Warum?" +
	"\x02Ladengeschäft\x02Digitale Güter\x02Onlineshop\x02Bestellservice\x02O" +
	"nlinedruckerei\x02Kontakt & News\x02Kontakt\x02Deutschland\x02Geschlosse" +
	"n: 8.-16. Juni 2024\x02Neue Öffnungszeiten!\x02Mo+Do 14-18 Uhr\x02Di+Mi+" +
	"Fr+Sa 10-14 Uhr\x02Fehler oder Hinweise? Schreib uns!\x02Versand\x02Abho" +
	"lung im Ladengeschäft\x02Abholung aus Schließfach im Ladengeschäft\x02DH" +
	"L-Paket handschriftlich frankiert\x02DHL-Paket digital frankiert\x02UPS-" +
	"Paket direkt von der Druckerei\x02Barzahlung im Ladengeschäft\x02PayPal" +
	"\x02Konzept\x02Unterstütze uns\x02Hinweise zur Batterieentsorgung\x02Bar" +
	"zahlung im Ladengeschäft\x02Bargeld per Post"

var en_USIndex = []uint32{ // 39 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x0000001b, 0x0000002a,
	0x00000032, 0x00000046, 0x0000004e, 0x00000053,
	0x00000071, 0x00000084, 0x00000097, 0x000000b1,
	0x000000b6, 0x000000c2, 0x000000d0, 0x000000dc,
	0x000000ea, 0x000000fa, 0x00000109, 0x00000114,
	0x0000011c, 0x0000013d, 0x00000150, 0x00000160,
	0x00000179, 0x000001a8, 0x000001b1, 0x000001cb,
	0x000001f1, 0x00000211, 0x0000022f, 0x00000257,
	// Entry 20 - 3F
	0x00000271, 0x00000278, 0x00000280, 0x0000028b,
	0x0000029c, 0x000002b8, 0x000002c5,
} // Size: 180 bytes

const en_USData string = "" + // Size: 709 bytes
	"\x02Legal\x02Terms and Conditions\x02Privacy policy\x02Imprint\x02Cancel" +
	"lation Policy\x02Payment\x02Cash\x02Cash by mail in 20 currencies\x02Mon" +
	"ero and Bitcoin\x02SEPA bank transfer\x02All Services and Projects\x02Wh" +
	"y?\x02Local Store\x02Digital Goods\x02Online shop\x02Order Service\x02On" +
	"line printing\x02Contact & News\x02Contact us\x02Germany\x02Closed: June" +
	" 8 – June 16, 2024\x02New opening hours!\x02Mon+Thu 2pm-6pm\x02Tue+Wed+F" +
	"ri+Sat 10am-2pm\x02Got an idea or found an error? Drop us a note!\x02Del" +
	"ivery\x02Pickup in our local store\x02Pickup from locker in our local st" +
	"ore\x02DHL parcel, franked handwritten\x02DHL parcel, franked digitally" +
	"\x02UPS parcel directly from our contractor\x02Cash payment in our store" +
	"\x02PayPal\x02Concept\x02Support us\x02Battery Disposal\x02Cash payment " +
	"in local store\x02Cash by mail"

	// Total table size 1799 bytes (1KiB); checksum: 167BA110
