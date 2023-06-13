# Inteligência em Colaboração em Pesquisa e Desenvolvimento Tecnológico Pasteur-Fiocruz (IAPF)

Este é o repositório do projeto IAPF, uma aplicação que realiza Graph Machine Learning a partir da ingestão de dados de diversas fontes para realizar inferências que permitam delinear propostas de projetos de pesquisa e de desenvolvimento tecnológico, otimizando o potencial transformador com base nas análises de Graph Machine Learning como suporte a colaboração científica e de desenvolvimento tecnológico entre os Institutos Pasteur da França e a Fundação Oswaldo Cruz do Brasil.

## Estrutura do Projeto

Em termos de Engenharia de Solução, este projeto adota as práticas de Clean Architecture e está organizado da seguinte forma:

- `frontend/`: Esta pasta contém a aplicação frontend escrita com Node.js e Express.
  - `public/`: Arquivos estáticos que são servidos diretamente, como CSS, JavaScript e imagens.
  - `routes/`: Arquivos que lidam com rotas do Express.
  - `views/`: Templates que o Express usa para gerar o HTML.
  - `scripts/`: Scripts adicionais para o ambiente de desenvolvimento.

- `backend/`: Esta pasta contém a aplicação backend escrita em Go.
  - `cmd/`: O ponto de partida para os binários do aplicativo.
  - `pkg/`: Bibliotecas e pacotes que podem ser usados por aplicações externas.
  - `internal/`: Códigos privados que são usados apenas dentro deste projeto.
  - `api/`: Definições de APIs e protocolos.
  - `web/`: Códigos para gerenciamento de web (templating, routing, etc.)
  - `scripts/`: Scripts adicionais para o ambiente de desenvolvimento.

- `comum/`: Esta pasta contém os códigos comuns ou compartilhados entre frontend e backend.

- `docker/`: Esta pasta contém os arquivos Dockerfile para criar imagens Docker.

- `kubernetes/`: Esta pasta contém os arquivos de configuração do Kubernetes.

## Contribuindo

Quando contribuir para este repositório, por favor, discuta primeiro a alteração que deseja fazer via issue, email, ou qualquer outro método com os proprietários deste repositório antes de fazer a alteração.

Durante a fase de engenharia de software, adotaremos o Domain-Driven Design (DDD) para delinear os contextos limitados, domínios e subdomínios. As funcionalidades surgirão dessas divisões e serão executadas por módulos e/ou microsserviços.

Lembre-se de sempre aderir à lógica da Clean Architecture ao expandir a aplicação, para garantir coesão entre toda equipe de desenvolvimento siga as melhores práticas da Clean Architecture, seguindo sempre as seguintes convenções e diretrizes para orientar o desenvolvimento usando Clean Architecture, e predominantemente Go no backend e React no frontend:

## Estrutução dos Módulos da Aplicação

Para implementar um novo módulo na aplicação, você deve seguir a estrutura de pastas e a lógica da Clean Architecture. Aqui está um guia sobre como usar a estrutura de pastas para cada módulo:

1. **Backend (Go)**:

    - Cada módulo deve estar contido em seu próprio diretório dentro da pasta `backend/`. Por exemplo, para um módulo chamado `User`, você criaria uma nova pasta em `backend/User`.

    - Dentro do diretório do módulo, crie arquivos separados para diferentes responsabilidades, como `handler.go` para manipulação de solicitações HTTP, `service.go` para lógica de negócios, e `repository.go` para acesso ao banco de dados.

    - O ponto de entrada para o módulo deve ser um arquivo `main.go` no diretório `cmd/User`.

    - As dependências comuns que podem ser reutilizadas em todo o projeto devem ser colocadas na pasta `pkg/`. 

    - Todo o código privado usado apenas dentro deste módulo deve estar no diretório `internal/`.

