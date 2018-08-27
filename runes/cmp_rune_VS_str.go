//THIS WAS COPIED FROM STACK OVERFLOW i do not claim ownership


//Differences - string & []rune:
//
//string is a read-only byte slice, encoded with utf8. Each char in string actually takes 1 ~ 3 bytes, while each rune takes 4 bytes
//For string, both len() and index are based on bytes.
//For []rune, both len() and index are based on rune (or int32).
//Relationships - string & []rune:
//
//When convert from string to []rune, each utf8 char in string becomes a rune.
//Similarly, in the reverse conversion, when convert from []rune to string, each rune becomes a utf8 char in the string.
//Tips:
//
//You can convert between string and []rune, but still they are different, in both type & overall size.


package main

import "fmt"

// string & rune compare,
func stringAndRuneCompare() {
	// string,
	s := "hello你好"

	fmt.Printf("%s, type: %T, len: %d\n", s, s, len(s))
	fmt.Printf("s[%d]: %v, type: %T\n", 0, s[0], s[0])
	li := len(s) - 1 // last index,
	fmt.Printf("s[%d]: %v, type: %T\n\n", li, s[li], s[li])

	// []rune
	rs := []rune(s)
	fmt.Printf("%v, type: %T, len: %d\n", rs, rs, len(rs))
}

func main() {
	stringAndRuneCompare()
}

//The string hello你好 has length 11, because first 5 chars each take 1 byte only, while the last 2 Chinese chars each takes 3 bytes.
//
//Thus, total bytes = 5 * 1 + 2 * 3 = 11
//Since len() on string is based on bytes, thus the first line printed len: 11
//Since index on string is also based on bytes, thus the following 2 lines print values of type uint8 (since byte is an alias type of uint8, in go).
//When convert the string to []rune, it found 7 utf8 chars, thus 7 runes.
//
//Since len() on []rune is based on rune, thus the last line printed len: 7.
//If you operate []rune via index, it will access base on rune.
//Since each rune is from a utf8 char in the original string, thus you can also say both len() and index operation on []rune are based on
// utf8 chars.