package main
import "io"
import "fmt"
import "log/slog"
import "strings"
import "os"
func countLetters(r io.Reader) (map[string]int, error){
	buf := make([]byte, 2048)
	out := map[string]int{}
	
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z'){
				out[string(b)]++
			}
		}
		
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func simpleCountLetters() error{
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil{
		return err
	}
	fmt.Println(counts)
	return nil
}

func fileCountLetters() error{
	r, err := os.Open("hello.txt")
	if err != nil{
		return err
	}
	
	counts, err := countLetters(r)
	if err != nil{
		return err
	}
	
	fmt.Println(counts)
	return nil
}

func main(){
	err := simpleCountLetters()
	if err != nil {
		slog.Error("error with simpleCountLetters", "msg", err)
	}
	
	err = fileCountLetters()
	if err != nil{
		slog.Error("error with fileCountLetters", "msg", err)
	}
}