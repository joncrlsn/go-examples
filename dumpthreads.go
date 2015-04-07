//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	templ "text/template"
)

// This script creates an SSH session and remotely runs a bash script
var dumpThreadsScript = `
#!/bin/bash
# Remotely dumps Java threads a number of times
ssh -o StrictHostKeyChecking=no {{.host}} 'bash -s' <<-END
#!/bin/bash
COUNT={{.dumpCount}}
# Write the thread dumps to a particular location
FILE="/home/{{.pidOwner}}/logs/threads.\$(date +%Y-%m-%d_%H%M%S).{{.host}}"
PID=\$(ps aux | grep -P '(central|blue).*java' | grep -v grep | grep -v flock | egrep -v 'su (central|blue)' | awk '{print \$2}')
#echo "\$FILE"
#echo "\$PID"
for (( c=1; c<=COUNT; c++ )) ; do
    sudo su {{.pidOwner}} -- -c "touch \${FILE}; jstack -l \$PID >> \${FILE}"
    echo "Threads dumped... to \$FILE.  Sleeping for {{.intervalSeconds}} seconds..."
    sleep {{.intervalSeconds}}
done
echo done
END
`

// main generates a script to dumps the Java threads on the given host for the given number of times
// This will need modification for monitoring a Windows-based resource because it creates a
// bash script that is executed remotely.
func main(host, user string, dumpCount int, intervalSeconds int) error {

	// Save script file
	filename := host + "_dumpThreadScript.sh"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	fmt.Println(" Writing script to file : " + filename)

	t := templ.New("Dump threads script")
	templ.Must(t.Parse(dumpThreadsScript))
	context := map[string]string{
		"pidOwner":        user, // user is the account process is running as
		"dumpCount":       strconv.Itoa(dumpCount),
		"intervalSeconds": strconv.Itoa(intervalSeconds),
		"host":            host}
	err = t.Execute(file, context)
	if err != nil {
		return err
	}
	file.Close()

	cmdStr := "./" + host + "_dumpThreadScript.sh"
	log.Printf("Executing: %s \n", cmdStr)
	sshCmd := exec.Command(cmdStr)
	bytes, err := sshCmd.CombinedOutput()
	fmt.Printf("Output:\n %s \n", string(bytes))
	if err != nil {
		return err
	}

	return nil
}
