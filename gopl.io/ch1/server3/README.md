# 包的管理

跟据书中的提示这一节我花费了很长时间： 因为书中说让浏览器中使用之前的lissajous程序，但是因为导入的包中不能有main包，所以我无法运行。

后来我使用：

module gopl.io/ch1/server3

go 1.22.1

replace gopl.io/ch1/lissajous => ../lissajous

require gopl.io/ch1/lissajous v0.0.0-00010101000000-000000000000

这样的结构并再程序中使用

lissajous.Lissajous(w)

终于成功了
