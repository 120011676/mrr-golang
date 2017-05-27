#!/bin/bash
if [[ `docker-compose ps | grep -c ''` == 2 ]]; then
	docker-compose up -d --build
else
	docker-compose start
fi