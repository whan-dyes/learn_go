```go
rm删除一个目录或文件后，提交时用：
git commit -am  "commit message"

在这里说一下git rm和rm的区别，虽然觉得这个问题有点肤浅，但对于刚接触git不久的朋友来说还是有必要的。

用 git rm 来删除文件，同时还会将这个删除操作记录下来；
用 rm 来删除文件，仅仅是删除了物理文件，没有将其从 git 的记录中剔除。

直观的来讲，git rm 删除过的文件，执行 git commit -m "abc" 提交时，会自动将删除该文件的操作提交上去。

而用 rm 命令直接删除的文件，单纯执行 git commit -m "abc" 提交时，则不会将删除该文件的操作提交上去，需要在执行commit的时候，多加一个-a参数，
即rm删除后，需要使用git commit -am "abc"提交才会将删除文件的操作提交上去。



比如：
1）删除文件test.file
# git rm test.file
# git commit -m "delete test.file"
# git push

或者
# rm test.file
# git commit -am "delete test.file"
# git push

2）删除目录work
# git rm work -r -f
# git commit -m "delete work"
# git push
```
