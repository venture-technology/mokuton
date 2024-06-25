<h1 align="center"> 🌬️ Venture </h1>

<h1 align="center"> Somos segurança, velocidade e tecnologia. Somos Venture </h1>

<p align="center">
  <img src="https://i.imgur.com/yieDOSJ.png"/>
</p>

## ⚙️ API Endpoints

Todas as rotas possuem `api/v1`. Antecendo, como prefixo da rota.

### GET /ping

Retorna uma simples mensagem de "pong" para validar funcionamento da aplicação.

**Resposta**

```json
{
    "message": "pong"
}
```
---

### POST /payout

Crie um mock de json para simular a resposta de um microserviço que será acessado via http.

**Parâmetros**

| Nome     | Local | Tipo   | Descrição            |
|----------|-------|--------|----------------------|
| `route`     | body  | string   | rota do microserviço que você está mockando.        |
| `payout`   | body  | json | body response esperado e mockado.      |


**Resposta**

```json
{
    "id": "01904cc9-e236-7dae-b9d5-b91a8abcff1b",
    "payout": {
        "invite": "teste"
    },
    "route": "/invite"
}
```

---

### GET /payout/:uuid

Faça a requisição na rota que acabou de realizar a criação.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `uuid` | uri | string  | UUID do mock. |     

**Resposta**

```json
{
    "id": "01904cc9-e236-7dae-b9d5-b91a8abcff1b",
    "payout": {
        "invite": "teste"
    },
    "route": "/invite"
}
```

---

### DELETE /school

Deleta um mock específico.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `uuid` | uri | string  | UUID do mock. |  

**Resposta**

```json
{
    "message": "payout deleted: 01904cbc-4e8a-76f2-b7a1-2bd37b2df6f0"
}
```

   