
无论是服务集成的metrics，或者业务的metrics，尽量包含 labels 信息。

## labels 命名
labels 的值只能为英文+"-"

* company: (easemob/bluemoon)
* cloud: (online/ssy/vpc)
* department: (im/cec)
* branch: (prod/sdb)
* region:（公有云/vpc为服务器所在区域，私有云统一为客户公司名称,如bluemoon）
* component:（组件和集群，如kafka-statistics/redis-nginx/rest-group1）
* api: api名称。可选，涉及到具体API的需包括 api labe，如{company:"easemob", cloud: "online", department:"cec", branch:"prod", region:"beijing", component:"webapp-a"}


## 服务 metrics 命名:
服务 metrics 只能为英文+数字+"_"

* 服务名，如 rest/webapp/redis
* 业务名，如 login, list(for redis)
* 指标名:
    1. 必须带单位，且单位需统一。单位为复数形式，如 seconds/minutes/bytes
    2. 最好用基本单位，如 seconds/bytes
    3. 如果为累加值，则结尾使用 total，如 webapp_login_total, webapp_login_seconds_total
    4. 如果为状态值，则使用表征的状态的名称，如 length, percent, bytes, 
    5. 如果表征速率，需加上单位，单位在状态前，如 senconds_fail_percent
    
例子： service_redis_list_length{company:"easemob", cloud: "online", department:"cec", branch:"prod", region:"beijing", component:"redis-webapp", name:"xxxxx"} 1000.0 


## 集成可得的 metrics
### http metrics:
* http_server_requests_seconds_count
    ~~~
    # HELP http_server_requests_seconds  
    # TYPE http_server_requests_seconds summary
    http_server_requests_seconds_count{exception="None",method="GET",status="200",uri="root",} 2.0
    http_server_requests_seconds_count{exception="None",method="GET",status="401",uri="root",} 100000.0
    http_server_requests_seconds_count{exception="None",method="POST",status="404",uri="NOT_FOUND",} 1.0
    http_server_requests_seconds_count{exception="None",method="GET",status="307",uri="REDIRECTION",} 6.0
    http_server_requests_seconds_count{exception="None",method="GET",status="404",uri="NOT_FOUND",} 3.0
    ~~~
* http_server_requests_seconds_sum
    ~~~
    http_server_requests_seconds_sum{exception="None",method="GET",status="200",uri="root",} 1.213212096
    http_server_requests_seconds_sum{exception="None",method="GET",status="401",uri="root",} 265.577053146
    http_server_requests_seconds_sum{exception="None",method="POST",status="404",uri="NOT_FOUND",} 0.064696446
    http_server_requests_seconds_sum{exception="None",method="GET",status="307",uri="REDIRECTION",} 0.134987473
    http_server_requests_seconds_sum{exception="None",method="GET",status="404",uri="NOT_FOUND",} 0.105154881
    ~~~
* http_server_requests_seconds_max
    ~~~
    http_server_requests_seconds_max{exception="None",method="GET",status="200",uri="root",} 0.0
    http_server_requests_seconds_max{exception="None",method="GET",status="401",uri="root",} 0.0
    http_server_requests_seconds_max{exception="None",method="POST",status="404",uri="NOT_FOUND",} 0.0
    http_server_requests_seconds_max{exception="None",method="GET",status="307",uri="REDIRECTION",} 0.0
    http_server_requests_seconds_max{exception="None",method="GET",status="404",uri="NOT_FOUND",} 0.0
    ~~~

### jvm metrics:

* jvm_buffer_total_capacity_bytes
    ~~~
    # HELP jvm_buffer_total_capacity_bytes An estimate of the total capacity of the buffers in this pool
    # TYPE jvm_buffer_total_capacity_bytes gauge
    jvm_buffer_total_capacity_bytes{id="direct",} 3.9994909E7
    jvm_buffer_total_capacity_bytes{id="mapped",} 0.0
    ~~~
