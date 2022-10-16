package ApprovalService

import (
	"MessageFilter/utility"
	"fmt"
	"io"
)

// Simulated Approval Service, 50% chance of Approval
func ApprovalService(w io.Writer, message utility.Message) {
	if utility.RandomBool() == true {
		fmt.Fprintf(w, "IMAGE [%v] APPROVED \n%v\n", message.UID, message.Body)
	} else {
		fmt.Fprintf(w, "\n # IMAGE [%v] REJECTED\n", message.UID)
	}

}
