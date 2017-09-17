#!/bin/bash

promfile=node_lvs_precent.prom

LVS="/sbin/lvs --noheadings -o lv_name,data_percent,metadata_percent"

while read -r line; do
  read -a arr <<< $line
  if [[ ! -z "${arr[1]}" && ! -z "${arr[2]}" ]]; then
            echo "node_lvs_precent_used{instance=\"data\", type=\"thinpool\", pool=\"${arr[0]}\"} ${arr[1]}" > $promfile
            echo "node_lvs_precent_used{instance=\"metadata\", type=\"thinpool\", pool=\"${arr[0]}\"} ${arr[2]}" >> $promfile
  fi

done < <($LVS)

