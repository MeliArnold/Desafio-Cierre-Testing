.PHONY: test coverage coverage-html total-coverage lint

test:
	@echo "Ejecutando pruebas..."
	go test ./...

coverage:
	@echo "Generando reporte de cobertura..."
	go test -coverprofile=coverage.out ./...

coverage-html:
	@echo "Generando reporte de cobertura en HTML..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

total-coverage:
	@echo "Obteniendo cobertura total del proyecto..."
	go test -cover ./...

lint:
	@echo "Ejecutando Linter..."
	golint ./...



# make "nombre del comando"