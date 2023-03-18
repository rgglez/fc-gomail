/*
BSD 3-Clause License

Copyright (c) 2023, Rodolfo González González

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conf := config{}
		cfg := conf.load()

		smtpHost := fmt.Sprint(cfg["SENDMETRIC"].(map[string]interface{})["HOST"])
		smtpPort, _ := strconv.Atoi(fmt.Sprint(cfg["SENDMETRIC"].(map[string]interface{})["PORT"]))
		smtpLogin := fmt.Sprint(cfg["SENDMETRIC"].(map[string]interface{})["LOGIN"])
		smtpPassword := fmt.Sprint(cfg["SENDMETRIC"].(map[string]interface{})["PASSWORD"])

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var mnsMessage map[string]interface{}
		err = json.Unmarshal([]byte(body), &mnsMessage)
		if err != nil {
			panic(err)
		}
		mnsJson := map[string]string{}
		err = json.Unmarshal([]byte(mnsMessage["Message"].(string)), &mnsJson)
		if err != nil {
			panic(err)
		}

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", mnsJson["From"])
		mailer.SetHeader("To", mnsJson["To"])
		mailer.SetHeader("Subject", mnsJson["Subject"])
		mailer.SetBody("text/html", mnsJson["Body"])

		n := gomail.NewDialer(smtpHost, smtpPort, smtpLogin, smtpPassword)
		if err := n.DialAndSend(mailer); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	fmt.Println("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