2. **Frontend (React)**:

    - Cada módulo deve estar contido em seu próprio diretório dentro da pasta `frontend/`. Por exemplo, para um módulo chamado `UserProfile`, você criaria uma nova pasta em `frontend/UserProfile`.

    - Dentro do diretório do módulo, crie arquivos separados para diferentes responsabilidades, como `UserProfile.js` para o componente React, `UserProfile.test.js` para testes do componente, e `UserProfile.css` para estilos específicos do módulo.

    - Os componentes devem ser desacoplados e reutilizáveis. Se um componente for usado em vários lugares, considere movê-lo para um diretório `components/` na raiz do `frontend/`.

3. **Docker/Kubernetes**:

    - Se o módulo precisar de seus próprios serviços Docker ou Kubernetes, você pode adicionar os arquivos de configuração correspondentes nas pastas `docker/` e `kubernetes/` respectivamente.

Lembre-se, o princípio fundamental da Clean Architecture é que o código deve ser organizado de tal forma que as dependências de código apontem para dentro e as regras de negócio sejam isoladas das preocupações externas, como frameworks e bibliotecas específicas. Isto é crucial para garantir que o sistema seja fácil de manter, testar e expandir.


## Princípios de Design da Clean Architecture:

    - Separe as preocupações: o código que muda por diferentes razões deve ser colocado em partes separadas do sistema.
    - Dependências de código devem apontar para dentro: as camadas externas podem depender das internas, mas nunca o contrário.
    - Abstrações não devem depender de detalhes: detalhes devem depender de abstrações.
    - Isolamento de frameworks e tecnologias: o código do aplicativo deve ser independente de qualquer tecnologia específica para facilitar a manutenção e os testes.

### Nomenclatura de arquivos e uso do Case: 

    - Use nomes significativos e descritivos para variáveis, funções, módulos e arquivos. O nome deve descrever o propósito ou a funcionalidade do elemento nomeado.
    - No Go, use CamelCase para nomes de funções, variáveis e métodos. Para constantes, use ALL_CAPS. Para nomes de pacotes, use letras minúsculas sem sublinhado.
    - Em React, use CamelCase para nomes de componentes e instâncias. Para arquivos de componentes, use PascalCase.

### Funções:

    - Uma função deve fazer apenas uma coisa e fazê-la bem. Isso torna o código mais legível, reutilizável e fácil de testar.
    - Cada função deve ser curta e concisa. Se uma função está se tornando muito grande ou complexa, é um sinal de que ela deve ser dividida em funções menores.
    - Em Go, use uma função de retorno único sempre que possível. Se uma função precisa retornar vários valores, considere encapsulá-los em uma estrutura ou interface.
    - Em React, crie componentes funcionais sempre que possível, pois eles são mais leves e fáceis de testar do que os componentes de classe.

### Módulos:

    - Os módulos devem ser autocontidos e encapsular uma funcionalidade específica. Eles devem ter uma interface clara para interagir com o resto do sistema.
    - Organize os módulos por funcionalidade, não por tipo. Por exemplo, todos os arquivos relacionados a uma funcionalidade específica (como controllers, views, testes, etc.) devem estar no mesmo módulo.
    - Em Go, use pacotes para criar módulos. O nome do pacote deve refletir o propósito do pacote.
    - Em React, use componentes para criar módulos. Cada componente deve ter seu próprio arquivo.

### Disposição dos Arquivos no Repositório:

    - Cada arquivo deve ter um propósito específico e conter código relacionado a uma funcionalidade específica.
    - Separe as funções lógicas em arquivos diferentes. Por exemplo, em Go, você pode ter arquivos separados para manipulação de banco de dados, lógica de negócios, manipulação de API, etc.
    - Em React, cada componente deve estar em seu próprio arquivo. Além disso, separe o código relacionado a funções auxiliares, lógica de estado, estilos, etc., em arquivos diferentes.


Essas diretrizes ajudarão a manter o código limpo, manutenível e testável. Lembre-se de que as diretrizes devem ser flexíveis e adaptáveis às necessidades específicas do projeto e da equipe.


## Licença

Por favor, veja o arquivo [LICENSE](LICENSE) para detalhes.