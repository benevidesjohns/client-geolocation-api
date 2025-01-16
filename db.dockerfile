FROM mysql:8.0

# ADICIONA E EXECUTA O SCRIPT DE CRIAÇÃO DA TABELA "CLIENTS" AO BUILDAR O CONTAINER
COPY internal/infra/database/mysql/create-table.sql /docker-entrypoint-initdb.d/