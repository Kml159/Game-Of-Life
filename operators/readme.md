# Definitions

Let there be a $n \times n$ matrix $S$. 

$S^{ij}_t$ denotes $i$’th row’s $j$’th column in $t$’th generation.

$$
S^{ij}_t =
\begin{cases}
1 & \text{alive} \\
0 & \text{dead}
\end{cases}
$$

$N(S^{ij}_t)$
: Returns the neighbors of $S^{ij}_t$.

$L(S^{ij}_t)$
: Returns the live neighbor amount of $S^{ij}_t$.

# Rules


1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
    
    $$
    \forall i, j, \quad \text{if } S^{ij}_t = 1 \text{ and } L(S^{ij}_t) < 2, \\ \text{ then } S^{ij}_{t+1} = 0
    $$
    
2. Any live cell with two or three live neighbours lives on to the next generation.
    
    $$
    \forall i, j, \quad \text{if } S^{ij}_t = 1 \text{ and } L(S^{ij}_t) \in \{2, 3\}, \\ \text{ then } S^{ij}_{t+1} = 1
    $$
    
3. Any live cell with more than three live neighbours dies, as if by overpopulation.
    
    $$
    \forall i, j, \quad \text{if } S^{ij}_t = 1 \text{ and } L(S^{ij}_t) > 3, \text{ then } S^{ij}_{t+1} = 0
    $$
    
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
    
    $$
    \forall i, j, \quad \text{if } S^{ij}_t = 0 \text{ and } L(S^{ij}_t) = 3, \text{ then } S^{ij}_{t+1} = 1
    $$


