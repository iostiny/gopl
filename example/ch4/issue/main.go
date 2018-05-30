// Issues prints a table of Github issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iostiny/gopl/example/ch4/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issue:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

// !-
// Args: repo:golang/go is:open json decode
/* Output
0 issue:
#23331  mspiegel proposal: encoding/json: export the offset method of th
#11046     kurin encoding/json: Decoder internally buffers full input
#19469 chengzhic runtime: temporary object is not garbage collected
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#22369     Splik encoding/json: add the full path to the field in Unmars
#16212 josharian encoding/json: do all reflect work before decoding
#5901        rsc encoding/json: allow override type marshaling
#22752  buyology proposal: encoding/json: add access to the underlying d
#7872  extempora encoding/json: Encoder internally buffers full output
#17609 nathanjsw encoding/json: ambiguous fields are marshalled
#7213  davechene cmd/compile: escape analysis oddity
#22816 ganelon13 encoding/json: include field name in unmarshal error me
#14750 cyberphon encoding/json: parser ignores the case of member names
#20528  jvshahid net/http: connection reuse does not work happily with n
#21092  trotha01 encoding/json: unmarshal into slice reuses element data
#20754       rsc encoding/xml: unmarshal only processes first XML elemen
#24768 billziss- runtime: Fatal Error in heapBitsForObject during GC on
#20206 markdryan encoding/base64: encoding is slow
#25426 josharian cmd/compile: revisit statement boundaries CL peformance
#15808 randall77 cmd/compile: loads/constants not lifted out of loop
#19109  bradfitz proposal: cmd/go: make fuzzing a first class citizen, l
#17244       adg proposal: decide policy for sub-repositories
*/
