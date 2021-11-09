package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	// "net/url"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
    "strings"
	"strconv"
	"time"
	"sync"

	// "github.com/gin-gonic/gin"

	"digipos/models"
	// elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

func GetHistoryPurchaseCronjobES() {

	startDate := time.Now().Add(-2*time.Hour)
	endDate := time.Now().Add(-1*time.Hour).Add(-1*time.Second)

	dealer_code := []string{
		"11129", "22123", "41149", "12132", "31131", "11133", "11116", "12149", "12120", "13110", 
		"13135", "21108", "23111", "24102", "22101", "22119", "22114", "31133", "31127", "32135", 
		"32101", "32130", "32103", "41114", "41152", "41116", "42132", "42119", "42151", "42112", 
		"42149", "43129", "43117", "43124", "43127", "41103", "11108", "13127", "31132", "41126", 
		"11122", "11104", "11113", "12129", "12146", "13119", "13136", "13132", "21132", "23107", 
		"23129", "24109", "22134", "31105", "31122", "32102", "32138", "32124", "33101", "33126"} 
		// "33124", "41113", "41124", "41137", "42115", "42130", "42134", "42150", "42122", "42138", 
		// "43104", "43118", "43105", "43110", "41128", "13113", "42146", "11103", "11130", "11135", 
		// "11105", "12144", "12151", "13124", "13103", "21130", "23123", "22120", "22127", "22133", 
		// "31124", "31136", "32131", "32120", "32134", "33114", "41151", "42127", "32139", "00301", 
		// "12153", "12152", "23133", "31137", "42152", "41154", "00501", "41157", "24111", "21134", 
		// "23134", "11137", "13138", "13139", "21133", "24110", "22135", "32140", "33127", "41155", "41156", "43130"}

	client := &http.Client{}

	sliceLength := len(dealer_code)
	var wg sync.WaitGroup
	wg.Add(sliceLength)

	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			code:=dealer_code[i]
			var jsonData = []byte(`{
				"creditParty": "`+ code +`",
				"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
				"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`",
				"limit": 1000,
				"page": 1,
				"serviceName": "10003438"
			}`)

			req, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyPurchase", bytes.NewBuffer(jsonData))
			req.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Accept", "application/octet-stream")
		
			resp, error := client.Do(req)
			if error != nil {
				panic(error)
			}

			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}

				var result models.Response
    			if err := json.Unmarshal(bodyBytes, &result); err != nil {  // Parse []byte to the go struct pointer
        			fmt.Println("Can not unmarshal JSON")
    			}

				for _, rec := range result.Data {
					var jsonDataRes = []byte(`{
						"link_aja_no":  "`+ rec.LinkAjaNo +`",
						"initiation_time": "`+ rec.InitiationTime +`",
						"service_name": "`+ rec.ServiceName +`",
						"initiator_party": "`+ rec.InitiatorParty +`",
						"credit_party": "`+ rec.CreditParty +`",
						"debit_party": "`+ rec.DebitParty +`",
						"transaction_status": "`+ rec.TransactionStatus +`",
						"transaction_amount": `+ strconv.Itoa(rec.TransactionAmount) +`,
						"receipt_no": "`+ rec.ReceiptNo +`",
						"transaction_flag": "`+ rec.TransactionFlag +`"
					}`)

					reqRes, _ := http.NewRequest("POST", "http://localhost:9200/digipos/history_purchases/", bytes.NewBuffer(jsonDataRes))
					reqRes.Header.Add("Content-Type", "application/json")
				
					_, error := client.Do(reqRes)
					if error != nil {
						panic(error)
					}
				}
			}

			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func GetHistoryPurchaseCronjobES2() {

	startDate := time.Now().Add(-2*time.Hour)
	endDate := time.Now().Add(-1*time.Hour).Add(-1*time.Second)

	dealer_code := []string{
		"11129", "22123", "41149", "12132", "31131", "11133", "11116", "12149", "12120", "13110", 
		"13135", "21108", "23111", "24102", "22101", "22119", "22114", "31133", "31127", "32135", 
		"32101", "32130", "32103", "41114", "41152", "41116", "42132", "42119", "42151", "42112", 
		"42149", "43129", "43117", "43124", "43127", "41103", "11108", "13127", "31132", "41126", 
		"11122", "11104", "11113", "12129", "12146", "13119", "13136", "13132", "21132", "23107", 
		"23129", "24109", "22134", "31105", "31122", "32102", "32138", "32124", "33101", "33126", 
		"33124", "41113", "41124", "41137", "42115", "42130", "42134", "42150", "42122", "42138", 
		"43104", "43118", "43105", "43110", "41128", "13113", "42146", "11103", "11130", "11135", 
		"11105", "12144", "12151", "13124", "13103", "21130", "23123", "22120", "22127", "22133", 
		"31124", "31136", "32131", "32120", "32134", "33114", "41151", "42127", "32139", "00301", 
		"12153", "12152", "23133", "31137", "42152", "41154", "00501", "41157", "24111", "21134", 
		"23134", "11137", "13138", "13139", "21133", "24110", "22135", "32140", "33127", "41155", "41156", "43130"}

	client := &http.Client{}

	sliceLength := len(dealer_code)
	var wg sync.WaitGroup
	wg.Add(sliceLength)

	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			code:=dealer_code[i]
			var jsonData2 = []byte(`{
				"creditParty": "`+ code +`",
				"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
				"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`",
				"limit": 1000,
				"page": 1,
				"serviceName": "10003465"
			}`)

			req2, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyPurchase", bytes.NewBuffer(jsonData2))
			req2.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
			req2.Header.Add("Content-Type", "application/json")
			req2.Header.Add("Accept", "application/octet-stream")
		
			resp2, error := client.Do(req2)
			if error != nil {
				panic(error)
			}

			if resp2.StatusCode == http.StatusOK {
				bodyBytes2, err := ioutil.ReadAll(resp2.Body)
				if err != nil {
					fmt.Println(err)
				}

				var result2 models.Response
    			if err := json.Unmarshal(bodyBytes2, &result2); err != nil {  // Parse []byte to the go struct pointer
        			fmt.Println("Can not unmarshal JSON")
    			}

				for _, rec := range result2.Data {
					var jsonDataRes2 = []byte(`{
						"link_aja_no":  "`+ rec.LinkAjaNo +`",
						"initiation_time": "`+ rec.InitiationTime +`",
						"service_name": "`+ rec.ServiceName +`",
						"initiator_party": "`+ rec.InitiatorParty +`",
						"credit_party": "`+ rec.CreditParty +`",
						"debit_party": "`+ rec.DebitParty +`",
						"transaction_status": "`+ rec.TransactionStatus +`",
						"transaction_amount": `+ strconv.Itoa(rec.TransactionAmount) +`,
						"receipt_no": "`+ rec.ReceiptNo +`",
						"transaction_flag": "`+ rec.TransactionFlag +`"
					}`)

					reqRes2, _ := http.NewRequest("POST", "http://localhost:9200/digipos/history_purchases/", bytes.NewBuffer(jsonDataRes2))
					reqRes2.Header.Add("Content-Type", "application/json")
				
					_, error := client.Do(reqRes2)
					if error != nil {
						panic(error)
					}
				}
			}

			defer wg.Done()	
		}(i)
	}
	wg.Wait()
}

func GetHistoryPurchaseCronjobES3() {

	startDate := time.Now().Add(-2*time.Hour)
	endDate := time.Now().Add(-1*time.Hour).Add(-1*time.Second)

	dealer_code := []string{
		"11129", "22123", "41149", "12132", "31131", "11133", "11116", "12149", "12120", "13110", 
		"13135", "21108", "23111", "24102", "22101", "22119", "22114", "31133", "31127", "32135", 
		"32101", "32130", "32103", "41114", "41152", "41116", "42132", "42119", "42151", "42112", 
		"42149", "43129", "43117", "43124", "43127", "41103", "11108", "13127", "31132", "41126", 
		"11122", "11104", "11113", "12129", "12146", "13119", "13136", "13132", "21132", "23107", 
		"23129", "24109", "22134", "31105", "31122", "32102", "32138", "32124", "33101", "33126", 
		"33124", "41113", "41124", "41137", "42115", "42130", "42134", "42150", "42122", "42138", 
		"43104", "43118", "43105", "43110", "41128", "13113", "42146", "11103", "11130", "11135", 
		"11105", "12144", "12151", "13124", "13103", "21130", "23123", "22120", "22127", "22133", 
		"31124", "31136", "32131", "32120", "32134", "33114", "41151", "42127", "32139", "00301", 
		"12153", "12152", "23133", "31137", "42152", "41154", "00501", "41157", "24111", "21134", 
		"23134", "11137", "13138", "13139", "21133", "24110", "22135", "32140", "33127", "41155", "41156", "43130"}

	client := &http.Client{}

	sliceLength := len(dealer_code)
	var wg sync.WaitGroup
	wg.Add(sliceLength)

	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			code:=dealer_code[i]
			var jsonData3 = []byte(`{
				"debitParty": "`+ code +`",
				"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
				"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`",
				"limit": 1000,
				"page": 1
			}`)

			req3, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyDeposit", bytes.NewBuffer(jsonData3))
			req3.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
			req3.Header.Add("Content-Type", "application/json")
			req3.Header.Add("Accept", "application/octet-stream")
		
			resp3, error := client.Do(req3)
			if error != nil {
				panic(error)
			}

			if resp3.StatusCode == http.StatusOK {
				bodyBytes3, err := ioutil.ReadAll(resp3.Body)
				if err != nil {
					fmt.Println(err)
				}

				var result3 models.Response
    			if err := json.Unmarshal(bodyBytes3, &result3); err != nil {  // Parse []byte to the go struct pointer
        			fmt.Println("Can not unmarshal JSON")
    			}

				for _, rec := range result3.Data {
					var jsonDataRes3 = []byte(`{
						"link_aja_no":  "`+ rec.LinkAjaNo +`",
						"initiation_time": "`+ rec.InitiationTime +`",
						"service_name": "`+ rec.ServiceName +`",
						"initiator_party": "`+ rec.InitiatorParty +`",
						"credit_party": "`+ rec.CreditParty +`",
						"debit_party": "`+ rec.DebitParty +`",
						"transaction_status": "`+ rec.TransactionStatus +`",
						"transaction_amount": `+ strconv.Itoa(rec.TransactionAmount) +`,
						"receipt_no": "`+ rec.ReceiptNo +`",
						"transaction_flag": "`+ rec.TransactionFlag +`"
					}`)

					reqRes3, _ := http.NewRequest("POST", "http://localhost:9200/digipos/history_deposits/", bytes.NewBuffer(jsonDataRes3))
					reqRes3.Header.Add("Content-Type", "application/json")
				
					_, error := client.Do(reqRes3)
					if error != nil {
						panic(error)
					}
				}
			}

			defer wg.Done()	
		}(i)
	}
	wg.Wait()
}

func GetHistoryPurchaseCronjob() {

	startDate := time.Now().Add(-4*time.Hour)
	endDate := time.Now().Add(-3*time.Hour).Add(-1*time.Second)

	dealer_code := []string{
		"11129", "22123", "41149", "12132", "31131", "11133", "11116", "12149", "12120", "13110", 
		"13135", "21108", "23111", "24102", "22101", "22119", "22114", "31133", "31127", "32135", 
		"32101", "32130", "32103", "41114", "41152", "41116", "42132", "42119", "42151", "42112", 
		"42149", "43129", "43117", "43124", "43127", "41103", "11108", "13127", "31132", "41126", 
		"11122", "11104", "11113", "12129", "12146", "13119", "13136", "13132", "21132", "23107", 
		"23129", "24109", "22134", "31105", "31122", "32102", "32138", "32124", "33101", "33126", 
		"33124", "41113", "41124", "41137", "42115", "42130", "42134", "42150", "42122", "42138", 
		"43104", "43118", "43105", "43110", "41128", "13113", "42146", "11103", "11130", "11135", 
		"11105", "12144", "12151", "13124", "13103", "21130", "23123", "22120", "22127", "22133", 
		"31124", "31136", "32131", "32120", "32134", "33114", "41151", "42127", "32139", "00301", 
		"12153", "12152", "23133", "31137", "42152", "41154", "00501", "41157", "24111", "21134", 
		"23134", "11137", "13138", "13139", "21133", "24110", "22135", "32140", "33127", "41155", "41156", "43130"}

	client := &http.Client{}

	sliceLength := len(dealer_code)
	var wg sync.WaitGroup
	wg.Add(sliceLength)

	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			code:=dealer_code[i]
			var jsonData = []byte(`{
				"creditParty": "`+ code +`",
				"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
				"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`",
				"serviceName": "10003438"
			}`)

			req, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyPurchaseDownload", bytes.NewBuffer(jsonData))
			req.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Accept", "application/octet-stream")
		
			resp, error := client.Do(req)
			if error != nil {
				panic(error)
			}

			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				bodyString := string(bodyBytes)
				results := strings.ReplaceAll(bodyString, "&quot;", " ")
				reader := csv.NewReader(strings.NewReader(results))
				reader.Comma = ';'
				reader.LazyQuotes = true
				data, err := reader.ReadAll()
				if err != nil {
					fmt.Println("error "+ code)
					// panic(err)
				}
				
				if data != nil {
					for idx, row := range data {
						// skip header
						if idx == 0 {
							continue
						}
						// statusTrans := strings.TrimSpace(row[6])
						// if(statusTrans!= "Completed"){
						// 	continue
						// }
				
						amount, _ := strconv.Atoi(strings.TrimSpace(row[7]))
						addHistoryPurchase := models.HistoryPurchase{LinkAjaNo: strings.TrimSpace(row[0]), InitiationTime: strings.TrimSpace(row[1]), ServiceName: strings.TrimSpace(row[2]), InitiatorParty: strings.TrimSpace(row[3]), CreditParty: strings.TrimSpace(row[4]), DebitParty: strings.TrimSpace(row[5]), TransactionStatus: strings.TrimSpace(row[6]), TransactionAmount: amount, ReceiptNo: strings.TrimSpace(row[8]),}
						if err := models.MPosGORM.Create(&addHistoryPurchase).Error; err != nil {
							fmt.Printf("error add History Purchase: %3v \n", err)
							return
						}
					}
				}
			}

			var jsonDataOuter = []byte(`{
				"creditParty": "`+ code +`",
				"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
				"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`",
				"serviceName": "10003465"
			}`)

			req2, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyPurchaseDownload", bytes.NewBuffer(jsonDataOuter))
			req2.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
			req2.Header.Add("Content-Type", "application/json")
			req2.Header.Add("Accept", "application/octet-stream")
		
			resp2, error := client.Do(req2)
			if error != nil {
				panic(error)
			}

			if resp2.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp2.Body)
				if err != nil {
					fmt.Println(err)
				}
				bodyString := string(bodyBytes)
				results := strings.ReplaceAll(bodyString, "&quot;", " ")
				reader := csv.NewReader(strings.NewReader(results))
				reader.Comma = ';'
				reader.LazyQuotes = true
				data, err := reader.ReadAll()
				if err != nil {
					fmt.Println("error "+ code)
					// panic(err)
				}
				
				if data != nil {
					for idx, row := range data {
						// skip header
						if idx == 0 {
							continue
						}
						// statusTrans := strings.TrimSpace(row[6])
						// if(statusTrans!= "Completed"){
						// 	continue
						// }
				
						amount, _ := strconv.Atoi(strings.TrimSpace(row[7]))
						addHistoryPurchase := models.HistoryPurchase{LinkAjaNo: strings.TrimSpace(row[0]), InitiationTime: strings.TrimSpace(row[1]), ServiceName: strings.TrimSpace(row[2]), InitiatorParty: strings.TrimSpace(row[3]), CreditParty: strings.TrimSpace(row[4]), DebitParty: strings.TrimSpace(row[5]), TransactionStatus: strings.TrimSpace(row[6]), TransactionAmount: amount, ReceiptNo: strings.TrimSpace(row[8]),}
						if err := models.MPosGORM.Create(&addHistoryPurchase).Error; err != nil {
							fmt.Printf("error add History Purchase: %3v \n", err)
							return
						}
					}
				}
			}

			var jsonDataDeposit = []byte(`{
				"debitParty": "`+ code +`",
				"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
				"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`"
				
			}`)

			reqDeposit, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyDepositDownload", bytes.NewBuffer(jsonDataDeposit))
			reqDeposit.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
			reqDeposit.Header.Add("Content-Type", "application/json")
			reqDeposit.Header.Add("Accept", "application/octet-stream")
		
			respDeposit, error := client.Do(reqDeposit)
			if error != nil {
				panic(error)
			}

			if respDeposit.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(respDeposit.Body)
				if err != nil {
					fmt.Println(err)
				}
				bodyString := string(bodyBytes)
				results := strings.ReplaceAll(bodyString, "&quot;", " ")
				reader := csv.NewReader(strings.NewReader(results))
				reader.Comma = ';'
				reader.LazyQuotes = true
				data, err := reader.ReadAll()
				if err != nil {
					fmt.Println("error "+ code)
					// panic(err)
				}

				for idx, row := range data {
					// skip header
					if idx == 0 {
						continue
					}

					// statusTrans := strings.TrimSpace(row[6])
					// if(statusTrans != "Completed"){
					// 	continue
					// }
					
					amount, _ := strconv.Atoi(strings.TrimSpace(row[7]))
					addHistoryDeposit := models.HistoryDeposit{LinkAjaNo: strings.TrimSpace(row[0]), InitiationTime: strings.TrimSpace(row[1]), ServiceName: strings.TrimSpace(row[2]), InitiatorParty: strings.TrimSpace(row[3]), CreditParty: strings.TrimSpace(row[4]), DebitParty: strings.TrimSpace(row[5]), TransactionStatus: strings.TrimSpace(row[6]), TransactionAmount: amount, ReceiptNo: strings.TrimSpace(row[8]),}
					if err := models.MPosGORM.Create(&addHistoryDeposit).Error; err != nil {
						fmt.Printf("error add History Deposit: %3v \n", err)
						return
					}
				}
			}	
			defer wg.Done()	
		}(i)
	}
	wg.Wait()
}

func GetHistoryDepositCronjob() {

	startDate := time.Now().Add(-4*time.Hour)
	endDate := time.Now().Add(-3*time.Hour).Add(-1*time.Second)

	dealer_code := []string{"11129", "22123", "41149", "12132", "31131", "11133", "11116", "12149", "12120", "13110", "13135", "21108", "23111", "24102", "22101", "22119", "22114", "31133", "31127", "32135", "32101", "32130", "32103", "41114", "41152", "41116", "42132", "42119", "42151", "42112", "42149", "43129", "43117", "43124", "43127", "41103", "11108", "13127", "31132", "41126", "11122", "11104", "11113", "12129", "12146", "13119", "13136", "13132", "21132", "23107", "23129", "24109", "22134", "31105", "31122", "32102", "32138", "32124", "33101", "33126", "33124", "41113", "41124", "41137", "42115", "42130", "42134", "42150", "42122", "42138", "43104", "43118", "43105", "43110", "41128", "13113", "42146", "11103", "11130", "11135", "11105", "12144", "12151", "13124", "13103", "21130", "23123", "22120", "22127", "22133", "31124", "31136", "32131", "32120", "32134", "33114", "41151", "42127", "32139", "00301", "12153", "12152", "23133", "31137", "42152", "41154", "00501", "41157", "24111", "21134", "23134", "11137", "13138", "13139", "21133", "24110", "22135", "32140", "33127", "41155", "41156", "43130"}

	client := &http.Client{}

	for _, code := range dealer_code {
		var jsonData = []byte(`{
			"debitParty": "`+ code +`",
			"startDate": "`+ startDate.Format("2006-01-02 15:04:05") +`",
			"endDate": "`+ endDate.Format("2006-01-02 15:04:05") +`"
			
		}`)

		req, _ := http.NewRequest("POST", "https://partner.linkaja.com/apidbx/v1/historyDepositDownload", bytes.NewBuffer(jsonData))
		req.Header.Add("Authorization", "Basic ZGlnaXBvczo3SDdOYVQ0eWhEbkR0ekRVNTdVRlA0NEdS")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/octet-stream")
	
		resp, error := client.Do(req)
		if error != nil {
			panic(error)
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			bodyString := string(bodyBytes)
			results := strings.ReplaceAll(bodyString, "&quot;", " ")
			reader := csv.NewReader(strings.NewReader(results))
			reader.Comma = ';'
			reader.LazyQuotes = true
			data, err := reader.ReadAll()
			if err != nil {
				fmt.Println("error "+ code)
				// panic(err)
			}

			for idx, row := range data {
				// skip header
				if idx == 0 {
					continue
				}
				
				amount, _ := strconv.Atoi(strings.TrimSpace(row[7]))
				addHistoryDeposit := models.HistoryDeposit{LinkAjaNo: strings.TrimSpace(row[0]), InitiationTime: strings.TrimSpace(row[1]), ServiceName: strings.TrimSpace(row[2]), InitiatorParty: strings.TrimSpace(row[3]), CreditParty: strings.TrimSpace(row[4]), DebitParty: strings.TrimSpace(row[5]), TransactionStatus: strings.TrimSpace(row[6]), TransactionAmount: amount, ReceiptNo: strings.TrimSpace(row[8]),}
				if err := models.MPosGORM.Create(&addHistoryDeposit).Error; err != nil {
					fmt.Printf("error add History Deposit: %3v \n", err)
					return
				}
			}
		}
	}
}