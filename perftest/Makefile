docker-run:
	docker-compose up --remove-orphans -d

perftest:
	docker-compose run k6 -e API_HOSTNAME=http://host-1906398250-port-58250.proxy.infralabs.cs.ui.ac.id/ run --include-system-env-vars --out influxdb=http://influxdb:8086/k6db /k6-script/oauth-scenario.js

stop:
	docker-compose down

volume-prune:
	docker volume prune

clean: 
	stop volume-prune