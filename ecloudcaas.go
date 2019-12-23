/*
 * @Author: your name
 * @Date: 2019-12-04 13:51:37
 * @LastEditTime : 2019-12-23 09:54:46
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \undefinedd:\gohome\src\github.com\beiertiao\ecloudcaas\ecloudcaas.go
 */
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "branch dev")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
