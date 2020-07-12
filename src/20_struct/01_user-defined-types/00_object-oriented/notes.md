Go 是面向对象语言

(1)封装
状态-属性
行为-方法
导出/未导出

(2)可重用性
继承-嵌入式

(3)多态
接口

(4)晋升
覆盖

在Golang：
-你创建的不是类，而是类型
-你无需实例化，只需创建类型的值
----------------------
Go is Object Oriented

(1) Encapsulation
state ("fields")
behavior ("methods")
exported / un-exported

(2) Reusability
inheritence ("embedded types")

(3) Polymorphism
interfaces

(4) Overriding
"promotion"

//////////////
Traditional OOP

Classes
-- data structure describing a type of object
-- you can then create "instances"/"objects" from the class/blue-print
-- classes hold both:
==== state / data / fields
==== behavior / methods
-- public / private

Inheritence

//////////////
In Go:
- you don't create classes, you create a type
- you don't instantiate, you create a value of a type