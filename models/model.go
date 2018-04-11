//package models
//Author: Sriram Kaushik
//Date: 04/10/2018

package models

import (
	"context"
	"errors"
	"io/ioutil"
	"os/exec"
	"time"
)

type Data struct {
	RawData []byte
}

func (data *Data) RunBySS(output *Data) error {

	//write the input data contents to a file in /var/tmp
	if err := ioutil.WriteFile("/var/tmp/testing-input.txt", data.RawData, 0755); err != nil {
		return errors.New("Could not write to an input file")
	}

	//input file is ready, now we can start stream splitter.

	//set a context

	ctx, cancel := context.WithTimeout(context.Background(), 3100*time.Second)
	defer cancel()

	// set the stream splitter start command.

	cmd := exec.CommandContext(ctx, "service", "stream-splitter", "restart")

	err := cmd.Run() //Run will immediately return (while the command is running on a fork.)

	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("ERROR: Some issue with starting stream-splitter. Configs may be wrong or some plugins are malformed")
	}

	if err != nil {
		//the cmd.Run() returned with non zero exit code.

		return errors.New("ERROR: SS started and exited with non-zero exit code. Is SS running as privileged user?")
	}

	//IF there are no errors then most likely our output file has data that we need to pass back. Let's read that.

	output.RawData, err = ioutil.ReadFile("/var/tmp/testing-output.txt")

	if err != nil {
		//some problem reading the output
		return errors.New("ERROR: Some issue reading the output file. Check in the server")
	}

	if len(output.RawData) < 1 {
		return errors.New("WARNING: NO DATA in the OUTPUT file")
	}

	return nil
}
