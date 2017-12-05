package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	//"strconv"
	//"io/ioutil"
)

type Node_lvs_precent_used struct {
	lvmname string
	data    string
	meta    string
}

func main() {
	cmd := exec.Command("/sbin/lvs", "--noheadings", "-o", "lv_name,data_percent,metadata_percent")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Run()
	output := cmdOutput.Bytes()
	//fmt.Println(output)
	parseOutput(output)
}

func parseOutput(outs []byte) {
	var average = regexp.MustCompile(`docker-([a-z]+)(\s+)(\d+.\d+)(\s+)(\d+.\d+)`)
	result := average.FindAllString(string(outs), -1)
	s := strings.Fields(result[0])
	lvmname, data, meta := s[0], s[1], s[2]

	//lvstat := Node_lvs_precent_used{lvmname: s[0], data: s[1], meta: s[2]}
	//fmt.Println(lvstat)
	fmt.Printf("node_lvs_precent_used{instance=\"metadata\", type=\"thinpool\", pool=\"%v\"} %v\n", lvmname, meta)
	fmt.Printf("node_lvs_precent_used{instance=\"data\", type=\"thinpool\", pool=\"%v\"} %v\n", lvmname, data)
	//b := []byte(fmt.Sprintf("node_lvs_precent_used{instance=\"data\", type=\"thinpool\", pool=\"%v\"} %v\n", lvmname, data))
	//ioutil.WriteFile("output.txt", b, 0644)
}

// https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
