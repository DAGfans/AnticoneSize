# AnticoneSize
Calculate anti-cone size by  propagation delay , block rate and security threshold

- Equation:

![Equation](https://latex.codecogs.com/svg.latex?k%28D_%7Bmax%7D%2C%5Cdelta%20%29%20%3A%3D%20min%20%5Cleft%20%5C%7B%20%5Chat%7Bk%7D%5Cin%20N%3A%5Cleft%20%28%20%5Csum%20_%7Bj%3D%5Chat%7Bk%7D&plus;1%7D%5E%7B%5Cinfty%20%7D%7Be%5E%7B-2%20%5Ccdot%20D_%7Bmax%7D%20%5Ccdot%20%5Clambda%20%7D%20%5Ccdot%20%5Cfrac%7B%5Cleft%20%28%20-2%20%5Ccdot%20D_%7Bmax%7D%20%5Ccdot%20%5Clambda%5Cright%20%29%5E%7Bj%7D%7D%7Bj%21%7D%7D%20%5Cright%20%29%3C%20%5Cdelta%20%5Cright%20%5C%7D)


- Equation simplification:  
https://godag.github.io/2018/05/12/dag-throughput

- Paper: 
https://eprint.iacr.org/2018/104.pdf

- 解释:

Phantom基于的假设是在网络的大多数算力被诚实节点控制的情况下, 假设某个矿工在时间点t挖出一个区块, 则t之前一个传播周期以及之后一个传播周期总共两个传播周期内, 网络中最多能允许出现k个区块. 如果超出k个, 说明网络被攻击的风险会增大, 因为很难形成一个比较稳定的最大蓝色集合, 即在DAG中的某个最大的子DAG, 其中每个顶点最多能确认k+1个父顶点, 这个子DAG里面的区块被认为是诚实的.所以这个不等式就是把在挖出至少一个区块的情况下, 网络上出现超过k个区块的概率限定在一个很小的概率δ内.
由于这个概率是建立在至少挖出一个区块的前提下, 所以这个概率是一个条件概率, 设为 P(B|A).
其中, A为至少挖出一个区块的概率, B为挖出k以上个区块的总概率

根据条件概率公式P(B|A)= P(A∩B)/P(A)
由于挖出k个以上的概率必定满足至少挖出一个区块, 所以P(A∩B)/P(A)=P(B)/P(A)

由于2Dmaxλ是平均两个传播周期内的出块数, 满足柏松分布, Possion~P(X=x)=(E^n)/(n!\*e^E)

注意: 
- 左边第二个操作数的求和是从k+1开始的, 因为k以及以下都是正常情况
