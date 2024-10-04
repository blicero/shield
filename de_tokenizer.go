// /home/krylon/go/src/github.com/blicero/shield/de_tokenizer.go
// -*- mode: go; coding: utf-8; -*-
// Created on 11. 10. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-10-11 21:43:11 krylon>
//
// Copyright (c) 2005, Jacques Savoy.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package shield

import (
	"regexp"
	"strings"
)

type deTokenizer struct {
}

func NewGermanTokenizer() Tokenizer {
	return &deTokenizer{}
}

func (t *deTokenizer) Tokenize(text string) (words map[string]int64) {
	words = make(map[string]int64)
	for _, w := range splitPat.Split(text, -1) {
		if len(w) > 2 {
			words[strings.ToLower(w)]++
		}
	}
	return
}

var splitPat = regexp.MustCompile(`(\W+|ab|aber|ach|acht|achte|achten|achter|achtes|ag|alle|allein|allem|allen|aller|allerdings|alles|allgemeinen|als|also|am|an|andere|anderen|andern|anders|au|auch|auf|aus|ausser|ausserdem|außer|außerdem|bald|bei|beide|beiden|beim|beispiel|bekannt|bereits|besonders|besser|besten|bin|bis|bisher|bist|d.h|da|dabei|dadurch|dafür|dagegen|daher|dahin|dahinter|damals|damit|danach|daneben|dank|dann|daran|darauf|daraus|darf|darfst|darin|darum|darunter|darüber|das|dasein|daselbst|dass|dasselbe|davon|davor|dazu|dazwischen|daß|dein|deine|deinem|deiner|dem|dementsprechend|demgegenüber|demgemäss|demgemäß|demselben|demzufolge|den|denen|denn|denselben|der|deren|derjenige|derjenigen|dermassen|dermaßen|derselbe|derselben|des|deshalb|desselben|dessen|deswegen|dich|die|diejenige|diejenigen|dies|diese|dieselbe|dieselben|diesem|diesen|dieser|dieses|dir|doch|dort|drei|drin|dritte|dritten|dritter|drittes|du|durch|durchaus|durfte|durften|dürfen|dürft|eben|ebenso|ehrlich|ei|eigen|eigene|eigenen|eigener|eigenes|ein|einander|eine|einem|einen|einer|eines|einige|einigen|einiger|einiges|einmal|eins|elf|en|ende|endlich|entweder|er|ernst|erst|erste|ersten|erster|erstes|es|etwa|etwas|euch|früher|fünf|fünfte|fünften|fünfter|fünftes|für|gab|ganz|ganze|ganzen|ganzer|ganzes|gar|gedurft|gegen|gegenüber|gehabt|gehen|geht|gekannt|gekonnt|gemacht|gemocht|gemusst|genug|gerade|gern|gesagt|geschweige|gewesen|gewollt|geworden|gibt|ging|gleich|gott|gross|grosse|grossen|grosser|grosses|groß|große|großen|großer|großes|gut|gute|guter|gutes|habe|haben|habt|hast|hat|hatte|hatten|heisst|her|heute|hier|hin|hinter|hoch|hätte|hätten|ich|ihm|ihn|ihnen|ihr|ihre|ihrem|ihren|ihrer|ihres|im|immer|in|indem|infolgedessen|ins|irgend|ist|ja|jahr|jahre|jahren|je|jede|jedem|jeden|jeder|jedermann|jedermanns|jedoch|jemand|jemandem|jemanden|jene|jenem|jenen|jener|jenes|jetzt|kam|kann|kannst|kaum|kein|keine|keinem|keinen|keiner|kleine|kleinen|kleiner|kleines|kommen|kommt|konnte|konnten|kurz|können|könnt|könnte|lang|lange|leicht|leide|lieber|los|machen|macht|machte|mag|magst|mahn|man|manche|manchem|manchen|mancher|manches|mann|mehr|mein|meine|meinem|meinen|meiner|meines|mensch|menschen|mich|mir|mit|mittel|mochte|mochten|morgen|muss|musst|musste|mussten|muß|möchte|mögen|möglich|mögt|müssen|müsst|na|nach|nachdem|nahm|natürlich|neben|nein|neue|neuen|neun|neunte|neunten|neunter|neuntes|nicht|nichts|nie|niemand|niemandem|niemanden|noch|nun|nur|ob|oben|oder|offen|oft|ohne|ordnung|recht|rechte|rechten|rechter|rechtes|richtig|rund|sa|sache|sagt|sagte|sah|satt|schlecht|schluss|schon|sechs|sechste|sechsten|sechster|sechstes|sehr|sei|seid|seien|sein|seine|seinem|seinen|seiner|seines|seit|seitdem|selbst|sich|sie|sieben|siebente|siebenten|siebenter|siebentes|sind|so|solang|solche|solchem|solchen|solcher|solches|soll|sollen|sollte|sollten|sondern|sonst|sowie|später|statt|tag|tage|tagen|tat|teil|tel|tritt|trotzdem|tun|uhr|um|und|und?|uns|unser|unsere|unserer|unter|vergangenen|viel|viele|vielem|vielen|vielleicht|vier|vierte|vierten|vierter|viertes|vom|von|vor|wahr?|wann|war|waren|wart|warum|was|wegen|weil|weit|weiter|weitere|weiteren|weiteres|welche|welchem|welchen|welcher|welches|wem|wen|wenig|wenige|weniger|weniges|wenigstens|wenn|wer|werde|werden|werdet|wessen|wie|wieder|will|willst|wir|wird|wirklich|wirst|wo|wohl|wollen|wollt|wollte|wollten|worden|wurde|wurden|während|währenddem|währenddessen|wäre|würde|würden|z.b|zehn|zehnte|zehnten|zehnter|zehntes|zeit|zu|zuerst|zugleich|zum|zunächst|zur|zurück|zusammen|zwanzig|zwar|zwei|zweite|zweiten|zweiter|zweites|zwischen|zwölf|über|überhaupt|übrigens)`)
