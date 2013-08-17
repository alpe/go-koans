package go_koans

import "bytes"
import "io"

func aboutCommonInterfaces() {
  {
    in := new(bytes.Buffer)
    in.WriteString("hello world")

    out := new(bytes.Buffer)

	io.Copy(out, in)
    /*
       Your code goes here.
       Hint, use these resources:

       $ godoc -http=:8081
       $ open http://localhost:8081/pkg/io/
       $ open http://localhost:8081/pkg/bytes/
    */

    assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
  }

  {
    in := new(bytes.Buffer)
    in.WriteString("hello world")

    out := new(bytes.Buffer)

    io.CopyN(out, in, 5)


	  assert(out.String() == "hello") // duplicate only a portion of the io.Reader
  }
}
