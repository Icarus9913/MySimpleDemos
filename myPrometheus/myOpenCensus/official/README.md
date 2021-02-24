#OpenCensus


![histogram指标](histogram.png)

> 1. 其中的bucket表示范围,设置bucket=[1,5,10]  
>   
> 2. observe表示采样点落在该bucket中的数量,即落在[~,1]的样点数是2, 落在[1,5]的样点数是3, 落在[5,10]的样点数为1
>   
> 3. write是得到的最终结果(histogram的最终结果bucket的计数是向下包含的,如最下面的是2,倒数第二个是2+3=5,最上面那个是2+3+1=6):  
> &nbsp;&nbsp;[basename]_bucket{le=“1”} = 2  
> &nbsp;&nbsp;[basename]_bucket{le=“5”} =3  
> &nbsp;&nbsp;[basename]_bucket{le=“10”} =6  
> &nbsp;&nbsp;[basename]_bucket{le="+Inf"} = 6  
> &nbsp;&nbsp;[basename]_count =6  
> &nbsp;&nbsp;[basename]_sum =18.8378745    
> 最后count值就是write的值,也就是样点数量的总和.  sum表示的是实际采样点的数据值总和  
>   
> histogram并不会保存数据采样点值，每个bucket只有个记录样本数的counter（float64），即histogram存储的是区间的样本数统计值，因此客户端性能开销相比 Counter 和 Gauge 而言没有明显改变，适合高并发的数据收集。



##Prometheus
* 配置文件
```shell
# my global config
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).
  

# Alertmanager configuration
alerting:
  alertmanagers:
  - static_configs:
    - targets:
      # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'prometheus'

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
    - targets: ['localhost:9090']
```
