user defined types - we declare a new type, foo
 the underlying type of foo: int

 conversion:int(myAge)
 converting type foo to type int

 THIS CODE IS ONLY FOR EXAMPLE
 IT IS A BAD PRACTICE TO ALIAS TYPES
 one exception: if you need to attach methods to a type
 see the time package for an example of this
     godoc.org/time
     type Duration int64
 Duration has methods attached to it
 
 -------------
 用户定义类型——我们声明一个新的类型，foo
 底层类型foo: int
 转换:int (myAge)
 转换类型foo到类型int
 这段代码只是一个例子
 使用别名类型是一种不好的做法
 有一个例外:如果需要将方法附加到类型
 请参阅time包以获得这方面的示例
 godoc.org/time
 type Duration int64
 Duration有附加到它的方法