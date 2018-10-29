package variable

import(
  . "fmt"
  "unsafe"
  "sort"
)
// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
  ID string
  Name string
  Address string
}
func Mymap(){
  Println("--------------------------Mymap-------------------------------------")
  map1()
  map2()
  map3()
  map4()
}
func map1(){
  Println(".........map1..........")
  //声明并初始化一个 map
  myMap := map[string] PersonInfo{
    "1234": PersonInfo{"1", "Jack", "Room 101,..."},
  }
  Println("myMap=",myMap)
  Println()
}
func map2(){
  Println(".........map2..........")
  var m1 map[string]string
  var m2 map[string]string = map[string]string{}      // m2 := map[string]string{}
  var m3 map[string]string = make(map[string]string, 10)  // m3 := make(map[string]string)

  //m1["1"] = "1"   // panic: assignment to entry in nil map
  m2["2"] = "2"
  m3["3"] = "3"

  Println("m1:",m1," m2:",m2," m3:",m3)

  for key, value := range m1 { Println("Key:", key, "Value:", value) }
  for key, value := range m2 { Println("Key:", key, "Value:", value) }
  for key, value := range m3 { Println("Key:", key, "Value:", value) }

  s1 := m1["1"]
  s2 := m2["2"]
  s3 := m3["3"]

  Printf("val=%s,%s,%s\n", s1, s2, s3)
  Printf("len=%d,%d,%d\n", len(m1), len(m2), len(m3))
  Printf("size=%d,%d,%d\n", unsafe.Sizeof(m1), unsafe.Sizeof(m2), unsafe.Sizeof(m3))

  // PrintMemory(unsafe.Pointer(&m1), 8)
  // PrintMemory(unsafe.Pointer(&m2), 8)
  // PrintMemory(unsafe.Pointer(&m3), 8)
}

func map3(){
  Println("............map3.............")
  var personDB map[string] PersonInfo
  personDB = make(map[string] PersonInfo)
  // 往这个map里插入几条数据
  personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
  personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
  // 从这个map查找键为"1234"的信息
  person, ok := personDB["12345"]
  // ok是一个返回的bool型，返回true表示找到了对应的数据
  if ok {
    Println("Found person", person.Name, "with ID 12345.")
    } else {
      Println("Did not find person with ID 1234.")
    }

    delete(personDB,"12345")
    person, ok = personDB["12345"]
    // ok是一个返回的bool型，返回true表示找到了对应的数据
    if ok {
      Println("Found person", person.Name, "with ID 12345.")
      } else {
        Println("Did not find person with ID 12345.")
      }

      var m map[int]map[int]string
      m=make(map[int]map[int]string)
      a,ok:=m[1][1]
      if !ok{
        m[1]=make(map[int]string)
      }
      m[1][1]="GOOD"
      a,ok=m[1][1]
      Println(a,ok)
    }
    func map4(){
      Println("............map4.............")
      // var m=map[int]int{}
      var m=make(map[int]int)

      for i:=0;i<10;i++{
        m[i]=i+10
      }
      Println(m)

      for k,v:=range m {
        Print(k,":",v," ")
      }
      Println()

      for i:=0;i<5;i++{
        delete(m,i)
      }
      Println(m)
      s:=make([]int,len(m))
      i:=0
      for _,v:=range m{
        s[i]=v
        i++
      }
      Println(s)
      sort.Ints(s)
      Println(s)
    }
