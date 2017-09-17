# lvs docker thinpool output fot node_exporter collector

$ /node_exporter --collector.textfile.directory textfile_collector


```prom
node_lvs_precent_used{instance="data", type="thinpool", pool="docker-pool"} 0.00
node_lvs_precent_used{instance="metadata", type="thinpool", pool="docker-pool"} 0.03
```
