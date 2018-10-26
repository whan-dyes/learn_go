package variable
import "fmt"
func GetName() (firstName, lastName, nickName string) {
  return "May", "Chan", "Chibi Maruko"
}

func AnonymityVar(){
var nickName string

  _,_,nickName=GetName()
  fmt.Println("nickName=",nickName)
}
