package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/wigo_engine_timer_cron/configs"
	"bitbucket.org/isbtotogroup/wigo_engine_timer_cron/db"
	"bitbucket.org/isbtotogroup/wigo_engine_timer_cron/helpers"
	"bitbucket.org/isbtotogroup/wigo_engine_timer_cron/models"
	"github.com/buger/jsonparser"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/nleeper/goment"
)

var invoice = ""
var invoice_agen = ""
var time_status = "LOCK"
var game_status = "OFFLINE"
var game_result = ""
var operator_status = "N"
var data_send = ""
var data_send_agen = ""
var data_send_savetransaksi = ""

func main() {
	local, err_local := time.LoadLocation("Asia/Jakarta")
	if err_local != nil {
		local = time.UTC
	}
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	initRedis := helpers.RedisHealth()

	if !initRedis {
		panic("cannot load redis")
	}
	db.Init()
	envCompany := os.Getenv("DB_CONF_COMPANY")
	envCurr := os.Getenv("DB_CONF_CURR")
	time_game := 0
	fieldconfig_redis := strings.ToLower(envCompany) + ":12D30S:CONFIG_TIMER"

	type Configure struct {
		Time        int    `json:"time"`
		Maintenance string `json:"maintenance"`
		Operator    string `json:"operator"`
	}
	var obj Configure

	resultredis, flag_config := helpers.GetRedis(fieldconfig_redis)
	jsonredis := []byte(resultredis)
	timeRD, _ := jsonparser.GetInt(jsonredis, "time")
	maintenanceRD, _ := jsonparser.GetString(jsonredis, "maintenance")
	operatorRD, _ := jsonparser.GetString(jsonredis, "operator")

	if !flag_config {
		fmt.Println("ENGINE TIMER CONFIG DATABASE")
		time_game_DB, game_status_DB, operator_DB := _GetConf(envCompany)
		obj.Time = time_game_DB
		obj.Maintenance = game_status_DB
		obj.Operator = operator_DB
		helpers.SetRedis(fieldconfig_redis, obj, 60*time.Minute)
		time_game = time_game_DB
		game_status = game_status_DB
		operator_status = operator_DB

	} else {
		fmt.Println("ENGINE TIMER CONFIG CACHE")
		time_game = int(timeRD)
		game_status = maintenanceRD
		operator_status = operatorRD
	}
	if operator_status == "N" {
		invoice = Save_transaksi(strings.ToLower(envCompany), envCurr)
		invoice_agen = invoice
	} else {
		time_game = 0
	}
	s := gocron.NewScheduler(local)
	s.Every(1).Seconds().Do(func() {
		resultredis, flag_config := helpers.GetRedis(fieldconfig_redis)
		jsonredis := []byte(resultredis)
		maintenanceRD, _ := jsonparser.GetString(jsonredis, "maintenance")
		operatorRD, _ := jsonparser.GetString(jsonredis, "operator")

		if !flag_config {
			fmt.Println("CONFIG DATABASE")
			time_game_DB, game_status_DB, operator_DB := _GetConf(envCompany)
			obj.Time = time_game_DB
			obj.Maintenance = game_status_DB
			obj.Operator = operator_DB
			helpers.SetRedis(fieldconfig_redis, obj, 60*time.Minute)
			game_status = game_status_DB
			operator_status = operatorRD

		} else {
			fmt.Println("CONFIG CACHE")
			game_status = maintenanceRD
			operator_status = operatorRD
		}

		if game_status == "ONLINE" {
			if time_game < 0 { //IDLE
				flag_compiledata := false
				time_game = 0
				time_status = "LOCK"
				if invoice != "" {
					invoice_agen = invoice
				}
				invoice = ""
				data_send = invoice + "|0|" + time_status + "|" + game_status + "|" + game_result
				data_send_agen = invoice_agen + "|OPEN"
				fmt.Printf("%s:%.2d:%s:%s:%s\n", invoice, time_game%60, time_status, game_status, game_result)
				senddata(data_send, envCompany)
				senddata_agen(data_send_agen, envCompany)

				if operator_status == "N" { // tanpa operator
					fmt.Println("ONLINE AUTOMATION")
					if game_result == "" {
						game_result = helpers.Shuffle_nomor()
						data_send = invoice + "|0|" + time_status + "|" + game_status + "|" + game_result
						data_send_agen = invoice_agen + "|OPEN"
						fmt.Printf("%s:%.2d:%s:%s:%s\n", invoice, time_game%60, time_status, game_status, game_result)
						senddata(data_send, envCompany)

						flag_compiledata = Update_transaksi(strings.ToLower(envCompany), game_result)
						time.Sleep(3 * time.Second)
						// time.Sleep(3 * time.Second)
						if flag_compiledata {
							invoice = Save_transaksi(strings.ToLower(envCompany), envCurr)
							invoice_agen = invoice

							resultredis, flag_config := helpers.GetRedis(fieldconfig_redis)
							jsonredis := []byte(resultredis)
							timeRD, _ := jsonparser.GetInt(jsonredis, "time")
							maintenanceRD, _ := jsonparser.GetString(jsonredis, "maintenance")
							operatorRD, _ := jsonparser.GetString(jsonredis, "operator")

							if !flag_config {
								fmt.Println("CONFIG DATABASE")
								time_game_DB, game_status_DB, operator_DB := _GetConf(envCompany)
								obj.Time = time_game_DB
								obj.Maintenance = game_status_DB
								obj.Operator = operator_DB
								helpers.SetRedis(fieldconfig_redis, obj, 60*time.Minute)
								time_game = time_game_DB
								game_status = game_status_DB

							} else {
								fmt.Println("CONFIG CACHE")
								time_game = int(timeRD)
								game_status = maintenanceRD
								operator_status = operatorRD
							}
							game_result = ""
						}
					}
				} else { // dengan operator
					fmt.Println("ONLINE OPERATOR")
					fieldconfigoperator_redis := strings.ToLower(envCompany) + ":12D30S:TIMER"
					resultredis_timer, flag_timer := helpers.GetRedis(fieldconfigoperator_redis)
					jsonredis_timer := []byte(resultredis_timer)
					resultRD, _ := jsonparser.GetString(jsonredis_timer, "result")
					timeRD, _ := jsonparser.GetInt(jsonredis_timer, "time")
					invoiceRD, _ := jsonparser.GetString(jsonredis_timer, "invoice")

					if flag_timer {
						fmt.Println("TIMER_ CACHE")

						data_send = "|0|" + time_status + "|" + game_status + "|" + resultRD
						data_send_agen = invoice_agen + "|OPEN"
						fmt.Printf("%s:%.2d:%s:%s:%s\n", invoice, time_game%60, time_status, game_status, game_result)
						senddata(data_send, envCompany)
						time.Sleep(3 * time.Second)

						time_game = int(timeRD)
						invoice = invoiceRD
						invoice_agen = ""
						val_transaksi2d30s2 := helpers.DeleteRedis(fieldconfigoperator_redis)
						fmt.Printf("Redis Delete TIMER : %d", val_transaksi2d30s2)

						data_send_agen = invoice_agen + "|CLOSED"
						senddata_agen(data_send_agen, envCompany)
					} else {
						invoice = Save_transaksi(strings.ToLower(envCompany), envCurr)
						invoice_agen = invoice
						data_send_agen = invoice_agen + "|OPEN"
						senddata_agen(data_send_agen, envCompany)
					}
				}

			} else {
				if invoice == "" {
					time_status = "LOCK"
				} else {
					time_status = "OPEN"
					game_result = ""
				}
				if time_game < 6 {
					time_status = "LOCK"
				} else {
					time_status = "OPEN"
					game_result = ""
				}
				//invoice|time|status
				data_send = invoice + "|" + strconv.Itoa(time_game%60) + "|" + time_status + "|" + game_status + "|" + game_result
				fmt.Printf("%s:%.2d:%s:%s:%s\n", invoice, time_game%60, time_status, game_status, game_result)
				senddata(data_send, envCompany)
			}
			time_game--
		} else {
			resultredis, flag_config := helpers.GetRedis(fieldconfig_redis)
			jsonredis := []byte(resultredis)
			timeRD, _ := jsonparser.GetInt(jsonredis, "time")
			maintenanceRD, _ := jsonparser.GetString(jsonredis, "maintenance")
			operatorRD, _ := jsonparser.GetString(jsonredis, "operator")

			if !flag_config {
				fmt.Println("CONFIG DATABASE")
				time_game_DB, game_status_DB, operator_DB := _GetConf(envCompany)
				obj.Time = time_game_DB
				obj.Maintenance = game_status_DB
				obj.Operator = operator_DB
				helpers.SetRedis(fieldconfig_redis, obj, 60*time.Minute)
				time_game = time_game_DB
				game_status = game_status_DB
				operator_status = operatorRD

			} else {
				fmt.Println("CONFIG CACHE")
				time_game = int(timeRD)
				game_status = maintenanceRD
			}

			data_send = "|0|LOCK|" + game_status
			fmt.Printf("%s:%.2d:%s:%s\n", "", 0, "LOCK", game_status)
			senddata(data_send, envCompany)
		}

	})
	s.Every(2).Minute().Do(func() {
		log.Println("RUNNING 2 MINUTE CHECK DB")
		loop_statusrunning(envCompany)
	})
	s.StartBlocking()
}
func loop_statusrunning(idcompany string) {
	con := db.CreateCon()
	ctx := context.Background()
	invoice := ""

	_, _, tbl_trx_transaksidetail, _ := models.Get_mappingdatabase(idcompany)

	sql_select_detail := `SELECT 
					idtransaksidetail,idtransaksi, nomor,tipebet, bet, multiplier, username_client 
					FROM ` + tbl_trx_transaksidetail + `  
					WHERE status_transaksidetail='RUNNING'  
					`

	row, err := con.QueryContext(ctx, sql_select_detail)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			bet_db                                                                         int
			multiplier_db                                                                  float64
			idtransaksidetail_db, idtransaksi_db, nomor_db, tipebet_db, username_client_db string
		)

		err = row.Scan(&idtransaksidetail_db, &idtransaksi_db, &nomor_db, &tipebet_db, &bet_db, &multiplier_db, &username_client_db)
		helpers.ErrorCheck(err)
		prize_2D := _GetInvoiceInfo(strings.ToLower(idcompany), idtransaksi_db)

		if prize_2D != "" {
			invoice = idtransaksi_db

			data_send_savetransaksi = ""
			fmt.Println("Send Data ke engine save transaksi")
			data_send_savetransaksi = invoice + "|" + prize_2D + "|" + idcompany
			fmt.Printf("%s:%s:%s\n", invoice, prize_2D, idcompany)
			senddata_enginesave(data_send_savetransaksi, idcompany)
		}

	}
	defer row.Close()

}

