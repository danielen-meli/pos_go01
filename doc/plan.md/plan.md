# Plano de execução:

### Server.go:

* Criar um servidor HTTP básico na porta 8080 usando o pacote net/http. (OK)
* Configurar um handler para o endpoint /cotacao.
* Realizar uma requisição GET na API https://economia.awesomeapi.com.br/json/last/USD-BRL.
* Configurar um contexto com timeout de 200ms
* Na resposta da API, extrair o valor do campo "bid".
* Estruturar o JSON para ser enviado ao client.go.
* Usar o pacote database/sql para conectar e manipular o banco de dados.
* Criar um contexto de 10ms para inserir a cotação no banco.

### Client.go:

* Configurar uma requisição HTTP GET para o endpoint /cotacao do server.go. com
  contexto com timeout de 300ms para essa operação.
* Capturar a resposta e extraia o valor do câmbio do JSON.
* Salvar em Arquivo: o valor cotacao.txt no formato: Dólar: {valor}.
* Adicionar lógica para capturar, gerenciar e logar erros para cada estágio onde um timeout pode ocorrer.
