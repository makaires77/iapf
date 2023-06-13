O  arquivo neo4j.go é responsável pela comunicação do seu aplicativo com o banco de dados Neo4j. Aqui está um exemplo de como isso pode ser estruturado, da seguinte forma implementada em código:

A função NewNeo4j é usada para criar uma nova instância de um cliente Neo4j. Ela faz isso criando um novo driver Neo4j com a URI, o nome de usuário e a senha fornecidos. Se algo der errado ao criar o driver, a função retorna um erro.

A função Close é usada para fechar o driver Neo4j quando você terminar de usá-lo. É importante sempre fechar o driver quando você terminar de usá-lo para liberar recursos.

A função ExecuteQuery é usada para executar uma consulta Cypher no Neo4j. Ela faz isso iniciando uma nova sessão, executando a consulta na transação de gravação e retornando o resultado. Se algo der errado ao executar a consulta, a função registra o erro e o retorna.

ExecuteQueryWithRetry: Esta função tentará executar uma consulta fornecida três vezes antes de retornar um erro. Ela usa um backoff exponencial para aumentar o tempo de espera entre cada tentativa.

CreateNode: Esta função cria um nó com um rótulo e propriedades fornecidos.

CreateRelationship: Esta função cria um relacionamento de um tipo especificado entre dois nós. Ele usa os IDs dos nós para identificá-los.

MatchNodes: Esta função busca nós com um rótulo e propriedades específicos.

Cada função usa a função ExecuteQueryWithRetry para executar suas consultas, o que significa que elas se beneficiarão da lógica de re-tentativa. Além disso, se algo der errado ao executar a consulta, a função registra o erro e o retorna.