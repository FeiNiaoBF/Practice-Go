> 注意：所有的文本内容来自于 [jwasham/coding-interview-university](https://github.com/jwasham/coding-interview-university/blob/main/translations/README-cn.md)

>我只是将其fork下来，方便自己学习。

## 让我们开始吧

好了，说得够多了，让我们学习吧!

但在学习的同时，不要忘记做上面的编码问题！

## 算法复杂度 / Big-O / 渐进分析法

- 这里没有什么需要实施的，你只是在观看视频并记笔记！耶！
- 这里有很多视频，只要看到你理解为止就好了，你随时可以回来复习。
- 如果你不理解背后的所有数学，不要担心。
- 你只需要理解如何用大O表示法来表达算法的复杂度。
- [x] [哈佛大学CS50 - 渐进符号（视频）](https://www.youtube.com/watch?v=iOq5kSKqeR4)
- [ ] [大O符号（通用快速教程）（视频）](https://www.youtube.com/watch?v=V6mKVRU1evU)
- [ ] [大O符号（以及Ω和Θ）- 最佳数学解释（视频）](https://www.youtube.com/watch?v=ei-A_wy5Yxw&index=2&list=PL1BaGV1cIH4UhkL8a9bJGG356covJ76qN)
- [ ] [Skiena（视频）](https://www.youtube.com/watch?v=z1mkCe3kVUA)
- [ ] [加州大学伯克利分校关于大O符号（视频）](https://archive.org/details/ucberkeley_webcast_VIS4YDpuP98)
- [ ] [摊还分析（视频）](https://www.youtube.com/watch?v=B3SpQZaAZP4&index=10&list=PL1BaGV1cIH4UhkL8a9bJGG356covJ76qN)
- [ ] TopCoder（包括递归关系和主定理）：
  - [计算复杂性：第1部分](https://www.topcoder.com/thrive/articles/Computational%20Complexity%20part%20one)
  - [计算复杂性：第2部分](https://www.topcoder.com/thrive/articles/Computational%20Complexity%20part%20two)
- [ ] [速查表](http://bigocheatsheet.com/)
- [x] [[回顾] 在 18 分钟内分析算法（播放列表）（视频）](https://www.youtube.com/playlist?list=PL9xmBV_5YoZMxejjIyFHWa-4nKg6sdoIv)

好吧，差不多就到这里了。

当你阅读《破解编程面试》时，有一个章节专门讲述此事，并在最后进行了一次测验，
以测试你是否能够确定不同算法的运行时间复杂度。这是一个非常全面的复习和测试。

## 数据结构

- ### 数组（Arrays）

  - [ ] 介绍：
    - [数组 CS50 哈佛大学](https://www.youtube.com/watch?v=tI_tIZFyKBw&t=3009s)
    - [数组（视频）](https://www.coursera.org/lecture/data-structures/arrays-OsBSF)
    - [加州大学伯克利分校CS61B - 线性和多维数组（视频）](https://archive.org/details/ucberkeley_webcast_Wp8oiO_CZZE)（从15分32秒开始）
    - [动态数组（视频）](https://www.coursera.org/lecture/data-structures/dynamic-arrays-EwbnV)
    - [嵌套数组（视频）](https://www.youtube.com/watch?v=1jtrQqYpt7g)
  - [ ] 实现一个动态数组（可自动调整大小的可变数组）：
    - [ ] 练习使用数组和指针去编码，并且指针是通过计算去跳转而不是使用索引
    - [ ] 通过分配内存来新建一个原生数据型数组
      - 可以使用 int 类型的数组，但不能使用其语法特性
      - 从大小为16或更大的数（使用2的倍数 —— 16、32、64、128）开始编写
    - [ ] size() —— 数组元素的个数
    - [ ] capacity() —— 可容纳元素的个数
    - [ ] is_empty()
    - [ ] at(index) —— 返回对应索引的元素，且若索引越界则愤然报错
    - [ ] push(item)
    - [ ] insert(index, item) —— 在指定索引中插入元素，并把后面的元素依次后移
    - [ ] prepend(item) —— 可以使用上面的 insert 函数，传参 index 为 0
    - [ ] pop() —— 删除在数组末端的元素，并返回其值
    - [ ] delete(index) —— 删除指定索引的元素，并把后面的元素依次前移
    - [ ] remove(item) —— 删除指定值的元素，并返回其索引（即使有多个元素）
    - [ ] find(item) —— 寻找指定值的元素并返回其中第一个出现的元素其索引，若未找到则返回 -1
    - [ ] resize(new_capacity) // 私有函数
      - 若数组的大小到达其容积，则变大一倍
      - 获取元素后，若数组大小为其容积的1/4，则缩小一半
  - [ ] 时间复杂度
    - 在数组末端增加/删除、定位、更新元素，只允许占 O(1) 的时间复杂度（平摊（amortized）去分配内存以获取更多空间）
    - 在数组任何地方插入/移除元素，只允许 O(n) 的时间复杂度
  - [ ] 空间复杂度
    - 因为在内存中分配的空间邻近，所以有助于提高性能
    - 空间需求 = （大于或等于 n 的数组容积）* 元素的大小。即便空间需求为 2n，其空间复杂度仍然是 O(n)

- ### 链表（Linked Lists）

  - [ ] 介绍：
    - [ ] [链表 CS50 哈佛大学](https://www.youtube.com/watch?v=2T-A_GFuoTo&t=650s) - 这样建立了直观感。
    - [ ] [单链表（视频）](https://www.coursera.org/lecture/data-structures/singly-linked-lists-kHhgK)
    - [ ] [CS 61B - 链表1（视频）](https://archive.org/details/ucberkeley_webcast_htzJdKoEmO0)
    - [ ] [CS 61B - 链表 2（视频）](https://archive.org/details/ucberkeley_webcast_-c4I3gFYe3w)
    - [ ] [[复习] 4分钟了解链表（视频）](https://youtu.be/F8AbOfQwl1c)
  - [ ] [C代码（视频）](https://www.youtube.com/watch?v=QN6FPiD0Gzo)
            - 不是整个视频，只是关于Node结构和内存分配的部分。
  - [ ] 链表 vs 数组：
    - [核心链表与数组（视频）](https://www.coursera.org/lecture/data-structures-optimizing-performance/core-linked-lists-vs-arrays-rjBs9)
    - [在现实世界中，链表与数组的比较（视频）](https://www.coursera.org/lecture/data-structures-optimizing-performance/in-the-real-world-lists-vs-arrays-QUaUd)
  - [ ] [为什么你需要避免使用链表（视频）](https://www.youtube.com/watch?v=YQs6IC-vgmo)
  - [ ] 的确：你需要关于“指向指针的指针”的相关知识：（因为当你传递一个指针到一个函数时，
      该函数可能会改变指针所指向的地址）该页只是为了让你了解“指向指针的指针”这一概念。
      但我并不推荐这种链式遍历的风格。因为，这种风格的代码，其可读性和可维护性太低。
    - [指向指针的指针](https://www.eskimo.com/~scs/cclass/int/sx8.html)
  - [ ] 实现（我实现了使用尾指针以及没有使用尾指针这两种情况）：
    - [ ] size() —— 返回链表中数据元素的个数
    - [ ] empty() —— 若链表为空则返回一个布尔值 true
    - [ ] value_at(index) —— 返回第 n 个元素的值（从0开始计算）
    - [ ] push_front(value) —— 添加元素到链表的首部
    - [ ] pop_front() —— 删除首部元素并返回其值
    - [ ] push_back(value) —— 添加元素到链表的尾部
    - [ ] pop_back() —— 删除尾部元素并返回其值
    - [ ] front() —— 返回首部元素的值
    - [ ] back() —— 返回尾部元素的值
    - [ ] insert(index, value) —— 插入值到指定的索引，并把当前索引的元素指向到新的元素
    - [ ] erase(index) —— 删除指定索引的节点
    - [ ] value_n_from_end(n) —— 返回倒数第 n 个节点的值
    - [ ] reverse() —— 逆序链表
    - [ ] remove_value(value) —— 删除链表中指定值的第一个元素
  - [ ] 双向链表
    - [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/jpGKD/doubly-linked-lists)
    - 并不需要实现

- ### 堆栈（Stack）

  - [ ] [堆栈（视频）](https://www.coursera.org/learn/data-structures/lecture/UdKzQ/stacks)
  - [ ] [[Review] Stacks in 3 minutes (video)](https://youtu.be/KcT3aVgrrpU)
  - [ ] 可以不实现，因为使用数组来实现是微不足道的事

- ### 队列（Queue）

  - [ ] [队列（视频）](https://www.coursera.org/learn/data-structures/lecture/EShpq/queue)
  - [ ] [原型队列/先进先出（FIFO）](https://en.wikipedia.org/wiki/Circular_buffer)
  - [ ] [[Review] Queues in 3 minutes (video)](https://youtu.be/D6gu-_tmEpQ)
  - [ ] 使用含有尾部指针的链表来实现:
    - enqueue(value) —— 在尾部添加值
    - dequeue() —— 删除最早添加的元素并返回其值（首部元素）
    - empty()
  - [ ] 使用固定大小的数组实现：
    - enqueue(value) —— 在可容的情况下添加元素到尾部
    - dequeue() —— 删除最早添加的元素并返回其值
    - empty()
    - full()
  - [ ] 花销：
    - 在糟糕的实现情况下，使用链表所实现的队列，其入列和出列的时间复杂度将会是 O(n)。
        因为，你需要找到下一个元素，以致循环整个队列
    - enqueue：O(1)（平摊（amortized）、链表和数组 [探测（probing）]）
    - dequeue：O(1)（链表和数组）
    - empty：O(1)（链表和数组）

- ### 哈希表（Hash table）

  - [ ] 视频：
    - [ ] [链式哈希表（视频）](https://www.youtube.com/watch?v=0M_kIqhwbFo&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=8)
    - [ ] [Table Doubling 和 Karp-Rabin（视频）](https://www.youtube.com/watch?v=BRO7mVIFt08&index=9&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
    - [ ] [Open Addressing 和密码型哈希（Cryptographic Hashing）（视频）](https://www.youtube.com/watch?v=rvdJDijO2Ro&index=10&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
    - [ ] [PyCon 2010：强大的字典（视频）](https://www.youtube.com/watch?v=C4Kc8xzcA68)
    - [ ] [PyCon 2017：字典更强大（视频）](https://www.youtube.com/watch?v=66P5FMkWoVU)
    - [ ] [(高级) 随机化：通用和完美哈希（视频）](https://www.youtube.com/watch?v=z0lJ2k0sl1g&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&index=11)
    - [ ] [(进阶)完美哈希（Perfect hashing）（视频）](https://www.youtube.com/watch?v=N0COwN14gt0&list=PL2B4EEwhKD-NbwZ4ezj7gyc_3yNrojKM9&index=4)
    - [ ] [[复习]4分钟了解哈希表（视频）](https://youtu.be/knV86FlSXJ8)

  - [ ] 在线课程：
    - [ ] [核心哈希表（视频）](https://www.coursera.org/lecture/data-structures-optimizing-performance/core-hash-tables-m7UuP)
    - [ ] [数据结构（视频）](https://www.coursera.org/learn/data-structures/home/week/4)
    - [ ] [电话簿问题（视频）](https://www.coursera.org/lecture/data-structures/phone-book-problem-NYZZP)
    - [ ] 分布式哈希表：
      - [Dropbox中的即时上传和存储优化（视频）](https://www.coursera.org/lecture/data-structures/instant-uploads-and-storage-optimization-in-dropbox-DvaIb)
      - [分布式哈希表（视频）](https://www.coursera.org/lecture/data-structures/distributed-hash-tables-tvH8H)

  - [ ] 使用线性探测法的数组实现
    - hash(k, m) - m是哈希表的大小
    - add(key, value) - 如果键已存在，则更新值
    - exists(key) - 检查键是否存在
    - get(key) - 获取给定键的值
    - remove(key) - 删除给定键的值

## 树（Trees）

- ### 树-介绍

  - [ ] [树的介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/95qda/trees)
  - [ ] [树遍历（视频）](https://www.coursera.org/lecture/data-structures/tree-traversal-fr51b)
  - [ ] [BFS（广度优先搜索）和DFS（深度优先搜索）（视频）](https://www.youtube.com/watch?v=uWL6FJhq5fM)
    - BFS 笔记
      - 层次遍历（BFS，使用队列）
      - 时间复杂度： O(n)
      - 空间复杂度：最佳情况：O(1)，最坏情况：O(n/2)=O(n)
    - DFS 笔记：
      - 时间复杂度：O(n)
      - 空间复杂度：
        - 最好情况：O(log n) - 树的平均高度
        - 最坏情况：O(n)
      - 中序遍历（DFS：左、节点本身、右）
      - 后序遍历（DFS：左、右、节点本身）
      - 先序遍历（DFS：节点本身、左、右）
  - [ ] [[复习]4分钟内的广度优先搜索（视频）](https://youtu.be/HZ5YTanv5QE)
  - [ ] [[复习] 4分钟内的深度优先搜索（视频）](https://youtu.be/Urx87-NMm6c)
  - [ ] [[复习]11分钟内的树遍历（播放列表）（视频）](https://www.youtube.com/playlist?list=PL9xmBV_5YoZO1JC2RgEi04nLy6D-rKk6b)

- ### 二叉查找树（Binary search trees）：BSTs

  - [ ] [二叉搜索树复习（视频）](https://www.youtube.com/watch?v=x6At0nzX92o&index=1&list=PLA5Lqm4uh9Bbq-E0ZnqTIa8LRaL77ica6)
  - [ ] [介绍（视频）](https://www.coursera.org/learn/data-structures/lecture/E7cXP/introduction)
  - [ ] [MIT（视频）](https://www.youtube.com/watch?v=9Jry5-82I68)
  - C/C++:
    - [ ] [二叉查找树 —— 在 C/C++ 中实现（视频）](https://www.youtube.com/watch?v=COZK7NATh4k&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=28)
    - [ ] [BST 的实现 —— 在堆栈和堆中的内存分配（视频）](https://www.youtube.com/watch?v=hWokyBoo0aI&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=29)
    - [ ] [在二叉查找树中找到最小和最大的元素（视频）](https://www.youtube.com/watch?v=Ut90klNN264&index=30&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [寻找二叉树的高度（视频）](https://www.youtube.com/watch?v=_pnqMz5nrRs&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=31)
    - [ ] [二叉树的遍历 —— 广度优先和深度优先策略（视频）](https://www.youtube.com/watch?v=9RHO6jU--GU&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=32)
    - [ ] [二叉树：层序遍历（视频）](https://www.youtube.com/watch?v=86g8jAQug04&index=33&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [二叉树的遍历：先序、中序、后序（视频）](https://www.youtube.com/watch?v=gm8DUJJhmY4&index=34&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [判断一棵二叉树是否为二叉查找树（视频）](https://www.youtube.com/watch?v=yEwSGhSsT0U&index=35&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
    - [ ] [从二叉查找树中删除一个节点（视频）](https://www.youtube.com/watch?v=gcULXE7ViZw&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=36)
    - [ ] [二叉查找树中序遍历的后继者（视频）](https://www.youtube.com/watch?v=5cPbNCrdotA&index=37&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
  - [ ] 实现：
    - [ ] [insert    // 将值插入树中](https://leetcode.com/problems/insert-into-a-binary-search-tree/submissions/987660183/)
    - [ ] get_node_count // 查找树上的节点数
    - [ ] print_values // 从小到大打印树中节点的值
    - [ ] delete_tree
    - [ ] is_in_tree // 如果值存在于树中则返回 true
    - [ ] [get_height // 以节点为单位返回高度（单个节点的高度为1）](https://www.geeksforgeeks.org/find-the-maximum-depth-or-height-of-a-tree/)
    - [ ] get_min   // 返回树上的最小值
    - [ ] get_max   // 返回树上的最大值
    - [ ] is_binary_search_tree
    - [ ] delete_value
    - [ ] get_successor // 返回给定值的后继者，若没有则返回-1

- ### 堆（Heap） / 优先级队列（Priority Queue） / 二叉堆（Binary Heap）

  - 以树形结构可视化，但通常在存储上是线性的（数组、链表）
  - [ ] [堆（Heap）](https://en.wikipedia.org/wiki/Heap_(data_structure))
  - [ ] [堆简介（视频）](https://www.coursera.org/lecture/data-structures/introduction-2OpTs)
  - [ ] [二叉树（视频）](https://www.coursera.org/learn/data-structures/lecture/GRV2q/binary-trees)
  - [ ] [树高度备注（视频）](https://www.coursera.org/learn/data-structures/supplement/S5xxz/tree-height-remark)
  - [ ] [基本操作（视频）](https://www.coursera.org/learn/data-structures/lecture/0g1dl/basic-operations)
  - [ ] [完全二叉树（视频）](https://www.coursera.org/learn/data-structures/lecture/gl5Ni/complete-binary-trees)
  - [ ] [伪代码（视频）](https://www.coursera.org/learn/data-structures/lecture/HxQo9/pseudocode)
  - [ ] [堆排序 - 跳转到开始部分（视频）](https://youtu.be/odNJmw5TOEE?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3291)
  - [ ] [堆排序（视频）](https://www.coursera.org/lecture/data-structures/heap-sort-hSzMO)
  - [ ] [构建堆（视频）](https://www.coursera.org/lecture/data-structures/building-a-heap-dwrOS)
  - [ ] [MIT：堆和堆排序（视频）](https://www.youtube.com/watch?v=B7hVxCmfPtM&index=4&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [CS 61B Lecture 24：优先队列（视频）](https://archive.org/details/ucberkeley_webcast_yIUFT6AKBGE)
  - [ ] [线性时间构建堆（大顶堆）](https://www.youtube.com/watch?v=MiyLo8adrWw)
  - [ ] [[复习] 13分钟了解堆（视频）](https://www.youtube.com/playlist?list=PL9xmBV_5YoZNsyqgPW-DNwUeT8F8uhWc6)
  - [ ] 实现一个大顶堆：
    - [ ] insert
    - [ ] sift_up —— 用于插入元素
    - [ ] get_max —— 返回最大值但不移除元素
    - [ ] get_size() —— 返回存储的元素数量
    - [ ] is_empty() —— 若堆为空则返回 true
    - [ ] extract_max —— 返回最大值并移除
    - [ ] sift_down —— 用于获取最大值元素
    - [ ] remove(i) —— 删除指定索引的元素
    - [ ] heapify —— 构建堆，用于堆排序
    - [ ] heap_sort() —— 拿到一个未排序的数组，然后使用大顶堆或者小顶堆进行就地排序

## 图（Graphs）

图表可以用来表示计算机科学中的许多问题，所以这一部分很长，就像树和排序一样。

- 笔记:
  - 有4种基本方式在内存里表示一个图:
    - 对象和指针
    - 邻接矩阵
    - 邻接表
    - 邻接图
  - 熟悉以上每一种图的表示法，并了解各自的优缺点
  - 广度优先搜索和深度优先搜索：知道它们的计算复杂度和设计上的权衡以及如何用代码实现它们
  - 遇到一个问题时，首先尝试基于图的解决方案，如果没有再去尝试其他的。

- MIT（视频）：
  - [广度优先搜索](https://www.youtube.com/watch?v=s-CYnVz-uh4&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=13)
  - [深度优先搜索](https://www.youtube.com/watch?v=AfSk24UTFS8&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=14)

- [ ] Skiena 教授的课程 - 很不错的介绍:
  - [ ] [CSE373 2012 - 课程 11 - 图的数据结构（视频）](https://www.youtube.com/watch?v=OiXxhDrFruw&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=11)
  - [ ] [CSE373 2012 - 课程 12 - 广度优先搜索（视频）](https://www.youtube.com/watch?v=g5vF8jscteo&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=12)
  - [ ] [CSE373 2012 - 课程 13 - 图的算法（视频）](https://www.youtube.com/watch?v=S23W6eTcqdY&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b&index=13)
  - [ ] [CSE373 2012 - 课程 14 - 图的算法 (1)（视频）](https://www.youtube.com/watch?v=WitPBKGV0HY&index=14&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
  - [ ] [CSE373 2012 - 课程 15 - 图的算法 (2)（视频）](https://www.youtube.com/watch?v=ia1L30l7OIg&index=15&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)
  - [ ] [CSE373 2012 - 课程 16 - 图的算法 (3)（视频）](https://www.youtube.com/watch?v=jgDOQq6iWy8&index=16&list=PLOtl7M3yp-DV69F32zdK7YJcNXpTunF2b)

- [ ] 图 (复习和其他):

  - [ ] [6.006 单源最短路径问题（视频）](https://www.youtube.com/watch?v=Aa2sqUhIn-E&index=15&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [6.006 Dijkstra算法（视频）](https://www.youtube.com/watch?v=NSHizBK9JD8&t=1731s&ab_channel=MITOpenCourseWare)
  - [ ] [6.006 Bellman-Ford算法（视频）](https://www.youtube.com/watch?v=f9cVS_URPc0&ab_channel=MITOpenCourseWare)
  - [ ] [6.006 加速Dijkstra算法（视频）](https://www.youtube.com/watch?v=CHvQ3q_gJ7E&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb&index=18)
  - [ ] [Aduni：图算法 I - 拓扑排序，最小生成树，Prim算法 - 讲座6（视频）](https://www.youtube.com/watch?v=i_AQT_XfvD8&index=6&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
  - [ ] [Aduni：图算法 II - DFS，BFS，Kruskal算法，Union Find数据结构 - 讲座7（视频）](https://www.youtube.com/watch?v=ufj5_bppBsA&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=7)
  - [ ] [Aduni：图算法 III：最短路径 - 讲座8（视频）](https://www.youtube.com/watch?v=DiedsPsMKXc&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=8)
  - [ ] [Aduni：图算法 IV：几何算法入门 - 讲座9（视频）](https://www.youtube.com/watch?v=XIAQRlNkJAw&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=9)
  - [ ] [CS 61B 2014：加权图（视频）](https://archive.org/details/ucberkeley_webcast_zFbq8vOZ_0k)
  - [ ] [贪婪算法：最小生成树（视频）](https://www.youtube.com/watch?v=tKwnms5iRBU&index=16&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
  - [ ] [强连通分量Kosaraju算法图算法（视频）](https://www.youtube.com/watch?v=RpgcYiky7uw)
  - [ ] [[复习] 最短路径算法（播放列表）16分钟（视频）](https://www.youtube.com/playlist?list=PL9xmBV_5YoZO-Y-H3xIC9DGSfVYJng9Yw)
  - [ ] [[复习] 最小生成树（播放列表）4分钟（视频）](https://www.youtube.com/playlist?list=PL9xmBV_5YoZObEi3Hf6lmyW-CBfs7nkOV)

- 完整的 Coursera 课程:
  - [ ] [图的算法（视频）](https://www.coursera.org/learn/algorithms-on-graphs/home/welcome)

- 我会实现:
  - [ ] DFS 邻接表 (递归)
  - [ ] DFS 邻接表 (栈迭代)
  - [ ] DFS 邻接矩阵 (递归)
  - [ ] DFS 邻接矩阵 (栈迭代)
  - [ ] BFS 邻接表
  - [ ] BFS 邻接矩阵
  - [ ] 单源最短路径问题 (Dijkstra)
  - [ ] 最小生成树
  - 基于 DFS 的算法 (根据上文 Aduni 的视频):
    - [ ] 检查环 (我们会先检查是否有环存在以便做拓扑排序)
    - [ ] 拓扑排序
    - [ ] 计算图中的连通分支
    - [ ] 列出强连通分量
    - [ ] 检查双向图
