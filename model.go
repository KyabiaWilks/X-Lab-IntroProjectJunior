package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type DeleteRequest struct {
	ID int `json:"id"`
}

var total int64
var invertOrder bool
var mutex = &sync.Mutex{}

func getComments(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1 // 默认值为 1
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || (size < 1 && size != -1) {
		size = 10 // 默认值为 10
	}
	fmt.Printf("page&size: %d %d\n", page, size)
	var comments []Comment
	db.Model(&Comment{}).Count(&total)

	if invertOrder {
		if size == -1 {
			db.Order("id desc").Find(&comments)
		} else {
			db.Order("id desc").Offset((page - 1) * size).Limit(size).Find(&comments)
		}
	} else {
		if size == -1 {
			db.Find(&comments)
		} else {
			db.Offset((page - 1) * size).Limit(size).Find(&comments)
		}
	}

	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: map[string]interface{}{
			"total":    total,
			"comments": comments,
		},
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func addComment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var comment Comment
	var input struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//comment.ID = total
	comment.Name = input.Name
	comment.Content = input.Content
	db.Create(&comment)

	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: comment,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var req DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 确保 id 是有效的
	if req.ID <= 0 {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	// 删除指定 ID 的评论
	db.Delete(&Comment{}, req.ID)
	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}

	json.NewEncoder(w).Encode(resp)
}

func switchOrder(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	invertOrder = !invertOrder

	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}

	json.NewEncoder(w).Encode(resp)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
