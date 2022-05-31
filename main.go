// Copyright 2022 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
import (
	"bytes"
)

// Config is a struct for storing the parsed config.json
type Config struct {
	FbToken   string `json:"fb_token"`
	PageToken string `json:"page_token"`
}

// MessageJSON is the message that is sent back to the user
type MessageJSON struct {
	AccessToken string `json:"access_token"`
	Recipient   struct {
		ID int64 `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

// Response is the message that is sent from from Facebook
type Response struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        int64 `json:"id"`
		Time      int64 `json:"time"`
		Messaging []struct {
			Sender struct {
				ID int64 `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID int64 `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message   struct {
				Mid  string `json:"mid"`
				Seq  int    `json:"seq"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}

type WP struct {
	client  *http.Client
	baseURL string
	token   string
}

func NewWP(token string) *WP {
	return &WP{
		client:  http.DefaultClient,
		baseURL: "https://www.workplace.com/scim/v1/",
		token:   token,
	}
}

type thread struct {
	Recipient struct {
		Ids []int `json:"ids"`
	} `json:"recipient"`
	Message string `json:"message"`
}

func (w *WP) CreateThread(th thread) error {
	body, err := json.Marshal(th)
	if err != nil {
		return err
	}

	resp, err := w.client.Post(w.baseURL+"/me/message?access_token=?"+w.token, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(buf))

	return nil
}

//func sendMessage(sender int64, messageData string, cfg Config) {
//	m := MessageJSON{}
//	m.AccessToken = cfg.PageToken
//	m.Recipient.ID = sender
//	m.Message.Text = messageData
//
//	b := new(bytes.Buffer)
//	json.NewEncoder(b).Encode(m)
//
//	resp, err := http.Post("https://graph.facebook.com/v2.6/me/messages", "application/json", b)
//	if err != nil {
//		log.Println(err)
//	}
//	defer resp.Body.Close()
//
//	all, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//	}
//
//	log.Println(string(all))
//
//}

// https://developers.facebook.com/docs/workplace/reference/graph-api/reported-content
// https://github.com/fbsamples/workplace-platform-samples/blob/main/SampleAPIEndpoints/Postman/Workplace_Graph_Collection.json
// https://developers.facebook.com/docs/workplace/reference/graph-api/community

func main() {
	wp := NewWP("")

	err := wp.CreateThread(thread{
		Recipient: struct {
			Ids []int `json:"ids"`
		}{
			Ids: []int{1},
		},
		Message: "Hey there, I'm a bot!",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

// Ainsley - 100032732146773
// Jon - 100028939157999
// Kirk - 100024237102328
// Alex - 100070928680832
// Holly - 100056882813301
// Jack - 100020455205110

// Dev & Design - t_3774250762598367
