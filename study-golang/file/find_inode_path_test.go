package file

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_find_inode_path(t *testing.T) {

	var inode uint64 = 493098715

	path, err := os.Readlink(fmt.Sprintf("/proc/self/fd/%d", inode))
	if err != nil {
		log.Fatalf("Failed to find file with inode %d, error : %s", inode, err)
	}

	fmt.Println("File path:", path)

}
