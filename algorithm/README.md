> 注意：所有的文本内容来自于 [jwasham/coding-interview-university](https://github.com/jwasham/coding-interview-university/blob/main/translations/README-cn.md)
> 我只是将其fork下来，方便自己学习。

## algorithm 算法

- ### 二分查找（Binary search）

  - [ ] [二分查找（视频）](https://www.youtube.com/watch?v=D5SrAga1pno)
  - [ ] [二分查找（视频）](https://www.khanacademy.org/computing/computer-science/algorithms/binary-search/a/binary-search)
  - [ ] [详情](https://www.topcoder.com/community/data-science/data-science-tutorials/binary-search/)
  - [ ] [蓝图](https://leetcode.com/discuss/general-discussion/786126/python-powerful-ultimate-binary-search-template-solved-many-problems)
  - [ ] [【复习】四分钟二分查找(视频)](https://youtu.be/fDKIpRe8GW4)
  - [ ] 实现：
    - 二分查找（在一个已排序好的整型数组中查找）
    - 迭代式二分查找

- ### 按位运算（Bitwise operations）

  - [ ] [Bits 速查表](https://github.com/jwasham/coding-interview-university/blob/main/extras/cheat%20sheets/bits-cheat-sheet.pdf) ── 你需要知道大量2的幂数值（从2^1 到 2^16 及 2^32）
  - [ ] 好好理解位操作符的含义：&、|、^、~、>>、<<
    - [ ] [字码（words）](https://en.wikipedia.org/wiki/Word_(computer_architecture))
    - [ ] 好的介绍：
            [位操作（视频）](https://www.youtube.com/watch?v=7jkIUgLC29I)
    - [ ] [C 语言编程教程 2-10：按位运算（视频）](https://www.youtube.com/watch?v=d0AwjSpNXR0)
    - [ ] [位操作](https://en.wikipedia.org/wiki/Bit_manipulation)
    - [ ] [按位运算](https://en.wikipedia.org/wiki/Bitwise_operation)
    - [ ] [Bithacks](https://graphics.stanford.edu/~seander/bithacks.html)
    - [ ] [位元抚弄者（The Bit Twiddler）](http://bits.stephan-brumme.com/)
    - [ ] [交互式位元抚弄者（The Bit Twiddler Interactive）](http://bits.stephan-brumme.com/interactive.html)
    - [ ] [位操作技巧（Bit Hacks）（视频）](https://www.youtube.com/watch?v=ZusiKXcz_ac)
    - [ ] [练习位操作](https://pconrad.github.io/old_pconrad_cs16/topics/bitOps/)
  - [ ] 一补数和补码
    - [二进制：利 & 弊（为什么我们要使用补码）（视频）](https://www.youtube.com/watch?v=lKTsv6iVxV4)
    - [一补数（1s Complement）](https://en.wikipedia.org/wiki/Ones%27_complement)
    - [补码（2s Complement）](https://en.wikipedia.org/wiki/Two%27s_complement)
  - [ ] 计算置位（Set Bits）
    - [计算一个字节中置位（Set Bits）的四种方式（视频）](https://youtu.be/Hzuzo9NJrlc)
    - [计算比特位](https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan)
    - [如何在一个 32 位的整型中计算置位（Set Bits）的数量](http://stackoverflow.com/questions/109023/how-to-count-the-number-of-set-bits-in-a-32-bit-integer)
  - [ ] 交换值：
    - [交换（Swap）](http://bits.stephan-brumme.com/swap.html)
  - [ ] 绝对值：
    - [绝对整型（Absolute Integer）](http://bits.stephan-brumme.com/absInteger.html)

## 排序（Sorting）

- [ ] 笔记:
  - 实现各种排序，知道每种排序的最坏、最好和平均的复杂度分别是什么场景:
    - 不要用冒泡排序 - 效率太差 - 时间复杂度 O(n^2), 除非 n <= 16
  - [ ] 排序算法的稳定性 ("快排是稳定的么?")
    - [排序算法的稳定性](https://en.wikipedia.org/wiki/Sorting_algorithm#Stability)
    - [排序算法的稳定性](http://stackoverflow.com/questions/1517793/stability-in-sorting-algorithms)
    - [排序算法的稳定性](http://www.geeksforgeeks.org/stability-in-sorting-algorithms/)
    - [排序算法 - 稳定性](http://homepages.math.uic.edu/~leon/cs-mcs401-s08/handouts/stability.pdf)
  - [ ] 哪种排序算法可以用链表？哪种用数组？哪种两者都可？
    - 并不推荐对一个链表排序，但归并排序是可行的.
    - [链表的归并排序](http://www.geeksforgeeks.org/merge-sort-for-linked-list/)

- 关于堆排序，请查看前文堆的数据结构部分。堆排序很强大，不过是非稳定排序。

- [ ] [Sedgewick ── 归并排序（5个视频）](https://www.coursera.org/learn/algorithms-part1/home/week/3)
  - [ ] [1. 归并排序（Mergesort）](https://www.coursera.org/lecture/algorithms-part1/mergesort-ARWDq)
  - [ ] [2. 自底向上的归并排序（Bottom up Mergesort）](https://www.coursera.org/learn/algorithms-part1/lecture/PWNEl/bottom-up-mergesort)
  - [ ] [3. 排序复杂性（Sorting Complexity）](https://www.coursera.org/lecture/algorithms-part1/sorting-complexity-xAltF)
  - [ ] [4. 比较器（Comparators）](https://www.coursera.org/lecture/algorithms-part1/comparators-9FYhS)
  - [ ] [5. 稳定性（Stability）](https://www.coursera.org/learn/algorithms-part1/lecture/pvvLZ/stability)

- [ ] [Sedgewick ── 快速排序（4个视频）](https://www.coursera.org/learn/algorithms-part1/home/week/3)
  - [ ] [1. 快速排序（Quicksort）](https://www.coursera.org/lecture/algorithms-part1/quicksort-vjvnC)
  - [ ] [2. 选择排序（Selection）](https://www.coursera.org/lecture/algorithms-part1/selection-UQxFT)
  - [ ] [3. 重复键（Duplicate Keys）](https://www.coursera.org/lecture/algorithms-part1/duplicate-keys-XvjPd)
  - [ ] [4. 系统排序（System Sorts）](https://www.coursera.org/lecture/algorithms-part1/system-sorts-QBNZ7)

- [ ] 加州大学伯克利分校：
  - [ ] [CS 61B Lecture 29：排序 I（视频）](https://archive.org/details/ucberkeley_webcast_EiUvYS2DT6I)
  - [ ] [CS 61B Lecture 30：排序 II（视频）](https://archive.org/details/ucberkeley_webcast_2hTY3t80Qsk)
  - [ ] [CS 61B Lecture 32：排序 III（视频）](https://archive.org/details/ucberkeley_webcast_Y6LOLpxg6Dc)
  - [ ] [CS 61B Lecture 33：排序 V（视频）](https://archive.org/details/ucberkeley_webcast_qNMQ4ly43p4)
  - [ ] [CS 61B 2014-04-21：基数排序（视频）](https://archive.org/details/ucberkeley_webcast_pvbBMd-3NoI)

- [ ] [冒泡排序（视频）](https://www.youtube.com/watch?v=P00xJgWzz2c&index=1&list=PL89B61F78B552C1AB)
- [ ] [冒泡排序分析（视频）](https://www.youtube.com/watch?v=ni_zk257Nqo&index=7&list=PL89B61F78B552C1AB)
- [ ] [插入排序 & 归并排序（视频）](https://www.youtube.com/watch?v=Kg4bqzAqRBM&index=3&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
- [ ] [插入排序（视频）](https://www.youtube.com/watch?v=c4BRHC7kTaQ&index=2&list=PL89B61F78B552C1AB)
- [ ] [归并排序（视频）](https://www.youtube.com/watch?v=GCae1WNvnZM&index=3&list=PL89B61F78B552C1AB)
- [ ] [快排（视频）](https://www.youtube.com/watch?v=y_G9BkAm6B8&index=4&list=PL89B61F78B552C1AB)
- [ ] [选择排序（视频）](https://www.youtube.com/watch?v=6nDMgr0-Yyo&index=8&list=PL89B61F78B552C1AB)

- [ ] 归并排序代码：
  - [ ] [使用外部数组（C语言）](http://www.cs.yale.edu/homes/aspnes/classes/223/examples/sorting/mergesort.c)
  - [ ] [使用外部数组（Python语言）](https://github.com/jwasham/practice-python/blob/master/merge_sort/merge_sort.py)
  - [ ] [对原数组直接排序（C++）](https://github.com/jwasham/practice-cpp/blob/master/merge_sort/merge_sort.cc)
- [ ] 快速排序代码：
  - [ ] [实现（C语言）](http://www.cs.yale.edu/homes/aspnes/classes/223/examples/randomization/quick.c)
  - [ ] [实现（C语言）](https://github.com/jwasham/practice-c/blob/master/quick_sort/quick_sort.c)
  - [ ] [实现（Python语言）](https://github.com/jwasham/practice-python/blob/master/quick_sort/quick_sort.py)

- [ ] [[Review] Sorting (playlist) in 18 minutes](https://www.youtube.com/playlist?list=PL9xmBV_5YoZOZSbGAXAPIq1BeUf4j20pl)
  - [ ] [Quick sort in 4 minutes (video)](https://youtu.be/Hoixgm4-P4M)
  - [ ] [Heap sort in 4 minutes (video)](https://youtu.be/2DmK_H7IdTo)
  - [ ] [Merge sort in 3 minutes (video)](https://youtu.be/4VqmGXwpLqc)
  - [ ] [Bubble sort in 2 minutes (video)](https://youtu.be/xli_FI7CuzA)
  - [ ] [Selection sort in 3 minutes (video)](https://youtu.be/g-PGLbMth_g)
  - [ ] [Insertion sort in 2 minutes (video)](https://youtu.be/JU767SDMDvA)

- [ ] 实现:
  - [ ] 归并：平均和最差情况的时间复杂度为 O(n log n)。
  - [ ] 快排：平均时间复杂度为 O(n log n)。
  - 选择排序和插入排序的最坏、平均时间复杂度都是 O(n^2)。
  - 关于堆排序，请查看前文堆的数据结构部分。

- [ ] 有兴趣的话，还有一些补充，但并不是必须的:
  - [Sedgewick──基数排序 (6个视频)](https://www.coursera.org/learn/algorithms-part2/home/week/3)
    - [ ] [1. Java 中的字符串](https://www.coursera.org/learn/algorithms-part2/lecture/vGHvb/strings-in-java)
    - [ ] [2. 键值索引计数（Key Indexed Counting）](https://www.coursera.org/learn/algorithms-part2/lecture/2pi1Z/key-indexed-counting)
    - [ ] [3. Least Significant Digit First String Radix Sort](https://www.coursera.org/learn/algorithms-part2/lecture/c1U7L/lsd-radix-sort)
    - [ ] [4. Most Significant Digit First String Radix Sort](https://www.coursera.org/learn/algorithms-part2/lecture/gFxwG/msd-radix-sort)
    - [ ] [5. 3中基数快速排序](https://www.coursera.org/learn/algorithms-part2/lecture/crkd5/3-way-radix-quicksort)
    - [ ] [6. 后继数组（Suffix Arrays）](https://www.coursera.org/learn/algorithms-part2/lecture/TH18W/suffix-arrays)
  - [ ] [基数排序](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#radixSort)
  - [ ] [基数排序（视频）](https://www.youtube.com/watch?v=xhr26ia4k38)
  - [ ] [基数排序, 计数排序 (线性时间内)（视频）](https://www.youtube.com/watch?v=Nz1KZXbghj8&index=7&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
  - [ ] [随机算法: 矩阵相乘, 快排, Freivalds' 算法（视频）](https://www.youtube.com/watch?v=cNB2lADK3_s&index=8&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)
  - [ ] [线性时间内的排序（视频）](https://www.youtube.com/watch?v=pOKy3RZbSws&list=PLUl4u3cNGP61hsJNdULdudlRL493b-XZf&index=14)

总结一下，这是[15种排序算法](https://www.youtube.com/watch?v=kPRA0W1kECg)的可视化表示。
如果你需要有关此主题的更多详细信息，请参阅“[一些主题的额外内容](#一些主题的额外内容)”中的“排序”部分。

- ### 递归（Recursion）

  - [ ] Stanford 大学关于递归 & 回溯的课程:
    - [ ] [课程 8 | 抽象编程（视频）](https://www.youtube.com/watch?v=gl3emqCuueQ&list=PLFE6E58F856038C69&index=8)
    - [ ] [课程 9 | 抽象编程（视频）](https://www.youtube.com/watch?v=uFJhEPrbycQ&list=PLFE6E58F856038C69&index=9)
    - [ ] [课程 10 | 抽象编程（视频）](https://www.youtube.com/watch?v=NdF1QDTRkck&index=10&list=PLFE6E58F856038C69)
    - [ ] [课程 11 | 抽象编程（视频）](https://www.youtube.com/watch?v=p-gpaIGRCQI&list=PLFE6E58F856038C69&index=11)
  - 什么时候适合使用
  - 尾递归会更好么?
    - [ ] [什么是尾递归以及为什么它如此糟糕?](https://www.quora.com/What-is-tail-recursion-Why-is-it-so-bad)
    - [ ] [尾递归（视频）](https://www.coursera.org/lecture/programming-languages/tail-recursion-YZic1)
  - [ ] [解决任何递归问题的5个简单步骤（视频）](https://youtu.be/ngCos392W4w)

    回溯蓝图: [Java](https://leetcode.com/problems/combination-sum/discuss/16502/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning))
 [Python](https://leetcode.com/problems/combination-sum/discuss/429538/General-Backtracking-questions-solutions-in-Python-for-reference-%3A)

- ### 动态规划（Dynamic Programming）

  - 在你的面试中或许没有任何动态规划的问题，
    但能够知道一个题目可以使用动态规划来解决是很重要的。
  - 这一部分会有点困难，每个可以用动态规划解决的问题都必须先定义出递推关系，要推导出来可能会有点棘手。
  - 我建议先阅读和学习足够多的动态规划的例子，以便对解决 DP 问题的一般模式有个扎实的理解。
  - [ ] 视频:
    - [ ] [Skiena：CSE373 2020 - 讲座19 - 动态规划简介（视频）](https://www.youtube.com/watch?v=wAA0AMfcJHQ&list=PLOtl7M3yp-DX6ic0HGT0PUX_wiNmkWkXx&index=18)
    - [ ] [Skiena：CSE373 2020 - 讲座20 - 编辑距离（视频）](https://www.youtube.com/watch?v=T3A4jlHlhtA&list=PLOtl7M3yp-DX6ic0HGT0PUX_wiNmkWkXx&index=19)
    - [ ] [Skiena：CSE373 2020 - 讲座20 - 编辑距离（续）（视频）](https://www.youtube.com/watch?v=iPnPVcZmRbE&list=PLOtl7M3yp-DX6ic0HGT0PUX_wiNmkWkXx&index=20)
    - [ ] [Skiena：CSE373 2020 - 讲座21 - 动态规划（视频）](https://www.youtube.com/watch?v=2xPE4Wq8coQ&list=PLOtl7M3yp-DX6ic0HGT0PUX_wiNmkWkXx&index=21)
    - [ ] [Skiena：CSE373 2020 - 讲座22 - 动态规划和复习（视频）](https://www.youtube.com/watch?v=Yh3RzqQGsyI&list=PLOtl7M3yp-DX6ic0HGT0PUX_wiNmkWkXx&index=22)
    - [ ] [Simonson：动态规划 0（从59:18开始）（视频）](https://youtu.be/J5aJEcOr6Eo?list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&t=3558)
    - [ ] [Simonson：动态规划 I - 第11讲（视频）](https://www.youtube.com/watch?v=0EzHjQ_SOeU&index=11&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm)
    - [ ] [Simonson：动态规划 II - 第12讲（视频）](https://www.youtube.com/watch?v=v1qiRwuJU7g&list=PLFDnELG9dpVxQCxuD-9BSy2E7BWY3t5Sm&index=12)
    - [ ] 单独的动态规划问题列表（每个都很短）:
            [动态规划（视频）](https://www.youtube.com/playlist?list=PLrmLmBdmIlpsHaNTPP_jHHDx_os9ItYXr)
  - [ ] 耶鲁课程笔记:
    - [ ] [动态规划](http://www.cs.yale.edu/homes/aspnes/classes/223/notes.html#dynamicProgramming)
  - [ ] Coursera 课程:
    - [ ] [RNA 二级结构问题（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/80RrW/the-rna-secondary-structure-problem)
    - [ ] [动态规划算法（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/PSonq/a-dynamic-programming-algorithm)
    - [ ] [DP 算法描述（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/oUEK2/illustrating-the-dp-algorithm)
    - [ ] [DP 算法的运行时间（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/nfK2r/running-time-of-the-dp-algorithm)
    - [ ] [DP vs 递归实现（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/M999a/dp-vs-recursive-implementation)
    - [ ] [全局成对序列排列（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/UZ7o6/global-pairwise-sequence-alignment)
    - [ ] [本地成对序列排列（视频）](https://www.coursera.org/learn/algorithmic-thinking-2/lecture/WnNau/local-pairwise-sequence-alignment)
