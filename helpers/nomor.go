package helpers

import (
	"math/rand"
)

type nomor_result_data struct {
	nomor_id         string
	nomor_flag       bool
	nomor_css        string
	nomor_gangen     string
	nomor_besarkecil string
	nomor_line       string
	nomor_redblack   string
}

func Shuffle_nomor() string {
	var cards = []nomor_result_data{
		{nomor_id: "00", nomor_flag: false, nomor_css: "btn btn-error", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "01", nomor_flag: false, nomor_css: "btn", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "02", nomor_flag: false, nomor_css: "btn btn-error", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "03", nomor_flag: false, nomor_css: "btn", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "04", nomor_flag: false, nomor_css: "btn btn-error", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "05", nomor_flag: false, nomor_css: "btn", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "06", nomor_flag: false, nomor_css: "btn btn-error", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "07", nomor_flag: false, nomor_css: "btn", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "08", nomor_flag: false, nomor_css: "btn btn-error", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "09", nomor_flag: false, nomor_css: "btn", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "10", nomor_flag: false, nomor_css: "btn btn-error", nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "11", nomor_flag: false, nomor_css: "btn", nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"}}
	min := 0
	max := len(cards)
	var n = rand.Intn(max-min) + min
	return cards[n].nomor_id
}
