/* Алгоритм
Шаг 1 − Сначала нам нужно импортировать пакеты fmt и math. Затем определите структуру ребра для представления взвешенного ребра в графе с полями для начального узла, конечного узла и веса ребра.

Шаг 2 − Определите функцию дейкстры(n int, edges []Edge) [][]int, которая принимает количество узлов n и фрагмент реберных объектов в качестве входных данных и возвращает двумерный фрагмент, представляющий матрицу расстояний между всеми парами узлов в графе.

Шаг 3 − Инициализируйте матрицу смежности adj размером n x n и установите все записи в бесконечность, за исключением диагональных записей, которые установлены в 0. Заполните матрицу смежности на основе ребер во входном срезе.

Шаг 4 − Инициализируйте матрицу расстояний dist размером n x n и скопируйте в нее значения из матрицы смежности.

Шаг 5 − Вычислите кратчайший путь между всеми парами узлов, используя три вложенных цикла. Внешний цикл повторяется по всем узлам k, а внутренние циклы повторяются по всем парам узлов i и j.

Шаг 6 − Для каждой пары узлов проверьте, является ли расстояние от i до k плюс расстояние от k до j меньше текущего расстояния от i до j. Если это так, обновите матрицу расстояний новым расстоянием кратчайшего пути и верните матрицу расстояний.

Шаг 7 − Теперь запустите функцию main(). Внутри main() инициализируйте массив, передав ему ребра графика в качестве значений.

Шаг 8 − Затем вызовите функцию dijkastra() и передайте массив, инициализированный выше, в качестве аргументов функции и сохраните полученный результат в переменной.

Шаг 9 − Наконец, выведите полученный результат, который является кратчайшим путем между всеми парами узлов в графе.

Пример
Следующий пример представляет матрицу расстояний между всеми парами узлов в графе, где dist[i][j] представляет расстояние кратчайшего пути от узла i до узла j. Например, dist[0][1] равно 5, что означает, что кратчайший путь от узла 0 к узлу 1 имеет вес 5. Аналогично, dist[2][3] равно 1, что означает, что кратчайший путь от узла 2 к узлу 3 имеет вес 1. */

package main

import (
   "fmt"
   "math"
)

type Edge struct {
   from int
   to   int
   cost int
}

func dijkstra(n int, edges []Edge) [][]int {
   // Initialize adjacency matrix
   adj := make([][]int, n)
   for i := 0; i < n; i++ {
      adj[i] = make([]int, n)
      for j := 0; j < n; j++ {
         if i == j {
            adj[i][j] = 0
         } else {
            adj[i][j] = math.MaxInt32
         }
      }
   }

   // Build adjacency matrix
   for _, e := range edges {
      adj[e.from][e.to] = e.cost
   }

   // Initialize distance matrix
   dist := make([][]int, n)
   for i := 0; i < n; i++ {
      dist[i] = make([]int, n)
      copy(dist[i], adj[i])
   }

   // Compute shortest path between all pairs
   for k := 0; k < n; k++ {
      for i := 0; i < n; i++ {
         for j := 0; j < n; j++ {
            if dist[i][k]+dist[k][j] < dist[i][j] {
               dist[i][j] = dist[i][k] + dist[k][j]
            }
         }
      }
   }
   return dist
}

func main() {
   n := 4
   edges := []Edge{
      {0, 1, 5},
      {0, 2, 10},
      {1, 2, 3},
      {2, 3, 1},
      {3, 0, 2},
   }
   fmt.Println("The given vertices of graph are:", edges)
   dist := dijkstra(n, edges)
   fmt.Println("The shortest distance between all pairs of nodes are:", dist)
}
/* Output
The given vertices of graph are: [{0 1 5} {0 2 10} {1 2 3} {2 3 1} {3 0 2}]
The shortest distance between all pairs of nodes are: [[0 5 8 9] [6 0 3 4] [3 8 0 1] [2 7 10 0]]
Conclusion
We have successfully compiled and executed a go language program to find the shortest path between all pairs of nodes in a graph by using the dijkastra Algorithm. Dijkstra's Algorithm is a powerful tool for computing shortest paths in graphs, and the Go implementation presented in this article provides a clear and simple way to apply the Algorithm to find the shortest path between all pairs of nodes in a graph.*/