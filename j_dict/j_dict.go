package j_dict

import (
	"encoding/xml"
	"fmt"
	"j_study_blog/dictionary"
	"j_study_blog/repository"
	"log"
	"os"
	"strings"
)

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Text    string   `xml:",chardata"`
	EntSeq  string   `xml:"ent_seq"`
	KEle    struct {
		Text  string   `xml:",chardata"`
		Keb   string   `xml:"keb"`
		KePri []string `xml:"ke_pri"`
	} `xml:"k_ele"`
	REle []struct {
		Text  string   `xml:",chardata"`
		Reb   string   `xml:"reb"`
		RePri []string `xml:"re_pri"`
	} `xml:"r_ele"`
	Sense []struct {
		Text  string   `xml:",chardata"`
		Pos   []string `xml:"pos"`
		Gloss []struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"gloss"`
	} `xml:"sense"`
}

func Import(repo repository.IVocabRepo) {
	xmlFile, err := os.Open("j_dict/JMdict_e")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.Strict = false
	var curElement string
	var vocabs []dictionary.Vocab
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			curElement = se.Name.Local
			if curElement == "entry" {
				var entry Entry
				err := decoder.DecodeElement(&entry, &se)
				if err != nil {
					log.Fatal(err)
				}

				vocab := dictionary.Vocab{
					KanjiReading: strings.Trim(entry.KEle.Keb, "><"),
					Meanings:     []dictionary.VocabMeaning{},
					KanaReadings: []string{},
				}

				for _, s := range entry.Sense {
					for _, m := range s.Gloss {
						vocab.Meanings = append(vocab.Meanings, dictionary.VocabMeaning{Text: strings.Trim(m.Text, "><"), Language: strings.Trim(m.Lang, "><")})
					}
				}

				for _, r := range entry.REle {
					vocab.KanaReadings = append(vocab.KanaReadings, strings.Trim(r.Reb, "><"))
				}

				vocabs = append(vocabs, vocab)
				if len(vocabs) == 100 {
					repo.InsertMany(vocabs)
					vocabs = []dictionary.Vocab{}
				}
			}
		default:
		}
	}
}