* jvm_buffer_count
    ~~~
    # HELP jvm_buffer_count An estimate of the number of buffers in the pool
    # TYPE jvm_buffer_count gauge
    jvm_buffer_count{id="direct",} 141.0
    jvm_buffer_count{id="mapped",} 0.0
    ~~~
* jvm_gc_max_data_size_bytes
    ~~~
    # HELP jvm_gc_max_data_size_bytes Max size of old generation memory pool
    # TYPE jvm_gc_max_data_size_bytes gauge
    jvm_gc_max_data_size_bytes 8.589934592E9
    ~~~
* jvm_threads_peak
    ~~~
    # HELP jvm_threads_peak The peak live thread count since the Java virtual machine started or peak was reset
    # TYPE jvm_threads_peak gauge
    jvm_threads_peak 1251.0
    ~~~
* jvm_gc_memory_promoted_bytes_total
    ~~~
    # HELP jvm_gc_memory_promoted_bytes_total Count of positive increases in the size of the old generation memory pool before GC to after GC
    # TYPE jvm_gc_memory_promoted_bytes_total counter
    jvm_gc_memory_promoted_bytes_total 1.16569808E8
    ~~~
* jvm_gc_live_data_size_bytes
    ~~~
    # HELP jvm_gc_live_data_size_bytes Size of old generation memory pool after a full GC
    # TYPE jvm_gc_live_data_size_bytes gauge
    jvm_gc_live_data_size_bytes 1.15906832E8
    ~~~
* jvm_threads_daemon
    ~~~
    # HELP jvm_threads_daemon The current number of live daemon threads
    # TYPE jvm_threads_daemon gauge
    jvm_threads_daemon 30.0
    ~~~
* jvm_memory_committed_bytes
    ~~~
    # HELP jvm_memory_committed_bytes The amount of memory in bytes that is committed for  the Java virtual machine to use
    # TYPE jvm_memory_committed_bytes gauge
    jvm_memory_committed_bytes{area="nonheap",id="Code Cache",} 3.5848192E7
    jvm_memory_committed_bytes{area="nonheap",id="Metaspace",} 7.9167488E7
    jvm_memory_committed_bytes{area="nonheap",id="Compressed Class Space",} 1.0223616E7
    jvm_memory_committed_bytes{area="heap",id="G1 Eden Space",} 5.372903424E9
    jvm_memory_committed_bytes{area="heap",id="G1 Survivor Space",} 3.7748736E7
    jvm_memory_committed_bytes{area="heap",id="G1 Old Gen",} 3.179282432E9
    ~~~
* jvm_memory_max_bytes  
    ~~~
    # HELP jvm_memory_max_bytes The maximum amount of memory in bytes that can be used for memory management
    # TYPE jvm_memory_max_bytes gauge
    jvm_memory_max_bytes{area="nonheap",id="Code Cache",} 2.5165824E8
    jvm_memory_max_bytes{area="nonheap",id="Metaspace",} -1.0
    jvm_memory_max_bytes{area="nonheap",id="Compressed Class Space",} 1.073741824E9
    jvm_memory_max_bytes{area="heap",id="G1 Eden Space",} -1.0
    jvm_memory_max_bytes{area="heap",id="G1 Survivor Space",} -1.0
    jvm_memory_max_bytes{area="heap",id="G1 Old Gen",} 8.589934592E9
    ~~~
* jvm_memory_used_bytes
    ~~~
    # HELP jvm_memory_used_bytes The amount of used memory
    # TYPE jvm_memory_used_bytes gauge
    jvm_memory_used_bytes{area="nonheap",id="Code Cache",} 3.5118464E7
    jvm_memory_used_bytes{area="nonheap",id="Metaspace",} 7.742316E7
    jvm_memory_used_bytes{area="nonheap",id="Compressed Class Space",} 9932704.0
    jvm_memory_used_bytes{area="heap",id="G1 Eden Space",} 1.501560832E9
    jvm_memory_used_bytes{area="heap",id="G1 Survivor Space",} 3.7748736E7
    jvm_memory_used_bytes{area="heap",id="G1 Old Gen",} 1.15906832E8
    ~~~
