# Go-expert labs concorrencia

## Executando o projeto
Para subir o projeto, a partir da raiz do projeto execute o comando abaixo. Por padrão o fechamento automático de um leilão está setado para 10 minutos. Caso queira diminuir ou aumentar esse tempo, antes de subir o ambiente edite o arquivo .env (cmd/auction) e altere a variável de ambiente AUCTION_INTERVAL.
```bash
docker compose up -d
```
## Testando a api de leilão
Para facilicar os testes use o arquivo api.http (api/api.http), neste arquivo você encontrará os endpoints para:
- Criar um usuário,
- Criar um leilão,
- Ofertar lances para um leilão, enquanto o mesmo ainda estiver ativo.