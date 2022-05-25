### net pkg

* use net/http

Response Struct:    
1. Status(string)
2. StatusCode(int)
3. Body(io.ReadCloser) 
    => io.ReadCloser interface
        => Reader interface => io.Reader interface => Read([]byte) (int, err)
        => Closer interface => io.Closer interface => Close() (error)