* jvm_buffer_memory_used_bytes
    ~~~
    # HELP jvm_buffer_memory_used_bytes An estimate of the memory that the Java virtual machine is using for this buffer pool
    # TYPE jvm_buffer_memory_used_bytes gauge
    jvm_buffer_memory_used_bytes{id="direct",} 3.9994909E7
    jvm_buffer_memory_used_bytes{id="mapped",} 0.0
    ~~~
* jvm_threads_live
    ~~~
    # HELP jvm_threads_live The current number of live threads including both daemon and non-daemon threads
    # TYPE jvm_threads_live gauge
    jvm_threads_live 1251.0
    ~~~

* jvm_gc_memory_allocated_bytes_total
    ~~~
    # HELP jvm_gc_memory_allocated_bytes_total Incremented for an increase in the size of the young generation memory pool after one GC to before the next
    # TYPE jvm_gc_memory_allocated_bytes_total counter
    jvm_gc_memory_allocated_bytes_total 8.5840625664E10
    ~~~
* jvm_gc_pause_seconds_count
    ~~~
    # HELP jvm_gc_pause_seconds Time spent in GC pause
    # TYPE jvm_gc_pause_seconds summary
    jvm_gc_pause_seconds_count{action="end of minor GC",cause="G1 Evacuation Pause",} 18.0
    jvm_gc_pause_seconds_sum{action="end of minor GC",cause="G1 Evacuation Pause",} 3.244
    jvm_gc_pause_seconds_max{action="end of minor GC",cause="G1 Evacuation Pause",} 0.0
    ~~~
* jvm_gc_pause_seconds_sum
    ~~~
    jvm_gc_pause_seconds_sum{action="end of minor GC",cause="G1 Evacuation Pause",} 3.244
    ~~~
* jvm_gc_pause_seconds_max
    ~~~
    jvm_gc_pause_seconds_max{action="end of minor GC",cause="G1 Evacuation Pause",} 0.0
    ~~~

### 进程
* system_cpu_usage
    ~~~
    # HELP system_cpu_usage The "recent cpu usage" for the whole system
    # TYPE system_cpu_usage gauge
    system_cpu_usage 0.008894536213468869
    ~~~
* process_start_time_seconds
    ~~~
    # HELP process_start_time_seconds The start time of the Java virtual machine
    # TYPE process_start_time_seconds gauge
    process_start_time_seconds 1.51741692209E9
    ~~~
* process_uptime_seconds
    ~~~
    # HELP process_uptime_seconds The uptime of the Java virtual machine
    # TYPE process_uptime_seconds gauge
    process_uptime_seconds 66530.512
    ~~~
* system_load_average_1m
    ~~~
    # HELP system_load_average_1m The sum of the number of runnable entities queued to available processors and the number of runnable entities running on the available processors averaged over a period of time
    # TYPE system_load_average_1m gauge
    system_load_average_1m 0.0
    ~~~

* logback_events_total
    ~~~
    # HELP logback_events_total Number of error level events that made it to the logs
    # TYPE logback_events_total counter
    logback_events_total{level="error",} 2.0
    logback_events_total{level="warn",} 100025.0
    logback_events_total{level="info",} 100397.0
    logback_events_total{level="debug",} 0.0
    logback_events_total{level="trace",} 0.0
    ~~~
* system_cpu_count
    ~~~
    # HELP system_cpu_count The number of processors available to the Java virtual machine
    # TYPE system_cpu_count gauge
    system_cpu_count 4.0
    ~~~
* process_cpu_usage
    ~~~
    # HELP process_cpu_usage The "recent cpu usage" for the Java Virtual Machine process
    # TYPE process_cpu_usage gauge
    process_cpu_usage 0.0076335877862595426
    ~~~


