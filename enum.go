package p2pubapi

//go:generate stringer -type=Status

type Status int

const (
	None          Status = iota // not defined
	InPreparation               // contract
	InService                   // contract
	Stopped                     // vm
	Configuring                 // vm, storage, fw/lb, mog
	Starting                    // vm, fw/lb
	Running                     // vm, fw/lb, mog
	Stopping                    // vm
	Locked                      // vm, fw/lb
	Attached                    // storage
	NotAttached                 // storage
	Initializing                // storage
	Archiving                   // storage
	Initialized                 // fw/lb
	Configured                  // fw/lb
	Updating                    // fw/lb
)
