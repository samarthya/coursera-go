/*
func main() {
  x := []int {4, 8, 5}
  y := -1
  for _, elt := range x {
    if elt > y {
      y = elt
    }
  }
  fmt.Print(y)
}
*/
/*
 func main() {
  x := [...]int {4, 8, 5}
  y := x[0:2]
  z := x[1:3]
  y[0] = 1
  z[1] = 3
  fmt.Print(x)
}
*/
/*
func main() {
  x := [...]int {1, 2, 3, 4, 5}
  y := x[0:2]
  z := x[1:4]
  fmt.Print(len(y), cap(y), len(z), cap(z))
}
*/
/*
func main() {
  x := map[string]int {
    "ian": 1, "harris": 2}
  for i, j := range x {
    if i == "harris" {
      fmt.Print(i, j)
    }
  }
}
*/
/*
type P struct {
    x string
y int
}
func main() {
  b := P{"x", -1}
  a := [...]P{P{"a", 10},
        P{"b", 2},
        P{"c", 3}}
  for _, z := range a {
    if z.y > b.y {
      b = z
    }
  }
  fmt.Println(b.x)
}
*/
/*
func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}*/
/*
package main;
import (
  "fmt"
  "encoding/json"
)

func main() {
  fmt.Println("Please enter a name");
  var name string;
  fmt.Scan(&name);


  fmt.Println("Please enter a address");
  var addr string;
  fmt.Scan(&addr);

  person :=  map[string]string {
    "name": name,
    "address": addr,
  };

  // fmt.Printf("name %s", name);
  // fmt.Printf("addr %s", addr);
  // fmt.Println(person);

  var data []byte;
  data, err := json.Marshal(person);

  if err != nil {
    fmt.Println(err);
  } else {
    fmt.Println(string(data));
  }

}
*/

/*
package main;
import (
  "fmt"
  "os"
  "strings"
  "bufio"
  "log"
)

func main() {
  fmt.Println("Please enter a file name");
  var file_name string;
  fmt.Scan(&file_name);

  fd, err := os.Open(file_name);

  if err != nil {
    log.Fatal(err);
  }
  defer fd.Close();

  type Person struct {
    first_name string;
    last_name string;
  }
  people := make([]Person, 0);

  scanner := bufio.NewScanner(fd);
  for scanner.Scan() {
    line := scanner.Text();

    nameArr := strings.Split(line, " ");
    people = append(people, Person {
      first_name: nameArr[0],
      last_name: nameArr[1],
    })
  }

  if scanner.Err() != nil {
    log.Fatal(scanner.Err());
  }

  for _, person := range people {
    // fmt.Println(person);
    fmt.Printf("first_name: %s, last_name: %s\n", person.first_name, person.last_name);
  }
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
*/
/*
name has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters)
*/
/*
type name struct {
	fname []byte
	lname []byte
}

func main() {
	fmt.Print("File? ")

	var filename string
	fmt.Scanln(&filename)

	file, _ := os.Open(filename)
	defer file.Close()

	names := make([]name, 0)

	s := bufio.NewScanner(file)
	for s.Scan() {
		s := strings.Split(s.Text(), " ")
		n := name{fname: []byte(s[0]), lname: []byte(s[1])}

		names = append(names, n)
	}

	for _, n := range names {
		fmt.Printf("first name: %s, last name: %s\n", n.fname, n.lname)
	}

}
*/
/*
package main

import "fmt"

func slove( s string) {
	var Size = len(s)
	if(Size<=2){
		fmt.Println("Not Found!")
	}else{
		for i:=0;i<Size;i++{
			if(i==0){
				if(s[i]=='i'){
					continue
				}else if(s[i]=='I'){
					continue
				}else{
					fmt.Println("Not Found!")
					return
				}
			}else if(i<Size-1){
				if(s[i]=='a'){
					continue
				}
			}else if(i==Size-1){
				if(s[i]=='n'){
					fmt.Println("Found!")
					return
				}else if(s[i]=='N'){
					fmt.Println("Found!")
					return
				}else{
					fmt.Println("Not Found!")
					return
				}

			}

		}
		fmt.Println("Not Found!")
	}
}

func main(){
	var str1, str2 string
	fmt.Println("Enter two strings and open them with spaces : ")
	fmt.Scanf("%s %s",&str1,&str2)
	fmt.Println("the first string is : ")
	slove(str1)
	fmt.Println("the second string is : ")
	slove(str2)
}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}

	var filename string
	fmt.Print("Enter a filename which contains a list of names: ")
	_, err := fmt.Scan(&filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	nameSlice := make([]Name, 0, 1)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name := strings.Fields(scanner.Text())
		nameSlice = append(nameSlice, Name{fname: name[0], lname: name[1]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, n := range nameSlice {
		fmt.Printf("Name: %s %s\n", n.fname, n.lname)
	}
}
