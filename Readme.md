### 寻路算法

在游戏场景中目标棋子和障碍物都处于运动中，所以在棋子行走过程中，每走一步需要重新计算路径。如果采用标准寻路算法的话，浪费服务器性能。因此采用的是非最优解寻路，最坏情况下找不到路径。

#### 寻路步骤
1. 基准点可以直达目标点，直接返回
2. 取基准点周围8个格子，计算各自权重(格子到目标点距离越近权重越高)
3. 基准点相邻的格子可以直达目标点，直接返回
4. 选择权重最高的格子，判断所在方向每个格子是否直达目标点，如果不能直达，返回上一步寻找其他格子，记录最接近格子
5. 如果没有可达路径，跳到最接近格子

#### 坐标体系

![](https://z3.ax1x.com/2021/08/27/hMoSUg.png)