@echo off
if docker-compose ps | find /v /c "" == "2" (
	docker-compose up -d --build
) else (
	docker-compose start
)
