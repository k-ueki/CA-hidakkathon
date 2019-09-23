req-login:
	curl -XPOST http://localhost:8080/auth/login -d 'user_id=aaliyah_reichel2696&password=aaliyah_reichel2696' 
req-login2:
	curl -XPOST http://localhost:8080/auth/login -d 'user_id=&password=aaliyah_reichel2696' 

req-event-detail:
	curl -v -XGET http://localhost:8080/event/2 
req-top:
	curl -XGET http://localhost:8080/top\?target_year=2019\&target_month=2

req-top-lack:
	curl -XGET http://localhost:8080/top\?target_year=\&target_month=2

req-event-img:
	curl -XGET http://localhost:8080/event/image/1481 -v

req-del:
	curl -XDELETE http://localhost:8080/event/1481 -v


build :
	docker-compose build

up:
	 docker-compose up -d

run: 
	go run main.go

