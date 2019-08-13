/*
* @Author: scottxiong
* @Date:   2019-08-14 01:57:01
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-14 02:07:53
 */
package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/scott-x/gcmd"
)

func main() {
	d := color.New(color.FgCyan, color.Bold)
	e := color.New(color.FgRed, color.Bold)

	gcmd.AddQuestion("name", "select the method to share the revisions: ", "Please input correct number: ", "^[12]$")
	d.Println("1, ProofHQ")
	d.Println("2, Email")
	answers := gcmd.Exec()
	fmt.Println("")
	switch answers["name"] {
	case "1":
		e.Println(`Hi xxx,

Please click "go to proof" to see the comments from our proofing team and make revisions accordingly. Once you finished, please send back the requested files to our team in this address:

service@benchmarkdesign.org`)
		fmt.Println("")
	case "2":
		e.Println(`Hi XXX,

Please check the comments made in attached file and make revisions accordingly, once finished, please send back to us for approval.	`)
		fmt.Println("")
	}
}
