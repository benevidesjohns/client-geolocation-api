**Arquitetura e decisões de design**

Para a arquitetura, decidi utilizar a arquitetura limpa (em camadas), parar separar as responsabilidades e permitir uma melhor manutenibilidade.

Também busquei ao máximo seguir os padrões e as boas práticas de desenvolvimento de APIs RESTful no Golang, bem como o SOLID.

Fiz a implementação de um DI Container, para centralizar a criação e injeção de dependências e permitir a troca de dependências (como banco de dados, interfaces, usecases, repositórios e handlers).

**Bibliotecas de terceiros utilizadas**

As bibliotecas externas que utilizei foram:

- github.com/go-sql-driver/mysql v1.8.1, para criar e estabelecer a conexão.

- github.com/gorilla/mux v1.8.1, para facilitar na configuração das rotas para as requisições HTTP e também por conta do tempo.

- github.com/joho/godotenv v1.5.1, para obter as variáveis de ambientes e utilizá-las para conexão do banco de dados.

- github.com/rs/cors v1.11.1, para definir regras de CORS na API com mais precisão.

- Air, que é uma ferramenta que faz recarregamento em tempo real no Golang.

Utilizei o Air pra acelerar o processo de desenvolvimento, pois essa ferramenta eliminou a necessidade de compiilar e reiniciar manualmente os containers da aplicação sempre que havia alterações no código. Por isso, removi o comando "go build -o" do Dockerfile, visto que o Air iria cuidar da recompilação durante o desenvolvimento.

Em relação ao front-end, implementei alguns componentes, que eram necessários para a manipulação dos dados, como: 

- Formulário para criação e edição;
- Tabela para visualização dos dados;
- Filtros de busca (por ID e cidade).
