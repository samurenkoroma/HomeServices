repeat=1000
for n in $(seq $repeat); 
    do
        curl -X POST -d '{"url":"https://google.com"}' http://localhost:8081/link
    done