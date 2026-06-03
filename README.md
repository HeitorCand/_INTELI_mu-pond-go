# Diário de Bordo — API de Figurinhas Copa 2026

---

## Antes de tudo

Boa tarde, professor! Este é o meu diário de bordo do projeto da API de figurinhas. Nele, vou falar sobre as escolhas que fiz, os desafios que enfrentei e o que aprendi durante a implementação. Espero que seja útil para entender meu processo de pensamento e as decisões técnicas que tomei. Se tiver alguma dúvida ou quiser discutir algum ponto, estou à disposição!

---

## Por que Go?

Escolhi Go para este projeto por conta de uma um ponto sendo bem sincero: queria praticar, queria aprender mais sobre GO, e já estava fazendo a atividade que voce passou em aula usando Go, então achei que fazia sentido continuar com a mesma linguagem. Além disso, Go tem uma sintaxe simples e é conhecida por sua performance e facilidade de desenvolvimento de APIs. A comunidade é ativa e tem ótimas bibliotecas para trabalhar com HTTP e bancos de dados, o que facilitou bastante a implementação.


---

## A arquitetura em camadas

Quando li o enunciado e vi "Domain, Repository, Service, Handler", minha primeira reação foi "mas isso não é só uma API REST normal? Por que tanta camada?", mas depois de começar a implementar, fui entendendo o valor disso.

A ideia central é que cada camada só conhece a de baixo dela, nunca pula. O `handler` não sabe que existe um banco de dados. O `service` não sabe que existe HTTP. Isso parece burocracia no começo, mas quando você precisa trocar o banco ou escrever um teste sem subir servidor, você agradece.

A estrutura ficou assim:

```
domain/       -> a linguagem do problema: o que é uma Figurinha, quais tipos existem
repository/   -> a única coisa que fala com o banco
service/      -> onde as regras de negócio vivem
handler/      -> traduz HTTP para o service e o service de volta pra HTTP
```

Obs: a arquiteura em camadas ajuda muito mas em projetos maiores, fica um pouco chato pois para fazer uma coisa simples, tem que passar por várias camadas, mas acho que é um preço justo pra ter um código mais organizado e fácil de manter.

---

## Injeção de dependência manual

Não usei nenhum framework de DI. Tudo é montado no `main.go` na mão:

```go
repo := repository.NewFigurinhaRepository(db)
svc  := service.NewFigurinhaService(repo)
h    := handler.NewFigurinhaHandler(svc)
```

Essa abordagem é super transparente. Você vê exatamente quem depende de quem, e não tem mágica acontecendo. Em projetos pequenos, isso é mais do que suficiente. Em projetos maiores, talvez um framework de DI ajude a organizar as coisas, mas para esse projeto achei que a simplicidade era melhor.

---

## Os erros de domínio

Uma coisa que achei elegante foi separar os erros do service dos status HTTP. O service devolve `ErrFigurinhaNotFound`, e só o handler sabe que isso vira um `404`. A camada de regra de negócio não precisa conhecer o protocolo.

Isso fez bastante sentido na prática. Teve um momento que eu estava implementando o endpoint de atualização e percebi que precisava lidar com o caso de "figurinha não encontrada". Se eu tivesse misturado a lógica de negócio com HTTP, teria que pensar em status code ali, o que me tiraria do fluxo de pensamento sobre a regra de negócio. Com os erros de domínio, eu só me preocupei em devolver um erro claro, e o handler se encarregou de traduzir isso para HTTP. Foi uma separação de preocupações que me ajudou a manter o foco.

> Daqui para baixo foi gerado por IA

## Como rodar

```bash
git clone <url-do-repo>
cd mu-pond-go

go mod tidy
go run main.go
```

O servidor sobe na porta `8080`. O banco `figurinhas.db` é criado automaticamente.

---

## Referência rápida da API

| Método | Rota | O que faz |
|---|---|---|
| `POST` | `/figurinha` | Cria figurinha |
| `GET` | `/figurinha` | Lista todas (aceita `?tipo=` e `?posicao=`) |
| `GET` | `/figurinha/:id` | Busca por ID |
| `PUT` | `/figurinha/:id` | Atualiza |
| `DELETE` | `/figurinha/:id` | Remove |

**Tipos:** `comum` · `brilhante` · `legends_ouro` · `legends_bronze`  
**Posições:** `Goleiro` · `Zagueiro` · `Meio-campista` · `Atacante`

**Exemplo de body (POST/PUT):**
```json
{
  "numero": "BRA 15",
  "tipo": "comum",
  "posicao": "Atacante"
}
```

**Exemplo de resposta:**
```json
{
  "id": 1,
  "numero": "BRA 15",
  "tipo": "comum",
  "posicao": "Atacante",
  "updated_at": "2026-06-03T20:00:00Z",
  "created_at": "2026-06-03T20:00:00Z"
}
```
