# Challenge DNS Resolver

O desafio consiste em construir seu próprio servidor DNS.

## Passo 1

Um resolver DNS envia uma mensagem DNS através do protocolo UDP para diferentes servidores. O formato da mensagem é definido [Seção 4.1 do documento RFC](https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.1).

Cada mensagem é dividida em 5 seções diferentes:

```
    +---------------------+
    |        Header       |
    +---------------------+
    |       Question      |
    +---------------------+
    |        Answer       |
    +---------------------+
    |      Authority      |
    +---------------------+
    |      Additional     |
    +---------------------+
```

## Header

O formato do header é definido [Seção 4.1.1 do documento RFC](https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.1). Ele tem o seguinte formato:

```

                                    1  1  1  1  1  1
      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                      ID                       |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    QDCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    ANCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    NSCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    ARCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
```

Onde cada campo:

1. ID : Um ID de 16 bits da consulta
2. Flags: Uma área de 16 bits na memória após o ID que armazena vários indicadores (determinam se a consulta é uma resposta, se houve erro, etc.)
3. QDCOUNT: Um número inteiro de 16 bits que indica a quantidade de perguntas.
4. ANCOUNT: Um número inteiro de 16 bits que indica a quantidade de respostas.
5. NSCOUNT: Um número inteiro de 16 bits que indica a quantidade de autoridades.
6. ARCOUNT: Um número inteiro de 16 bits que indica a quantidade de adicionais.

### OPCODE

O campo `OPCODE` é um valor de 4 bits que especifica o tipo de operação da mensagem DNS. Ele é definido pelo originador da consulta e copiado na resposta.

Valores possíveis:

- `0` = `QUERY` (consulta padrão)
- `1` = `IQUERY` (consulta inversa)
- `2` = `STATUS` (solicitação de status do servidor)
- `3` a `15` = reservados para uso futuro

No campo de flags, o `OPCODE` ocupa os bits 14 a 11, imediatamente após o bit `QR`:

```text
|QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
```

### RCODE

O campo `RCODE` é um conjunto de 4 bits localizado no final dos flags do cabeçalho DNS. Ele indica o resultado da operação solicitada.

- `0` = `NOERROR` (sucesso)
- `1` = `FORMERR` (consulta mal formatada)
- `2` = `SERVFAIL` (falha no servidor)
- `3` = `NXDOMAIN` (domínio inexistente)
- `4` = `NOTIMP` (operação não implementada)
- `5` = `REFUSED` (consulta recusada)

## Pergunta

O formato da pergunta é definido [Seção 4.1.2 do documento RFC](https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.2). Ele tem o seguinte formato:

```
                                    1  1  1  1  1  1
      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                                               |
    /                     QNAME                     /
    /                                               /
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                     QTYPE                     |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                     QCLASS                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
```

Onde cada campo:

1. QNAME: Um espaço de memória de tamanho variável que armazena o nome de domínio que precisa ser resolvido.
2. QTYPE: Um número inteiro de 16 bits que indica o tipo de registro, como A, NS, CNAME, etc.
3. QCLASS: Um número inteiro de 16 bits que indica a classe da consulta.

### Criando a mensagem

Para enviar esses dados, precisamos converter o cabeçalho e a pergunta em sua representação em bytes e enviá-los pela rede. Não precisamos adicionar as outras partes da mensagem, pois elas são recebidas na resposta do servidor.

## Comandos disponíveis

```bash
# executa a consulta DNS com um domínio informado
make run DOMAIN=google.com

# roda a suíte de testes
make test

# compila o binário
make build
```
