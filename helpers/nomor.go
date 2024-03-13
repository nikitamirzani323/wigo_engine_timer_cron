package helpers

import (
	"math/rand"
)

type nomor_result_data struct {
	nomor_id         string
	nomor_gangen     string
	nomor_besarkecil string
	nomor_line       string
	nomor_redblack   string
}

func Shuffle_nomor() string {
	var cards = []nomor_result_data{
		{nomor_id: "00", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "01", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "02", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "03", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "04", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "05", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "06", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "07", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "08", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "09", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "10", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "11", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "12", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "13", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "14", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "15", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "16", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "17", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "18", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "19", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "20", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "21", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "22", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "23", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "24", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "25", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "26", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "27", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "28", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "29", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "30", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "31", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "32", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "33", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "34", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "35", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "36", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "37", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "38", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "39", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "40", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "41", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "42", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "43", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "44", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "45", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "46", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "47", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "48", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "49", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "50", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "51", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "52", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "53", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "54", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "55", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "56", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "57", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "58", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "59", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "60", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "61", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "62", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "63", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "64", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "65", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "66", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "67", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "68", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "69", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "70", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "71", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "72", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "73", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "74", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "75", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "76", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "77", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "78", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "79", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "80", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE5", nomor_redblack: "BLACK"},
		{nomor_id: "81", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "82", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "83", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "84", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "85", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "86", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "87", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "88", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "89", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "90", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "91", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "92", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "93", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "94", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "95", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "96", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "97", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"},
		{nomor_id: "98", nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "BLACK"},
		{nomor_id: "99", nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE4", nomor_redblack: "RED"}}
	min := 0
	max := len(cards)
	var n = rand.Intn(max-min) + min
	return cards[n].nomor_id
}
