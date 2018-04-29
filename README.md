# 从TS到Go
该系列教程记录我学习go语言中的知识点。
在熟悉运用node进行开发以后，深刻的认识到了node不足。不擅长CPU密集运行，多线程支持的缺乏。虽然有child-porcess和cluster模块，但并不是从语言上对多线程的支持，虽然好处在于免除了对各种复杂锁的学习，却也失去了编程的内在魅力。在处理复杂的应用时个人感觉还是有着性能的瓶颈。因此需要学习一门强类型支持多线程的语言。思考了有以下几种语言选型
1. **c++** 毕竟不是科班出身，也没有ACM经验，过多的指针以及没有自动垃圾回收，对于我们这种菜鸡还是算了
2. **JAVA** 之前也有看过，但是目前JAVA都出到10了，我估摸着下半辈子研究这玩意儿都追不上大佬们的角度了
3. **PYTHON** 机器学习和大数据的热门语言，然而搞这些门槛很高，无奈渣渣大学，下辈子重新来过吧
4. **php** 别闹 我是正经人

因此思来想去最后选择了go语言，2009年go语言发布开源，截至今日发布的是1.10版本，它很年轻充满着活力与朝气，同样也带着天坑。目前相关的教程较少，大家还处于开荒阶段，还有着许多的东西等着拓展。（其实就是知识点少，背起来快 《go语言程序设计》这本书薄薄的一本比其他书页数少多了）。另一方面go语言在容器技术 微服务大放光彩，著名的docker以及k8s。而我对微服务的如同对夏日少女黑丝的喜爱一般，想一层层揭开她的神秘面纱。因此go这门车可以上。

因为是前端入门，所以对JS的掌握比较好，在此基础上又学习了Typescript，补充完整了类型语言的学习，涉及了OOP 泛型 反射等相关知识，以此为基础，以类比的形式去学习go的相关理念。

## 基础
* [从TS到Go:部署Go开发及调试环境——00](./docs/部署Go开发及调试环境.md)
* [从TS到Go:数组切片映射——01](./docs/数组切片映射.md)