func Save_transaksi(idcompany, idcurr string) string {
	tglnow, _ := goment.New()
	id_invoice := _GetInvoice(idcompany)
	if id_invoice == "" {
		_, tbl_trx_transaksi, _, _ := models.Get_mappingdatabase(idcompany)
		sql_insert := `
			insert into
			` + tbl_trx_transaksi + ` (
				idtransaksi , idcurr, idcompany, datetransaksi, status_transaksi, 
				create_transaksi, createdate_transaksi 
			) values (
				$1, $2, $3, $4, $5,  
				$6, $7  
			)
		`

		field_column := tbl_trx_transaksi + tglnow.Format("YYYY") + tglnow.Format("MM") + tglnow.Format("DD")
		idrecord_counter := models.Get_counter(field_column)
		idrecrodparent_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		date_transaksi := tglnow.Format("YYYY-MM-DD HH:mm:ss")

		flag_insert, msg_insert := models.Exec_SQL(sql_insert, tbl_trx_transaksi, "INSERT",
			idrecrodparent_value, idcurr, idcompany, date_transaksi, "OPEN",
			"SYSTEM", date_transaksi)

		if flag_insert {
			id_invoice = idrecrodparent_value

		} else {
			fmt.Println(msg_insert)
		}
	}

	return id_invoice
}
func Update_transaksi(idcompany, prize2d string) bool {
	tglnow, _ := goment.New()
	id_invoice := _GetInvoice(idcompany)

	flag_compile := false

	if id_invoice != "" {
		_, tbl_trx_transaksi, _, _ := models.Get_mappingdatabase(idcompany)
		// UPDATE RESULT PARENT
		sql_update := `
				UPDATE 
				` + tbl_trx_transaksi + `  
				SET resultwigo=$1, status_transaksi=$2, 
				update_transaksi=$3, updatedate_transaksi=$4           
				WHERE idtransaksi=$5          
			`

		flag_update, msg_update := models.Exec_SQL(sql_update, tbl_trx_transaksi, "UPDATE",
			prize2d, "CLOSED",
			"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"), id_invoice)

		if flag_update {
			data_send_savetransaksi = ""
			fmt.Println("Send Data ke engine save transaksi")
			data_send_savetransaksi = id_invoice + "|" + prize2d + "|" + idcompany
			fmt.Printf("%s:%s:%s\n", id_invoice, prize2d, idcompany)
			senddata_enginesave(data_send_savetransaksi, idcompany)
			flag_compile = true
		} else {
			fmt.Println(msg_update)
		}
	}
	return flag_compile
}
func senddata(data, company string) {
	key := "payload" + "_" + strings.ToLower(company)
	helpers.SetPublish(key, data)
}
func senddata_agen(data, company string) {
	key := "payload" + "_agen_" + strings.ToLower(company)
	helpers.SetPublish(key, data)
}
func senddata_enginesave(data, company string) {
	key := "payload" + "_enginesave_" + strings.ToLower(company)
	helpers.SetPublish(key, data)
}
func _GetConf(idcompany string) (int, string, string) {
	con := db.CreateCon()
	ctx := context.Background()

	time := 0
	maintenance := "N"
	operator := "N"

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "conf_2digit_30_time, conf_2digit_30_maintenance, conf_2digit_30_operator "
	sql_select += "FROM " + configs.DB_tbl_mst_company_config + " "
	sql_select += "WHERE idcompany='" + idcompany + "' "
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&time, &maintenance, &operator); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	if maintenance == "Y" {
		maintenance = "OFFLINE"
	} else {
		maintenance = "ONLINE"
	}
	return time, maintenance, operator
}
func _GetInvoice(idcompany string) string {
	con := db.CreateCon()
	ctx := context.Background()

	_, tbl_trx_transaksi, _, _ := models.Get_mappingdatabase(idcompany)

	idtransaksi := ""

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idtransaksi "
	sql_select += "FROM " + tbl_trx_transaksi + " "
	sql_select += "WHERE resultwigo='' "
	sql_select += "AND status_transaksi='OPEN' "
	sql_select += "ORDER BY idtransaksi DESC LIMIT 1"

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&idtransaksi); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return idtransaksi
}
func _GetInvoiceInfo(idcompany, idinvoice string) string {
	con := db.CreateCon()
	ctx := context.Background()

	_, tbl_trx_transaksi, _, _ := models.Get_mappingdatabase(idcompany)

	result := ""

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "resultwigo "
	sql_select += "FROM " + tbl_trx_transaksi + " "
	sql_select += "WHERE idtransaksi='" + idinvoice + "' "
	sql_select += "AND resultwigo !='' "

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&result); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return result
}
