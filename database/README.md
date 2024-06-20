Golang migrate is error:
```bash
migrate -path databases/migration -database "postgres://habi:habi123@localhost:5432/triva?sslmode=disable" -verbose up
INFO[0000]log.go:104 gosnowflake.(*defaultLogger).Infof reset OCSP cache file. /home/habi/.cache/snowflake/ocsp_response_cache.json 
INFO[0000]log.go:104 gosnowflake.(*defaultLogger).Infof reading OCSP Response cache file. /home/habi/.cache/snowflake/ocsp_response_cache.json 
ERRO[0000]log.go:120 gosnowflake.(*defaultLogger).Errorf failed to open. Ignored. open /home/habi/.cache/snowflake/ocsp_response_cache.json: no such file or directory 
2024/06/13 19:50:54 no change
2024/06/13 19:50:54 Finished after 880.557µs
2024/06/13 19:50:54 Closing source and database
```

For now, don't use migration