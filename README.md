# VirtualBank Backend

O **VirtualBank Backend** é a parte do sistema responsável por gerenciar as operações financeiras de um banco virtual em constante evolução. Desenvolvido em **Go**, oferece uma estrutura robusta e flexível para lidar com cadastros, transações financeiras, metas e gestão do patrimônio do usuário.

## Funcionalidades

- **Cadastro de Usuários:** Permita que os usuários se cadastrem no sistema, fornecendo informações essenciais como nome, e-mail, senha e número de contato.

- **Transações Financeiras:** Realize transferências seguras entre contas, garantindo uma experiência eficiente e confiável para as operações financeiras dos usuários.

- **Estabelecimento de Metas:** Ajude os usuários a definirem metas financeiras e acompanhem seu progresso, incentivando a gestão financeira responsável.

- **Gestão do Patrimônio:** Proporcione aos usuários uma visão clara do seu patrimônio, incluindo saldos, transações e histórico financeiro.

## Tecnologias Utilizadas

- **Go (Golang):** Linguagem de programação eficiente e moderna, proporcionando um desenvolvimento rápido e confiável.

- **MySQL:** Banco de dados relacional para armazenar e gerenciar os dados de forma segura.

- **RESTful API:** Arquitetura de API que permite a comunicação eficiente entre o backend e o frontend.

## Estrutura do Projeto

- **controllers:** Contém os controladores responsáveis por receber as requisições HTTP e chamar as funções adequadas para processar as operações.

- **repositories:** Abstração para o acesso ao banco de dados, fornecendo métodos para interagir com as tabelas e entidades.

- **models:** Define a estrutura dos modelos de dados usados pelo sistema.

- **services:** Lógica de negócios que coordena as operações entre os controladores e os repositórios.

## Docker Compose

O projeto inclui um arquivo `docker-compose.yml` para simplificar a configuração do banco de dados MySQL usado pelo backend. Certifique-se de ter o Docker instalado e execute:

```bash
docker-compose up -d
